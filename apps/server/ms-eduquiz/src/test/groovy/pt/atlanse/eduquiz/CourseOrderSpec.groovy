package pt.atlanse.eduquiz

import io.micronaut.test.extensions.spock.annotation.MicronautTest
import io.micronaut.validation.validator.Validator
import jakarta.inject.Inject
import pt.atlanse.eduquiz.domain.CourseEntity
import pt.atlanse.eduquiz.domain.CourseOrderEntity
import pt.atlanse.eduquiz.domain.ModulesEntity
import pt.atlanse.eduquiz.repositories.CourseOrderRepository
import pt.atlanse.eduquiz.repositories.CourseRepository
import pt.atlanse.eduquiz.repositories.ModulesRepository
import spock.lang.Shared
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class CourseOrderSpec extends Specification {
    @Inject
    CourseOrderRepository repository

    @Inject
    CourseRepository courseRepository

    @Inject
    ModulesRepository modulesRepository

    @Inject
    Validator validator

    @Shared
    CourseEntity course = new CourseEntity()

    @Shared
    ModulesEntity module = new ModulesEntity()

    def setup() {
        // reset the shared entity before each test method
        course = new CourseEntity(
            title: 'New Course',
            description: 'New Description',
            status: 'New State',
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        module = new ModulesEntity(
            title: 'New Module',
            description: 'New Description',
            status: 'New Status',
            categories: 'New Category',
            extras: '{isSponsor: true}',
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        courseRepository.save(course)
        modulesRepository.save(module)
    }

    def "test is CourseOrder not null"() {
        given:
        CourseOrderEntity courseOrder = new CourseOrderEntity()

        expect:
        courseOrder != null
    }

    def "test creating a new CourseOrder"() {
        given:
        def author = 'Spock Test'
        def courseOrder = new CourseOrderEntity(
            course: course,
            module: module,
            position: 10,
            createdBy: author,
            updatedBy: author
        )

        when:
        def saved = repository.save(courseOrder)

        then:
        saved.id != null
        saved.getModule().id == module.id
        saved.getCourse().id == course.id
        saved.getCreatedBy() == author
    }

    def "test list all course order"() {
        given:
        createCourseOrder(1)
        createCourseOrder(2)

        when:
        def result = repository.findAll()

        then:
        result.size() == 2
    }

    def "test find a course order by id"() {
        given:
        def co = createCourseOrder(1)

        when:
        def result = repository.findById(co.id)

        then:
        result.isPresent()
        result.get().course.id == co.course.id
    }

    def "test should reject new CourseOrder with null Course"() {
        given:
        def author = 'Author'
        def co = new CourseOrderEntity(
            createdBy: author,
            updatedBy: author,
            position: 1,
            module: module
        )

        when:
        def result = validator.validate(co)

        then:
        result.size() == 1
        result.message == ['CourseId can not be null']
    }

    def "test delete CourseOrder by id"() {
        given:
        def co = createCourseOrder(1)

        when:
        repository.delete(co)

        then:
        !repository.findById(co.id).isPresent()
    }

    def createCourseOrder(Long position) {
        repository.save(new CourseOrderEntity(
            course: course,
            module: module,
            position: position,
            createdBy: 'Author',
            updatedBy: 'Author'
        ))
    }
}
