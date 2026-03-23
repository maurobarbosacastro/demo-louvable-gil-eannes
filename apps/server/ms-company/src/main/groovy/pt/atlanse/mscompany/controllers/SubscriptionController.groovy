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
import io.swagger.v3.oas.annotations.tags.Tag
import jakarta.inject.Inject
import jakarta.validation.Valid
import pt.atlanse.mscompany.domains.SubscriptionEntity
import pt.atlanse.mscompany.dtos.SocialDTO
import pt.atlanse.mscompany.dtos.SocialParams
import pt.atlanse.mscompany.dtos.SubscriptionDTO
import pt.atlanse.mscompany.dtos.SubscriptionParams
import pt.atlanse.mscompany.models.CustomException
import pt.atlanse.mscompany.services.SocialsService
import pt.atlanse.mscompany.services.SubscriptionService

@Slf4j
@ExecuteOn(TaskExecutors.IO)
@Controller("/api/subscription")
@Tag(name = "Subscription")
class SubscriptionController {

    @Inject
    SubscriptionService service


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
    MutableHttpResponse<Page<SubscriptionDTO>> findAll(SubscriptionParams params, Pageable pageable) {
        HttpResponse.ok(service.findAll(params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse<SubscriptionDTO> find(@NonNull UUID id) {
        HttpResponse.ok(service.findById(id))
    }

    @Post
    MutableHttpResponse<SubscriptionEntity> create(@Body @Valid SubscriptionDTO dto) {
        HttpResponse.created(service.create(dto))
    }

    @Put("/{id}")
    MutableHttpResponse<SubscriptionEntity> update(@NonNull UUID id, @Body @Valid SubscriptionDTO dto) {
        HttpResponse.ok(service.update(id, dto))
    }

    @Patch("/{id}")
    MutableHttpResponse<SubscriptionEntity> patch(@NonNull UUID id, @Body @Valid SubscriptionDTO dto) {
        HttpResponse.ok(service.patch(id, dto))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(@NonNull UUID id) {
        service.delete(id)
        HttpResponse.noContent()
    }
}
