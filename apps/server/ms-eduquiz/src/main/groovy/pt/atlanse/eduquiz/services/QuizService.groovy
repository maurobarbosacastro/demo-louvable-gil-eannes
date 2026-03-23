package pt.atlanse.eduquiz.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.eduquiz.DTO.QuizDTO
import pt.atlanse.eduquiz.DTO.QuizParams
import pt.atlanse.eduquiz.domain.CategoryEntity

import pt.atlanse.eduquiz.domain.QuestionsEntity
import pt.atlanse.eduquiz.domain.QuizEntity
import pt.atlanse.eduquiz.domain.QuizQuestionsEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.repositories.CategoryRepository

import pt.atlanse.eduquiz.repositories.QuestionsRepository
import pt.atlanse.eduquiz.repositories.QuizQuestionsRepository
import pt.atlanse.eduquiz.repositories.QuizRepository
import pt.atlanse.eduquiz.specifications.QuizSpecification

@Slf4j
@Singleton
class QuizService {

    @Inject
    QuizRepository quizRepository

    @Inject
    CategoryRepository categoryRepository

    @Inject
    QuestionsRepository questionRepository

    @Inject
    QuizQuestionsRepository quizQuestionsRepository

    /**
     * Find a single {@link QuizEntity} by Id
     * @param Long id
     * @return An entire {@link QuizEntity}
     * */
    QuizEntity findById(String quizId) {
        return quizRepository.findById(UUID.fromString(quizId)).orElseThrow {
            new CustomException("Quiz not found",
                "The quiz with id $quizId was not found",
                HttpStatus.NOT_FOUND)
        }
    }

    /**
     * Finds all {@link QuizEntity}
     * @return List of Page<{@link QuizEntity}>
     * */
    Page<QuizEntity> findAll(QuizParams params, Pageable pageable) {
        log.info "Using pageable arguments: Page_number: ${ pageable.offset }; Amount_of_articles: ${ pageable.size }"
        return applyFilters(params, pageable)
    }

    /**
     * Build the pages of type {@link QuizEntity}
     * @param params Other parameters. Class type {@link QuizParams}. Allows usage of filters for category
     * @param pageable ({@link Pageable}) obtained from the parameters and used for the pagination (e.g., size and page)
     * @return An entire {@link Page} book made of type {@link QuizEntity}
     * */
    private Page<QuizEntity> applyFilters(QuizParams params, Pageable pageable) {
        PredicateSpecification<QuizEntity> specification = QuizSpecification.createQueryBySpecification(params)
        return quizRepository.findAll(specification, pageable)
    }

    // Fetching a Category by the ID and throwing an exception if it is not found.
    private CategoryEntity findCategory(String id) {
        CategoryEntity category = categoryRepository.findById(id).orElseThrow {
            new CustomException("Error fetching category",
                "Error fetching category $id; Not Found",
                HttpStatus.BAD_REQUEST)
        }
        return category
    }

    // Fetching a Category by the ID and throwing an exception if it is not found.
    private QuestionsEntity findQuestion(String id) {
        QuestionsEntity question = questionRepository.findById(UUID.fromString(id)).orElseThrow {
            new CustomException("Error fetching question",
                "Error fetching question $id; Not Found",
                HttpStatus.BAD_REQUEST)
        }
        return question
    }

    /**
     * Creating a new quiz
     * @param {@link QuizDTO}, String author that created
     * */
    void create(QuizDTO dto, String author) {
        log.info "Creating a new Quiz"
        // 1. Build and save a new Quiz
        QuizEntity quiz = new QuizEntity(title: dto.title,
            description: dto.description,
            random: dto.random,
            category: findCategory(dto.categoryId),
            createdBy: author,
            updatedBy: author)
        quizRepository.save(quiz)
    }

    /**
     * Updating an existent Quiz.
     * @param {@link QuizDTO}
     * @param String author that changed
     * @param Long id of the quiz to update
     * */
    QuizEntity update(String id, QuizDTO dto, String author) {
        QuizEntity quizToUpdate = findById(id)
        if (dto.categoryId) {
            quizToUpdate.category = findCategory(dto.categoryId)
        }
        if (dto.title) {
            quizToUpdate.title = dto.title
        }
        if (dto.description) {
            quizToUpdate.description = dto.description
        }
        if (dto.random) {
            quizToUpdate.random = dto.random
        }
        quizToUpdate.updatedBy = author
        quizRepository.update(quizToUpdate)
    }

    /**
     * Deleting the Quiz by ID
     * @param String id of the quiz to delete
     */
    void delete(String id) {
        QuizEntity quizToDelete = findById(id)
        quizRepository.delete(quizToDelete)
    }

    /**
     * Deleting a Quiz Question
     * @params String quizQuestionId
     */
    void deleteQuizQuestion(String quizQuestionId) {
        QuizQuestionsEntity questionToDelete = quizQuestionsRepository.findById(UUID.fromString(quizQuestionId)).orElseThrow {
            new CustomException("Error fetching question",
                "Error fetching quiz question $quizQuestionId; Not Found",
                HttpStatus.BAD_REQUEST)
        }
        quizQuestionsRepository.delete(questionToDelete)
    }

    /**
     * Adding a question to a quiz
     * @param quizId
     * @param questionId
     * @param author
     */
    void addQuestion(String quizId, String questionId, String author) {
        QuizEntity quiz = findById(quizId)
        QuestionsEntity question = findQuestion(questionId)
        QuizQuestionsEntity quizQuestion = new QuizQuestionsEntity(quiz: quiz,
            questions: question,
            createdBy: author,
            updatedBy: author)
        quizQuestionsRepository.save(quizQuestion)
    }

    /**
     * Build the pages of Questions {@link QuestionsEntity}
     * @param Long QuizId
     * @param pageable ({@link Pageable}) obtained from the parameters and used for the pagination (e.g., size and page)
     * @return An entire {@link Page} book made of type {@link QuestionsEntity}
     * */
    Page<QuestionsEntity> findAllQuizQuestions(String quizId, Pageable pageable) {
        QuizEntity quiz = findById(quizId)
        quizQuestionsRepository.findAllByQuiz(quiz, pageable).map {
            return it.questions
        }
    }

    /**
     *
     * @param params {@link QuizParams} filter by category
     * @return Long total number of Quiz
     */
    Long countTotalQuiz(QuizParams params) {
        if (params.categoryId) {
            CategoryEntity category = findCategory(params.categoryId)
            return quizRepository.countByCategory(category)
        }
        return quizRepository.count()
    }
}
