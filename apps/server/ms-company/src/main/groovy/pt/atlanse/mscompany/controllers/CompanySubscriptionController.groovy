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
import pt.atlanse.mscompany.domains.CompanySubscriptionEntity
import pt.atlanse.mscompany.dtos.CompanySubscriptionDTO
import pt.atlanse.mscompany.dtos.CompanySubscriptionParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.services.CompanySubscriptionService

@Slf4j
@ExecuteOn(TaskExecutors.IO)
@Controller("/api/company/{uuid}/subscription")
@Tag(name = "Company Subscription")
class CompanySubscriptionController {

    @Inject
    CompanySubscriptionService service


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
    MutableHttpResponse<Page<CompanySubscriptionDTO>> findAll(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        CompanySubscriptionParams params,
        Pageable pageable) {
        HttpResponse.ok(service.findAll(uuid, params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse<CompanySubscriptionDTO> find(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        HttpResponse.ok(service.findById(uuid, id))
    }

    @Post
    MutableHttpResponse<CompanySubscriptionEntity> create(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @Body @Valid CompanySubscriptionDTO dto) {
        HttpResponse.created(service.create(uuid, dto))
    }

    @Put("/{id}")
    MutableHttpResponse<CompanySubscriptionEntity> update(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id,
        @Body @Valid CompanySubscriptionDTO dto) {
        HttpResponse.ok(service.update(uuid, id, dto))
    }

    @Patch("/{id}")
    MutableHttpResponse<CompanySubscriptionEntity> patch(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id,
        @Body @Valid CompanySubscriptionDTO dto) {
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
