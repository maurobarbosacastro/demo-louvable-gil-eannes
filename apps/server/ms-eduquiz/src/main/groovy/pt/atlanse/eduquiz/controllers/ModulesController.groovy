package pt.atlanse.eduquiz.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.HttpStatus
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
import pt.atlanse.eduquiz.DTO.ModulesDTO
import pt.atlanse.eduquiz.DTO.ModulesParams
import pt.atlanse.eduquiz.domain.LessonEntity
import pt.atlanse.eduquiz.domain.ModulesEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.models.Lesson
import pt.atlanse.eduquiz.services.ElearningBeans
import pt.atlanse.eduquiz.services.LessonService
import pt.atlanse.eduquiz.services.ModulesService
import java.security.Principal
import io.micronaut.http.annotation.Error

@Slf4j
@ExecuteOn(TaskExecutors.IO)
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller("/api/modules")
class ModulesController {

    @Inject
    ElearningBeans elearning
    @Inject
    ModulesService modulesService
    @Inject
    LessonService lessonService

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    // todo move to service
    @Post("/{moduleId}/lessons/{lessonId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse addLessonToModule(String lessonId, String moduleId, @Nullable Principal principal) {
        log.debug("User ${principal.name} attempting to add a lesson $lessonId to a module $moduleId")
        try {
            LessonEntity lesson = elearning.lessons.findById(UUID.fromString(lessonId)).orElseThrow {
                new CustomException(
                    "Lesson not found",
                    "The lesson with id $lessonId was not found",
                    HttpStatus.NOT_FOUND
                )
            }
            return HttpResponse.created(modulesService.addLessonToModule(moduleId, lesson, principal ? principal.name : "change-me"))
        } catch (Exception e) {
            return HttpResponse.badRequest().body({
                message: "Unable to process request"
                reason:
                e.message
            })
        }
    }

    @Post
    @Secured("keycloak-administrator")
    MutableHttpResponse createModule(@Body @Valid ModulesDTO module, @Nullable Principal principal) {
        log.debug("User ${principal?.name} attempting to create a new module; $module")
        try {
            return HttpResponse.created(modulesService.create(module, principal ? principal.name : "change-me"))
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
    MutableHttpResponse<Page<ModulesEntity>> findAll(ModulesParams params, Pageable pageable) {
        return HttpResponse.ok(modulesService.findAll(params, pageable))
    }

    @Get("/{moduleId}")
    @Secured("read-eduquiz")
    MutableHttpResponse<ModulesEntity> getModule(String moduleId) {
        return HttpResponse.ok(modulesService.findById(moduleId))
    }

    @Get("/{moduleId}/lessons")
    @Secured("read-eduquiz")
    MutableHttpResponse<Page<Lesson>> getModuleLessons(String moduleId, Pageable pageable) {
        try {
            return HttpResponse.ok().body(lessonService.findAllByModule(moduleId, pageable))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Patch("/{moduleId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse<ModulesEntity> updateModule(String moduleId, @Body ModulesDTO dto, @Nullable Principal principal) {
        try {
            return HttpResponse.ok(modulesService.update(moduleId, dto, principal ? principal.name : 'change-me'))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Delete("/{moduleId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse<ModulesEntity> deleteModule(String moduleId) {
        try {
            return HttpResponse.ok().body(modulesService.delete(moduleId))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

}
