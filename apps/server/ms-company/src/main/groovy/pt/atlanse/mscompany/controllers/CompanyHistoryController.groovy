package pt.atlanse.mscompany.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.*
import io.micronaut.scheduling.TaskExecutors
import io.micronaut.scheduling.annotation.ExecuteOn
import io.swagger.v3.oas.annotations.Parameter
import io.swagger.v3.oas.annotations.enums.ParameterIn
import io.swagger.v3.oas.annotations.tags.Tag
import jakarta.inject.Inject
import jakarta.validation.Valid
import pt.atlanse.mscompany.domains.CompanyHistoryEntity
import pt.atlanse.mscompany.dtos.CompanyHistoryDTO
import pt.atlanse.mscompany.dtos.CompanyHistoryParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.services.CompanyHistoryService

@Slf4j
@ExecuteOn(TaskExecutors.IO)
@Controller("/api/company/{uuid}/history")
@Tag(name = "Company History")
class CompanyHistoryController {

    @Inject
    CompanyHistoryService companyHistoryService


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
    MutableHttpResponse<Page<CompanyHistoryDTO>> findAll(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        CompanyHistoryParams params,
        Pageable pageable) {
        HttpResponse.ok(companyHistoryService.findAll(uuid, params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse<CompanyHistoryDTO> find(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,

        @NonNull UUID id) {
        HttpResponse.ok(companyHistoryService.findById(uuid, id))
    }

    @Post
    MutableHttpResponse<CompanyHistoryEntity> create(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @Body @Valid CompanyHistoryDTO dto) {
        HttpResponse.created(companyHistoryService.create(uuid, dto))
    }

    @Put("/{id}")
    MutableHttpResponse<CompanyHistoryEntity> update(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid CompanyHistoryDTO dto) {
        HttpResponse.ok(companyHistoryService.update(uuid, id, dto))
    }

    @Patch("/{id}")
    MutableHttpResponse<CompanyHistoryEntity> patch(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid CompanyHistoryDTO dto) {
        HttpResponse.ok(companyHistoryService.patch(uuid, id, dto))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        companyHistoryService.delete(uuid, id)
        HttpResponse.noContent()
    }
}
