package pt.atlanse.products.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpStatus
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
import pt.atlanse.products.domains.CategoryEntity

import pt.atlanse.products.dtos.Category
import pt.atlanse.products.dtos.CategoryParams
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.CategoryRepository
import pt.atlanse.products.repositories.ProductRepository
import pt.atlanse.products.utils.ExceptionService



@Slf4j
@Singleton
class CategoryService {

    CategoryRepository categories

    @Inject
    ProductRepository productRepository

    @PersistenceContext
    EntityManager entityManager

    CategoryService(CategoryRepository categories, EntityManager entityManager) {
        log.debug "Injecting Product repository"
        this.categories = categories
        this.entityManager = entityManager
    }

    /**
     * Retrieves the category entity with the specified ID.
     *
     * @param id The ID of the brand entity to retrieve
     * @return The category entity with the specified ID
     * @throws CustomException if the category entity with the specified ID does not exist
     */
    CategoryEntity getById(UUID id) throws CustomException {
        categories.findById(id).orElseThrow(ExceptionService::CategoryNotFoundException)
    }

    /**
     * <b>FIND ALL</b>
     * Find categories using pagination
     * @param pageable {@link Pageable} value
     * @return page of {@link CategoryEntity}
     * */
    @Transactional
    Page<Category> findAll(CategoryParams params, Pageable pageable) {
        log.info("Get all categories")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<CategoryEntity> query = cb.createQuery(CategoryEntity.class)
        Root<CategoryEntity> root = query.from(CategoryEntity.class)
        CriteriaQuery<CategoryEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.name ? cb.like(cb.lower(root.get("name")), "%" + params.name.toLowerCase() + "%") : null)
        predicates << (params.state ? cb.equal(cb.lower(root.get("state")), params.state.toLowerCase()) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<CategoryEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<CategoryEntity> categoryEntities = typedQuery.getResultList()

        return Page.of(
            categoryEntities.stream().map {
                Long numProducts = productRepository.countByCategory(getById(it.id))

                Category category = new Category(
                    id: it.id,
                    name: it.name,
                    description: it.description,
                    state: it.state,
                    numProducts: numProducts,
                )

                if (it.image) {
                    category.image = contentService.get(it.image.id)
                }

                return category
            }.toList(),
            pageable,
            categories.count()
        )
    }

    /**
     * <b>FIND</b>
     * Find brand by ID
     * @param id {@link UUID} value
     * @return void
     * */
    Category findById(UUID id) {
        log.info("Get category with id ${ id }")
        CategoryEntity category = getById(id)
        parse(category)
    }

    /**
     * <b>CREATE</b>
     * Create a new category for the products
     * @param payload {@link Category} object
     * @return void
     * */
    //todo remove anonymous author
    void create(Category payload, String image, String author = "anonymous") {
        log.info "Creating new category"

        try {
            // 1. Create brand entity
            CategoryEntity category = new CategoryEntity(
                name: payload.name,
                description: payload.description,
                state: payload.state,
                createdBy: author,
                updatedBy: author
            )

            if (image) {
                category.image = image
            }

            // 2. Save brand using repository
            categories.save(category)
        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${ e.message }"
            throw new CustomException(
                "Error creating category",
                "Error happened while trying to create new Category ${ payload.toString() }; Reason: ${ e.message }",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    /**
     * <b>OVERWRITE</b>
     * Overwrite category or create if it does not exist
     * @param id {@link UUID} string
     * @param payload {@link Category} object
     * @param image {@link String} object
     * @return void
     * */
    void overwrite(UUID id, Category payload, String image, String author = "anonymous") {
        // 1. Fetch brand
        Optional<CategoryEntity> category = categories.findById(id)

        // 2. If the brand does not exist, create a new one
        if (category.isEmpty()) {
            create(payload, image)
            return
        }

        // 3. Try to overwrite existing brand
        try {
            category.get().with {
                it.name = payload.name
                it.description = payload.description
                it.state = payload.state
                it.image = image
                it.updatedBy = author

                // 4. Update existing brand
                categories.update(it)
            }
        } catch (Exception e) {
            // 5. Handle exceptions found
            log.error "Unhandled exception found while attempting to update brand $id; Reason ${ e.message }"
        }
    }

    /**
     * <b>PARTIAL UPDATE</b>
     * Update category
     * @param id {@link UUID} string
     * @param payload {@link Category} object
     * @param image {@link String} object
     * @return void
     * */
    void partialUpdate(UUID id, Category payload, String image, String author = "anonymous") {
        log.info("Updating category with id ${ id }")

        // 1. Fetch category
        CategoryEntity category = getById(id)

        // 2. Update category fields if payload is not null
        payload?.with {
            category.name = name ?: category.name
            category.description = description ?: category.description
            category.state = state ?: category.state
        }

        // 3. Update other fields if they are not null
        category.image = image ?: category.image
        category.updatedBy = author

        // 4. Update existing category
        categories.update(category)
    }

    /**
     * <b>DELETE</b>
     * Delete category
     * @param id {@link UUID} string
     * @return void
     * */
    void delete(UUID id) {
        // 1. Fetch brand
        CategoryEntity category = getById(id)

        // 2. Delete brand
        categories.delete(category)
    }

    private Category parse(CategoryEntity category) {
        Long numProducts = productRepository.countByCategory(getById(category.id))

        Category cat = new Category(
            id: category.id,
            name: category.name,
            description: category.description,
            state: category.state,
            numProducts: numProducts
        )

        if (category.image) {
            cat.image = contentService.get(category.image.id)
        }

        return cat
    }

}
