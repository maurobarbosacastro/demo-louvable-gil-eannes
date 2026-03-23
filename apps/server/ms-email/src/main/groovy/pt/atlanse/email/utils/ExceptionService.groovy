package pt.atlanse.email.utils

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpStatus
import pt.atlanse.email.exceptions.CustomException

@Slf4j
class ExceptionService {
    static CustomException TemplateNotFoundException() {
        new CustomException(
            "Template not found",
            "Supplied IDs do not correspond with any of the available templates",
            HttpStatus.NOT_FOUND
        )
    }
}
