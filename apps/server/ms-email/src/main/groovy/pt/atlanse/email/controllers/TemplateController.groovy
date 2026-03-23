package pt.atlanse.email.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.*
import jakarta.inject.Inject
import jakarta.validation.Valid
import jakarta.validation.constraints.NotBlank
import pt.atlanse.email.dto.TemplateDto
import pt.atlanse.email.dto.TemplatePatchDto
import pt.atlanse.email.exceptions.CustomException
import pt.atlanse.email.services.TemplateService

@Slf4j
@Controller("/api/templates")
class TemplateController {

    @Inject
    TemplateService templateService

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Get
    HttpResponse getAll(Pageable pageable) {
        log.info "User requests all templates"
        return HttpResponse.ok(templateService.getAll(pageable))
    }

    @Get("/{id}")
    HttpResponse getById(@NonNull @NotBlank String id) {
        log.info "User requests template with id: $id"
        return HttpResponse.ok(templateService.getById(id))
    }

    @Post
    HttpResponse create(@Body @Valid TemplateDto template) {
        log.info "User requested to create a new template: $template"
        return HttpResponse.ok(templateService.createTemplate(template))
    }

    @Patch("/{id}")
    HttpResponse update(@NonNull @NotBlank String id, @Body @Valid TemplatePatchDto templateDto) {
        log.info "User requests to update template with id: $id"
        return HttpResponse.ok(templateService.update(id, templateDto))
    }

    @Delete("/{id}")
    HttpResponse delete(@NonNull @NotBlank String id) {
        log.info "User requests to delete template with id: $id"
        return HttpResponse.ok(templateService.delete(id))
    }
}
