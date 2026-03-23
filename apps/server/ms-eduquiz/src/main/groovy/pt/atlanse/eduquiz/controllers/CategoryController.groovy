package pt.atlanse.eduquiz.controllers

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Get
import io.micronaut.security.annotation.Secured
import io.micronaut.security.rules.SecurityRule
import jakarta.inject.Inject
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.services.CategoryService
import io.micronaut.http.annotation.Error

@Slf4j
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller("/api/categories")
class CategoryController {

    @Inject
    CategoryService categoryService

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Get("{?params*}")
    @Secured("read-eduquiz")
    MutableHttpResponse getCategories(Map params, Pageable pageable) {
        return HttpResponse.ok(categoryService.findAll(params, pageable))
    }
}
