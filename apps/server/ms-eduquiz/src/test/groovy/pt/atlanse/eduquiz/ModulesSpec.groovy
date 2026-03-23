package pt.atlanse.eduquiz

import io.micronaut.data.model.Pageable
import io.micronaut.data.model.Sort
import io.micronaut.test.extensions.spock.annotation.MicronautTest
import io.micronaut.validation.validator.Validator
import jakarta.inject.Inject
import pt.atlanse.eduquiz.DTO.ModulesDTO
import pt.atlanse.eduquiz.domain.ModulesEntity
import pt.atlanse.eduquiz.repositories.ModulesRepository
import pt.atlanse.eduquiz.services.ModulesService
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class ModulesSpec extends Specification {
    @Inject
    ModulesRepository repository

    @Inject
    ModulesService service

    @Inject
    Validator validator

    def "test creating a new Module"() {
        given:
        def module = new ModulesEntity(
            title: 'title',
            description: 'description',
            status: 'Draft',
            createdBy: 'author',
            updatedBy: 'author'
        )

        when:
        def saved = repository.save(module)
        def retrieved = repository.findById(saved.id)

        then:
        retrieved.isPresent()
        retrieved.get().getTitle() == module.title
    }

    def "test list all modules"() {
        given:
        createModule('Module A', 'Description A')
        createModule('Module B', 'Description B')

        when:
        def result = repository.findAll()

        then:
        result.size() == 2
        result*.title == ['Module A', 'Module B']
    }

    def "test update module"() {
        given:
        def saved = createModule('Module A', 'Description A')
        def dto = new ModulesDTO(
            title: 'New Title',
            description: 'New Description'
        )

        when:
        def updated = service.update(saved.id.toString(), dto, 'New Author')

        then:
        updated.id == saved.id
        updated.updatedBy == 'New Author'
        updated.title == dto.title
    }

    def "test find a module by id"() {
        given:
        def module = createModule('Module A', 'Description A')

        when:
        def result = repository.findById(module.id)

        then:
        result.isPresent()
        result.get().title == 'Module A'
    }

    def "should reject a module with null title"() {
        given:
        def module = new ModulesEntity(
            description: 'description',
            status: 'Draft',
            createdBy: 'author',
            updatedBy: 'author'
        )

        when:
        def result = validator.validate(module)

        then:
        result.size() == 2
    }

    def "test delte a module by id"() {
        given:
        def module = createModule('Module A', 'Description A')

        when:
        repository.delete(module)

        then:
        !repository.findById(module.id).isPresent()
    }

    def "should return all modules sorted by title"() {
        given:
        createModule('Module B', 'Description A')
        createModule('Module C', 'Description A')
        createModule('Module A', 'Description A')

        when:
        def result = service.findAll(null, Pageable.from(0, 10, Sort.of(Sort.Order.asc('title'))))

        then:
        result.size() == 3
        result[0].title == 'Module A'
        result[1].title == 'Module B'
        result[2].title == 'Module C'
    }

    def createModule(String title, String description) {
        def author = 'Author'
        repository.save(new ModulesEntity(
            title: title,
            description: description,
            status: 'Draft',
            createdBy: author,
            updatedBy: author
        ))
    }
}
