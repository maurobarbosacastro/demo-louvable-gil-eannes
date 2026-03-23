package pt.atlanse.email.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Error
import io.micronaut.http.annotation.Post
import jakarta.inject.Inject
import jakarta.validation.Valid
import jakarta.validation.constraints.NotBlank
import pt.atlanse.email.domains.TemplateEntity
import pt.atlanse.email.services.TemplateService
import pt.atlanse.email.dto.EmailDto
import pt.atlanse.email.exceptions.CustomException
import pt.atlanse.email.services.MessageService


@Slf4j
@Controller("/api/emails")
class EmailController {

    @Inject
    MessageService messages

    @Inject
    TemplateService templates

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ ex.toString() }"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Post("/{template}")
    HttpResponse send(@Body EmailDto payload, @NonNull String template) {
        log.info "User requested email using template: $template"

        TemplateEntity templateEntity = templates.getByCode(template)
        if (!templateEntity) {
            return HttpResponse.notFound()
        }

        //parse template and title
        String body = templates.parseBody(templateEntity, payload)

        messages.send(templateEntity.name, body,  payload, true)

        return HttpResponse.ok()
    }

}
