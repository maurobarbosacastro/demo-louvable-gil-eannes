package pt.atlanse.products.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.HttpStatus
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.http.annotation.Error
import io.micronaut.http.annotation.Get
import jakarta.inject.Inject
import jakarta.validation.constraints.NotBlank
import pt.atlanse.products.models.CustomException

import pt.atlanse.products.services.ProductService
import pt.atlanse.products.services.RatingService
import pt.atlanse.products.services.ReviewService


@Slf4j
@Controller("/api/reviews")
class ReviewController {

    @Inject
    ReviewService reviews

    @Inject
    RatingService ratings

    @Inject
    ProductService products

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ ex.toString() }"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Get
    MutableHttpResponse findAll(Pageable pageable) {
        // 1. Find all reviews using pagination
        HttpResponse.ok(reviews.findAll(pageable))
    }

    @Get("/{id}")
    MutableHttpResponse find(@NonNull @NotBlank UUID id) {
        // 1. Find review  by id and return with OK status
        HttpResponse.ok(reviews.findById(id))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(@NonNull @NotBlank UUID id) {
        // 1. Delete review
        reviews.delete(id)

        // 2. Return deleted status
        HttpResponse.status(HttpStatus.OK)
    }
}
