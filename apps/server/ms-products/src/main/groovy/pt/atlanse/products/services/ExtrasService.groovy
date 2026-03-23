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
import pt.atlanse.products.domains.ExtrasEntity
import pt.atlanse.products.dtos.Extras
import pt.atlanse.products.dtos.ExtrasParams
import pt.atlanse.products.dtos.GetProduct
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.ExtrasRepository
import pt.atlanse.products.utils.ExceptionService



@Slf4j
@Singleton
class ExtrasService {

    @Inject
    ExtrasRepository extrasRepository

    @Inject
    ImagesClientService imagesClientService

    @PersistenceContext
    EntityManager entityManager

    /**
     * Retrieves the Extras entity with the specified ID.
     *
     * @param id The ID of the extras entity to retrieve
     * @return The extras entity with the specified ID
     * @throws CustomException if the extras entity with the specified ID does not exist
     */
    Extras findById(UUID id, boolean getProduct = false) throws CustomException {
        parseExtra(extrasRepository.findById(id).orElseThrow(ExceptionService::ExtrasNotFoundException), getProduct)
    }

    /**
     * <b>FIND ALL</b>
     * Find extras using pagination
     * @param pageable {@link Pageable} value
     * @return page of {@link ExtrasEntity}
     * */
    @Transactional
    Page<Extras> findAll(ExtrasParams params, Pageable pageable) {
        log.info("Get all Extras")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<ExtrasEntity> query = cb.createQuery(ExtrasEntity.class)
        Root<ExtrasEntity> root = query.from(ExtrasEntity.class)
        CriteriaQuery<ExtrasEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.name ? cb.like(cb.lower(root.get("name")), "%" + params.name.toLowerCase() + "%") : null)
        predicates << (params.state ? cb.equal(cb.lower(root.get("state")), params.state.toLowerCase()) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<ExtrasEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<ExtrasEntity> extrasEntities = typedQuery.getResultList()
        List<Extras> extrasList = new ArrayList<>()
        extrasEntities.each { it -> extrasList.add(parseExtra(it)) }

        return Page.of(
            extrasList,
            pageable,
            extrasRepository.count()
        )
    }


    /**
     * <b>CREATE</b>
     * Create a new Extra for the products
     * @param payload {@link Extras} object
     * @return void
     * */
    //todo remove anonymous author
    ExtrasEntity create(Extras payload) {
        log.info "Creating new Extra"

        try {
            // 1. Create extras entity
            ExtrasEntity extrasEntity = new ExtrasEntity(
                name: payload.name,
                description: payload.description,
                image: payload.image ? UUID.fromString(imagesClientService.create(payload.image)) : null,
                state: payload.state,
                price: payload.price
            )

            return extrasRepository.save(extrasEntity)
        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error creating Extras",
                "Error happened while trying to create new Extras ${payload.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    /**
     * <b>OVERWRITE</b>
     * Overwrite Extras or create if it does not exist
     * @param id {@link UUID} string
     * @param payload {@link Extras} object
     * @return void
     * */
    ExtrasEntity update(UUID id, Extras payload) {
        // 1. Fetch extras
        Optional<ExtrasEntity> extras = extrasRepository.findById(id)

        // 2. If the extras does not exist, create a new one
        if (extras.isEmpty()) {
            create(payload)
            return
        }

        // 3. Try to overwrite existing extras
        try {
            extras.get().with {
                it.name = payload.name
                it.description = payload.description
                it.price = payload.price
                it.state = payload.state
                it.image = UUID.fromString(imagesClientService.create(payload.image))

                // 4. Update existing extras
                return extrasRepository.update(it)
            }
        } catch (Exception e) {
            // 5. Handle exceptions found
            log.error "Unhandled exception found while attempting to update Extra $id; Reason ${e.message}"
        }
    }


    /**
     * <b>PARTIAL UPDATE</b>
     * Update Extras
     * @param id {@link UUID} string
     * @param payload {@link Extras} object
     * @return void
     * */
    ExtrasEntity partialUpdate(UUID id, Extras payload) {
        log.info("Updating Extras with id ${id}")

        try{
            // 1. Fetch Extras
            ExtrasEntity extras = extrasRepository.findById(id).orElseThrow()

            // 2. Update extras fields if payload is not null
            payload?.with {
                extras.name = name ?: extras.name
                extras.description = description ?: extras.description
                extras.price = price ?: extras.price
                extras.state = state ?: extras.state
                extras.image = image ? UUID.fromString( imagesClientService.create(image) ): extras.image
            }

            //3. Save
            return extrasRepository.update(extras)
        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error partial updating Extras",
                "Error happened while trying to partial updating a Extras ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }


    /**
     * <b>DELETE</b>
     * Delete extras
     * @param id {@link UUID} string
     * @return void
     * */
    void delete(UUID id, boolean force) {
        // 1. Fetch extras
        ExtrasEntity extrasEntity = extrasRepository.findById(id).orElseThrow()

        try{
            if (force) {
                extrasRepository.deleteProductRelation(id)
            }
            // 2. Delete extras and relation
            extrasRepository.delete(extrasEntity)
        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error Deleting Extras",
                "Error happened while trying to delete a Extras ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }


    static parseExtra(ExtrasEntity entity, boolean listProducts = false) {
        return new Extras(
            name: entity.name,
            description: entity.description,
            productsCount: entity.products.size(),
            price: entity.price,
            state: entity.state,
            imageId: entity.image,
            products: listProducts ? entity.products.collect { it ->
                new GetProduct(
                    id: it.id,
                    name: it.name,
                    description: it.description,
                    brand: it.brand ? [id: it.brand.id, name: it.brand.name] : null,
                    category: it.category ? [id: it.category.id, name: it.category.name] : null,
                    price: it.price,
                    state: it.state,
                    rating: it.rating ? it.rating.amount : 0,
                    numReviews: it.reviews ? it.reviews.size() : 0,
                    image: it.image
                )
            } : null
        )
    }
}
