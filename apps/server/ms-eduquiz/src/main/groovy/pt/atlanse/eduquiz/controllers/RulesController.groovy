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
import jakarta.inject.Inject
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Patch
import io.micronaut.http.annotation.Post
import jakarta.validation.Valid
import pt.atlanse.eduquiz.DTO.RulesDTO
import pt.atlanse.eduquiz.DTO.RulesParams
import pt.atlanse.eduquiz.domain.RulesEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.services.RulesService
import java.security.Principal
import io.micronaut.http.annotation.Error

@ExecuteOn(TaskExecutors.IO)
@Slf4j
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller("/api/rules")
class RulesController {
    @Inject
    RulesService rulesService

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([message: ex.title,
                                                    details: ex.details,
                                                    link   : request.path])
    }

    @Post
    @Secured("keycloak-administrator")
    MutableHttpResponse createRule(RulesDTO dto, @Nullable Principal principal = null) {
        log.debug("User ${principal.name} attempting to create a new rule; $dto")
        try {
            return HttpResponse.created(rulesService.create(dto, principal ? principal.name : "change-me"))
        } catch (Exception e) {
            return HttpResponse.badRequest().body({
                message: "Unable to process request"
                reason:
                e.message
            })
        }
    }

    @Get("{?params*}")
    @Secured("read-eduquiz")
    MutableHttpResponse<Page<RulesEntity>> getRules(RulesParams params, @Valid Pageable pageable) {
        return HttpResponse.ok().body(rulesService.findAll(params, pageable))
    }

    @Get("/{rulesId}")
    @Secured("read-eduquiz")
    MutableHttpResponse<RulesEntity> getRule(String rulesId) {
        try {
            return HttpResponse.ok().body(rulesService.findById(rulesId))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Patch("/{rulesId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse<RulesEntity> updateRule(String rulesId, @Body RulesDTO dto, @Nullable Principal principal = null) {
        try {
            return HttpResponse.ok().body(rulesService.update(rulesId, dto, principal ? principal.name : "change-me"))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Delete("/{rulesId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse<RulesEntity> deleteRule(String rulesId) {
        try {
            return HttpResponse.ok().body(rulesService.delete(rulesId))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }
}
