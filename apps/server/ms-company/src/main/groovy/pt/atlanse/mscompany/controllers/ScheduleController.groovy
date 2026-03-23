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
import pt.atlanse.mscompany.domains.ScheduleEntity
import pt.atlanse.mscompany.dtos.ScheduleDTO
import pt.atlanse.mscompany.dtos.ScheduleParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.services.ScheduleService

@Slf4j
@ExecuteOn(TaskExecutors.IO)
@Controller("/api/company/{uuid}/schedule")
@Tag(name = "Company Schedule")
class ScheduleController {

    @Inject
    ScheduleService scheduleService


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
    MutableHttpResponse<Page<ScheduleDTO>> findAll(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        ScheduleParams params, Pageable pageable) {
        HttpResponse.ok(scheduleService.findAll(uuid, params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse<ScheduleDTO> find(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        HttpResponse.ok(scheduleService.findById(uuid, id))
    }

    @Post
    MutableHttpResponse<ScheduleEntity> create(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @Body @Valid ScheduleDTO dto) {
        HttpResponse.created(scheduleService.create(uuid, dto))
    }

    @Put("/{id}")
    MutableHttpResponse<ScheduleEntity> update(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid ScheduleDTO dto) {
        HttpResponse.ok(scheduleService.update(uuid, id, dto))
    }

    @Patch("/{id}")
    MutableHttpResponse<ScheduleEntity> patch(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid ScheduleDTO dto) {
        HttpResponse.ok(scheduleService.patch(uuid, id, dto))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        scheduleService.delete(uuid, id)
        HttpResponse.noContent()
    }
}
