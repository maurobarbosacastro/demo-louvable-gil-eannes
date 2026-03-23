package pt.atlanse.eduquiz.controllers

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.security.annotation.Secured
import io.micronaut.security.rules.SecurityRule
import jakarta.inject.Inject
import pt.atlanse.eduquiz.domain.ModulesEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.services.ModulesOrderService
import io.micronaut.http.annotation.Error

@Slf4j
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller("/api/modules-order")
class ModuleOrderController {
    @Inject
    ModulesOrderService modulesOrderService

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Delete("/{moduleOrderId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse<ModulesEntity> deleteModuleLesson(String moduleOrderId) {
        HttpResponse.ok(modulesOrderService.delete(moduleOrderId))
    }
}
