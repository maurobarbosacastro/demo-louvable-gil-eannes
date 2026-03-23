package pt.atlanse.products.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.HttpStatus
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.http.annotation.Error
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Patch
import io.micronaut.http.annotation.Post
import io.micronaut.http.annotation.Put
import jakarta.inject.Inject
import jakarta.validation.Valid
import jakarta.validation.constraints.NotBlank
import pt.atlanse.products.dtos.SellingStore
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.services.SellingStoreService


@Slf4j
@Controller("/api/selling-stores")
class SellingStoresController {

    @Inject
    SellingStoreService sellingStores

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
        // Find all sellingStores using pagination
        HttpResponse.ok(sellingStores.findAll(pageable))
    }

    @Get("/{id}")
    MutableHttpResponse find(@NonNull @NotBlank UUID id) {
        // 1. Find sellingStore by id and return with OK status
        HttpResponse.ok(sellingStores.findById(id))
    }

    @Post
    MutableHttpResponse add(@Body @Valid SellingStore sellingStore) {
        // 2. Create sellingStore using the request's body and content created
        sellingStores.create(sellingStore)

        // 3. Return final response
        HttpResponse.status(HttpStatus.CREATED)
    }

    @Put("/{id}")
    MutableHttpResponse overwrite(@NonNull @NotBlank UUID id, @Body @Valid SellingStore sellingStore) {
        // 1. Overwrite existing sellingStore
        sellingStores.overwrite(id, sellingStore)

        // 2. Return OK response
        HttpResponse.ok()
    }

    @Patch("/{id}")
    MutableHttpResponse partialUpdate(@NonNull @NotBlank UUID id, @Body @Valid SellingStore sellingStore) {
        // 1. Partial update the existing sellingStore
        sellingStores.partialUpdate(id, sellingStore)

        // 2. Ok result
        HttpResponse.ok()
    }

    @Delete("/{id}")
    MutableHttpResponse delete(@NonNull @NotBlank UUID id) {
        // 1. Delete sellingStore
        sellingStores.delete(id)

        // 2. Return deleted status
        HttpResponse.status(HttpStatus.OK)
    }
}
