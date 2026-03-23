package pt.atlanse.products.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.HttpStatus
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.*
import io.micronaut.serde.annotation.Serdeable
import jakarta.inject.Inject
import jakarta.validation.Valid
import jakarta.validation.constraints.NotBlank
import pt.atlanse.products.dtos.Extras
import pt.atlanse.products.dtos.ExtrasParams
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.services.ExtrasService


@Slf4j
@Controller("/api/extras")
class ExtrasController {

    @Inject
    ExtrasService extrasService

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
    MutableHttpResponse findAll(ExtrasParams params, Pageable pageable) {
        // Find all brands using pagination
        HttpResponse.ok(extrasService.findAll(params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse find(@NonNull @NotBlank UUID id) {
        // 1. Find brand by id and return with OK status
        HttpResponse.ok(extrasService.findById(id))
    }

    @Get("/{id}/product")
    MutableHttpResponse getExtraWithProducts(@NonNull @NotBlank UUID id) {
        // 1. Find brand by id and return with OK status
        HttpResponse.ok(extrasService.findById(id, true))
    }

    @Post
    MutableHttpResponse add(@Body @Valid Extras extras) {
        HttpResponse.created(extrasService.create(extras))
    }

    @Put("/{id}")
    MutableHttpResponse overwrite(@NonNull @NotBlank UUID id, @Body @Valid Extras extras) {
        HttpResponse.ok( extrasService.update(id, extras))
    }

    @Patch("/{id}")
    MutableHttpResponse partialUpdate(@NonNull @NotBlank UUID id, @Body @Valid Extras extras) {
        HttpResponse.ok(extrasService.partialUpdate(id, extras))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(@NotBlank UUID id, @QueryValue(defaultValue = "false") boolean force ) {
        extrasService.delete(id, force)
        HttpResponse.noContent()
    }
}
