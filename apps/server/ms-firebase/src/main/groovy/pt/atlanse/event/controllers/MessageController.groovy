package pt.atlanse.event.controllers

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Post
import io.micronaut.security.annotation.Secured
import pt.atlanse.event.DTO.MessageDTO
import pt.atlanse.event.service.FirebaseService

@Slf4j
@Controller('/api/messages')
class MessageController {
    private final FirebaseService firebaseService

    MessageController(FirebaseService firebaseService) {
        this.firebaseService = firebaseService
    }

    @Post('topic/{group}')
    @Secured("keycloak-administrator")
    MutableHttpResponse sendGroupMessage(MessageDTO message, String group) {
        log.info "Arrived at /api/messages/topic"
        HttpResponse.ok(firebaseService.sendGroupMessage(group, message))
    }

    @Post('token/{group}')
    @Secured("keycloak-administrator")
    MutableHttpResponse sendIndividualMessage(MessageDTO message, String group) {
        log.info "Arrived at /api/messages/token"
        HttpResponse.ok(firebaseService.sendIndividualMessage(group, message))
    }

}






