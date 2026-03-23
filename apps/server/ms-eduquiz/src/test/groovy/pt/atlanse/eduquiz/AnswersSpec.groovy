package pt.atlanse.eduquiz

import io.micronaut.data.model.Pageable
import io.micronaut.data.model.Sort
import io.micronaut.test.extensions.spock.annotation.MicronautTest
import io.micronaut.validation.validator.Validator
import jakarta.inject.Inject
import pt.atlanse.eduquiz.DTO.AnswersDTO
import pt.atlanse.eduquiz.DTO.AnswersParams
import pt.atlanse.eduquiz.domain.AnswersEntity
import pt.atlanse.eduquiz.domain.QuestionsEntity
import pt.atlanse.eduquiz.repositories.AnswersRepository
import pt.atlanse.eduquiz.repositories.QuestionsRepository
import pt.atlanse.eduquiz.services.AnswersService
import spock.lang.Shared
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class AnswersSpec extends Specification {
    @Inject
    AnswersRepository repository

    @Inject
    AnswersService service

    @Inject
    QuestionsRepository questionsRepository

    @Inject
    Validator validator

    @Shared
    QuestionsEntity question = new QuestionsEntity()

    def setup() {
        // reset the shared entity before each test method
        question = new QuestionsEntity(
            type: 'Type 1',
            description: 'Description A',
            points: 1,
            createdBy: 'Creation Author',
            updatedBy: 'Question Author'
        )
        questionsRepository.save(question)
    }

    def "test AnswersEntity is not null"() {
        given:
        AnswersEntity answer = new AnswersEntity()

        expect:
        answer != null
    }

    def "test creating a new answers entity"() {
        given:
        def answer = new AnswersEntity(
            question: question,
            content: 'This is the content',
            isCorrect: false,
            points: 1,
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def saved = repository.save(answer)

        then:
        saved.id != null
        saved.getQuestion().id == question.id
        saved.getContent() == answer.content
        saved.getIsCorrect() == answer.isCorrect
    }

    def "test updating an Answer Entity"() {
        given:
        def answer = createAnswers(true, 10, 'Content')

        def dto = new AnswersDTO(
            content: 'This is the new content',
            questionId: question.id,
            isCorrect: true,
            points: 1
        )

        when:
        def updated = service.update(answer.id.toString(), dto, 'new Author')

        then:
        updated.getId() == answer.id
        updated.getPoints() == dto.points
        updated.getUpdatedBy() == 'new Author'
        updated.getCreatedBy() == answer.createdBy
        updated.getContent() == dto.content
    }

    def "test list all answers"() {
        given:
        createAnswers(true, 50, 'Content')
        createAnswers(true, 100, 'Content 2')

        when:
        def result = repository.findAll()

        then:
        result.size() == 2
        result*.points == [50, 100]
    }

    def "test find an answer by id"() {
        given:
        def answer = createAnswers(true, 50, 'Content')

        when:
        def result = repository.findById(answer.id)

        then:
        result.isPresent()
        result.get().points == 50
    }

    def "test should reject answer with null points"() {
        given:
        def answer = new AnswersEntity(
            points: null,
            isCorrect: false,
            content: 'Content',
            question: question,
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def result = validator.validate(answer)

        then:
        result.size() == 1
        result.message == ['Points can not be null']
    }

    def "test delete Answer by id"() {
        given:
        def answer = createAnswers(true, 100, 'Content')

        when:
        repository.delete(answer)

        then:
        !repository.findById(answer.id).isPresent()
    }

    def "should return all answers by points desc"() {
        given:
        createAnswers(true, 60, 'Content 1')
        createAnswers(true, 70, 'Content 2')
        createAnswers(true, 50, 'Content 3')

        when:
        def result = service.findAll(null, Pageable.from(0, 10, Sort.of(Sort.Order.asc('points'))))

        then:
        result.size() == 3
        result[0].points == 50
        result[1].points == 60
        result[2].points == 70
    }

    def "should return all answers of a given question"() {
        given:
        createAnswers(true, 50, 'Content 1')
        createAnswers(true, 50, 'Content 2')
        createAnswers(true, 50, 'Content 3')

        when:
        def result = service.findAll(new AnswersParams(questionId: question.id), Pageable.from(0, 10, Sort.of(Sort.Order.asc('points'))))

        then:
        result.size() == 3
    }

    def createAnswers(boolean isCorrect, Long points, String content) {
        AnswersEntity answer = new AnswersEntity(
            question: question,
            content: content,
            points: points,
            isCorrect: isCorrect,
            createdBy: 'Author',
            updatedBy: 'Author'
        )
        repository.save(answer)
    }

}
