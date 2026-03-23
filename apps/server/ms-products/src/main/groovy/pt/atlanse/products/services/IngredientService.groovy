package pt.atlanse.products.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.products.domains.IngredientEntity
import pt.atlanse.products.domains.ProductEntity
import pt.atlanse.products.domains.compositeIds.ProductIngredientRelationEntity
import pt.atlanse.products.dtos.Ingredient
import pt.atlanse.products.dtos.IngredientParams
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.IngredientRepository
import pt.atlanse.products.repositories.ProductIngredientRepository
import pt.atlanse.products.repositories.ProductRepository
import pt.atlanse.products.utils.ExceptionService

@Slf4j
@Singleton
class IngredientService {

    IngredientRepository ingredients

    ProductIngredientRepository productIngredientRepository

    @Inject
    ProductRepository productRepository

    @Inject
    ProductService productService

    IngredientService(IngredientRepository ingredients, ProductIngredientRepository productIngredientRepository) {
        log.debug "Injecting ingredients repository"
        this.ingredients = ingredients
        this.productIngredientRepository = productIngredientRepository
    }

    /**
     * Returns the ingredient with the specified id, or throws a CustomException if the ingredient does not exist.
     *
     * @param id The id of the ingredient to retrieve
     * @return The IngredientEntity with the specified id
     * @throws CustomException if the ingredient does not exist
     */
    private IngredientEntity getById(UUID id) throws CustomException {
        log.info("Getting ingredient with id: ${ id }")
        Optional<IngredientEntity> ingredient = ingredients.findById(id)
        ingredient.isPresent() ? ingredient.get() : ExceptionService.IngredientNotFound(id as String)
    }

    /**
     * Returns a Page of IngredientEntity objects, according to the provided Pageable parameters.
     *
     * @param pageable The Pageable parameters to use when querying for ingredients
     * @return A Page of IngredientEntity objects
     */
    Page<IngredientEntity> findAll(Pageable pageable) {
        log.info("Getting all ingredients")
        ingredients.findAll(pageable)
    }

    /**
     * Returns a Page of IngredientEntity objects, according to the provided Pageable parameters.
     *
     * @param pageable The Pageable parameters to use when querying for ingredients
     * @return A Page of IngredientEntity objects
     */
    Page<IngredientEntity> findAllByName(IngredientParams params, Pageable pageable) {
        if (params.name) {
            ingredients.findAllByNameContains(params.name, pageable)
        } else {
            ingredients.findAll(pageable)
        }
    }

    /**
     * Returns a Page of IngredientEntity objects that belong to the specified ProductEntity, according to the provided Pageable parameters.
     *
     * @param product The ProductEntity to filter by
     * @param pageable The Pageable parameters to use when querying for ingredients
     * @return A Page of IngredientEntity objects that belong to the specified ProductEntity
     */
    Page<IngredientEntity> findAll(ProductEntity product, Pageable pageable) {
        log.info("Getting all ingredients from Product with productId: ${ product.id }")
        productIngredientRepository.findByProduct(product, pageable).map {
            return it.ingredient
        }
    }

    /**
     * Returns the IngredientEntity with the specified id, or throws a CustomException if the ingredient does not exist.
     *
     * @param id The id of the ingredient to retrieve
     * @return The IngredientEntity with the specified id
     * @throws CustomException if the ingredient does not exist
     */
    IngredientEntity findById(UUID id) {
        getById(id)
    }

    /**
     * Creates a new IngredientEntity with the specified name and author.
     *
     * @param name The name of the new ingredient
     * @param author The author who is creating the new ingredient
     * @return The newly created IngredientEntity
     */
    //todo remove anonymous author
    IngredientEntity create(String name, String author = "anonymous") {
        log.info("Creating new ingredient")

        try {
            // 1. Create ingredient entity
            IngredientEntity ingredient = new IngredientEntity(
                name: name,
                createdBy: author,
                updatedBy: author
            )

            // 2. Save ingredient using repository
            ingredients.save(ingredient)
        } catch (Exception e) {
            ExceptionService.IngredientCreatingException(e, name)
        }
    }

    /**
     * Overwrites the IngredientEntity with the specified id with a new IngredientEntity with the specified name and author.
     *
     * @param id The id of the IngredientEntity to overwrite
     * @param name The new name for the IngredientEntity
     * @param author The author who is overwriting the IngredientEntity
     * @throws CustomException if the IngredientEntity does not exist
     */
    void overwrite(UUID id, String name, String author = "anonymous") {
        // 1. Fetch ingredient
        Optional<IngredientEntity> ingredient = ingredients.findById(id)

        // 2. If the ingredient does not exist, create a new one
        if (ingredient.isEmpty()) {
            create(name)
            return
        }

        // 3. Try to overwrite existing ingredient
        try {
            ingredient.get().with {
                it.name = name
                it.updatedBy = author

                // 4. Update existing ingredient
                ingredients.update(it)
            }
        } catch (Exception e) {
            // 5. Handle exceptions found
            log.error "Unhandled exception found while attempting to update ingredient $id; Reason ${ e.message }"
        }
    }

    /**
     * Partially updates the IngredientEntity with the specified id by setting the new name and author, if provided.
     *
     * @param id The id of the IngredientEntity to partially update
     * @param name The new name for the IngredientEntity, or null if not provided
     * @param author The author who is updating the IngredientEntity, or null if not provided
     * @throws CustomException if the IngredientEntity does not exist
     */
    void partialUpdate(UUID id, String name, String author = "anonymous") {
        log.info("Updating ingredient with id: ${ id }")
        // 1. Fetch ingredient
        IngredientEntity ingredient = getById(id)

        // 3. Update existing ingredient fields
        ingredient.with {
            it.name = name
            it.updatedBy = author

            // 4. Update existing ingredient
            ingredients.update(it)
        }

    }

    /**
     * Deletes the IngredientEntity with the specified id.
     *
     * @param id The id of the IngredientEntity to delete
     * @throws CustomException if the IngredientEntity does not exist
     */
    void delete(UUID id) {
        log.info("Deleting ingredient with id: ${ id }")
        // 1. Fetch ingredient
        getById(id).with {
            List<ProductIngredientRelationEntity> l = productIngredientRepository.findByIngredient(it, Pageable.UNPAGED).content

            l.each {
                productService.productRepository.update(it.product)
            }

            productIngredientRepository.deleteAll(l)
            ingredients.delete(it)
        }

    }

    void deleteFromProduct(UUID id, ProductEntity product) {
        log.info("Deleting ingredient with id: ${ id } from products (not from DB)")
        // 1. Fetch ingredient
        IngredientEntity ingredient = getById(id)
        ingredient.products.remove(product)
        ingredients.update(ingredient)

        // 2. Fetch product
        product.ingredients.remove(ingredient)
        productRepository.update(product)

        // 3. Fetch relation
        List<ProductIngredientRelationEntity> l = productIngredientRepository.findByProductAndIngredient(product, ingredient, Pageable.UNPAGED).content
        productIngredientRepository.deleteAll(l)

    }

    // Function for ingredient management (create, add and delete them)
    List<IngredientEntity> manageIngredients(List<Ingredient> ingredients, ProductEntity product = null) {
        List<IngredientEntity> ingredientList = new ArrayList<>()

        ingredients.each {
            if (it.action == 'create') {
                IngredientEntity ingredient = create(it.content)
                ingredientList.add(ingredient)
            }

            if (it.action == 'add' && (!product || !product.ingredients.stream().anyMatch { iti -> iti.id === UUID.fromString(it.content) })) {
                ingredientList.add(findById(UUID.fromString(it.content)))
            }

            if (it.action == 'delete') {
                deleteFromProduct(UUID.fromString(it.content), product)
            }
        }
        return ingredientList
    }

    void addProductToIngredient(ProductEntity product, IngredientEntity ingredient) {

        ingredient.each {
            it.products.add(product)
            ingredients.update(it)
        }

    }

}
