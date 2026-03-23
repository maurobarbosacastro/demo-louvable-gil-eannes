package pt.atlanse.products.services

import groovy.util.logging.Slf4j
import jakarta.inject.Singleton
import pt.atlanse.products.clients.ClientOperations
import pt.atlanse.products.dtos.Media


@Slf4j
@Singleton
class ImagesClientService {
    ClientOperations images

    ImagesClientService(ClientOperations images) {
        this.images = images
    }

    String create(Media payload) {
        try {
            return images.create(payload.base64, payload.fileName).body()['id']
        } catch (Exception ignored) {
            log.warn "Error saving image on ms-images; Payload: ${payload.extension}"
            return null
        }
    }

}
