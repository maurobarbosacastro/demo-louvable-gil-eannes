package pt.atlanse.eduquiz

import io.micronaut.data.model.Pageable
import io.micronaut.data.model.Sort
import io.micronaut.test.extensions.spock.annotation.MicronautTest
import io.micronaut.validation.validator.Validator
import jakarta.inject.Inject
import pt.atlanse.eduquiz.DTO.LessonDTO
import pt.atlanse.eduquiz.domain.LessonEntity
import pt.atlanse.eduquiz.repositories.LessonRepository
import pt.atlanse.eduquiz.services.LessonService
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class LessonSpec extends Specification {
    @Inject
    LessonRepository repository

    @Inject
    LessonService service

    @Inject
    Validator validator

    def "test creating a new Lesson"() {
        given:
        def lesson = new LessonEntity(
            title: 'title',
            subtitle: 'Subtitle',
            type: 'Type',
            conclusion: 'Conclusion',
            status: 'Published',
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def saved = repository.save(lesson)
        def retrievedLesson = repository.findById(saved.getId())

        then:
        retrievedLesson.isPresent()
        retrievedLesson.get().getTitle() == lesson.title
        retrievedLesson.get().getStatus() == lesson.status
    }

    def "test list all Lessons"() {
        given:
        createLesson('Title 1')
        createLesson('Title 2')

        when:
        def result = repository.findAll()

        then:
        result.size() == 2
        result*.title == ['Title 1', 'Title 2']
    }

    def "test update lesson title"() {
        given:
        def lesson = createLesson('Lesson 1')

        when:
        def updated = service.update(lesson.id.toString(), new LessonDTO(
            title: 'Title',
            subtitle: 'Subtitle',
            type: 'type',
            status: 'status'
        ), 'Updated By')

        then:
        updated.id == lesson.id
        updated.title == 'Title'
        updated.updatedBy == 'Updated By'
    }

    def "test find a lesson by id"() {
        given:
        def lesson = createLesson('Lesson')

        when:
        def result = repository.findById(lesson.id)

        then:
        result.isPresent()
        result.get().title == 'Lesson'
    }

    def "should reject a lesson with null title"() {
        given:
        def lesson = new LessonEntity(
            subtitle: 'Subtitle',
            type: 'Type',
            conclusion: 'Conclusion',
            status: 'Published',
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def result = validator.validate(lesson)

        then:
        result.size() == 2
    }

    def "test delete lesson by id"() {
        given:
        def lesson = createLesson('Lesson A')

        when:
        repository.delete(lesson)

        then:
        !repository.findById(lesson.id).isPresent()
    }

    def "should return all lessons sorted by title in asc"() {
        given:
        createLesson('Lesson A')
        createLesson('Lesson C')
        createLesson('Lesson B')

        when:
        def result = service.findAll(Pageable.from(0, 10, Sort.of(Sort.Order.asc('title'))))

        then:
        result.size() == 3
        result[0].title == 'Lesson A'
        result[1].title == 'Lesson B'
        result[2].title == 'Lesson C'
    }

    def createLesson(String title) {
        repository.save(new LessonEntity(
            title: title,
            subtitle: 'Subtitle',
            type: 'Type',
            conclusion: 'Conclusion',
            status: 'Published',
            createdBy: 'Author',
            updatedBy: 'Author'
        ))
    }
}
