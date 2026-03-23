package pt.atlanse.eduquiz

import io.micronaut.test.extensions.spock.annotation.MicronautTest
import io.micronaut.validation.validator.Validator
import jakarta.inject.Inject
import pt.atlanse.eduquiz.domain.ModulesEntity
import pt.atlanse.eduquiz.domain.ModulesOrderEntity
import pt.atlanse.eduquiz.domain.QuizEntity
import pt.atlanse.eduquiz.repositories.ModulesOrderRepository
import pt.atlanse.eduquiz.repositories.ModulesRepository
import pt.atlanse.eduquiz.repositories.QuizRepository
import spock.lang.Shared
import spock.lang.Specification

import javax.transaction.Transactional

@MicronautTest
@Transactional
class ModulesOrderSpec extends Specification {
    @Inject
    ModulesOrderRepository repository

    @Inject
    Validator validator

    @Inject
    QuizRepository quizRepository

    @Inject
    ModulesRepository modulesRepository

    @Shared
    QuizEntity quiz = new QuizEntity()

    @Shared
    ModulesEntity module = new ModulesEntity()

    def setup() {
        // reset the shared entity before each test
        quiz = new QuizEntity(
            updatedBy: 'Author',
            createdBy: 'Author',
            title: 'Quiz Title',
            description: 'Quiz Description',
            random: false
        )

        module = new ModulesEntity(
            updatedBy: 'Author',
            createdBy: 'Author',
            title: 'Module Title',
            description: 'Module Description',
            status: 'Module Status'
        )

        quizRepository.save(quiz)
        modulesRepository.save(module)
    }

    def "test if module order is not null"() {
        given:
        ModulesOrderEntity mo = new ModulesOrderEntity()

        expect:
        mo != null
    }

    def "test creating a new modules order entity"() {
        given:
        def author = 'Author'
        def mo = new ModulesOrderEntity(
            createdBy: author,
            updatedBy: author,
            module: module,
            quiz: quiz,
            position: 1
        )

        when:
        def saved = repository.save(mo)
        def retrieved = repository.findById(saved.id)

        then:
        retrieved.isPresent()
        retrieved.get().getModule().id == module.id
        retrieved.get().getQuiz().id == quiz.id
    }

    def "test list all ModuleOrder"() {
        given:
        createModuleOrder(1)
        createModuleOrder(2)

        when:
        def result = repository.findAll()

        then:
        result.size() == 2
        result*.position == [1, 2]
    }

    def "test find lesson by id"() {
        given:
        def mo = createModuleOrder(1)

        when:
        def result = repository.findById(mo.id)

        then:
        result.isPresent()
        result.get().getId() == mo.id
    }

    def "test delete ModulesOrder by id"() {
        given:
        def mo = createModuleOrder(1)

        when:
        repository.delete(mo)

        then:
        !repository.findById(mo.id).isPresent()
    }

    def createModuleOrder(Long position) {
        def author = 'Author'
        repository.save(new ModulesOrderEntity(
            createdBy: author,
            updatedBy: author,
            module: module,
            quiz: quiz,
            position: position
        ))
    }
}
