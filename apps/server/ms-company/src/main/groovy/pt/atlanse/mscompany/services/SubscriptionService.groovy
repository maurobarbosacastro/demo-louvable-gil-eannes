package pt.atlanse.mscompany.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Singleton
import jakarta.persistence.EntityManager
import jakarta.persistence.PersistenceContext
import jakarta.persistence.TypedQuery
import jakarta.persistence.criteria.CriteriaBuilder
import jakarta.persistence.criteria.CriteriaQuery
import jakarta.persistence.criteria.Predicate
import jakarta.persistence.criteria.Root
import jakarta.transaction.Transactional
import pt.atlanse.mscompany.domains.SocialsEntity
import pt.atlanse.mscompany.domains.SubscriptionEntity
import pt.atlanse.mscompany.dtos.SocialDTO
import pt.atlanse.mscompany.dtos.SocialParams
import pt.atlanse.mscompany.dtos.SubscriptionDTO
import pt.atlanse.mscompany.dtos.SubscriptionParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.repositories.CompanyRepository
import pt.atlanse.mscompany.repositories.SocialsRepository
import pt.atlanse.mscompany.repositories.SubscriptionRepository
import pt.atlanse.mscompany.utils.ExceptionService


@Slf4j
@Singleton
class SubscriptionService {


    @Inject
    SubscriptionRepository repository

    @PersistenceContext
    EntityManager entityManager


    SubscriptionEntity getEntityById(UUID id){
        repository.findById(id).orElseThrow(ExceptionService::SubscriptionNotFoundException)
    }

    SubscriptionDTO findById(UUID id){
        parse(repository.findById(id).orElseThrow(ExceptionService::SubscriptionNotFoundException))
    }

    @Transactional
    Page<SubscriptionDTO> findAll(SubscriptionParams params, Pageable pageable){
        log.info("Get all subscriptions")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<SubscriptionEntity> query = cb.createQuery(SubscriptionEntity.class)
        Root<SubscriptionEntity> root = query.from(SubscriptionEntity.class)
        CriteriaQuery<SubscriptionEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.status ? cb.equal(root.get("status"), params.status) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<SubscriptionEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<SubscriptionEntity> entities = typedQuery.getResultList()
        List<SubscriptionDTO> dtos = new ArrayList<>()
        entities.forEach { it -> dtos.add(parse(it)) }

        return Page.of(
            dtos,
            pageable,
            repository.count()
        )
    }

    SubscriptionEntity create(SubscriptionDTO payload){
        try{


            SubscriptionEntity entity = new SubscriptionEntity(
                status: payload.status,
                name: payload.name,
                description: payload.description,
                price: payload.price
            )

            return repository.save(entity)

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error creating social",
                "Error happened while trying to create subcription id: ${payload.name}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    SubscriptionEntity update(UUID id, SubscriptionDTO payload ){

        SubscriptionEntity entity = repository.findById(id).orElseThrow(ExceptionService::SubscriptionNotFoundException)

        try{

            entity?.with {
                it.name = payload.name
                it.price = payload.price
                it.description = payload.description
                it.status = payload.status

                return repository.update(it)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating subscription",
                "Error happened while trying to update subscription id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }

    SubscriptionEntity patch( UUID id, SubscriptionDTO payload ){

        SubscriptionEntity entity = repository.findById(id).orElseThrow(ExceptionService::SubscriptionNotFoundException)

        try{

            payload?.with {
                entity.name = payload.name ?: entity.name
                entity.price = payload.price ?: entity.price
                entity.description = payload.description ?:entity.description
                entity.status = payload.status ?: entity.status
                return repository.update(entity)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating subscription",
                "Error happened while trying to update subscription id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    void delete(UUID id){
        try{
            repository.deleteById(id)
        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error deleting social",
                "Error happened while trying to delete social id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    static SubscriptionDTO parse (SubscriptionEntity entity){

        return new SubscriptionDTO(
            id: entity.id,
            status: entity.status,
            price: entity.price,
            name: entity.name,
            description: entity.description
        )
    }
}

