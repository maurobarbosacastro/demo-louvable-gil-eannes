package pt.atlanse.eduquiz.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.eduquiz.DTO.CourseDTO
import pt.atlanse.eduquiz.DTO.CourseParams
import pt.atlanse.eduquiz.domain.CourseEntity
import pt.atlanse.eduquiz.domain.CourseOrderEntity
import pt.atlanse.eduquiz.domain.ModulesEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.repositories.CourseOrderRepository
import pt.atlanse.eduquiz.repositories.CourseRepository
import pt.atlanse.eduquiz.repositories.ModulesRepository
import pt.atlanse.eduquiz.utils.ExceptionService

import jakarta.persistence.EntityManager
import jakarta.persistence.PersistenceContext
import jakarta.persistence.TypedQuery
import jakarta.persistence.criteria.CriteriaBuilder
import jakarta.persistence.criteria.CriteriaQuery
import jakarta.persistence.criteria.Predicate
import jakarta.persistence.criteria.Root
import jakarta.transaction.Transactional

@Slf4j
@Singleton
class CourseService {

    @Inject
    FileHandler files

    @Inject
    ImagesClientService imagesService

    @Inject
    ModulesService modulesService

    @PersistenceContext
    EntityManager entityManager

    CourseRepository courseRepository
    ModulesRepository modulesRepository
    CourseOrderRepository courseOrderRepository

    CourseService(CourseRepository courseRepository,
                  ModulesRepository modulesRepository,
                  CourseOrderRepository courseOrderRepository
    ) {
        this.courseRepository = courseRepository
        this.modulesRepository = modulesRepository
        this.courseOrderRepository = courseOrderRepository
    }

    CourseEntity findById(String courseId) {
        return courseRepository.findById(UUID.fromString(courseId)).orElseThrow {
            new CustomException(
                "Course not found",
                "The course with id $courseId was not found",
                HttpStatus.NOT_FOUND
            )
        }
    }

    CourseDTO getCourse(String courseId) {
        CourseEntity course = findById(courseId)
        return new CourseDTO(
            id: course.id,
            imageId: course.image,
            title: course.title,
            description: course.description,
            beginDate: course.beginDate,
            endDate: course.endDate,
            extras: course.extras,
            status: course.status,
            courseOrder: course.courseOrder
        )
    }

    ModulesEntity getModules(String moduleId) {
        modulesRepository.findById(UUID.fromString(moduleId)).orElseThrow {
            new CustomException(
                "Module not found",
                "The module with id $moduleId was not found",
                HttpStatus.NOT_FOUND
            )
        }
    }

    @Transactional
    Page<CourseEntity> applyFilters(CourseParams params, Pageable pageable) {
        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<CourseEntity> query = cb.createQuery(CourseEntity.class)
        Root<CourseEntity> root = query.from(CourseEntity.class)
        CriteriaQuery<CourseEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.title ? cb.like(cb.lower(root.get("title")), '%' + params.title.toLowerCase() + '%') : null)
        predicates << (params.status ? cb.like(root.get("status"), params.status.toLowerCase()) : null)
        predicates << (params.beginDate ? cb.greaterThanOrEqualTo(root.get("beginDate"), params.beginDate) : null)
        predicates << (params.endDate ? cb.lessThanOrEqualTo(root.get("endDate"), params.endDate) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<CourseEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<CourseEntity> courseEntities = typedQuery.getResultList()

        return Page.of(courseEntities, pageable, courseRepository.count())
    }

    Page<CourseEntity> getCourses(CourseParams params, Pageable pageable) {
        applyFilters(params, pageable)
    }

    /**
     * Create a new Course
     * @param {@link CourseDTO}
     * @param String author
     */
    CourseEntity createCourse(CourseDTO dto, String author) {
        log.info "Creating a new Course"

        // Check which fields are present
        List<CourseOrderEntity> addedModules = new ArrayList<>()

        // 1. Build and Save a New Course
        CourseEntity course = new CourseEntity(
            title: dto.title,
            description: dto.description ?: dto.title,
            status: dto.status,
            beginDate: dto.beginDate ?: null,
            endDate: dto.endDate ?: null,
            extras: dto.extras ?: null,
            createdBy: author,
            updatedBy: author,
            image: imagesService.create(dto.image)
        )

        try {
            courseRepository.save(course)

            // Add Modules to Course
            dto.moduleIds.eachWithIndex { it, index ->
                addedModules.add(addModuleToCourse(it, course, author, index))
            }

        } catch (Exception e) {
            // rollback the save operation
            courseRepository.delete(course)
            addedModules.each {
                courseOrderRepository.delete(it)
            }

            throw e // rethrow the exception to propagate it up the call stack
        }
        return course
    }

    // Add CourseOrder on Course Create
    CourseOrderEntity addModuleToCourse(String moduleId, CourseEntity course, String author, Long index) {
        ModulesEntity module = modulesService.findById(moduleId)

        courseOrderRepository.save(new CourseOrderEntity(
            module: module,
            course: course,
            position: index,
            createdBy: author,
            updatedBy: author
        ))

    }

    /**
     * Update a {@link CourseEntity}.
     * @param {@link CourseDTO}
     * @param String author that changed
     * @param Long id of the course to update
     * @return {@link CourseEntity}
     * */
    CourseEntity editCourse(String courseId, CourseDTO dto, String author) {
        List<CourseOrderEntity> addedModules = new ArrayList<>()
        CourseEntity course = findById(courseId)

        course.title = dto.title ?: course.title
        course.description = dto.description ?: course.description
        course.status = dto.status ?: course.status
        course.beginDate = dto.beginDate ?: course.beginDate
        course.endDate = dto.endDate ?: course.endDate
        course.extras = dto.extras ?: course.extras
        course.updatedBy = author
        course.image = dto.image ? imagesService.create(dto.image) : null

        if (dto.moduleIds) {
            // Remove all modules from this course
            List<CourseOrderEntity> courses = courseOrderRepository.findAllByCourse(course)
            courses.each {
                courseOrderRepository.delete(it)
            }

            // Add all modules to the course with correct index
            dto.moduleIds.eachWithIndex { it, index ->
                addedModules.add(addModuleToCourse(it, course, author, index))
            }
        }

        course = course.with {
            if (dto.moduleIds?.size()) {
                return it
            }
            it.courseOrder = []
            return it
        }

        return courseRepository.update(course)
    }

    void addModule(String courseId, String moduleId, String createdBy) {
        CourseEntity course = this.findById(courseId)

        CourseOrderEntity order = new CourseOrderEntity(
            module: modulesRepository.findById(UUID.fromString(moduleId)).get(),
            course: course,
            position: courseOrderRepository.countByCourse(course),
            createdBy: createdBy,
            updatedBy: createdBy
        )

        courseOrderRepository.save(order)
    }

    void addExistentModule(String courseId, String moduleId, String createdBy) {
        CourseEntity course = this.findById(courseId)
        ModulesEntity module = this.getModules(moduleId)

        CourseOrderEntity order = new CourseOrderEntity(
            module: module,
            course: course,
            position: courseOrderRepository.countByCourse(course) + 1,
            createdBy: createdBy,
            updatedBy: createdBy
        )

        courseOrderRepository.save(order)
    }

    void deleteCourse(String courseId) {
        CourseEntity course = this.findById(courseId)
        List<CourseOrderEntity> courses = courseOrderRepository.findAllByCourse(course)
        courseOrderRepository.deleteAll(courses)
        courseRepository.delete(course)
    }

    Map<String, Object> getActiveCourse() {
        List<CourseEntity> course = courseRepository.findByStatusIlike('active')

        if (course.isEmpty()) {
            throw ExceptionService.ParticipantNotFoundException(course)
        }

        return [
            before : courseRepository.findTop1ByEndDateLessThan(course.first().beginDate).orElse(null),
            current: course.first(),
            next   : courseRepository.findTop1ByBeginDateGreaterThan(course.first().endDate).orElse(null)
        ]
    }
}
