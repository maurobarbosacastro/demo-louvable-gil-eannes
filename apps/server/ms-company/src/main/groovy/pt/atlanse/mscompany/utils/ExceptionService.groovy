package pt.atlanse.mscompany.utils

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpStatus
import pt.atlanse.mscompany.dtos.Media
import pt.atlanse.mscompany.models.CustomException

@Slf4j
class ExceptionService {

    static CustomException CompanyNotFoundException() {
        new CustomException(
            "No Company was found",
            "Supplied IDs do not correspond with any of the available companies",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException SocialNotFoundException() {
        new CustomException(
            "No Social was found",
            "Supplied IDs do not correspond with any of the available socials ",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException SubscriptionNotFoundException() {
        new CustomException(
            "No Subscription was found",
            "Supplied IDs do not correspond with any of the available Subscriptions ",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException CompanySubscriptionNotFoundException() {
        new CustomException(
            "No Company Subscription was found",
            "Supplied IDs do not correspond with any of the available Company Subscriptions ",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException CompanyHistoryNotFoundException() {
        new CustomException(
            "No Company History was found",
            "Supplied IDs do not correspond with any of the available company historic  ",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException CompanyUserNotFoundException() {
        new CustomException(
            "No Company user was found",
            "Supplied IDs do not correspond with any of the available company user  ",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException CompanyMailInfoNotFoundException() {
        new CustomException(
            "No Company Mail Info was found",
            "Supplied IDs do not correspond with any of the available Company Mail Info",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException ScheduleNotFoundException() {
        new CustomException(
            "No Schedule was found",
            "Supplied IDs do not correspond with any of the available schedules ",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException ImageSavingException(Media payload) {
        new CustomException(
            "Error saving image",
            "Error saving image on ms-images; Extension: ${payload.extension}",
            HttpStatus.INTERNAL_SERVER_ERROR
        )
    }
}
