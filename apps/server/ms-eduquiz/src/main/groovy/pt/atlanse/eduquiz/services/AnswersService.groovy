package pt.atlanse.eduquiz.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.eduquiz.DTO.AnswersDTO
import pt.atlanse.eduquiz.DTO.AnswersParams
import pt.atlanse.eduquiz.domain.AnswersEntity
import pt.atlanse.eduquiz.domain.QuestionsEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.repositories.AnswersRepository
import pt.atlanse.eduquiz.repositories.QuestionsRepository
import pt.atlanse.eduquiz.specifications.AnswersSpecification

@Slf4j
@Singleton
class AnswersService {
    @Inject
    AnswersRepository answersRepository

    @Inject
    QuestionsRepository questionsRepository

    AnswersService(AnswersRepository answersRepository, QuestionsRepository questionsRepository) {
        this.answersRepository = answersRepository
        this.questionsRepository = questionsRepository
    }

    /**
     * Finds a single {@link AnswersEntity} by id
     * @params String id
     * @return an entire {@link AnswersEntity}
     */
    AnswersEntity findById(String id) {
        answersRepository.findById(UUID.fromString(id)).orElseThrow {
            new CustomException("Answered Question not found",
                "The answered question with id $id was not found",
                HttpStatus.NOT_FOUND)
        }
    }

    /**
     * Finds all {@link AnswersEntity}
     * @return List of Pages {@link AnswersEntity}
     */
    Page<AnswersEntity> findAll(AnswersParams params, Pageable pageable) {
        log.info "Using pageable arguments: Page_number: ${pageable.offset}; Amount_of_articles: ${pageable.size}"
        applyFilters(params, pageable)
    }

    /**
     * Create a new Answer
     * @params {@link AnswersDTO}
     * @params String author that created
     */
    UUID create(AnswersDTO dto, String author) {
        log.info "Creating a new answer"
        // 1. Build and save a new answer
        AnswersEntity answer = new AnswersEntity(question: findQuestion(dto.questionId),
            content: dto.content,
            isCorrect: dto.isCorrect,
            points: dto.points ? dto.points : 1,
            createdBy: author,
            updatedBy: author)

        answersRepository.save(answer)
        return answer.id
    }

    /**
     * Update an Answer
     * * @params {@link AnswersDTO},
     * @params String author that updated
     * @params String id of the answer to edit
     */
    AnswersEntity update(String id, AnswersDTO dto, String author) {
        AnswersEntity answerToUpdate = findById(id)
        if (dto.questionId) {
            answerToUpdate.question = findQuestion(dto.questionId)
        }
        if (dto.points) {
            answerToUpdate.points = dto.points
        }
        answerToUpdate.isCorrect = dto.isCorrect
        if (dto.content) {
            answerToUpdate.content = dto.content
        }
        answerToUpdate.updatedBy = author
        answersRepository.update(answerToUpdate)
    }

    /**
     * Delete Answer by Id
     * @param String id of the answer to delete
     */
    void delete(String id) {
        answersRepository.delete(findById(id))
    }

    /**
     * Build the pages of type {@link AnswersEntity}
     * @param params - Other parameters. Class type {@link AnswersParams}. Allows usage of filters
     * @param pageable ({@link Pageable}) obtained from the parameters and used for the pagination (e.g., size and page)
     * @return An entire {@link Page} book made of type {@link AnswersEntity}
     * */
    private Page<AnswersEntity> applyFilters(AnswersParams params, Pageable pageable) {
        PredicateSpecification<AnswersEntity> specification = AnswersSpecification.createQueryBySpecification(params)
        return answersRepository.findAll(specification, pageable)
    }

    // Find a Question by id or throw an exception if it is not found
    private QuestionsEntity findQuestion(String id) {
        questionsRepository.findById(UUID.fromString(id)).orElseThrow {
            new CustomException("Error fetching question",
                "Error fetching question with id $id; Not Found",
                HttpStatus.NOT_FOUND)
        }
    }

    AnswersDTO mapToDto(AnswersEntity answer) {
        new AnswersDTO(
            id: answer.id,
            content: answer.content,
            isCorrect: answer.isCorrect,
            points: answer.points,
            createdAt: answer.createdAt
        )
    }
}
