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
import pt.atlanse.mscompany.domains.SocialsEntity
import pt.atlanse.mscompany.dtos.SocialDTO
import pt.atlanse.mscompany.dtos.SocialParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.services.SocialsService

@Slf4j
@ExecuteOn(TaskExecutors.IO)
@Controller("/api/company/{uuid}/socials")
@Tag(name = "Company Socials")
class SocialsController {

    @Inject
    SocialsService socialsService


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
    MutableHttpResponse<Page<SocialDTO>> findAll(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        SocialParams params, Pageable pageable) {
        HttpResponse.ok(socialsService.findAll(uuid, params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse<SocialDTO> find(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        HttpResponse.ok(socialsService.findById(uuid, id))
    }

    @Post
    MutableHttpResponse<SocialsEntity> create(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @Body @Valid SocialDTO socials) {
        HttpResponse.created(socialsService.create(uuid, socials))
    }

    @Put("/{id}")
    MutableHttpResponse<SocialsEntity> update(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid SocialDTO socials) {
        HttpResponse.ok(socialsService.update(uuid, id, socials))
    }

    @Patch("/{id}")
    MutableHttpResponse<SocialsEntity> patch(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id, @Body @Valid SocialDTO socials) {
        HttpResponse.ok(socialsService.patch(uuid, id, socials))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(
        @Parameter(description = "UUID of the company", required = true, in = ParameterIn.PATH) UUID uuid,
        @NonNull UUID id) {
        socialsService.delete(uuid, id)
        HttpResponse.noContent()
    }
}
