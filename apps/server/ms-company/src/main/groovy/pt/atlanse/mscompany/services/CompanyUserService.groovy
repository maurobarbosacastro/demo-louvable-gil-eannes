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
import pt.atlanse.mscompany.domains.CompanyUserEntity
import pt.atlanse.mscompany.dtos.CompanyUserDTO
import pt.atlanse.mscompany.dtos.CompanyUserParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.repositories.CompanyUserRepository
import pt.atlanse.mscompany.utils.ExceptionService

@Slf4j
@Singleton
class CompanyUserService {


    @Inject
    CompanyUserRepository companyUserRepository

    @Inject
    CompanyService companyService

    @PersistenceContext
    EntityManager entityManager

    CompanyUserEntity getById(UUID companyUUID, UUID id) {
        log.info("Get CompanyUser id: ${id}")
        CompanyEntity companyEntity = companyService.getEntityById(companyUUID)
        companyUserRepository.findByIdAndCompany(id, companyEntity).orElseThrow(ExceptionService::CompanyUserNotFoundException)
    }

    CompanyUserDTO findById(UUID companyUUID, UUID id) {
        log.info("Get CompanyUser id: ${id}")
        parse(getById(companyUUID, id))
    }

    @Transactional
    Page<CompanyUserDTO> findAll(UUID companyUUID, CompanyUserParams params, Pageable pageable) {
        log.info("Get all company users for company id: ${companyUUID}")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<CompanyUserEntity> query = cb.createQuery(CompanyUserEntity.class)
        Root<CompanyUserEntity> root = query.from(CompanyUserEntity.class)
        CriteriaQuery<CompanyUserEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.keycloakUserId ? cb.equal(root.get("keycloakUserId"), params.keycloakUserId) : null)
        predicates << (companyUUID ? cb.equal(root.get("company"), companyService.getEntityById(companyUUID)) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<CompanyUserEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<CompanyUserEntity> companyEntities = typedQuery.getResultList()
        List<CompanyUserDTO> socialDTOArrayList = new ArrayList<>()
        companyEntities.forEach { it -> socialDTOArrayList.add(parse(it)) }

        return Page.of(
            socialDTOArrayList,
            pageable,
            companyUserRepository.count()
        )
    }

    CompanyUserEntity create(UUID companyUUID, CompanyUserDTO payload) {
        log.info("Create CompanyUser for company id: ${companyUUID}")
        try {
            CompanyUserEntity entity = new CompanyUserEntity(
                keycloakUserId: payload.keycloakUserId,
                company: companyService.getEntityById(companyUUID)
            )

            log.print(entity)

            return companyUserRepository.save(entity)

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error creating company user",
                "Error happened while trying to create company user; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    CompanyUserEntity update(UUID companyUUID, UUID id, CompanyUserDTO payload) {
        log.info("Update CompanyUser id: ${id}")
        CompanyUserEntity companyUser = getById(companyUUID, id)

        try {
            companyUser?.with {
                it.keycloakUserId = payload.keycloakUserId
                it.company = companyService.getEntityById(companyUUID)
                return companyUserRepository.update(it)
            }

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating company user",
                "Error happened while trying to update company user id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }

    CompanyUserEntity patch(UUID companyUUID, UUID id, CompanyUserDTO payload) {
        log.info("Patch CompanyUser id: ${id}")
        CompanyUserEntity companyUser = getById(companyUUID, id)

        try {
            payload?.with {
                companyUser.keycloakUserId = payload.keycloakUserId ?: companyUser.keycloakUserId
                companyUser.company = payload.company ? companyService.getEntityById(companyUUID) : companyUser.company
                return companyUserRepository.update(companyUser)
            }

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating company user",
                "Error happened while trying to update company user id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    void delete(UUID companyUUID, UUID id) {
        log.info("Delete CompanyUser id: ${id}")
        try {
            CompanyUserEntity entity = getById(companyUUID, id)
            companyUserRepository.delete(entity)
        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error deleting company user",
                "Error happened while trying to delete company user id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    static CompanyUserDTO parse(CompanyUserEntity companyUser) {

        return new CompanyUserDTO(
            company: companyUser.company.id,
            keycloakUserId: companyUser.keycloakUserId,
            id: companyUser.id
        )
    }
}

