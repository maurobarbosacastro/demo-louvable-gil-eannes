package pt.atlanse.mscompany.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.model.Page
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
import io.swagger.v3.oas.annotations.tags.Tag
import jakarta.inject.Inject
import jakarta.validation.Valid
import jakarta.validation.constraints.NotBlank
import pt.atlanse.mscompany.domains.CompanyEntity
import pt.atlanse.mscompany.dtos.CompanyDTO
import pt.atlanse.mscompany.dtos.CompanyParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.services.CompanyService


@Slf4j
@Controller("/api/company")
@Tag(name = "Company")
class CompanyController {

    @Inject
    CompanyService companyService


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
    MutableHttpResponse<Page<CompanyDTO>> findAll(CompanyParams params, Pageable pageable) {
        HttpResponse.ok(companyService.findAll(params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse<CompanyDTO> find(@NonNull UUID id) {
        HttpResponse.ok(companyService.findById(id))
    }

    @Post
    MutableHttpResponse<CompanyEntity> create(@Body @Valid CompanyDTO company) {
        HttpResponse.created(companyService.create(company))
    }

    @Put("/{id}")
    MutableHttpResponse<CompanyEntity> update(@NonNull UUID id, @Body @Valid CompanyDTO company) {
        HttpResponse.ok(companyService.update(id, company))
    }

    @Patch("/{id}")
    MutableHttpResponse<CompanyEntity> patch(@NonNull UUID id, @Body @Valid CompanyDTO company) {
        HttpResponse.ok(companyService.patch(id, company))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(@NonNull UUID id) {
        companyService.delete(id)
        HttpResponse.noContent()
    }
}
