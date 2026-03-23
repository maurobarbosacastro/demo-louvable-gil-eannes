package pt.atlanse.products.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpStatus
import jakarta.inject.Singleton
import pt.atlanse.products.domains.ProductEntity
import pt.atlanse.products.domains.SellingStoreEntity
import pt.atlanse.products.dtos.SellingStore
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.SellingStoreRepository
import pt.atlanse.products.utils.ExceptionService

@Slf4j
@Singleton
class SellingStoreService {

    SellingStoreRepository stores

    SellingStoreService(SellingStoreRepository stores) {
        log.debug "Injecting Stores repository"
        this.stores = stores
    }

    /**
     * Gets a SellingStoreEntity with the specified ID.
     *
     * @param id the ID of the SellingStoreEntity to retrieve
     * @return the SellingStoreEntity with the specified ID
     * @throws CustomException if the SellingStoreEntity with the specified ID is not found
     */
    private SellingStoreEntity getById(UUID id) throws CustomException {
        log.info("Getting sellingStore with id: ${ id }")
        Optional<SellingStoreEntity> store = stores.findById(id)

        if (store.isEmpty()) {
            throw new CustomException(
                "store not found",
                "store with id $id was not found",
                HttpStatus.NOT_FOUND
            )
        }

        store.get()
    }

    /**
     * Finds multiple SellingStoreEntity objects with the specified IDs.
     *
     * @param ids the IDs of the SellingStoreEntity objects to retrieve
     * @return a list of SellingStoreEntity objects with the specified IDs
     */
    List<SellingStoreEntity> findMultiple(List<UUID> ids) {
        List<SellingStoreEntity> entities = new ArrayList<>()

        ids.each {
            try {
                entities.add(getById(it))
            } catch (Exception ignored) {
                log.error "The specified store with id $it was not found... jumping to next iteration"
            }
        }

        entities
    }

    /**
     * Retrieves a page of SellingStoreEntity objects.
     *
     * @param pageable the Pageable object containing information about which page to retrieve
     * @return a page of SellingStoreEntity objects
     */
    Page<SellingStoreEntity> findAll(Pageable pageable) {
        log.info("Getting all sellingStores")
        stores.findAll(pageable)
    }

    /**
     * Gets a SellingStoreEntity with the specified ID.
     *
     * @param id the ID of the SellingStoreEntity to retrieve
     * @return the SellingStoreEntity with the specified ID
     */
    SellingStoreEntity findById(UUID id) {
        getById(id)
    }

    /**
     * Creates a new SellingStoreEntity object with the specified payload.
     *
     * @param payload the SellingStore object containing information to create a new SellingStoreEntity object
     * @param author the name of the author creating the SellingStoreEntity object (default is "anonymous")
     */
    //todo remove anonymous author
    SellingStoreEntity create(SellingStore payload, ProductEntity product, String author = "anonymous") {
        log.info "Creating new store"

        try {
            // 1. Create store entity
            SellingStoreEntity store = new SellingStoreEntity(
                name: payload.name,
                website: payload.website,
                product: product,
                createdBy: author,
                updatedBy: author
            )

            // 2. Save store using repository
            stores.save(store)
        } catch (Exception e) {
            ExceptionService.SellingStoreCreatingException(e, payload)
        }
    }

    /**
     * Overwrites an existing SellingStoreEntity object with the specified ID with the information in the payload.
     *
     * @param id the ID of the SellingStoreEntity object to overwrite
     * @param payload the SellingStore object containing information to overwrite the SellingStoreEntity object
     * @param author the name of the author overwriting the SellingStoreEntity object (default is "anonymous")
     */
    void overwrite(UUID id, SellingStore payload, String author = "anonymous") {
        // 1. Fetch store
        Optional<SellingStoreEntity> store = stores.findById(id)

        // 2. If the store does not exist, create a new one
        if (store.isEmpty()) {
            create(payload)
            return
        }

        // 3. Try to overwrite existing store
        try {
            store.get().with {
                it.name = payload.name
                it.website = payload.website
                it.updatedBy = author

                // 4. Update existing store
                stores.update(it)
            }
        } catch (Exception e) {
            // 5. Handle exceptions found
            log.error "Unhandled exception found while attempting to update store $id; Reason ${ e.message }"
        }
    }

    /**
     * Partially updates an existing SellingStoreEntity object with the specified ID with the information in the payload.
     *
     * @param id the ID of the SellingStoreEntity object to update
     * @param payload the SellingStore object containing information to update the SellingStoreEntity object
     * @param author the name of the author updating the SellingStoreEntity object (default is "anonymous")
     */
    void partialUpdate(UUID id, SellingStore payload, String author = "anonymous") {
        log.info("Updating sellingStore with id: ${ id }")
        // 1. Fetch store
        SellingStoreEntity store = getById(id)

        // 3. Update existing store fields
        store.with {
            it.name = payload.name
            it.website = payload.website
            it.updatedBy = author

            // 4. Update existing store
            stores.update(it)
        }

    }

    /**
     * Deletes a SellingStoreEntity object with the specified ID.
     *
     * @param id the ID of the SellingStoreEntity object to delete
     */
    void delete(UUID id) {
        log.info("Deleting sellingStore with id: ${ id }")
        // 1. Fetch brand
        SellingStoreEntity store = getById(id)

        // 2. Delete brand
        stores.delete(store)
    }

}
