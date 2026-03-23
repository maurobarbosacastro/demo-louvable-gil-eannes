package pt.atlanse.eduquiz.services

import io.micronaut.http.HttpStatus
import jakarta.inject.Singleton
import pt.atlanse.eduquiz.clients.ClientOperations
import pt.atlanse.eduquiz.models.CustomException
import pt.atlanse.eduquiz.models.Media

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
            new CustomException(
                "Error saving image",
                "Error saving image on ms-images; Extension: ${payload.extension}",
                HttpStatus.INTERNAL_SERVER_ERROR
            )
        }
    }

}
