package pt.atlanse.eduquiz.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.scheduling.TaskExecutors
import io.micronaut.scheduling.annotation.ExecuteOn
import io.micronaut.security.annotation.Secured
import io.micronaut.security.rules.SecurityRule
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Patch
import io.micronaut.http.annotation.Post
import jakarta.inject.Inject
import jakarta.validation.Valid
import pt.atlanse.eduquiz.DTO.QuizDTO
import pt.atlanse.eduquiz.DTO.QuizParams
import pt.atlanse.eduquiz.domain.QuestionsEntity
import pt.atlanse.eduquiz.domain.QuizEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.services.QuizService
import java.security.Principal
import io.micronaut.http.annotation.Error

@ExecuteOn(TaskExecutors.IO)
@Slf4j
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller("/api/quizzes")
class QuizController {
    @Inject
    QuizService quizService

    QuizController() {}

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Get("/{quizId}")
    @Secured("read-eduquiz")
    MutableHttpResponse<QuizEntity> getQuiz(String quizId) {
        HttpResponse.ok(quizService.findById(quizId))
    }

    @Get("{?params*}")
    @Secured("read-eduquiz")
    MutableHttpResponse<Page<QuizEntity>> getQuizzes(QuizParams params, @Valid Pageable pageable) {
        Page<QuizEntity> quizzes = quizService.findAll(params, pageable)
        return HttpResponse.ok(quizzes)
    }

    @Get("/{quizId}/questions")
    @Secured("read-eduquiz")
    MutableHttpResponse<Page<QuestionsEntity>> getQuizQuestions(String quizId, @Valid Pageable pageable) {
        HttpResponse.ok(quizService.findAllQuizQuestions(quizId, pageable))
    }

    @Get("/count{?params*}")
    @Secured("read-eduquiz")
    MutableHttpResponse<Long> countQuizzes(QuizParams params) {
        try {
            return HttpResponse.ok(quizService.countTotalQuiz(params))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Post
    @Secured("keycloak-administrator")
    MutableHttpResponse createQuiz(@Body @Valid QuizDTO dto, @Nullable Principal principal) {
        log.debug("User ${principal.name} attempting to create a new Quiz; $dto")
        try {
            return HttpResponse.created(quizService.create(dto, principal ? principal.name : 'change-me'))
        } catch (Exception e) {
            return HttpResponse.badRequest().body({
                message: "Unable to process request"
                reason:
                e.message
            })
        }
    }

    @Post("/{quizId}/question/{questionId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse addQuestionToQuiz(String quizId, String questionId, @Nullable Principal principal) {
        log.debug("User ${principal.name} attempting to add a question $questionId to the quiz $quizId")
        try {
            return HttpResponse.created(quizService.addQuestion(quizId, questionId, principal ? principal.name : 'change-me'))
        } catch (Exception e) {
            return HttpResponse.badRequest().body({
                message: "Unable to process request"
                reason:
                e.message
            })
        }
    }

    @Patch("/{quizId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse updateQuiz(String quizId, QuizDTO dto, @Nullable Principal principal) {
        log.debug("User ${principal.name} attempting to update the quiz $quizId")
        try {
            return HttpResponse.ok(quizService.update(quizId, dto, principal ? principal.name : 'change-me'))
        } catch (Exception e) {
            return HttpResponse.badRequest().body({
                message: "Unable to process request"
                reason:
                e.message
            })
        }
    }

    @Delete("/{quizId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse deleteQuiz(String quizId) {
        try {
            return HttpResponse.ok().body(quizService.delete(quizId))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Delete("/quiz-question/{id}")
    @Secured("keycloak-administrator")
    MutableHttpResponse deleteQuizQuestion(String id) {
        try {
            return HttpResponse.ok().body(quizService.deleteQuizQuestion(id))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }
}
