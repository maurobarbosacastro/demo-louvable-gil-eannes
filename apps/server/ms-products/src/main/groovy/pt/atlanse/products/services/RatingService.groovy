package pt.atlanse.products.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import jakarta.inject.Singleton
import pt.atlanse.products.domains.RatingEntity
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.RatingRepository
import pt.atlanse.products.utils.ExceptionService

/**
 * @deprecated
 * */
@Slf4j
@Singleton
class RatingService {

    RatingRepository ratings

    RatingService(RatingRepository ratings) {
        log.debug "Injecting ratings repository"
        this.ratings = ratings
    }

    /**
     * Returns a Page of all RatingEntity objects with pagination based on the specified Pageable object.
     *
     * @param pageable The Pageable object that defines the pagination parameters
     * @return A Page of RatingEntity objects
     */
    private RatingEntity getById(UUID id) throws CustomException {
        log.info("Getting rating with id: ${ id }")
        Optional<RatingEntity> ratings = ratings.findById(id)
        return ratings.isPresent() ? ratings.get() : ExceptionService.RatingNotFound(id as String)
    }

    /**
     * Returns the RatingEntity object with the specified ID.
     *
     * @param id The ID of the RatingEntity to retrieve
     * @return The RatingEntity object with the specified ID
     * @throws CustomException if a RatingEntity with the specified ID does not exist
     */
    Page<RatingEntity> findAll(Pageable pageable) {
        log.info("Getting all ratings")
        ratings.findAll(pageable)
    }

    /**
     * Creates a new RatingEntity object with the specified value and author.
     *
     * @param value The value of the rating (1-5)
     * @param author The username of the author creating the new RatingEntity (default: "anonymous")
     */
    RatingEntity findById(UUID id) {
        getById(id)
    }

    /**
     * Overwrites an existing RatingEntity with the specified ID with a new RatingEntity with the specified value and author.
     *
     * @param id The ID of the RatingEntity to overwrite
     * @param value The new value for the RatingEntity
     * @param author The username of the author overwriting the RatingEntity (default: "anonymous")
     * @throws CustomException if a RatingEntity with the specified ID does not exist
     */
    //todo remove anonymous author
    RatingEntity create(Long value, String author = "anonymous") {
        log.info("Creating new ratings")

        try {
            // 1. Create ratings entity
            RatingEntity rating = new RatingEntity(
                amount: value,
                createdBy: author,
                updatedBy: author
            )

            // 2. Save ratings using repository
            ratings.save(rating)
        } catch (Exception e) {
            ExceptionService.RatingCreatingException(e, value)
        }
    }

    /**
     * Partially updates an existing RatingEntity with the specified amount and author.
     *
     * @param id The ID of the RatingEntity to update
     * @param amount The amount to update the RatingEntity with
     * @param author The username of the author updating the RatingEntity (default: "anonymous")
     * @throws CustomException if a RatingEntity with the specified ID does not exist
     */
    void overwrite(UUID id, long value, String author = "anonymous") {
        // 1. Fetch ratings
        Optional<RatingEntity> rating = ratings.findById(id)

        // 2. If the ratings does not exist, create a new one
        if (rating.isEmpty()) {
            create(value)
            return
        }

        // 3. Try to overwrite existing ratings
        try {
            rating.get().with {
                it.amount
                it.updatedBy = author

                // 4. Update existing ratings
                ratings.update(it)
            }
        } catch (Exception e) {
            // 5. Handle exceptions found
            log.error "Unhandled exception found while attempting to update ratings $id; Reason ${ e.message }"
        }
    }

    /**
     * Deletes the RatingEntity with the specified ID.
     *
     * @param id The ID of the RatingEntity to delete
     * @throws CustomException if a RatingEntity with the specified ID does not exist
     */
    void partialUpdate(UUID id, Long amount, String author = "anonymous") {
        log.info("Updating rating with id: ${ id }")
        // 1. Fetch ratings
        RatingEntity rating = getById(id)

        // 3. Update existing ratings fields
        rating.with {
            it.amount = amount
            it.updatedBy = author

            // 4. Update existing ratings
            ratings.update(it)
        }

    }

    /**
     * Returns the RatingEntity with the specified ID.
     *
     * @param id The ID of the RatingEntity to retrieve
     * @return The RatingEntity with the specified ID
     * @throws CustomException if a RatingEntity with the specified ID does not exist
     */
    void delete(UUID id) {
        log.info("Deleting rating with id: ${ id }")
        // 1. Fetch ratings
        RatingEntity rating = getById(id)

        // 2. Delete ratings
        ratings.delete(rating)
    }

}
