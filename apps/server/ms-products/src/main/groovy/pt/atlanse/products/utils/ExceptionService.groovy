package pt.atlanse.products.utils

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpStatus
import pt.atlanse.products.dtos.Product
import pt.atlanse.products.dtos.Review
import pt.atlanse.products.dtos.SellingStore
import pt.atlanse.products.models.CustomException

@Slf4j
class ExceptionService {

    static CustomException BrandNotFoundException() {
        new CustomException(
            "No brand was found",
            "Supplied IDs do not correspond with any of the available brands",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException ReviewNotFoundException() {
        new CustomException(
            "No review was found",
            "Supplied IDs do not correspond with any of the available reviews",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException CategoryNotFoundException() {
        new CustomException(
            "No category was found",
            "Supplied IDs do not correspond with any of the available categories",
            HttpStatus.NOT_FOUND
        )
    }

    static void NoStoresFoundException() {
        throw new CustomException(
            "No stores were found",
            "Supplied IDs do not correspond with any of the available stores",
            HttpStatus.NOT_FOUND
        )
    }

    static void ProductNotFound(String id) {
        throw new CustomException(
            "Product not found",
            "Product with id $id was not found",
            HttpStatus.NOT_FOUND
        )
    }

    static CustomException ProductNotFound() {
        new CustomException(
            "Product not found",
            "Product was not found",
            HttpStatus.NOT_FOUND
        )
    }

    static void RatingNotFound(String id) {
        throw new CustomException(
            "Rating not found",
            "Rating with id $id was not found",
            HttpStatus.NOT_FOUND
        )
    }

    static void IngredientNotFound(String id) {
        throw new CustomException(
            "Ingredient not found",
            "Ingredient with id $id was not found",
            HttpStatus.NOT_FOUND
        )
    }

    static void ProductCreatingException(Exception e, Product payload) {
        log.error "Unhandled exception occured: Reason: ${ e.message }"
        throw new CustomException(
            "Error creating product",
            "Error happened while trying to create new Product ${ payload.toString() }; Reason: ${ e.message }",
            HttpStatus.BAD_REQUEST
        )
    }

    static void SellingStoreCreatingException(Exception e, SellingStore payload) {
        log.error "Unhandled exception occured: Reason: ${ e.message }"
        throw new CustomException(
            "Error creating store",
            "Error happened while trying to create new store ${ payload.toString() }; Reason: ${ e.message }",
            HttpStatus.BAD_REQUEST
        )
    }

    static void RatingCreatingException(Exception e, Long payload) {
        log.error "Unhandled exception occured: Reason: ${ e.message }"
        throw new CustomException(
            "Error creating rating",
            "Error happened while trying to create new rating $payload; Reason: ${ e.message }",
            HttpStatus.BAD_REQUEST
        )
    }

    static void ReviewCreatingException(Exception e, Review payload) {
        log.error "Unhandled exception occured: Reason: ${ e.message }"
        throw new CustomException(
            "Error creating review",
            "Error happened while trying to create new review ${ payload.toString() }; Reason: ${ e.message }",
            HttpStatus.BAD_REQUEST
        )
    }

    static void IngredientCreatingException(Exception e, String name) {
        log.error "Unhandled exception occured: Reason: ${ e.message }"
        throw new CustomException(
            "Error creating ingredient",
            "Error happened while trying to create new ingredient $name; Reason: ${ e.message }",
            HttpStatus.BAD_REQUEST
        )
    }

    static CustomException ExtrasNotFoundException() {
        new CustomException(
            "No Extras was found",
            "Supplied IDs do not correspond with any of the available extras",
            HttpStatus.NOT_FOUND
        )
    }
}
