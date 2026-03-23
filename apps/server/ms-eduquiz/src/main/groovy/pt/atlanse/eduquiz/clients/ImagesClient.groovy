package pt.atlanse.eduquiz.clients


import io.micronaut.http.HttpResponse
import io.micronaut.http.MediaType
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Header
import io.micronaut.http.annotation.Post
import io.micronaut.http.client.annotation.Client

interface ClientOperations {

    HttpResponse create(String file, String fileName)

    HttpResponse get(/*@Header(name = "Authorization") String authorization, */ String imageId)

    HttpResponse transform(@Header(name = "Authorization") String authorization, String keycloakId)
}

@Client("images")
interface ImagesClient extends ClientOperations {

    @Post(uri = "/api/image/base64", consumes = MediaType.APPLICATION_JSON, produces = MediaType.APPLICATION_JSON)
    HttpResponse create(String file, String fileName)

    @Get("/api/image/{imageId}")
    HttpResponse get(
        @Header(name = "Authorization") String authorization,
        String imageId)

    @Get("/api/image/{imageId}/free-transform")
    HttpResponse transform(@Header(name = "Authorization") String authorization, String imageId)
}

