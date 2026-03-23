package pt.atlanse.email.services

import groovy.util.logging.Slf4j
import io.micronaut.context.annotation.Value
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.email.domains.TemplateEntity
import pt.atlanse.email.dto.EmailDto
import pt.atlanse.email.dto.TemplateDto
import pt.atlanse.email.dto.TemplatePatchDto
import pt.atlanse.email.exceptions.CustomException
import pt.atlanse.email.repositories.TemplateRepository
import pt.atlanse.email.utils.ExceptionService
import pt.atlanse.email.utils.HtmlSanitizationService

import java.time.LocalDateTime
import java.util.function.Supplier

@Slf4j
@Singleton
class TemplateService {

    private final TemplateRepository templateRepository
    private final HtmlSanitizationService htmlSanitizationService

    @Value('${micronaut.application.name}')
    private String appName

    @Value('${email.template-folder}')
    private String templateFolder

    String author = 'admin'

    @Inject
    TemplateService(TemplateRepository templateRepository, HtmlSanitizationService htmlSanitizationService) {
        this.templateRepository = templateRepository
        this.htmlSanitizationService = htmlSanitizationService
    }

    Page<TemplateEntity> getAll(Pageable pageable) {
        log.info "Finding all templates"
        return templateRepository.findAll(pageable)
    }

    TemplateEntity getById(String id) {

        log.info("Finding template with id: $id")

        return templateRepository.findById(UUID.fromString(id)).orElseThrow(ExceptionService::TemplateNotFoundException() as Supplier<? extends Throwable>)
    }

    TemplateEntity getByCode(String code) {
        log.info("Finding template with code: $code")

        return templateRepository.findByCode(code).orElseThrow(ExceptionService::TemplateNotFoundException() as Supplier<? extends Throwable>)
    }

    TemplateEntity createTemplate(TemplateDto payload) {
        try {
            TemplateEntity template = new TemplateEntity()
            template.name = payload.name
            template.code = payload.code
            template.templateHtml = htmlSanitizationService.sanitize(payload.templateHtml)
            template.templateJson = payload.templateJson
            template.createdBy = author
            template.updatedBy = author
            template.createdAt = LocalDateTime.now()
            template.updatedAt = LocalDateTime.now()

            return templateRepository.save(template)
        }
        catch (e) {
            throw new CustomException(
                "Error creating template",
                "Error happened while trying to create new template with name ${payload.name}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    TemplateEntity update(String id, TemplatePatchDto payload) {
        try {
            TemplateEntity template = getById(id)
            template.name = payload.name ?: template.name
            template.templateHtml = payload.templateHtml ? htmlSanitizationService.sanitize(payload.templateHtml) : template.templateHtml
            template.templateJson = payload.templateJson ?: template.templateJson
            template.updatedBy = author
            template.updatedAt = LocalDateTime.now()

            return templateRepository.update(template)
        }
        catch (e) {
            throw new CustomException(
                "Error updating template",
                "Error happened while trying to update template with name ${payload.name}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    void delete(String id) {
        try {
            TemplateEntity template = getById(id)
            templateRepository.delete(template)
        }
        catch (e) {
            throw new CustomException(
                "Error deleting template",
                "Error happened while trying to delete template with id ${id}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }


    static String parseBody(TemplateEntity template, EmailDto payload) {

        String body = template.templateHtml

        // Replace each key in the dictionary with its corresponding value
        payload.dictionary.each { key, value ->
            body = body.replaceAll("\\{\\{${key}\\}\\}", value)
        }
        return body
    }
}
