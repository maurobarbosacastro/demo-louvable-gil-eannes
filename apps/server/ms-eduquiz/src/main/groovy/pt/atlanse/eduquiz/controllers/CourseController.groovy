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
import io.micronaut.scheduling.annotation.ExecuteOn
import io.micronaut.security.annotation.Secured
import io.micronaut.security.rules.SecurityRule
import jakarta.inject.Inject
import jakarta.validation.Valid
import pt.atlanse.eduquiz.DTO.CourseDTO
import pt.atlanse.eduquiz.DTO.CourseParams
import pt.atlanse.eduquiz.domain.CourseEntity
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.services.CourseService
import pt.atlanse.eduquiz.services.ModulesService

import java.security.Principal
import io.micronaut.http.annotation.Error

@ExecuteOn(TaskExecutors.IO)
@Slf4j
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller("/api/courses")
class CourseController {

    @Inject
    CourseService courseService
    @Inject
    ModulesService modulesService

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Get("/{courseId}")
    HttpResponse<CourseDTO> getCourse(String courseId) {
        return HttpResponse.ok(courseService.getCourse(courseId))
    }

    @Get("/{?params*}")
    @Secured("read-eduquiz")
    HttpResponse<Page<CourseEntity>> getCourses(CourseParams params, Pageable pageable) {
        Page<CourseEntity> courses = courseService.getCourses(params, pageable)
        return HttpResponse.ok(courses)
    }

    @Get("/active")
    @Secured("read-eduquiz")
    HttpResponse<Map<String, Object>> getActiveCourse() {
        HttpResponse.ok(courseService.getActiveCourse())
    }

    @Post
    @Secured("keycloak-administrator")
    MutableHttpResponse createCourses(@Body @Valid CourseDTO course, @Nullable Principal principal) {
        return HttpResponse.created(courseService.createCourse(course, principal ? principal.name : "change-me"))
    }

    @Post('/{courseId}/modules')
    @Secured("keycloak-administrator")
    MutableHttpResponse addModules(String courseId, @Body @Valid String module, @Nullable Principal principal) {
        return HttpResponse.ok(courseService.addModule(courseId, module, principal ? principal.name : "change-me"))
    }

    @Post('/{courseId}/modules/{moduleId}')
    @Secured("keycloak-administrator")
    MutableHttpResponse addExistentModule(String courseId, String moduleId, @Nullable Principal principal) {
        return HttpResponse.ok(courseService.addExistentModule(courseId, moduleId, principal ? principal.name : "change-me"))
    }

    @Patch("/{courseId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse patchContent(String courseId, @Body @Valid CourseDTO course, @Nullable Principal principal) {
        return HttpResponse.ok(courseService.editCourse(courseId, course, principal ? principal.name : "change-me"))
    }

    @Delete("/{courseId}")
    @Secured("keycloak-administrator")
    MutableHttpResponse deleteCourse(String courseId) {
        return HttpResponse.ok(courseService.deleteCourse(courseId))
    }

}


