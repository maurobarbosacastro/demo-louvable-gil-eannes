package pt.atlanse.eduquiz

import io.micronaut.data.model.Pageable
import io.micronaut.data.model.Sort
import io.micronaut.test.extensions.spock.annotation.MicronautTest
import io.micronaut.validation.validator.Validator
import jakarta.inject.Inject
import pt.atlanse.eduquiz.DTO.CourseDTO
import pt.atlanse.eduquiz.DTO.CourseParams
import pt.atlanse.eduquiz.domain.CourseEntity
import pt.atlanse.eduquiz.repositories.CourseRepository
import pt.atlanse.eduquiz.services.CourseService
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class CourseSpec extends Specification {

    @Inject
    CourseRepository repository

    @Inject
    Validator validator

    @Inject
    CourseService courseService

    def "test creating a new Course Entity"() {
        given:
        def title = "Title Spec"
        def description = "Description Spec"
        def status = "active"
        def author = 'SPOCK Test'

        when:
        def content = new CourseEntity(
            title: title,
            description: description,
            status: status,
            updatedBy: author,
            createdBy: author
        )

        CourseEntity saved = repository.save(content)
        def retrievedCourse = repository.findById(saved.getId())

        then:
        retrievedCourse.isPresent()
        retrievedCourse.get().getTitle() == title
        retrievedCourse.get().getDescription() == description
        retrievedCourse.get().getStatus() == status
        retrievedCourse.get().getCreatedBy() == author
    }

    def "test list all Course Entity"() {
        given:
        def course1 = new CourseEntity(title: 'Title 1', description: 'Description 1', status: 'state 1', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test')
        def course2 = new CourseEntity(title: 'Title 2', description: 'Description 2', status: 'state 2', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test')
        repository.save(course1)
        repository.save(course2)

        when:
        def result = repository.findAll()

        then:
        result.size() == 2
        result*.title == ['Title 1', 'Title 2']
    }

    def "test update course"() {
        given:
        def course1 = new CourseEntity(title: 'Title 1', description: 'Description 1', status: 'state 1', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test')
        def saved = repository.save(course1)

        when:
        def updated = courseService.editCourse(saved.id.toString(), new CourseDTO(title: 'New Title', description: 'New Description', status: 'New State'), 'New Author')

        then:
        updated.id == saved.id
        updated.title == 'New Title'
        updated.description == 'New Description'
        updated.status == 'New State'
    }

    def "test find Course Entity by Id"() {
        given:
        def course = new CourseEntity(title: 'Title 1', description: 'Description 1', status: 'state 1', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test')
        repository.save(course)

        when:
        def result = repository.findById(course.id)

        then:
        result.isPresent()
        result.get().title == 'Title 1'
    }

    def "should reject course with null title"() {
        given:
        def course = new CourseEntity(title: null, description: 'Description 1', status: 'state 1', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test')

        when:
        def result = validator.validate(course)

        then:
        result.size() == 2
    }

    def "test delete Course by Id"() {
        given:
        def course = new CourseEntity(title: 'Title 1', description: 'Description 1', status: 'state 1', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test')
        repository.save(course)

        when:
        repository.delete(course)

        then:
        !repository.findById(course.id).isPresent()
    }

    def "should return courses sorted by title in ascending order"() {
        given:
        def courses = [
            new CourseEntity(title: 'Title C', description: 'Description 1', status: 'state 1', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test'),
            new CourseEntity(title: 'Title B', description: 'Description 2', status: 'state 2', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test'),
            new CourseEntity(title: 'Title A', description: 'Description 3', status: 'state 3', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test')
        ]
        repository.saveAll(courses)

        when:
        def result = courseService.getCourses(null, Pageable.from(0, 10, Sort.of(Sort.Order.asc('title'))))

        then:
        result.size() == 3
        result[0].title == "Title A"
        result[1].title == "Title B"
        result[2].title == "Title C"
    }

    def "should return courses sorted by title of page 2"() {
        given:
        def courses = [
            new CourseEntity(title: 'Title C', description: 'Description 1', status: 'state 1', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test'),
            new CourseEntity(title: 'Title B', description: 'Description 2', status: 'state 2', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test'),
            new CourseEntity(title: 'Title A', description: 'Description 3', status: 'state 3', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test')
        ]
        repository.saveAll(courses)

        when:
        def result = courseService.getCourses(null, Pageable.from(1, 2, Sort.of(Sort.Order.asc('title'))))

        then:
        result.size() == 1
        result[0].title == "Title C"
    }

    def "should return courses filtered by title name"() {
        given:
        def courses = [
            new CourseEntity(title: 'Title C', description: 'Description 1', status: 'state 1', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test'),
            new CourseEntity(title: 'Title B', description: 'Description 2', status: 'state 2', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test'),
            new CourseEntity(title: 'Title A', description: 'Description 3', status: 'state 3', updatedBy: 'SPOCK Test', createdBy: 'SPOCK test')
        ]
        repository.saveAll(courses)

        when:
        def result = courseService.getCourses(new CourseParams('Title C'), Pageable.from(0, 10, Sort.of(Sort.Order.asc('title'))))

        then:
        result.size() == 1
    }
}
