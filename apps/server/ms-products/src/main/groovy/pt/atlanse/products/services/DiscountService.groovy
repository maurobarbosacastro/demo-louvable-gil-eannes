package pt.atlanse.products.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpStatus
import jakarta.inject.Singleton
import pt.atlanse.products.domains.DiscountEntity
import pt.atlanse.products.dtos.Discount
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.DiscountRepository

@Slf4j
@Singleton
class DiscountService {

    DiscountRepository discounts

    DiscountService(DiscountRepository discounts) {
        log.debug "Injecting Discount repository"
        this.discounts = discounts
    }

    /**
     * Retrieves the discount entity with the specified ID.
     *
     * @param id The ID of the discount entity to retrieve
     * @return The discount entity with the specified ID
     * @throws CustomException if the discount entity with the specified ID does not exist
     */
    private DiscountEntity getById(UUID id) throws CustomException {
        log.info("Getting discount with id: ${ id }")
        Optional<DiscountEntity> discount = discounts.findById(id)

        if (discount.isEmpty()) {
            throw new CustomException(
                "Discount not found",
                "Discount with id $id was not found",
                HttpStatus.NOT_FOUND
            )
        }

        discount.get()
    }

    /**
     * Retrieves a page of discount entities based on the provided pagination information.
     *
     * @param pageable The pagination information to use when retrieving the discount entities
     * @return A page of discount entities based on the provided pagination information
     */
    Page<DiscountEntity> findAll(Pageable pageable) {
        log.info("Getting all discounts")
        discounts.findAll(pageable)
    }

    /**
     * Retrieves the discount entity with the specified ID.
     *
     * @param id The ID of the discount entity to retrieve
     * @return The discount entity with the specified ID, or null if the entity does not exist
     */
    DiscountEntity findById(UUID id) {
        getById(id)
    }

    /**
     * Creates a new discount entity based on the provided payload and author.
     *
     * @param payload The payload to use for creating the new discount entity
     * @param author The author of the discount entity (default: "anonymous")
     */
    //todo remove anonymous author
    void create(Discount payload, String author = "anonymous") {
        log.info "Creating new discount"

        try {
            // 1. Create discount entity
            DiscountEntity discount = new DiscountEntity(
                code: payload.code,
                active: payload.active,
                amount: payload.amount,
                createdBy: author,
                updatedBy: author
            )

            // 2. Save discount using repository
            discounts.save(discount)
        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${ e.message }"
            throw new CustomException(
                "Error creating discount",
                "Error happened while trying to create new Discount ${ payload.toString() }; Reason: ${ e.message }",
                HttpStatus.BAD_REQUEST
            )
        }
    }

    /**
     * Overwrites an existing discount entity with the specified ID with the provided payload and author.
     *
     * @param id The ID of the discount entity to overwrite
     * @param payload The payload to use for overwriting the discount entity
     * @param author The author of the discount entity (default: "anonymous")
     * @throws CustomException if the discount entity with the specified ID does not exist
     */
    void overwrite(UUID id, Discount payload, String author = "anonymous") {
        // 1. Fetch discount
        Optional<DiscountEntity> discount = discounts.findById(id)

        // 2. If the discount does not exist, create a new one
        if (discount.isEmpty()) {
            create(payload)
            return
        }

        // 3. Try to overwrite existing discount
        try {
            discount.get().with {
                it.code = payload.code
                it.active = payload.active
                it.amount = payload.amount
                it.updatedBy = author

                // 4. Update existing discount
                discounts.update(it)
            }
        } catch (Exception e) {
            // 5. Handle exceptions found
            log.error "Unhandled exception found while attempting to update discount $id; Reason ${ e.message }"
        }
    }

    /**
     * Partially updates an existing discount entity with the specified ID with the provided payload and author.
     *
     * @param id The ID of the discount entity to update
     * @param payload The payload to use for partially updating the discount entity
     * @param author The author of the discount entity (default: "anonymous")
     * @throws CustomException if the discount entity with the specified ID does not exist
     */
    void partialUpdate(UUID id, Discount payload, String author = "anonymous") {
        log.info("Updating discount with id: ${ id }")
        // 1. Fetch discount
        DiscountEntity discount = getById(id)

        // 3. Update existing discount fields
        discount.with {
            it.code = payload.code
            it.active = payload.active
            it.amount = payload.amount
            it.updatedBy = author

            // 4. Update existing discount
            discounts.update(it)
        }

    }

    /**
     * Deletes the discount entity with the specified ID.
     *
     * @param id The ID of the discount entity to delete
     * @throws CustomException if the discount entity with the specified ID does not exist
     */
    void delete(UUID id) {
        log.info("Deleting discount with id: ${ id }")
        // 1. Fetch discount
        DiscountEntity discount = getById(id)

        // 2. Delete discount
        discounts.delete(discount)
    }

}
