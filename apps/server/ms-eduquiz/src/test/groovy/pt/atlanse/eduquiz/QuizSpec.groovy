package pt.atlanse.eduquiz

import io.micronaut.data.model.Pageable
import io.micronaut.test.extensions.spock.annotation.MicronautTest
import io.micronaut.validation.validator.Validator
import jakarta.inject.Inject
import pt.atlanse.eduquiz.DTO.QuizDTO
import pt.atlanse.eduquiz.DTO.QuizParams
import pt.atlanse.eduquiz.domain.QuizEntity
import pt.atlanse.eduquiz.repositories.QuizRepository
import pt.atlanse.eduquiz.services.QuizService
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class QuizSpec extends Specification {
    @Inject
    QuizRepository repository

    @Inject
    QuizService service

    @Inject
    Validator validator

    def "test creating a new Quiz"() {
        given:
        def quiz = new QuizEntity(
            title: 'Title',
            description: 'Description',
            random: 1,
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def saved = repository.save(quiz)
        def retrieved = repository.findById(saved.id)

        then:
        retrieved.isPresent()
        retrieved.get().getRandom() == quiz.random
        retrieved.get().getTitle() == quiz.title
    }

    def "test list all Quiz"() {
        given:
        createQuiz('Quiz 1', 'Description 1')
        createQuiz('Quiz 2', 'Description 2')
        createQuiz('Quiz 3', 'Description 3')

        when:
        def result = repository.findAll()

        then:
        result.size() == 3
        result*.title == ['Quiz 1', 'Quiz 2', 'Quiz 3']
    }

    def "test update quiz"() {
        given:
        def quiz = createQuiz('Quiz 1', 'Description 1')
        def dto = new QuizDTO(
            title: 'Title 1'
        )
        when:
        def updated = service.update(quiz.id as String, dto, 'Updated')

        then:
        updated.id == quiz.id
        updated.createdBy == quiz.createdBy
        updated.updatedBy == 'Updated'
        updated.title == dto.title
    }

    def "test find a quiz by id"() {
        given:
        def quiz = createQuiz('Quiz 1', 'Description 1')

        when:
        def result = repository.findById(quiz.id)

        then:
        result.isPresent()
        result.get().title == quiz.title
    }

    def "should reject a quiz with null title"() {
        given:
        def quiz = new QuizEntity(
            description: 'Description',
            random: 1,
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def result = validator.validate(quiz)

        then:
        result.size() == 2
    }

    def "test delete a quiz by id"() {
        given:
        def quiz = createQuiz('Quiz 1', 'Description 1')

        when:
        repository.delete(quiz)

        then:
        !repository.findById(quiz.id).isPresent()
    }

    def "test find all by title"() {
        given:
        createQuiz('Quiz 1', 'Description 1')
        createQuiz('Quiz 2', 'Description 2')
        createQuiz('Quiz 3', 'Description 3')
        createQuiz('Quiz 4', 'Description 4')

        when:
        def result = service.findAll(new QuizParams(title: 'Quiz 2'), Pageable.from(0, 10))

        then:
        result.size() == 1
        result*.description == ['Description 2']
    }

    def createQuiz(String title, String description) {
        repository.save(new QuizEntity(
            title: title,
            description: description,
            random: 1,
            createdBy: 'Author',
            updatedBy: 'Author'
        ))
    }
}
