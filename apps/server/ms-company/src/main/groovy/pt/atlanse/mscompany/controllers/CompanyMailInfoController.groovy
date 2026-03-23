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
import pt.atlanse.mscompany.domains.CompanyMailInfoEntity
import pt.atlanse.mscompany.dtos.CompanyMailInfoDTO
import pt.atlanse.mscompany.dtos.CompanyMailInfoParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.services.CompanyMailInfoService

@Slf4j
@Controller("/api/company/{uuid}/mail")
@Tag(name = "Company Mail")
class CompanyMailInfoController {

    @Inject
    CompanyMailInfoService service


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
    MutableHttpResponse<Page<CompanyMailInfoDTO>> findAll(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        CompanyMailInfoParams params, Pageable pageable) {
        HttpResponse.ok(service.findAll(uuid, params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse<CompanyMailInfoDTO> find(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        HttpResponse.ok(service.findById(uuid, id))
    }

    @Post
    MutableHttpResponse<CompanyMailInfoEntity> create(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @Body @Valid CompanyMailInfoDTO dto) {
        HttpResponse.created(service.create(uuid, dto))
    }

    @Put("/{id}")
    MutableHttpResponse<CompanyMailInfoEntity> update(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid CompanyMailInfoDTO dto) {
        HttpResponse.ok(service.update(uuid, id, dto))
    }

    @Patch("/{id}")
    MutableHttpResponse<CompanyMailInfoEntity> patch(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid CompanyMailInfoDTO dto) {
        HttpResponse.ok(service.patch(uuid, id, dto))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        service.delete(uuid, id)
        HttpResponse.noContent()
    }
}
