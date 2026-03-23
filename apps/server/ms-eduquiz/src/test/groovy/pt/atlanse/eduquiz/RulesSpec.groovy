package pt.atlanse.eduquiz

import io.micronaut.data.model.Pageable
import io.micronaut.test.extensions.spock.annotation.MicronautTest
import io.micronaut.validation.validator.Validator
import jakarta.inject.Inject
import pt.atlanse.eduquiz.DTO.RulesDTO
import pt.atlanse.eduquiz.DTO.RulesParams
import pt.atlanse.eduquiz.domain.RulesEntity
import pt.atlanse.eduquiz.repositories.RulesRepository
import pt.atlanse.eduquiz.services.RulesService
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class RulesSpec extends Specification {
    @Inject
    RulesRepository repository

    @Inject
    RulesService service

    @Inject
    Validator validator

    def "test creating a new Rule"() {
        given:
        def rule = new RulesEntity(
            code: 'code',
            value: 'value',
            title: 'title',
            description: 'Description',
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def saved = repository.save(rule)
        def retrieved = repository.findById(saved.id)

        then:
        retrieved.isPresent()
        retrieved.get().title == rule.title
        retrieved.get().description == rule.description
    }

    def "test list all Rules"() {
        given:
        createRules('Code', 'Value', 'Title')
        createRules('Code', 'Value', 'Title')
        createRules('Code', 'Value', 'Title')

        when:
        def result = repository.findAll()

        then:
        result.size() == 3
    }

    def "test find a rule by id"() {
        given:
        def rule = createRules('Code', 'Value', 'Title')

        when:
        def result = repository.findById(rule.id)

        then:
        result.isPresent()
        result.get().title == rule.title
    }

    def "test update rule"() {
        given:
        def rule = createRules('Code 1', 'Value 1', 'Title 1')
        def dto = new RulesDTO(
            code: 'New Code'
        )

        when:
        def updated = service.update(rule.id as String, dto, 'Updated')

        then:
        updated.id == rule.id
        updated.createdBy == rule.createdBy
        updated.updatedBy == 'Updated'
        updated.code == dto.code
    }

    def "should reject a rule without title"() {
        given:
        def rule = new RulesEntity(
            code: 'code',
            value: 'value',
            description: 'Description',
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def result = validator.validate(rule)

        then:
        result.size() == 2
    }

    def "test delete a rule by id"() {
        given:
        def rule = createRules('Code', 'Value', 'Title')

        when:
        repository.delete(rule)

        then:
        !repository.findById(rule.id).isPresent()
    }

    def "test find all rules by title"() {
        given:
        createRules('Code 1', 'Value 1', 'Title 1')
        createRules('Code 2', 'Value 2', 'Title 2')
        createRules('Code 3', 'Value 3', 'Title 3')

        when:
        def result = service.findAll(new RulesParams('title': 'Title 2'), Pageable.from(0, 10))

        then:
        result.size() == 1
        result*.value == ['Value 2']
    }

    def createRules(String code, String value, String title) {
        repository.save(new RulesEntity(
            code: code,
            value: value,
            title: title,
            description: 'Description',
            createdBy: 'Author',
            updatedBy: 'Author'
        ))
    }
}
