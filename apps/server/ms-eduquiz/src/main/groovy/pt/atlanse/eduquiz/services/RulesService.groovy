package pt.atlanse.eduquiz.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import io.micronaut.http.HttpStatus
import jakarta.inject.Singleton
import pt.atlanse.eduquiz.DTO.RulesDTO
import pt.atlanse.eduquiz.DTO.RulesParams
import pt.atlanse.eduquiz.domain.RulesEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.repositories.RulesRepository
import pt.atlanse.eduquiz.specifications.RulesSpecification

@Slf4j
@Singleton
class RulesService {
    RulesRepository rulesRepository

    RulesService(RulesRepository rulesRepository) {
        this.rulesRepository = rulesRepository
    }

    /**
     * Find a single {@link RulesEntity} by Id
     * @param Long id
     * @return An entire {@link RulesEntity}
     * */
    RulesEntity findById(String id) {
        rulesRepository.findById(UUID.fromString(id)).orElseThrow {
            new CustomException(
                "Rule was not found",
                "The rule with the id $id was not found",
                HttpStatus.NOT_FOUND
            )
        }
    }

    /**
     * Finds all {@link RulesEntity}
     * @return Page<{@link RulesEntity}>
     * */
    Page<RulesEntity> findAll(RulesParams params, Pageable pageable) {
        log.info "Using pageable arguments: Page_number: ${ pageable.offset }; Amount_of_articles: ${ pageable.size }"
        Page<RulesEntity> pages = applyFilters(params, pageable)
        return pages
    }

    /**
     * Build the pages of type {@link RulesEntity}
     * @param params Other parameters. Class type {@link RulesParams}. Allows usage of filters for image, status and text content
     * @param pageable ({@link Pageable}) obtained from the parameters and used for the pagination (e.g., size and page)
     * @return An entire {@link Page} book made of type {@link RulesEntity}
     * */
    Page<RulesEntity> applyFilters(RulesParams params, Pageable pageable) {
        PredicateSpecification<RulesEntity> specification = RulesSpecification.createQueryBySpecification(params)
        return rulesRepository.findAll(specification, pageable)
    }

    /**
     * Creating a new Rule.
     * @param {@link RulesDTO}, String author that requested
     * */
    void create(RulesDTO dto, String author) {
        log.info "Creating a new rule"
        // 1. Create a new rule
        RulesEntity newRule = new RulesEntity(
            code: dto.code,
            value: dto.value,
            title: dto.title,
            description: dto.description,
            createdBy: author,
            updatedBy: author
        )
        rulesRepository.save(newRule)
    }

    /**
     * Updating a new rule.
     * @param {@link RulesDTO}, String author that changed
     * @param String author that changed
     * @param Long id of the rule to update
     * */
    RulesEntity update(String id, RulesDTO dto, String author) {
        RulesEntity ruleToUpdate = findById(id)
        if (dto.title) {
            ruleToUpdate.title = dto.title
        }
        if (dto.description) {
            ruleToUpdate.description = dto.description
        }
        if (dto.code) {
            ruleToUpdate.code = dto.code
        }
        if (dto.value) {
            ruleToUpdate.value = dto.value
        }
        ruleToUpdate.updatedBy = author
        rulesRepository.update(ruleToUpdate)
    }

    /**
     * Deleting the rules by the id.
     * @param Long id of the rule to delete
     * */
    void delete(String id) {
        RulesEntity ruleToDelete = findById(id)
        rulesRepository.delete(ruleToDelete)
    }
}
