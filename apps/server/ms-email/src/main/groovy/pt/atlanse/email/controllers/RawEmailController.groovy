package pt.atlanse.email.controllers

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Error
import io.micronaut.http.annotation.Post
import jakarta.inject.Inject
import pt.atlanse.email.dto.SendEmailDto
import pt.atlanse.email.exceptions.CustomException
import pt.atlanse.email.services.MessageService


@Slf4j
@Controller("/api/raw")
class RawEmailController {

    @Inject
    MessageService messages

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ ex.toString() }"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Post("/")
    HttpResponse sendRaw(@Body SendEmailDto payload) {
        log.info "User requested raw email send to: ${payload.to}"
        messages.sendRaw(payload)
        return HttpResponse.ok()
    }

}
