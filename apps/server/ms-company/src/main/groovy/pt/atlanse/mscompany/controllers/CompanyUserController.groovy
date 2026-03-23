package pt.atlanse.mscompany.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.*
import io.swagger.v3.oas.annotations.Parameter
import io.swagger.v3.oas.annotations.enums.ParameterIn
import io.swagger.v3.oas.annotations.tags.Tag
import jakarta.inject.Inject
import jakarta.validation.Valid
import pt.atlanse.mscompany.domains.CompanyUserEntity
import pt.atlanse.mscompany.dtos.CompanyUserDTO
import pt.atlanse.mscompany.dtos.CompanyUserParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.services.CompanyUserService

@Slf4j
@Controller("/api/company/{uuid}/user")
@Tag(name = "Company User")
class CompanyUserController {

    @Inject
    CompanyUserService service


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
    MutableHttpResponse<Page<CompanyUserDTO>> findAll(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        CompanyUserParams params, Pageable pageable) {
        HttpResponse.ok(service.findAll(uuid, params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse<CompanyUserDTO> find(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        HttpResponse.ok(service.findById(uuid, id))
    }

    @Post
    MutableHttpResponse<CompanyUserEntity> create(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @Body @Valid CompanyUserDTO company) {
        HttpResponse.created(service.create(uuid, company))
    }

    @Put("/{id}")
    MutableHttpResponse<CompanyUserEntity> update(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid CompanyUserDTO company) {
        HttpResponse.ok(service.update(uuid, id, company))
    }

    @Patch("/{id}")
    MutableHttpResponse<CompanyUserEntity> patch(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid CompanyUserDTO company) {
        HttpResponse.ok(service.patch(uuid, id, company))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        service.delete(uuid, id)
        HttpResponse.noContent()
    }
}
