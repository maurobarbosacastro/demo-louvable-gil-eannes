package pt.atlanse.eduquiz

import io.micronaut.test.extensions.spock.annotation.MicronautTest
import jakarta.inject.Inject
import pt.atlanse.eduquiz.domain.CategoryEntity
import pt.atlanse.eduquiz.repositories.CategoryRepository
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class CategorySpec extends Specification {
    @Inject
    CategoryRepository repository

    def "test Category is not null"() {
        given:
        CategoryEntity category = new CategoryEntity()

        expect:
        category != null
    }

    def "test creating a new category"() {
        given:
        def category = new CategoryEntity(
            name: 'Category',
            createdBy: 'Author',
            updatedBy: 'Author'
        )

        when:
        def saved = repository.save(category)

        then:
        saved.id != null
        saved.getName() == category.name
    }

    def "test listing all category"() {
        given:
        createCategory()
        createCategory()
        createCategory()

        when:
        def result = repository.findAll()

        then:
        result.size() == 3
    }

    def "test find a category by id"() {
        given:
        def category = createCategory()

        when:
        def result = repository.findById(category.id)

        then:
        result.isPresent()
    }

    def "test delete a category by id"() {
        given:
        def category = createCategory()

        when:
        repository.delete(category)

        then:
        !repository.findById(category.id).isPresent()
    }

    def createCategory() {
        repository.save(new CategoryEntity(name: 'Category', createdBy: 'Author', updatedBy: 'Author'))
    }
}
