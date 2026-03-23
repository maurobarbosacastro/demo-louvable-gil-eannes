package pt.atlanse.eduquiz

import io.micronaut.data.model.Pageable
import io.micronaut.test.extensions.spock.annotation.MicronautTest
import io.micronaut.validation.validator.Validator
import jakarta.inject.Inject
import pt.atlanse.eduquiz.DTO.AnswersDTO
import pt.atlanse.eduquiz.DTO.QuestionsDTO
import pt.atlanse.eduquiz.DTO.QuestionsParams
import pt.atlanse.eduquiz.domain.QuestionsEntity
import pt.atlanse.eduquiz.repositories.QuestionsRepository
import pt.atlanse.eduquiz.services.QuestionsService
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class QuestionsSpec extends Specification {
    @Inject
    QuestionsRepository repository

    @Inject
    QuestionsService service

    @Inject
    Validator validator

    def "test creating a new QuestionsEntity"() {
        given:
        def question = new QuestionsEntity(
            type: 'Title',
            points: 35,
            description: "This is a description",
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def saved = repository.save(question)
        def retrieved = repository.findById(saved.id)

        then:
        retrieved.isPresent()
        retrieved.get().getType() == question.type
    }

    def "test list all questions"() {
        given:
        createQuestion('Type 1', 25)
        createQuestion('Type 1', 26)
        createQuestion('Type 1', 27)

        when:
        def result = repository.findAll()

        then:
        result.size() == 3
        result*.points == [25, 26, 27]
    }

    def "test update question"() {
        given:
        def question = createQuestion('Type 1', 25)
        def dto = new QuestionsDTO(
            type: 'TEXT'
        )

        when:
        def updated = service.update(question.id as String, dto, 'New Author')

        then:
        updated.id == question.id
        updated.createdBy == question.createdBy
        updated.updatedBy == 'New Author'
    }

    def "test find a question by id"() {
        given:
        def question = createQuestion('Test', 1)

        when:
        def result = repository.findById(question.id)

        then:
        result.isPresent()
        result.get().type == 'Test'
    }

    def "should reject a question with null description"() {
        given:
        def question = new QuestionsEntity(
            createdBy: 'Author',
            updatedBy: 'Author',
            points: 25
        )

        when:
        def result = validator.validate(question)

        then:
        result.size() == 2
    }

    def "should add an answer to a question"() {
        given:
        List<AnswersDTO> answers = new ArrayList<>()
        def question = new QuestionsDTO(
            type: 'TEXT',
            description: 'Description 1',
            answers: answers
        )
        answers.add(createAnswer('Answer 1', true))
        answers.add(createAnswer('Answer 2', false))
        answers.add(createAnswer('Answer 3', false))
        answers.add(createAnswer('Answer 4', false))

        when:
        def saved = service.create(question, 'author')
        def savedQuestion = service.find(saved.id.toString())

        then:
        savedQuestion.description == 'Description 1'
    }

    def "test delete a question by id"() {
        given:
        def question = createQuestion('Type 1', 25)

        when:
        repository.delete(question)

        then:
        !repository.findById(question.id).isPresent()
    }

    def "test find all by description"() {
        given:
        createQuestion('TEXT', 10)
        createQuestion('TEXT', 10)
        createQuestion('TEXT', 10)
        createQuestion('VIDEO', 10)

        when:
        def result = service.findAll(new QuestionsParams(description: 'TEXT'), Pageable.from(0, 10))

        then:
        result.size() == 3
    }

    def createQuestion(String type, Long points) {
        repository.save(new QuestionsEntity(
            type: type,
            points: points,
            description: type,
            createdBy: 'Author',
            updatedBy: 'Author'
        ))
    }

    def createAnswer(String value, boolean isCorrect) {
        return new AnswersDTO(
            content: value,
            isCorrect: isCorrect,
            points: 1
        )
    }
}
