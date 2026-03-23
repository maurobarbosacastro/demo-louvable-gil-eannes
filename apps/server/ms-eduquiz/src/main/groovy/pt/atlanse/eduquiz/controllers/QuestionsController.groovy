package pt.atlanse.eduquiz.controllers

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.scheduling.TaskExecutors
import io.micronaut.scheduling.annotation.ExecuteOn
import io.micronaut.security.annotation.Secured
import io.micronaut.security.rules.SecurityRule
import jakarta.inject.Inject
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Patch
import io.micronaut.http.annotation.Post
import jakarta.validation.Valid
import pt.atlanse.eduquiz.DTO.QuestionsDTO
import pt.atlanse.eduquiz.DTO.QuestionsParams
import pt.atlanse.eduquiz.domain.AnswersEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.services.QuestionsService
import java.security.Principal
import io.micronaut.http.annotation.Error

@ExecuteOn(TaskExecutors.IO)
@Slf4j
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller("/api/questions")
class QuestionsController {
    @Inject
    QuestionsService questionsService

    @Error(exception = CustomException.class)
    static MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Get("/{questionId}")
    @Secured("read-eduquiz")
    MutableHttpResponse<QuestionsDTO> getQuestion(String questionId) {
        try {
            return HttpResponse.ok().body(questionsService.find(questionId))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Get("{?params*}")
    @Secured("read-eduquiz")
    MutableHttpResponse<Page<QuestionsDTO>> getQuestions(QuestionsParams params, @Valid Pageable pageable) {
        try {
            return HttpResponse.ok(questionsService.findAll(params, pageable))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Get("/{questionId}/answers")
    @Secured("read-eduquiz")
    MutableHttpResponse<Page<AnswersEntity>> getQuestionAnswers(String questionId, @Valid Pageable pageable) {
        try {
            return HttpResponse.ok(questionsService.findAllAnswers(questionId, pageable))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Post
    @Secured("keycloak-administrator")
    MutableHttpResponse createQuestion(@Body @Valid QuestionsDTO dto, Principal principal) {
        try {
            return HttpResponse.created(questionsService.create(dto, principal.name))
        } catch (Exception e) {
            return HttpResponse.badRequest().body({
                message: "Unable to process request"
                reason:
                e.message
            })
        }
    }

    @Patch("/{questionId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse updateQuestion(String questionId, @Body QuestionsDTO dto, Principal principal) {
        try {
            return HttpResponse.ok(questionsService.update(questionId, dto, principal.name))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Delete("/{questionId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse deleteQuestion(String questionId) {
        try {
            return HttpResponse.ok(questionsService.delete(questionId))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

}
