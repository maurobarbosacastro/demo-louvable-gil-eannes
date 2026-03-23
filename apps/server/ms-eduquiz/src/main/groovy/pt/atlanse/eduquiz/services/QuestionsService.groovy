package pt.atlanse.eduquiz.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import jakarta.inject.Inject
import jakarta.persistence.EntityManager
import jakarta.persistence.PersistenceContext
import jakarta.persistence.TypedQuery
import jakarta.persistence.criteria.CriteriaBuilder
import jakarta.persistence.criteria.CriteriaQuery
import jakarta.persistence.criteria.Predicate
import jakarta.persistence.criteria.Root
import jakarta.transaction.Transactional
import pt.atlanse.eduquiz.DTO.AnswersDTO
import pt.atlanse.eduquiz.DTO.QuestionsDTO
import pt.atlanse.eduquiz.DTO.QuestionsParams
import pt.atlanse.eduquiz.domain.AnswersEntity
import pt.atlanse.eduquiz.domain.CategoryEntity
import pt.atlanse.eduquiz.domain.QuestionsEntity
import pt.atlanse.eduquiz.repositories.AnswersRepository
import pt.atlanse.eduquiz.repositories.CategoryRepository
import pt.atlanse.eduquiz.repositories.QuestionsRepository
import pt.atlanse.eduquiz.repositories.QuizRepository
import pt.atlanse.eduquiz.utils.ExceptionService

import java.time.LocalDateTime
import jakarta.inject.Singleton

@Slf4j
@Singleton
class QuestionsService {
    @Inject
    QuestionsRepository questionsRepository

    @Inject
    CategoryRepository categoryRepository

    @Inject
    QuizRepository quizRepository

    @Inject
    AnswersRepository answersRepository

    @Inject
    AnswersService answersService

    @PersistenceContext
    EntityManager entityManager

    @Inject
    ImagesClientService imagesClientService

    /**
     * Finds a single {@link QuestionsEntity} by Id
     * @param String id
     * @return An entire {@link QuestionsEntity}
     */
    QuestionsEntity findById(String questionId) {
        questionsRepository.findById(UUID.fromString(questionId)).orElseThrow {
            ExceptionService.QuestionNotFoundException(questionId)
        }
    }

    /**
     * Finds a single {@link QuestionsEntity} by Id
     * @param String id
     * @return An entire {@link QuestionsDTO}
     */
    QuestionsDTO find(String questionId) {
        QuestionsEntity question = findById(questionId)
        return mapToDto(question)
    }

    /**
     * Finds all {@link QuestionsDTO}
     *
     * @return List of Pages {@link QuestionsDTO}
     */
    Page<QuestionsDTO> findAll(QuestionsParams params, Pageable pageable) {
        log.info "Using pageable arguments: Page_number: ${ pageable.offset }; Amount_of_articles: ${ pageable.size }"
        return applyFilters(params, pageable)
    }

    /**
     * Find all answer of a question
     * @param String questionId
     * @return a Page<{@link AnswersEntity>
     */
    Page<AnswersEntity> findAllAnswers(String questionId, Pageable pageable) {
        QuestionsEntity question = findById(questionId)
        return answersRepository.findAllByQuestion(question, pageable)
    }

    /**
     * Create a new Question
     * @param {@link QuestionsDTO}, String author that created
     */
    QuestionsEntity create(QuestionsDTO dto, String author) {
        log.info "Creating a new question"

        // Check which fields are present
        def image = dto.content ? imagesClientService.create(dto.content) : null
        def category = dto.category ? findCategory(dto.category.id.toString()) : null
        def points = dto.points ? dto.points : null
        def type = dto.type ? dto.type : null

        List<UUID> addedAnswers = new ArrayList<>()

        // 1. Build and save a new question
        QuestionsEntity question = new QuestionsEntity(type: type,
            image: image,
            category: category,
            beginDate: dto.beginDate ? dto.beginDate : null,
            endDate: dto.endDate ? dto.endDate : null,
            extras: dto.extras ? dto.extras : null,
            description: dto.description,
            status: dto.status,
            points: points,
            createdBy: author,
            updatedBy: author)

        try {
            questionsRepository.save(question)
            // Add answers to the question
            dto.answers.each {
                addedAnswers.add(addAnswer(it, question.id, author))
            }
        } catch (Exception e) {
            // rollback the save operation
            questionsRepository.delete(question)

            addedAnswers.each {
                answersService.delete(it.toString())
            }

            throw e // rethrow the exception to propagate it up the call stack
        }

        return question
    }

    /**
     * Update a question
     * @param {@link QuestionsDTO}
     * @param String author that wants to update
     * @param long id - questionId
     */
    QuestionsEntity update(String id, QuestionsDTO dto, String author) {
        QuestionsEntity questionToUpdate = findById(id)
        List<AnswersEntity> addedAnswers = new ArrayList<>()
        questionToUpdate.image = dto.content ? imagesClientService.create(dto.content) : null

        if (dto.category) {
            questionToUpdate.category = findCategory(dto.category.id.toString())
        }

        if (dto.type) {
            questionToUpdate.type = dto.type
        }

        if (dto.beginDate) {
            questionToUpdate.beginDate = dto.beginDate
        }

        if (dto.endDate) {
            questionToUpdate.endDate = dto.endDate
        }

        if (dto.extras) {
            questionToUpdate.extras = dto.extras
        }

        if (dto.points) {
            questionToUpdate.points = dto.points
        }

        if (dto.description) {
            questionToUpdate.description = dto.description
        }

        if (dto.status) {
            questionToUpdate.status = dto.status
        }

        questionToUpdate.extras = dto.extras ? dto.extras : null

        if (dto.answers) {
            dto.answers.each {
                addedAnswers.add(updateAnswer(it, author))
            }
        }
        questionToUpdate.answers = addedAnswers
        questionToUpdate.updatedBy = author
        questionsRepository.update(questionToUpdate)
    }

    /**
     * Deleting a question by Id
     * @param String id of the question to delete
     */
    void delete(String id) {
        QuestionsEntity questionToDelete = findById(id)
        questionToDelete.answers.each {
            deleteAnswers(it.id)
        }

        questionsRepository.delete(questionToDelete)
    }

    /**
     * Deleting the answers of a question
     * @param String id of the answers to delete
     */
    void deleteAnswers(UUID id) {
        AnswersEntity answerToDelete = answersRepository.findById(id).orElseThrow {
            ExceptionService.AnswerNotFoundException(id)
        }
        answersRepository.delete(answerToDelete)
    }

    /**
     * Build the pages of type {@link QuestionsEntity}
     * @param params Other parameters. Class type {@link QuestionsParams}. Allows usage of filters
     * @param pageable ({@link Pageable}) obtained from the parameters and used for the pagination (e.g., size and page)
     * @return An entire {@link Page} book made of type {@link QuestionsEntity}
     * */
    @Transactional
    Page<QuestionsDTO> applyFilters(QuestionsParams params, Pageable pageable) {

        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<QuestionsEntity> query = cb.createQuery(QuestionsEntity.class)
        Root<QuestionsEntity> root = query.from(QuestionsEntity.class)
        CriteriaQuery<QuestionsEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (params.description ? cb.like(cb.lower(root.get("description")), "%" + params.description.toLowerCase() + "%") : null)
        predicates << (params.type ? cb.like(cb.lower(root.get("type")), "%" + params.type.toLowerCase() + "%") : null)
        predicates << (params.category ? cb.equal(root.get("category").get("id"), UUID.fromString(params.category.toLowerCase())) : null)
        predicates << (params.status ? cb.like(root.get("status"), params.status.toLowerCase()) : null)
        predicates << (params.beginDate ? cb.greaterThanOrEqualTo(root.get("beginDate"), LocalDateTime.parse(params.beginDate)) : null)
        predicates << (params.endDate ? cb.lessThanOrEqualTo(root.get("endDate"), LocalDateTime.parse(params.endDate)) : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<QuestionsEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<QuestionsEntity> questionsEntities = typedQuery.getResultList()

        // 6. Map to QuestionDTO
        List<QuestionsDTO> questionsDTOList = new ArrayList<>()
        questionsEntities.each {
            questionsDTOList.add(mapToDto(it))
        }

        return Page.of(questionsDTOList, pageable, questionsRepository.count())
    }

    // Fetching a Category by id and throwing an exception if it is not found
    private CategoryEntity findCategory(String id) {
        CategoryEntity category = categoryRepository.findById(id)
            .orElseThrow(ExceptionService::CategoryNotFoundException)
        return category
    }

    // Add Answers when we add the Question
    UUID addAnswer(AnswersDTO answer, UUID question, String author) {
        answer.questionId = question
        answersService.create(answer, author)
    }

    // Update Answers when we update the question
    AnswersEntity updateAnswer(AnswersDTO answer, String author) {
        answersService.update(answer.id.toString(), answer, author)
    }

    QuestionsDTO mapToDto(QuestionsEntity question) {
        // 5. Run query and parse results
        List<AnswersDTO> answersDTO = new ArrayList<>()

        QuestionsDTO dto = new QuestionsDTO(id: question.id,
            type: question.type,
            description: question.description,
            category: question.category,
            beginDate: question.beginDate,
            endDate: question.endDate,
            extras: question.extras,
            status: question.status,
            imageId: question.image)

        if (question.answers) {
            question.answers.each {
                answersDTO.add(answersService.mapToDto(it))
            }
            dto.answers = answersDTO.sort { a, b -> a.createdAt <=> b.createdAt }
        }
        return dto
    }
}
