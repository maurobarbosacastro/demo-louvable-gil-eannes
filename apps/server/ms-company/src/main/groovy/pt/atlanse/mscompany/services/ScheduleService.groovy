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
import pt.atlanse.mscompany.domains.ScheduleEntity
import pt.atlanse.mscompany.dtos.ScheduleDTO
import pt.atlanse.mscompany.dtos.ScheduleParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.repositories.CompanyRepository
import pt.atlanse.mscompany.repositories.ScheduleRepository
import pt.atlanse.mscompany.utils.ExceptionService

@Slf4j
@Singleton
class ScheduleService {


    @Inject
    ScheduleRepository scheduleRepository

    @Inject
    CompanyService companyService

    @PersistenceContext
    EntityManager entityManager

    ScheduleEntity getById(UUID companyUUID, UUID id) {
        log.info("Get Schedule id: ${id}")
        CompanyEntity companyEntity = companyService.getEntityById(companyUUID)
        scheduleRepository.findByIdAndCompany(id, companyEntity).orElseThrow(ExceptionService::ScheduleNotFoundException)
    }

    ScheduleDTO findById(UUID companyUUID, UUID id){
        log.info("Get Schedule id: ${id}")
        parse(getById(companyUUID, id))
    }

    @Transactional
    Page<ScheduleDTO> findAll(UUID companyUUID, ScheduleParams params, Pageable pageable){
        log.info("Get all Schedules for company id: ${companyUUID}")

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<ScheduleEntity> query = cb.createQuery(ScheduleEntity.class)
        Root<ScheduleEntity> root = query.from(ScheduleEntity.class)
        CriteriaQuery<ScheduleEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.type ? cb.equal(cb.lower(root.get("type")), params.type) : null)
        predicates << (companyUUID ? cb.equal(root.get("company"), companyService.getEntityById(companyUUID)) : null)
        predicates << (params.weekDay ? cb.equal(cb.lower(root.get("weekDay")), params.weekDay) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<ScheduleEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<ScheduleEntity> scheduleEntities = typedQuery.getResultList()
        List<ScheduleDTO> scheduleDTOArrayList = new ArrayList<>()
        scheduleEntities.forEach { it -> scheduleDTOArrayList.add(parse(it)) }

        return Page.of(
            scheduleDTOArrayList,
            pageable,
            scheduleRepository.count()
        )
    }

    ScheduleEntity create(UUID companyUUID, ScheduleDTO payload){
        log.info("Create Schedule for company id: ${companyUUID}")
        try{
            ScheduleEntity scheduleEntity = new ScheduleEntity(
                type: payload.type,
                company: companyService.getEntityById(companyUUID),
                endDate: payload.endDate,
                startDate: payload.startDate,
                weekDay: payload.weekDay
            )

            return scheduleRepository.save(scheduleEntity)

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error creating schedule",
                "Error happened while trying to create schedule; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    ScheduleEntity update(UUID companyUUID, UUID id, ScheduleDTO payload ){
        log.info("Update Schedule id: ${id}")
        ScheduleEntity socialsEntity = scheduleRepository.findById(id).orElseThrow(ExceptionService::ScheduleNotFoundException)

        try{
            socialsEntity?.with {
                it.endDate = payload.endDate
                it.type = payload.type
                it.startDate = payload.startDate
                it.weekDay = payload.weekDay
                it.company = companyService.getEntityById(companyUUID)

                return scheduleRepository.update(it)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating schedule",
                "Error happened while trying to update schedule id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }

    }

    ScheduleEntity patch(UUID companyUUID, UUID id, ScheduleDTO payload ){
        log.info("Patch Schedule id: ${id}")
        ScheduleEntity scheduleEntity = getById(companyUUID, id)

        try{

            payload?.with {
                scheduleEntity.type = payload.type ?: scheduleEntity.type
                scheduleEntity.endDate = payload.endDate ?: scheduleEntity.endDate
                scheduleEntity.weekDay = payload.weekDay ?: scheduleEntity.weekDay
                scheduleEntity.startDate = payload.startDate ?: scheduleEntity.startDate
                return scheduleRepository.update(scheduleEntity)
            }

        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error updating schedule",
                "Error happened while trying to update schedule id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    void delete(UUID companyUUID, UUID id){
        log.info("Delete Schedule id: ${id}")
        try{
            ScheduleEntity entity = getById(companyUUID, id)
            scheduleRepository.delete(entity)
        }catch (Exception e){
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error deleting schedule",
                "Error happened while trying to delete schedule id: ${id.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    static ScheduleDTO parse (ScheduleEntity schedule){

        return new ScheduleDTO(
            endDate: schedule.endDate,
            startDate: schedule.startDate,
            type: schedule.type,
            company: schedule.company.id,
            weekDay: schedule.weekDay,
            id: schedule.id
        )
    }
}

