package pt.atlanse.products.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import jakarta.inject.Singleton
import pt.atlanse.products.domains.ProductEntity
import pt.atlanse.products.domains.RatingEntity
import pt.atlanse.products.domains.ReviewEntity
import pt.atlanse.products.dtos.Review
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.ReviewRepository
import pt.atlanse.products.utils.ExceptionService

@Slf4j
@Singleton
class ReviewService {

    ReviewRepository reviews

    ReviewService(ReviewRepository reviews) {
        log.debug "Injecting review repository"
        this.reviews = reviews
    }

    /**
     * Returns the ReviewEntity with the specified ID.
     *
     * @param id The ID of the ReviewEntity to retrieve
     * @return The ReviewEntity with the specified ID
     * @throws CustomException if a ReviewEntity with the specified ID does not exist
     */
    private ReviewEntity getById(UUID id) throws CustomException {
        log.info("Getting review with id: ${ id }")
        reviews.findById(id).orElseThrow(ExceptionService::ReviewNotFoundException())
    }

    /**
     * Returns a paginated list of all reviews.
     *
     * @param pageable The pagination information
     * @return A page of ReviewEntity objects
     */
    Page<ReviewEntity> findAll(Pageable pageable) {
        log.info("Getting all reviews")
        reviews.findAll(pageable)
    }

    /**
     * Returns the ReviewEntity with the specified ID.
     *
     * @param id The ID of the ReviewEntity to retrieve
     * @return The ReviewEntity with the specified ID
     */
    ReviewEntity findById(UUID id) {
        getById(id)
    }

    /**
     * Creates a new ReviewEntity with the given payload, rating, and product.
     *
     * @param payload The review payload
     * @param rating The associated rating entity
     * @param product The associated product entity
     * @param author The author of the review (default "anonymous")
     * @return The newly created ReviewEntity
     */
    //todo remove anonymous author
    ReviewEntity create(Review payload, RatingEntity rating, ProductEntity product, String author = "anonymous") {
        log.info("Creating new review")

        try {
            // 1. Create review entity
            ReviewEntity review = new ReviewEntity(
                content: payload.content,
                createdBy: author,
                updatedBy: author,
                rating: rating,
                product: product
            )

            // 2. Save review using repository
            reviews.save(review)
        } catch (Exception e) {
            ExceptionService.ReviewCreatingException(e, payload)
        }
    }

    /**
     * Overwrites the ReviewEntity with the specified ID with the given payload, rating, and author.
     *
     * @param id The ID of the ReviewEntity to update
     * @param payload The new review payload
     * @param rating The associated rating entity
     * @param author The author of the review (default "anonymous")
     *
     * @deprecated
     */
    void overwrite(UUID id, Review payload, RatingEntity rating, String author = "anonymous") {
        // 1. Fetch review
        Optional<ReviewEntity> review = reviews.findById(id)

        // 2. If the review does not exist, create a new one
        if (review.isEmpty()) {
            create(payload, rating)
            return
        }

        // 3. Try to overwrite existing review
        try {
            review.get().with {
                it.content = payload.content
                it.rating = rating
                it.updatedBy = author

                // 4. Update existing review
                reviews.update(it)
            }
        } catch (Exception e) {
            // 5. Handle exceptions found
            log.error "Unhandled exception found while attempting to update review $id; Reason ${ e.message }"
        }
    }

    /**
     * Partially updates the ReviewEntity with the specified ID with the given payload, rating, and author.
     *
     * @param id The ID of the ReviewEntity to update
     * @param payload The new review payload
     * @param rating The associated rating entity
     * @param author The author of the review (default "anonymous")
     */
    void partialUpdate(UUID id, Review payload, RatingEntity rating, String author = "anonymous") {
        log.info("Updating review with id: ${ id }")
        // 1. Fetch review
        ReviewEntity review = getById(id)

        // 3. Update existing review fields
        review.with {
            it.content = payload.content
            it.rating = rating
            it.updatedBy = author

            // 4. Update existing review
            reviews.update(it)
        }

    }

    /**
     * Deletes the ReviewEntity with the specified ID.
     *
     * @param id The ID of the ReviewEntity to delete
     */
    void delete(UUID id) {
        log.info("Deleting review with id: ${ id }")
        // 1. Fetch review
        ReviewEntity review = getById(id)

        // 2. Delete review
        reviews.delete(review)
    }

    /**
     * Returns a paginated list of all reviews for the specified product.
     *
     * @param product The product to retrieve reviews for
     * @param pageable The pagination information
     * @return A page of ReviewEntity objects
     */
    Page<ReviewEntity> findAll(ProductEntity product, Pageable pageable) {
        reviews.findByProduct(product, pageable)
    }

}
