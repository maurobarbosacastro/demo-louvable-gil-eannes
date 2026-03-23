package pt.atlanse.eduquiz.models

import groovy.transform.ToString
import io.micronaut.http.HttpStatus

@ToString(includePackage = false, includeNames = true, includeFields = true)
class CustomException extends Exception {
    String title
    String details
    HttpStatus status

    CustomException() {}

    CustomException(String title, String details, HttpStatus status) {
        this.title = title
        this.details = details
        this.status = status
    }
}
