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
import io.micronaut.http.annotation.Patch
import io.micronaut.http.annotation.Post
import io.micronaut.http.annotation.Put
import jakarta.inject.Inject
import jakarta.validation.constraints.NotBlank
import pt.atlanse.products.dtos.IngredientParams
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.services.IngredientService
import pt.atlanse.products.services.ProductService


@Slf4j
@Controller("/api/ingredients")
class IngredientController {

    @Inject
    IngredientService ingredients

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

    @Get("{?params*}")
    MutableHttpResponse findAll(IngredientParams params, Pageable pageable) {
        // Find all brands using pagination
        HttpResponse.ok(ingredients.findAllByName(params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse find(@NonNull @NotBlank UUID id) {
        // 1. Find brand by id and return with OK status
        HttpResponse.ok(ingredients.findById(id))
    }

    @Post
    MutableHttpResponse add(@NonNull @NotBlank(message = "Ingredient must be specified") String ingredient) {
        // 1. Create brand using the request's body and content created
        ingredients.create(ingredient)

        // 2. Return final response
        HttpResponse.status(HttpStatus.CREATED)
    }

    @Put("/{id}")
    MutableHttpResponse overwrite(@NonNull @NotBlank UUID id, @NonNull @NotBlank(message = "Property name is required") String name) {
        // 1. Overwrite existing brand
        ingredients.overwrite(id, name)

        // 2. Return OK response
        HttpResponse.ok()
    }

    @Patch("/{id}")
    MutableHttpResponse partialUpdate(@NonNull @NotBlank UUID id, @NonNull @NotBlank(message = "Property name is required") String name) {
        // 1. Partial update the existing brand
        ingredients.partialUpdate(id, name)

        // 2. Ok result
        HttpResponse.ok()
    }

    @Delete("/{id}")
    MutableHttpResponse delete(@NonNull @NotBlank UUID id) {
        // 1. Delete brand
        ingredients.delete(id)

        // 2. Return deleted status
        HttpResponse.status(HttpStatus.OK)
    }
}
