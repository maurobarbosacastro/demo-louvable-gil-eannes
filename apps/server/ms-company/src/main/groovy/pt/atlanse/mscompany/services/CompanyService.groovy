package pt.atlanse.mscompany.services

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
import pt.atlanse.mscompany.domains.CompanyEntity
import pt.atlanse.mscompany.dtos.CompanyDTO
import pt.atlanse.mscompany.dtos.CompanyParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.repositories.CompanyRepository
import pt.atlanse.mscompany.utils.ExceptionService

@Slf4j
@Singleton
class CompanyService {

    @Inject
    CompanyRepository companyRepository

    @PersistenceContext
    EntityManager entityManager


    CompanyEntity getEntityById(UUID id) {

        log.info("Get Company id: ${id}")
        companyRepository.findById(id).orElseThrow(ExceptionService::CompanyNotFoundException)
    }

    CompanyDTO findById(UUID id) {
        log.info("Get Company id: ${id}")
        parse(companyRepository.findById(id).orElseThrow(ExceptionService::CompanyNotFoundException))
    }

    @Transactional
    Page<CompanyDTO> findAll(CompanyParams params, Pageable pageable) {
       log.info("Get all companies")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<CompanyEntity> query = cb.createQuery(CompanyEntity.class)
        Root<CompanyEntity> root = query.from(CompanyEntity.class)
        CriteriaQuery<CompanyEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.name ? cb.like(cb.lower(root.get("name")), "%" + params.name.toLowerCase() + "%") : null)
        predicates << (params.status ? cb.equal(root.get("state"), params.status) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<CompanyEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<CompanyEntity> companyEntities = typedQuery.getResultList()
        List<CompanyDTO> companyDTOList = new ArrayList<>()
        companyEntities.forEach { it -> companyDTOList.add(parse(it)) }

        return Page.of(
            companyDTOList,
            pageable,
            companyRepository.count()
        )
    }

    CompanyEntity create(CompanyDTO payload) {
        log.info "Creating new Company"

        try {
            CompanyEntity entity = new CompanyEntity(
                name: payload.name,
                emailAddress: payload.emailAddress,
                phoneNumber: payload.phoneNumber,
                status: payload.status,
                category: payload.category,
                vatNumber: payload.vatNumber,
                description: payload.description,
                legalName: payload.legalName,
                deliveryRate: payload.deliveryRate,
                deliveryQuantity: payload.deliveryQuantity,
            )

            companyRepository.save(entity);

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error creating company",
                "Error happened while trying to create new company ${payload.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }

    CompanyEntity update(UUID id, CompanyDTO payload) {

        log.info("Update Company id: ${id}")

        CompanyEntity companyEntity = companyRepository.findById(id).orElseThrow(ExceptionService::CompanyNotFoundException)

        try {
            companyEntity.with {
                it.name = payload.name
                it.description = payload.description
                it.vatNumber = payload.vatNumber
                it.category = payload.category
                it.status = payload.status
                it.phoneNumber = payload.phoneNumber
                it.emailAddress = payload.emailAddress
                it.deliveryQuantity = payload.deliveryQuantity
                it.deliveryRate = payload.deliveryRate
                it.legalName = payload.legalName

                return companyRepository.update(it)
            }

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error Updating company",
                "Error happened while trying to update company id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }


    CompanyEntity patch(UUID id, CompanyDTO payload) {

       log.info("Patch Company id: ${id}")

        CompanyEntity companyEntity = companyRepository.findById(id).orElseThrow(ExceptionService::CompanyNotFoundException)

        try {
            payload?.with {
                companyEntity.name = name ?: companyEntity.name
                companyEntity.description = description ?: companyEntity.description
                companyEntity.vatNumber = vatNumber ?: companyEntity.vatNumber
                companyEntity.category = category ?: companyEntity.category
                companyEntity.status = status ?: companyEntity.status
                companyEntity.phoneNumber = phoneNumber ?: companyEntity.phoneNumber
                companyEntity.emailAddress = emailAddress ?: companyEntity.emailAddress
                companyEntity.deliveryQuantity = deliveryQuantity ?: companyEntity.deliveryQuantity
                companyEntity.deliveryRate = deliveryRate ?: companyEntity.deliveryRate
                companyEntity.legalName = legalName ?: companyEntity.legalName
                return companyRepository.update(companyEntity)

            }

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error patching company",
                "Error happened while trying to patch company id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }


    void delete(UUID id) {
        try {
            log.info("Delete Company id: ${id}")
            companyRepository.deleteById(id)
        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error deleting company",
                "Error happened while trying to delete company id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    static CompanyDTO parse(CompanyEntity company) {

        CompanyDTO companyDTO = new CompanyDTO(
            id: company.id,
            name: company.name,
            emailAddress: company.emailAddress,
            phoneNumber: company.phoneNumber,
            status: company.status,
            category: company.category,
            vatNumber: company.vatNumber,
            description: company.description,
            legalName: company.legalName,
            deliveryRate: company.deliveryRate,
            deliveryQuantity: company.deliveryQuantity,
        )

        return companyDTO
    }
}
