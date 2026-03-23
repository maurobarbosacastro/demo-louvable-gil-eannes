package pt.atlanse.products.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.transaction.annotation.Transactional
import jakarta.inject.Inject
import jakarta.inject.Singleton
import jakarta.persistence.EntityManager
import jakarta.persistence.PersistenceContext
import jakarta.persistence.TypedQuery
import jakarta.persistence.criteria.CriteriaBuilder
import jakarta.persistence.criteria.CriteriaQuery
import jakarta.persistence.criteria.Predicate
import jakarta.persistence.criteria.Root
import pt.atlanse.products.domains.BrandEntity
import pt.atlanse.products.domains.CategoryEntity

import pt.atlanse.products.domains.DiscountEntity
import pt.atlanse.products.domains.IngredientEntity
import pt.atlanse.products.domains.ProductEntity
import pt.atlanse.products.domains.RatingEntity
import pt.atlanse.products.domains.ReviewEntity
import pt.atlanse.products.domains.SellingStoreEntity
import pt.atlanse.products.domains.compositeIds.ProductIngredientRelationEntity
import pt.atlanse.products.dtos.GetProduct
import pt.atlanse.products.dtos.Product
import pt.atlanse.products.dtos.ProductParams
import pt.atlanse.products.dtos.ProductSearchDTO
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.BrandRepository
import pt.atlanse.products.repositories.CategoryRepository
import pt.atlanse.products.repositories.ProductIngredientRepository
import pt.atlanse.products.repositories.ProductRepository
import pt.atlanse.products.repositories.SellingStoreRepository
import pt.atlanse.products.utils.ExceptionService



@Slf4j
@Singleton
class ProductService {

    ProductRepository productRepository
    BrandRepository brandRepository
    CategoryRepository categoryRepository

    ProductIngredientRepository productIngredients

    @Inject
    RatingService ratingService

    @Inject
    SellingStoreRepository sellingStoreRepository

    @PersistenceContext
    EntityManager entityManager

    ProductService(BrandRepository brandRepository,
                   CategoryRepository categoryRepository, ProductRepository productRepository, ProductIngredientRepository productIngredients, EntityManager entityManager) {
        log.debug "Injecting Product repository"
        this.productRepository = productRepository
        this.productIngredients = productIngredients
        this.entityManager = entityManager
        this.brandRepository = brandRepository
        this.categoryRepository = categoryRepository
    }

    ProductEntity getById(UUID id) throws CustomException {
        productRepository.findById(id).orElseThrow(ExceptionService::ProductNotFound)
    }

    /**
     * Retrieves a page of all products in the system.
     *
     * @param pageable The pagination information to apply to the results
     * @return A page of all products in the system
     */
    @Transactional
    Page<GetProduct> findAll(ProductParams params, Pageable pageable) {
        log.info("Get all products")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<ProductEntity> query = cb.createQuery(ProductEntity.class)
        Root<ProductEntity> root = query.from(ProductEntity.class)
        CriteriaQuery<ProductEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        if (params.brand) {
            brandRepository.findById(UUID.fromString(params.brand)).ifPresent {
                predicates << cb.equal(root.get("brand"), it)
            }
        }

        if (params.department) {
            categoryRepository.findById(UUID.fromString(params.department)).ifPresent {
                predicates << cb.equal(root.get("category"), it)
            }
        }

        predicates << (params.name ? cb.like(cb.lower(root.get("name")), "%" + params.name.toLowerCase() + "%") : null)
        predicates << (params.state ? cb.equal(cb.lower(root.get("state")), params.state.toLowerCase()) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<ProductEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<ProductEntity> productEntities = typedQuery.getResultList()

        List<GetProduct> parsedProducts = new ArrayList<>()

        productEntities.each {
            GetProduct parsed = new GetProduct(
                id: it.id,
                name: it.name,
                description: it.description,
                brand: [id: it.brand.id, name: it.brand.name],
                category: [id: it.category.id, name: it.category.name] ,
                price: it.price,
                state: it.state,
                rating: it.rating ? it.rating.amount : 0,
                numReviews: it.reviews ? it.reviews.size() : 0,
                image: it.image
            )

            parsedProducts.add(parsed)
        }

        return Page.of(parsedProducts, pageable, productRepository.count())
    }

    /**
     * Retrieves the product with the specified ID.
     *
     * @param id The ID of the product to retrieve
     * @return The product with the specified ID
     * @throws {@link pt.atlanse.products.models.CustomException} if the product with the specified ID does not exist
     */

    GetProduct findById(UUID id) {
        log.info("Get product with id ${ id }")
        ProductEntity products = getById(id)
        parse(products)
    }

    /**
     * @deprecated
     * */
    void addReview(UUID id, ReviewEntity review) {

        ProductEntity product = getById(id)

        if (!product.reviews) {
            product.reviews = new ArrayList<>()
        }

        product.reviews.add(review)
        productRepository.update(product)
    }

    /**
     * Adds the specified IngredientEntity to the ProductEntity with the specified id.
     *
     * @param id The id of the ProductEntity to add the IngredientEntity to
     * @param ingredient The IngredientEntity to add to the ProductEntity
     * @throws CustomException if the ProductEntity or IngredientEntity does not exist
     */
    void addIngredient(UUID id, IngredientEntity ingredient) {
        log.info("Adding ingredient with id: ${ ingredient.id } on product with id: ${ id } ")
        getById(id).with {
            it.ingredients.add(ingredient)
            productRepository.update(it)
        }
    }

    /**
     * Removes the specified IngredientEntity from the ProductEntity with the specified id.
     *
     * @param id The id of the ProductEntity to remove the IngredientEntity from
     * @param ingredient The IngredientEntity to remove from the ProductEntity
     * @throws CustomException if the ProductEntity or IngredientEntity does not exist
     */
    void deleteIngredient(UUID id, IngredientEntity ingredient) {
        log.info("Deleting ingredient with id: ${ ingredient.id } on product with id: ${ id } ")
        getById(id).with {
            List<ProductIngredientRelationEntity> l = productIngredients.findByProductAndIngredient(it, ingredient, Pageable.UNPAGED).content
            productIngredients.deleteAll(l)
            productRepository.update(it)
        }
    }

    /**
     * Creates a new ProductEntity with the specified payload, image, brand, and selling stores.
     * Optionally includes a category and discount for the new product.
     *
     * @param payload The payload data for the new ProductEntity
     * @param image The ContentEntity representing the image for the new product
     * @param brand The BrandEntity for the new product
     * @param stores The SellingStoreEntity objects representing the stores that will sell the new product
     * @param category (optional) The CategoryEntity for the new product
     * @param discount (optional) The DiscountEntity for the new product
     * @param author The username of the author creating the new product (default: "anonymous")
     * @throws CustomException if the image, brand, or selling stores do not exist
     */
    ProductEntity create(Product payload, String image, BrandEntity brand, List<SellingStoreEntity> stores, CategoryEntity category = null, DiscountEntity discount = null,
                         List<IngredientEntity> ingredientList, String author = "anonymous") {

        log.info "Creating new product"

        RatingEntity rating = ratingService.create(0)

        try {
            ProductEntity product = new ProductEntity(
                name: payload.name,
                description: payload.description,
                state: payload.state,
                brand: brand,
                sellingStores: stores,
                price: payload.price,
                rating: rating,
                ingredients: ingredientList,
                productionTime: payload.productionTime,
                createdBy: author,
                updatedBy: author
            )

            product.image = UUID.fromString(image) ?: product.image
            product.category = category ?: product.category
            product.discount = discount ?: product.discount

            productRepository.save(product)
        }
        catch (Exception e) {
            ExceptionService.ProductCreatingException(e, payload)
        }

    }

    /**
     * Partially updates an existing ProductEntity with the specified payload, image, and selling stores.
     *
     * @param id The id of the ProductEntity to update
     * @param payload The payload data to update the ProductEntity with
     * @param image The ContentEntity representing the new image for the product
     * @param stores The SellingStoreEntity objects representing the new stores that will sell the product
     * @param author The username of the author updating the ProductEntity (default: "anonymous")
     * @throws CustomException if the specified ProductEntity or ContentEntity do not exist
     */
    void partialUpdate(UUID id, Product payload = null, String image = null, List<SellingStoreEntity> stores = null, List<IngredientEntity> ingredientsUpdatedList = null,
                       BrandEntity brand, CategoryEntity category, String author = "anonymous") {
        log.info("Updating product with id ${ id }")

        // 1. Fetch product
        ProductEntity product = getById(id)

        // 2. Update product fields if payload is not null
        payload?.with {
            product.name = name ?: product.name
            product.description = description ?: product.description
            product.state = state ?: product.state
            product.price = price ?: product.price
            product.price = price ?: product.price
        }

        // 3. Update other fields if they are not null
        product.image = UUID.fromString(image) ?: product.image
        product.sellingStores = stores ?: product.sellingStores
        product.ingredients = ingredientsUpdatedList ?: product.ingredients
        product.category = category ?: product.category
        product.brand = brand ?: product.brand

        // 4. Update existing product
        productRepository.update(product)
    }

    /**
     * Deletes an entity by its ID.
     *
     * @param id The ID of the entity to delete
     */
    void delete(UUID id) {
        log.info "Deleting product $id"
        ProductEntity product = getById(id)
        productRepository.delete(product)
    }

    private GetProduct parse(ProductEntity product) {

        def productPayload = new GetProduct(
            id: product.id,
            name: product.name,
            description: product.description,
            brand:  [id: product.brand.id, name: product.brand.name],
            category: [id: product.category.id, name: product.category.name],
            price: product.price,
            state: product.state,
            rating: product.rating ? product.rating.amount : 0,
            numReviews: product.reviews ? product.reviews.size() : 0,
            ingredients: product.ingredients.collect { ing ->
                [id: ing.id, name: ing.name]
            },
            sellingStores: sellingStoreRepository.findByProduct(product).stream().map { store ->
                [id: store.id, name: store.name, website: store.website]
            },
            productionTime: product.productionTime,
            image: product.image
        )

        if (product.image) {
            productPayload.image = contentService.get(product.image.id)
        }

        return productPayload
    }

    List<GetProduct> productSearch(ProductSearchDTO payload, List<IngredientEntity> ingredients, Pageable pageable) {

        log.info("Received search payload: ${ payload.toString() }")

        List<ProductEntity> productsInPriceRange = []
        List<ProductEntity> products = new ArrayList<>()
        List<ProductEntity> finalList = []
        List<GetProduct> result = []

        products = ingredients ? searchProductIngredient(ingredients, pageable) : productRepository.findAll(pageable).content

        log.warn "Found ${ products.size() } products ! :)"

        if (payload.minPrice && payload.maxPrice) {
            productsInPriceRange = productRepository.findByPriceBetween(payload.minPrice, payload.maxPrice)
        }

        if (products) {
            if (productsInPriceRange) {
                products.each {
                    productsInPriceRange.each { price ->
                        if (it.id == price.id) {
                            finalList.add(it)
                        }
                    }
                }
            } else {
                products.each {
                    finalList.add(it)
                }
            }
        } else {
            productsInPriceRange.each {
                finalList.add(it)
            }
        }
        finalList.unique()

        finalList.each {
            GetProduct parsed = new GetProduct(
                id: it.id,
                name: it.name,
                description: it.description,
                brand: [id: it.brand.id, name: it.brand.name],
                category: [id: it.category.id, name: it.category.name],
                price: it.price,
                state: it.state,
                rating: it.rating ? it.rating.amount : 0,
                numReviews: it.reviews ? it.reviews.size() : 0,
                productionTime: it.productionTime,
                image: it.image
            )

            result.add(parsed)
        }

        // Get by state -> e.g.,  PUBLISHED, ARCHIVED, ...
        if (payload.state) {
            result.removeIf {

                log.info("Comparing states: ${ it.state.toLowerCase() } to ${ payload.state.toLowerCase() }")
                it.state.toLowerCase() != payload.state.toLowerCase()
            }
        }

        return result
    }

    // Aux function to search for product ingredients
    List<ProductEntity> searchProductIngredient(List<IngredientEntity> ingredients, Pageable pageable) {

        List<ProductIngredientRelationEntity> searchProdIngredient = []
        List<ProductEntity> products = []

        ingredients.each {

            searchProdIngredient = productIngredients.findByIngredient(it, pageable).getContent()

            // For each relation, get his product
            searchProdIngredient.each {
                products.add(it.product)
            }

        }

        return products
    }

}
