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
import pt.atlanse.products.domains.BrandEntity

import pt.atlanse.products.dtos.Brand
import pt.atlanse.products.dtos.BrandParams
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.BrandRepository
import pt.atlanse.products.repositories.ProductRepository
import pt.atlanse.products.repositories.RatingRepository
import pt.atlanse.products.utils.ExceptionService


@Slf4j
@Singleton
class BrandService {

    @Inject
    BrandRepository brands

    @Inject
    ProductRepository productRepository

    @Inject
    RatingRepository ratingRepository

    @PersistenceContext
    EntityManager entityManager

    BrandService(BrandRepository brands, EntityManager entityManager) {
        log.debug "Injecting Product repository"
        this.brands = brands
        this.entityManager = entityManager
    }

    /**
     * Retrieves the brand entity with the specified ID.
     *
     * @param id The ID of the brand entity to retrieve
     * @return The brand entity with the specified ID
     * @throws CustomException if the brand entity with the specified ID does not exist
     */
    BrandEntity getById(UUID id) throws CustomException {
        brands.findById(id).orElseThrow(ExceptionService::BrandNotFoundException)
    }

    /**
     * <b>FIND ALL</b>
     * Find brands using pagination
     * @param pageable {@link Pageable} value
     * @return page of {@link BrandEntity}
     * */
    @Transactional
    Page<Brand> findAll(BrandParams params, Pageable pageable) {
        log.info("Get all brands")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<BrandEntity> query = cb.createQuery(BrandEntity.class)
        Root<BrandEntity> root = query.from(BrandEntity.class)
        CriteriaQuery<BrandEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.name ? cb.like(cb.lower(root.get("name")), "%" + params.name.toLowerCase() + "%") : null)
        predicates << (params.state ? cb.equal(cb.lower(root.get("state")), params.state.toLowerCase()) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<BrandEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<BrandEntity> brandEntities = typedQuery.getResultList()

        return Page.of(
            brandEntities.stream().map {
                Optional<Long> sumQuery = ratingRepository.sumRatingByBrand(it.id)

                Long sumRating = sumQuery.isPresent() ? sumQuery.get() : 0
                Long numProducts = productRepository.countByBrand(it)
                Float averageRating = sumRating ? sumRating / numProducts : 0

                Brand parsed = new Brand(
                    name: it.name,
                    id: it.id,
                    description: it.description,
                    website: it.website,
                    state: it.state,
                    numProducts: numProducts,
                    averageRating: String.format("%.1f", averageRating),
                )

                if (it.image) {
                    parsed.image = contentService.get(it.image.id)
                }

                if (it.banner) {
                    parsed.banner = contentService.get(it.banner.id)
                }

                return parsed
            }.toList(),
            pageable,
            brands.count()
        )
    }

    /**
     * <b>FIND</b>
     * Find brand by ID
     * @param id {@link UUID} value
     * @return void
     * */
    Brand findById(UUID id) {
        log.info("Get brand with id ${ id }")
        BrandEntity brand = getById(id)
        parse(brand)
    }

    /**
     * <b>CREATE</b>
     * Create a new brand for the products
     * @param payload {@link Brand} object
     * @return void
     * */
    //todo remove anonymous author
    void create(Brand payload, String image, String banner, String author = "anonymous") {
        log.info "Creating new brand"

        try {
            // 1. Create brand entity
            BrandEntity brand = new BrandEntity(
                name: payload.name,
                description: payload.description,
                website: payload.website,
                state: payload.state,
                createdBy: author,
                updatedBy: author
            )

            if (image) {
                brand.image = image
            }

            if (banner) {
                brand.banner = banner
            }

            // 2. Save brand using repository
            brands.save(brand)
        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${ e.message }"
            throw new CustomException(
                "Error creating brand",
                "Error happened while trying to create new Brand ${ payload.toString() }; Reason: ${ e.message }",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    /**
     * <b>OVERWRITE</b>
     * Overwrite brand or create if it does not exist
     * @param id {@link UUID} string
     * @param payload {@link Brand} object
     * @param image {@link String} object
     * @return void
     * */
    void overwrite(UUID id, Brand payload, String image, String banner, String author = "anonymous") {
        // 1. Fetch brand
        Optional<BrandEntity> brand = brands.findById(id)

        // 2. If the brand does not exist, create a new one
        if (brand.isEmpty()) {
            create(payload, image, banner)
            return
        }

        // 3. Try to overwrite existing brand
        try {
            brand.get().with {
                it.name = payload.name
                it.description = payload.description
                it.website = payload.website
                it.state = payload.state
                it.image = image
                it.banner = banner
                it.updatedBy = author

                // 4. Update existing brand
                brands.update(it)
            }
        } catch (Exception e) {
            // 5. Handle exceptions found
            log.error "Unhandled exception found while attempting to update brand $id; Reason ${ e.message }"
        }
    }

    /**
     * <b>PARTIAL UPDATE</b>
     * Update brand
     * @param id {@link UUID} string
     * @param payload {@link Brand} object
     * @param image {@link String} object
     * @return void
     * */
    void partialUpdate(UUID id, Brand payload, String image, String banner, String author = "anonymous") {
        log.info("Updating brand with id ${ id }")

        // 1. Fetch brand
        BrandEntity brand = getById(id)

        // 2. Update brand fields if payload is not null
        payload?.with {
            brand.name = name ?: brand.name
            brand.description = description ?: brand.description
            brand.website = website ?: brand.website
            brand.state = state ?: brand.state

        }

        // 3. Update other fields if they are not null
        brand.image = image ?: brand.image
        brand.banner = banner ?: brand.banner
        brand.updatedBy = author

        // 4. Update existing brand
        brands.update(brand)

    }

    /**
     * <b>DELETE</b>
     * Delete brand
     * @param id {@link UUID} string
     * @return void
     * */
    void delete(UUID id) {
        // 1. Fetch brand
        BrandEntity brand = getById(id)

        // 2. Delete brand
        brands.delete(brand)
    }

    private Brand parse(BrandEntity brand) {

        Optional<Long> sumQuery = ratingRepository.sumRatingByBrand(brand.id)
        Long sumRating = sumQuery.isPresent() ? sumQuery.get() : 0
        Long numProducts = productRepository.countByBrand(findBrandById(brand.id))
        Float averageRating = sumRating ? sumRating / numProducts : 0

        Brand brandParsed = new Brand(
            name: brand.name,
            id: brand.id,
            description: brand.description,
            website: brand.website,
            state: brand.state,
            numProducts: numProducts,
            averageRating: String.format("%.1f", averageRating)
        )

        if (brand.image) {
            brandParsed.image = contentService.get(brand.image.id)
        }

        if (brand.banner) {
            brandParsed.banner = contentService.get(brand.banner.id)
        }

        return brandParsed
    }

    private BrandEntity findBrandById(UUID brandId) {
        brands.findById(brandId).orElseThrow {
            new CustomException("Error fetching question",
                "Error fetching brand with id $brandId; " +
                    "Not Found", HttpStatus.BAD_REQUEST)
        }
    }
}
