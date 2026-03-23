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
import pt.atlanse.mscompany.domains.ScheduleEntity
import pt.atlanse.mscompany.domains.SocialsEntity
import pt.atlanse.mscompany.dtos.SocialDTO
import pt.atlanse.mscompany.dtos.SocialParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.repositories.CompanyRepository
import pt.atlanse.mscompany.repositories.SocialsRepository
import pt.atlanse.mscompany.utils.ExceptionService


@Slf4j
@Singleton
class SocialsService {

    @Inject
    SocialsRepository socialsRepository

    @Inject
    CompanyService companyService

    @Inject
    ImagesClientService imagesClient

    @PersistenceContext
    EntityManager entityManager

    SocialsEntity getById(UUID companyUUID, UUID id) {
        log.info("Get Social id: ${id}")
        CompanyEntity companyEntity = companyService.getEntityById(companyUUID)
        socialsRepository.findByIdAndCompany(id, companyEntity).orElseThrow(ExceptionService::SocialNotFoundException)
    }

    SocialDTO findById(UUID companyUUID, UUID id){
        log.info("Get Social id: ${id}")
        parse(getById(companyUUID, id))
    }

    @Transactional
    Page<SocialDTO> findAll(UUID companyUUID, SocialParams params, Pageable pageable){
        log.info("Get all Socials for company id: ${companyUUID}")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<SocialsEntity> query = cb.createQuery(SocialsEntity.class)
        Root<SocialsEntity> root = query.from(SocialsEntity.class)
        CriteriaQuery<SocialsEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.type ? cb.equal(root.get("type"), params.type) : null)
        predicates << (companyUUID ? cb.equal(root.get("company"), companyService.getEntityById(companyUUID)) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<SocialsEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<SocialsEntity> companyEntities = typedQuery.getResultList()
        List<SocialDTO> socialDTOArrayList = new ArrayList<>()
        companyEntities.forEach { it -> socialDTOArrayList.add(parse(it)) }

        return Page.of(
            socialDTOArrayList,
            pageable,
            socialsRepository.count()
        )
    }

    SocialsEntity create(UUID companyUUID, SocialDTO payload){
        log.info("Create Social for company id: ${companyUUID}")
        try{
            String imageId = imagesClient.create(payload.image)

            SocialsEntity socialsEntity = new SocialsEntity(
                type: payload.type,
                company: companyService.getEntityById(companyUUID),
                image:  imageId,
                link:  payload.link
            )

            return socialsRepository.save(socialsEntity)

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating social",
                "Error happened while trying to create social id: ${payload.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    SocialsEntity update(UUID companyUUID, UUID id, SocialDTO payload ){
        log.info("Update Social id: ${id}")
        SocialsEntity socialsEntity = socialsRepository.findById(id).orElseThrow(ExceptionService::SocialNotFoundException)

        try{
            String imageId = imagesClient.create(payload.image)

            socialsEntity?.with {
                it.image = imageId
                it.type = payload.type
                it.link = payload.link
                it.company = companyService.getEntityById(companyUUID)

                return socialsRepository.update(it)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating social",
                "Error happened while trying to update social id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }

    SocialsEntity patch(UUID companyUUID, UUID id, SocialDTO payload ){
        log.info("Patch Social id: ${id}")
        SocialsEntity socialsEntity = getById(companyUUID, id)

        try{

            payload?.with {
                socialsEntity.type = payload.type ?: socialsEntity.type
                socialsEntity.link = payload.link ?: socialsEntity.link
                socialsEntity.image = payload.image ? imagesClient.create(payload.image) : socialsEntity.image
                 return socialsRepository.update(socialsEntity)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating social",
                "Error happened while trying to update social id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    void delete(UUID companyUUID, UUID id){
        log.info("Delete Social id: ${id}")
        try{
            SocialsEntity entity = getById(companyUUID, id)
            socialsRepository.delete(entity)
        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error deleting social",
                "Error happened while trying to delete social id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    static SocialDTO parse (SocialsEntity social){

        return new SocialDTO(
            company: social.company.id,
            link: social.link,
            type: social.type,
            imageId: social.image,
            id: social.id
        )
    }
}

