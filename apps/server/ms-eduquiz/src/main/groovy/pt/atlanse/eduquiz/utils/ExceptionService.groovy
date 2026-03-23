package pt.atlanse.eduquiz.utils

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpStatus
import pt.atlanse.eduquiz.models.CustomException

@Slf4j
class ExceptionService {
    static CustomException CategoryNotFoundException() {
        new CustomException(
            "No category was found",
            "Supplied IDs do not correspond with any of the available categories",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException ParticipantNotFoundException(Object course) {
        new CustomException(
            "The participant was not found",
            "The participant with the specified ID $course was not found",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException ModuleNotFoundException() {
        new CustomException(
            "No module was found",
            "Supplied IDs do not correspond with any of the available modules",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException LessonNotFoundException(UUID id = null) {
        new CustomException(
            "Lesson not found",
            "The lesson with id $id was not found",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException QuestionNotFoundException(String id = null) {
        new CustomException(
            "Lesson not found",
            "The question with id $id was not found",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException AnswerNotFoundException(UUID id = null) {
        new CustomException("Answer not found",
            "The question with id $id was not found",
            HttpStatus.NOT_FOUND)
    }
}
