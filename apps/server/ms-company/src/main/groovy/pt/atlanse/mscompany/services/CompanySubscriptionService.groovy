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
import pt.atlanse.mscompany.domains.CompanyEntity
import pt.atlanse.mscompany.domains.CompanySubscriptionEntity
import pt.atlanse.mscompany.domains.SubscriptionEntity
import pt.atlanse.mscompany.dtos.CompanySubscriptionDTO
import pt.atlanse.mscompany.dtos.CompanySubscriptionParams
import pt.atlanse.mscompany.dtos.SubscriptionDTO
import pt.atlanse.mscompany.dtos.SubscriptionParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.repositories.CompanySubscriptionRepository
import pt.atlanse.mscompany.utils.ExceptionService

@Slf4j
@Singleton
class CompanySubscriptionService {


    @Inject
    CompanySubscriptionRepository repository

    @Inject
    SubscriptionService subscriptionService

    @Inject
    CompanyService companyService

    @PersistenceContext
    EntityManager entityManager

    CompanySubscriptionEntity getById(UUID companyUUID, UUID id) {
        log.info("Get CompanySubscription id: ${id}")
        CompanyEntity entity = companyService.getEntityById(companyUUID)
        repository.findByIdAndCompany(id, entity).orElseThrow(ExceptionService::CompanySubscriptionNotFoundException)
    }

    CompanySubscriptionDTO findById(UUID companyUUID, UUID id){
        log.info("Get CompanySubscription id: ${id}")
        parse(getById(companyUUID, id))
    }

    @Transactional
    Page<CompanySubscriptionDTO> findAll(UUID companyUUID, CompanySubscriptionParams params, Pageable pageable){
        log.info("Get all company subscriptions")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<CompanySubscriptionEntity> query = cb.createQuery(CompanySubscriptionEntity.class)
        Root<CompanySubscriptionEntity> root = query.from(CompanySubscriptionEntity.class)
        CriteriaQuery<CompanySubscriptionEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (companyUUID ? cb.equal(root.get("company"), companyService.getEntityById(companyUUID)) : null)
        predicates << (params.status ? cb.equal(root.get("status"), params.status) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<CompanySubscriptionEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<CompanySubscriptionEntity> entities = typedQuery.getResultList()
        List<CompanySubscriptionDTO> dtos = new ArrayList<>()
        entities.forEach { it -> dtos.add(parse(it)) }

        return Page.of(
            dtos,
            pageable,
            repository.count()
        )
    }

    CompanySubscriptionEntity create(UUID companyUUID, CompanySubscriptionDTO payload){
        log.info("Create CompanySubscription for company id: ${companyUUID}")

        try{
            CompanySubscriptionEntity entity = new CompanySubscriptionEntity(
                company: companyService.getEntityById(companyUUID) ,
                subscription: subscriptionService.getEntityById(payload.subscription),
                expireDate: payload.expireDate,
                status: payload.status,
                price: payload.price,
                startDate: payload.startDate,
                buyDate: payload.buyDate
            )

            return repository.save(entity)

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error creating company subscription",
                "Error happened while trying to create company subcription id: ${payload}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    CompanySubscriptionEntity update(UUID companyUUID, UUID id, CompanySubscriptionDTO payload ){
        log.info("Update CompanySubscription id: ${id}")
        CompanySubscriptionEntity entity = getById(companyUUID, id)
        CompanyEntity companyEntity = companyService.getEntityById(companyUUID)
        SubscriptionEntity subscriptionEntity = subscriptionService.getEntityById(payload.subscription)
        try{

            entity?.with {
                it.price = payload.price
                it.startDate = payload.startDate
                it.status = payload.status
                it.expireDate = payload.expireDate
                it.company = companyEntity
                it.subscription = subscriptionEntity
                it.buyDate = payload.buyDate

                return repository.update(it)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating company subscription",
                "Error happened while trying to update company subscription id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }

    CompanySubscriptionEntity patch(UUID companyUUID, UUID id, CompanySubscriptionDTO payload ){
        log.info("Patch CompanySubscription id: ${id}")
        CompanySubscriptionEntity entity = getById(companyUUID, id)

        try{

            payload?.with {
                entity.price = payload.price ?: entity.price
                entity.subscription = payload.subscription ? subscriptionService.getEntityById(payload.subscription) : entity.subscription
                entity.expireDate = payload.expireDate ?: entity.expireDate
                entity.status = payload.status ?: entity.status
                entity.startDate = payload.startDate ?: entity.startDate
                entity.buyDate = payload.buyDate ?: entity.buyDate
                return repository.update(entity)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating company subscription",
                "Error happened while trying to update company subscription id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    void delete(UUID companyUUID, UUID id){
        log.info("Delete CompanySubscription id: ${id}")
        try{
           CompanySubscriptionEntity entity = getById(companyUUID, id)
            repository.delete(entity)
        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error deleting company subscription",
                "Error happened while trying to delete company subscritpion id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    static CompanySubscriptionDTO parse (CompanySubscriptionEntity entity){

        return new CompanySubscriptionDTO(
            id: entity.id,
            price: entity.price,
            company: entity.company.id,
            subscription: entity.subscription.id,
            status: entity.status,
            startDate: entity.startDate,
            expireDate: entity.expireDate,
            buyDate: entity.buyDate
        )
    }
}

