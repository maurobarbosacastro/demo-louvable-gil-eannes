package pt.atlanse.mscompany.services

import groovy.util.logging.Slf4j
import jakarta.inject.Singleton
import pt.atlanse.mscompany.clients.ClientOperations
import pt.atlanse.mscompany.dtos.Media
import pt.atlanse.mscompany.utils.ExceptionService


@Slf4j
@Singleton
class ImagesClientService {
    ClientOperations images
    private final String DEFAULT_EXTENSION = ".png"

    ImagesClientService(ClientOperations images) {
        this.images = images
    }

    String create(Media payload) {
        try {
            return images.create(payload.base64, payload.fileName).body()['id']
        } catch (Exception ignored) {
            log.print(ignored)
            log.error(ExceptionService.ImageSavingException(payload).toString())
            return null
        }
    }

    String upload(String encoded) {
        log.info "Uploading new image for ms-images..."

        String randomName = UUID.randomUUID().toString() + DEFAULT_EXTENSION

        // Return the image ID or Null if an exception occurred,
        // This approach will grant us time to deal with any problems that might originate
        // from the ms-images service
        try {
            return images.create(encoded, randomName).body()['id']
        } catch (Exception e) {
            log.error "Random exception occured on the server from the service ms-images... Details following: $e"
            return null
        }
    }

}
