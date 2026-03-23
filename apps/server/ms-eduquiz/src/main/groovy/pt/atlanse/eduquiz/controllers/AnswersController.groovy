package pt.atlanse.eduquiz.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Patch
import io.micronaut.http.annotation.Post
import io.micronaut.scheduling.TaskExecutors
import io.micronaut.http.annotation.Error
import io.micronaut.scheduling.annotation.ExecuteOn
import io.micronaut.security.annotation.Secured
import io.micronaut.security.rules.SecurityRule
import jakarta.inject.Inject
import jakarta.validation.Valid
import pt.atlanse.eduquiz.DTO.AnswersDTO
import pt.atlanse.eduquiz.DTO.AnswersParams
import pt.atlanse.eduquiz.domain.AnswersEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.services.AnswersService

import java.security.Principal

@ExecuteOn(TaskExecutors.IO)
@Slf4j
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller("/api/answers")
class AnswersController {
    @Inject
    AnswersService answersService

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Get("/{id}")
    @Secured("read-eduquiz")
    MutableHttpResponse<AnswersEntity> getAnswer(String id) {
        HttpResponse.ok(answersService.findById(id))
    }

    @Get("{?params*}")
    @Secured("read-eduquiz")
    MutableHttpResponse<Page<AnswersEntity>> getAnswers(AnswersParams params, Pageable pageable) {
        HttpResponse.ok(answersService.findAll(params, pageable))
    }

    @Post
    @Secured("keycloak-administrator")
    MutableHttpResponse createAnswer(@Body @Valid AnswersDTO dto, @Nullable Principal principal) {
        try {
            return HttpResponse.created(answersService.create(dto, principal ? principal.name : 'change-me'))
        } catch (Exception e) {
            return HttpResponse.badRequest().body({
                message: "Unable to process request"
                reason:
                e.message
            })
        }
    }

    @Patch("/{id}")
    @Secured("keycloak-administrator")
    MutableHttpResponse<AnswersEntity> updateAnswer(String id, AnswersDTO dto, @Nullable Principal principal) {
        HttpResponse.ok(answersService.update(id, dto, principal ? principal.name : 'change-me'))
    }

    @Delete("/{id}")
    @Secured("keycloak-administrator")
    MutableHttpResponse deleteAnswer(String id) {
        HttpResponse.ok(answersService.delete(id))
    }

}
