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
import pt.atlanse.mscompany.domains.CompanyMailInfoEntity
import pt.atlanse.mscompany.dtos.CompanyMailInfoDTO
import pt.atlanse.mscompany.dtos.CompanyMailInfoParams
import pt.atlanse.mscompany.dtos.CoordinatesDTO
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.repositories.CompanyMailInfoRepository
import pt.atlanse.mscompany.utils.ExceptionService

@Slf4j
@Singleton
class CompanyMailInfoService {


    @Inject
    CompanyMailInfoRepository companyMailInfoRepository

    @Inject
    CompanyService companyService

    @PersistenceContext
    EntityManager entityManager

    CompanyMailInfoEntity getById(UUID companyUUID, UUID id) {
        log.info("Get CompanyMailInfo id: ${id}")
        CompanyEntity companyEntity = companyService.getEntityById(companyUUID)
        companyMailInfoRepository.findByIdAndCompany(id, companyEntity).orElseThrow(ExceptionService::CompanyMailInfoNotFoundException)
    }


    CompanyMailInfoDTO findById(UUID companyUUID, UUID id) {
        log.info("Get CompanyMailInfo id: ${id}")
        parse(getById(companyUUID, id))
    }

    @Transactional
    Page<CompanyMailInfoDTO> findAll(UUID companyUUID, CompanyMailInfoParams params, Pageable pageable) {
        log.info("Get all company mail info for company id: ${companyUUID}")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<CompanyMailInfoEntity> query = cb.createQuery(CompanyMailInfoEntity.class)
        Root<CompanyMailInfoEntity> root = query.from(CompanyMailInfoEntity.class)
        CriteriaQuery<CompanyMailInfoEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.countryName ? cb.equal(root.get("countryName"), params.countryName) : null)
        predicates << (params.locality ? cb.equal(root.get("locality"), params.locality) : null)
        predicates << (companyUUID ? cb.equal(root.get("company"), companyService.getEntityById(companyUUID)) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<CompanyMailInfoEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<CompanyMailInfoEntity> entities = typedQuery.getResultList()
        List<CompanyMailInfoDTO> dtoList = new ArrayList<>()
        entities.forEach { it -> dtoList.add(parse(it)) }

        return Page.of(
            dtoList,
            pageable,
            companyMailInfoRepository.count()
        )
    }

    CompanyMailInfoEntity create(UUID companyUUID, CompanyMailInfoDTO payload) {
        log.info("Create CompanyMailInfo for company id: ${companyUUID}")
        try {

            CompanyMailInfoEntity companyHistoryEntity = new CompanyMailInfoEntity(
                address1: payload.address1,
                address2: payload.address2,
                countryName: payload.countryName,
                longitude: payload.coordinates.longitude,
                latitude: payload.coordinates.latitude,
                locality: payload.locality,
                postalCode: payload.postalCode,
                company: companyService.getEntityById(companyUUID),
            )

            return companyMailInfoRepository.save(companyHistoryEntity)

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error creating company mail info",
                "Error happened while trying to create company mail info; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    CompanyMailInfoEntity update(UUID companyUUID, UUID id, CompanyMailInfoDTO payload) {

        log.info("Update CompanyMailInfo id: ${id}")
        CompanyMailInfoEntity entity = getById(companyUUID, id)

        try {

            entity?.with {
                it.address1 = payload.address1
                it.address2 = payload.address2
                it.countryName = payload.countryName
                it.longitude = payload.coordinates.longitude
                it.latitude = payload.coordinates.latitude
                it.locality = payload.locality
                it.postalCode = payload.postalCode
                return companyMailInfoRepository.update(it)
            }

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating company mail info",
                "Error happened while trying to update company mail info id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }

    CompanyMailInfoEntity patch(UUID companyUUID, UUID id, CompanyMailInfoDTO payload) {
        log.info("Patch CompanyMailInfo id: ${id}")
        CompanyMailInfoEntity entity = getById(companyUUID, id)

        try {

            payload?.with {
                entity.address1 = payload.address1 ?: entity.address1
                entity.address2 = payload.address2 ?: entity.address2
                entity.countryName = payload.countryName ?: entity.countryName
                entity.longitude = payload.coordinates.longitude ?: entity.longitude
                entity.latitude = payload.coordinates.latitude ?: entity.latitude
                entity.locality = payload.locality ?: entity.locality
                entity.postalCode = payload.postalCode ?: entity.postalCode
                return companyMailInfoRepository.update(entity)
            }

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating company mail info",
                "Error happened while trying to update company mail info id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    void delete(UUID companyUUID, UUID id) {
        try {
            log.info("Delete CompanyMailInfo id: ${id}")
            CompanyMailInfoEntity entity = getById(companyUUID, id)
            companyMailInfoRepository.delete(entity)

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error deleting company mail info",
                "Error happened while trying to delete company mail info id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    static CompanyMailInfoDTO parse(CompanyMailInfoEntity entity) {

        return new CompanyMailInfoDTO(
            address1: entity.address1,
            address2: entity.address2,
            countryName: entity.countryName,
            coordinates: new CoordinatesDTO(latitude: entity.latitude, longitude: entity.longitude),
            locality: entity.locality,
            postalCode: entity.postalCode,
            company: entity.company.id,
            id: entity.id
        )
    }
}

