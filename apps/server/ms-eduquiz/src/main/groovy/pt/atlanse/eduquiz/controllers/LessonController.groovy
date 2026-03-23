package pt.atlanse.eduquiz.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.Nullable
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
import pt.atlanse.eduquiz.DTO.LessonDTO
import pt.atlanse.eduquiz.DTO.LessonParams
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.services.LessonService
import java.security.Principal
import io.micronaut.http.annotation.Error

@Slf4j
@ExecuteOn(TaskExecutors.IO)
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller("/api/lessons")
class LessonController {

    @Inject
    LessonService lessons

    LessonController() {}

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
    MutableHttpResponse getLesson(String id, @Nullable Principal principal = null) {
        return HttpResponse.ok(lessons.findById(id))
    }

    @Get("{?params*}")
    @Secured("read-eduquiz")
    MutableHttpResponse getLessons(LessonParams params, @Valid Pageable pageable, @Nullable Principal principal = null) {
        HttpResponse.ok(lessons.findAll(pageable))
    }

    @Post
    @Secured("keycloak-administrator")
    MutableHttpResponse addLesson(@Body @Valid LessonDTO lessonPayload, @Nullable Principal principal) {
        return HttpResponse.created(lessons.create(lessonPayload, principal ? principal.name : "change-me"))
    }

    @Patch("/{id}")
    @Secured("keycloak-administrator")
    MutableHttpResponse patchLesson(String id, @Body LessonDTO lessonPayload, @Nullable Principal principal) {
        return HttpResponse.ok(lessons.update(id, lessonPayload, principal ? principal.name : 'change-me'))
    }

    @Delete("/{id}")
    @Secured("keycloak-administrator")
    MutableHttpResponse deleteLesson(String id) {
        return HttpResponse.ok(lessons.delete(id))
    }

}
