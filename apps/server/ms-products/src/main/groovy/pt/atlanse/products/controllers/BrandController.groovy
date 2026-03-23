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
import pt.atlanse.products.dtos.Brand
import pt.atlanse.products.dtos.BrandParams
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.services.BrandService

import pt.atlanse.products.services.ImagesClientService
import pt.atlanse.products.services.ProductService



@Slf4j
@Controller("/api/brands")
class BrandController {

    @Inject
    BrandService brands

    @Inject
    ProductService products

    @Inject
    ImagesClientService imagesClientService

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
    MutableHttpResponse findAll(BrandParams params, Pageable pageable) {
        // Find all brands using pagination
        HttpResponse.ok(brands.findAll(params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse find(@NonNull @NotBlank UUID id) {
        // 1. Find brand by id and return with OK status
        HttpResponse.ok(brands.findById(id))
    }

    @Post
    MutableHttpResponse add(@Body @Valid Brand brand) {
        // 1. Create content from image object

        String brandImage = brand.image ? imagesClientService.create(brand.image) : null
        String brandBanner = brand.banner ? imagesClientService.create(brand.banner) : null

        // 2. Create brand using the request's body and content created
        brands.create(brand, brandImage, brandBanner)

        // 3. Return final response
        HttpResponse.status(HttpStatus.CREATED)
    }

    @Put("/{id}")
    MutableHttpResponse overwrite(@NonNull @NotBlank UUID id, @Body @Valid Brand brand) {
        // 1. Create content from image object
        String brandImage = brand.image ? imagesClientService.create(brand.image) : null
        String brandBanner = brand.banner ? imagesClientService.create(brand.banner) : null

        // 2. Overwrite existing brand
        brands.overwrite(id, brand, brandImage, brandBanner)

        // 3. Return OK response
        HttpResponse.ok()
    }

    @Patch("/{id}")
    MutableHttpResponse partialUpdate(@NonNull @NotBlank UUID id, @Body @Valid Brand brand) {
        // 1. Verify if the image is new
        String brandImage = brand.image ? imagesClientService.create(brand.image) : null
        String brandBanner = brand.banner ? imagesClientService.create(brand.banner) : null

        // 2. Partial update the existing brand
        brands.partialUpdate(id, brand, brandImage, brandBanner)

        // 3. Ok result
        HttpResponse.ok()
    }

    @Delete("/{id}")
    MutableHttpResponse delete(@NonNull @NotBlank UUID id) {
        // 1. Delete brand
        brands.delete(id)

        // 2. Return deleted status
        HttpResponse.status(HttpStatus.OK)
    }
}
