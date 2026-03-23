package pt.atlanse.blog.services.imageprocessing

import groovy.util.logging.Slf4j
import io.micronaut.context.annotation.Value
import jakarta.inject.Singleton
import pt.atlanse.blog.DTO.MediaDTO
import pt.atlanse.blog.config.ArticleConfiguration
import pt.atlanse.blog.domain.imageprocessing.MediaFileDoc
import pt.atlanse.blog.repository.imageprocessing.MediaFileDocRepository

import java.nio.file.Files
import java.nio.file.Paths

/**
 * @deprecated
 * */
@Slf4j
@Singleton
class MediaService {

    // Config from the application . yml file
    @Value('${articles.directory}')
    String rootDirectory
    MediaFileDocRepository mediaFileDocRepository

    MediaService(MediaFileDocRepository mediaFileDocRepository) {
        this.mediaFileDocRepository = mediaFileDocRepository
    }

    MediaFileDoc createMediaFile(MediaDTO media, ArticleConfiguration config) {
        try {
            FileOutputStream fos
            Base64.Decoder dec = Base64.getDecoder()
            byte[] bytes = dec.decode(media.getBase64())
            String path = rootDirectory + config.directory

            String uuid = UUID.randomUUID() as String
            String publicCode = UUID.randomUUID().toString() + uuid + UUID.randomUUID().toString()

            File newFile = new File("$path/$uuid.${ media.extension }")
            Files.createDirectories(Paths.get(path))
            fos = new FileOutputStream(newFile)
            fos.write(bytes)
            fos.close()

            MediaFileDoc doc = new MediaFileDoc(
                name: media.fileName,
                publicCode: publicCode,
                path: "$path/$uuid.${ media.extension }",
                hash: bytes.sha256()
            )

            return mediaFileDocRepository.save(doc)

        } catch (Exception e) {
            log.error "Error while trying to create doc file and/or entity"
            return null
        }
    }

    MediaFileDoc findByhash(String hash) {
        mediaFileDocRepository.findByHash(hash).ifPresent { return it }
        return null
    }

}
