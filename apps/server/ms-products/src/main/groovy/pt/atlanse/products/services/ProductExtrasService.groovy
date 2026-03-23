package pt.atlanse.products.services

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpStatus
import io.micronaut.transaction.annotation.Transactional
import jakarta.inject.Inject
import jakarta.inject.Singleton
import jakarta.persistence.EntityManager
import jakarta.persistence.PersistenceContext
import pt.atlanse.products.domains.ExtrasEntity
import pt.atlanse.products.dtos.Extras
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.ExtrasRepository
import pt.atlanse.products.repositories.ProductRepository


@Slf4j
@Singleton
class ProductExtrasService {

    @Inject
    ExtrasRepository extrasRepository

    @Inject
    ProductRepository productRepository

    @PersistenceContext
    EntityManager entityManager

    /**
     * <b>CREATE</b>
     * Create a new Extra for the products
     * @param payload {@link Extras} object
     * @return void
     * */
    //todo remove anonymous author
    @Transactional
    void create(UUID productId, List<UUID> extrasListPayload) {
        log.info "Creating new Product Extra"


        try {
            Set<ExtrasEntity> extrasList = new HashSet<>()

            productRepository.findById(productId).ifPresent {
                it ->
                    {
                        extrasListPayload.each { extras ->
                            extrasRepository.findById(extras).ifPresent { extra ->
                                {
                                    extrasList.add(extra)
                                }
                            }
                        }
                        it.extras.addAll(extrasList)
                        productRepository.update(it)
                    }
            }
        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error creating Product Extras",
                "Error happened while trying to create new Product Extras ${extrasListPayload.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }


    /**
     * <b>DELETE</b>
     * Delete extras
     * @param id {@link UUID} string
     * @return void
     * */
    @Transactional
    void delete(UUID productId, List<UUID> extrasListPayload) {
        try {
            productRepository.findById(productId).ifPresent {
                product ->
                    {
                        extrasListPayload.each { extras ->
                            extrasRepository.findById(extras).ifPresent { extra ->
                                {
                                    int isDeleted = productRepository.deleteExtraRelation(product.id, extra.id)
                                    log.print("Deleted ${isDeleted == 1}")
                                }
                            }
                        }
                    }
            }

        } catch (Exception e) {
            log.error "Unhandled exception occured: Reason: ${e.message}"
            throw new CustomException(
                "Error Deleting Product Extras",
                "Error happened while trying to delete Product Extras ${extrasListPayload.toString()}; Reason: ${e.message}",
                HttpStatus.BAD_REQUEST
            )
        }
    }
}
