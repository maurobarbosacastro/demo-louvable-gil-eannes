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
import pt.atlanse.mscompany.domains.CompanyHistoryEntity
import pt.atlanse.mscompany.domains.CompanyHistoryType
import pt.atlanse.mscompany.domains.SocialsEntity
import pt.atlanse.mscompany.dtos.CompanyDTO
import pt.atlanse.mscompany.dtos.CompanyHistoryDTO
import pt.atlanse.mscompany.dtos.CompanyHistoryParams
import pt.atlanse.mscompany.dtos.SocialDTO
import pt.atlanse.mscompany.dtos.SocialParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.repositories.CompanyHistoryRepository
import pt.atlanse.mscompany.repositories.CompanyRepository
import pt.atlanse.mscompany.repositories.SocialsRepository
import pt.atlanse.mscompany.utils.ExceptionService

import java.time.LocalDateTime


@Slf4j
@Singleton
class CompanyHistoryService {


    @Inject
    CompanyHistoryRepository companyHistoryRepository

    @Inject
    CompanyService companyService

    @PersistenceContext
    EntityManager entityManager

    CompanyHistoryEntity getById(UUID companyUUID, UUID id){
        log.info("Get CompanyHistory id: ${id}")
        CompanyEntity companyEntity = companyService.getEntityById(companyUUID)
        companyHistoryRepository.findByIdAndCompany(id, companyEntity ).orElseThrow(ExceptionService::CompanyHistoryNotFoundException)
    }

    CompanyHistoryDTO findById(UUID companyUUID, UUID id){
        log.info("Get CompanyHistory id: ${id}")
        parse(getById(companyUUID, id))
    }

    @Transactional
    Page<CompanyHistoryDTO> findAll(UUID companyUUID, CompanyHistoryParams params, Pageable pageable){
        log.info("Get all company history for company id: ${companyUUID}")
        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<CompanyHistoryEntity> query = cb.createQuery(CompanyHistoryEntity.class)
        Root<CompanyHistoryEntity> root = query.from(CompanyHistoryEntity.class)
        CriteriaQuery<CompanyHistoryEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.type ? cb.equal(root.get("changeType"), params.type) : null)
        predicates << (companyUUID ? cb.equal(root.get("company"), companyService.getEntityById(companyUUID)) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<CompanyHistoryEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<CompanyHistoryEntity> entities = typedQuery.getResultList()
        List<CompanyHistoryDTO> dtos = new ArrayList<>()
        entities.forEach { it -> dtos.add(parse(it)) }

        return Page.of(
            dtos,
            pageable,
            companyHistoryRepository.count()
        )
    }

    CompanyHistoryEntity create(UUID companyUUID, CompanyHistoryDTO payload){
        log.info("Create CompanyHistory for company id: ${companyUUID}")
        try{
            CompanyHistoryEntity companyHistoryEntity = new CompanyHistoryEntity(
                changeType: payload.changeType,
                company: companyService.getEntityById(companyUUID),
                description: payload.description
            )

            return companyHistoryRepository.save(companyHistoryEntity)

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error creating company history",
                "Error happened while trying to create company history; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    CompanyHistoryEntity update(UUID companyUUID, UUID id, CompanyHistoryDTO payload ){

        log.info("Update CompanyHistory id: ${id}")
        CompanyHistoryEntity companyHistoryEntity = getById(companyUUID, id)

        try{

            companyHistoryEntity?.with {
                it.changeType = payload.changeType
                it.description = payload.description
                it.company = companyService.getEntityById(companyUUID)
                return companyHistoryRepository.update(it)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating company history",
                "Error happened while trying to update company history id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }

    CompanyHistoryEntity patch(UUID companyUUID, UUID id, CompanyHistoryDTO payload ){

        log.info("Patch CompanyHistory id: ${id}")
        CompanyHistoryEntity companyHistoryEntity = getById(companyUUID, id)

        try{

            payload?.with {
                companyHistoryEntity.changeType = payload.changeType ?: companyHistoryEntity.changeType
                companyHistoryEntity.description = payload.description ?: companyHistoryEntity.description
                return companyHistoryRepository.update(companyHistoryEntity)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating company history",
                "Error happened while trying to update company history id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    void delete(UUID companyUUID, UUID id){
        try{
            log.info("Delete CompanyHistory id: ${id}")
            companyHistoryRepository.delete(getById(id, companyUUID))
        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error deleting company history",
                "Error happened while trying to delete company history id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    static CompanyHistoryDTO parse (CompanyHistoryEntity companyHistory){

        return new CompanyHistoryDTO(
            company: companyHistory.company.id,
            changeType: companyHistory.changeType,
            description: companyHistory.description,
            changedBy: companyHistory.updatedBy,
            changed_date: companyHistory.updatedAt,
            id: companyHistory.id
        )
    }
}

