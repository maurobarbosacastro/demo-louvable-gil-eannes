package pt.atlanse.eduquiz.services

import groovy.util.logging.Slf4j
import io.micronaut.context.annotation.Value
import jakarta.inject.Named
import pt.atlanse.eduquiz.config.FileConfiguration
import pt.atlanse.eduquiz.models.Media
import jakarta.inject.Singleton
import java.nio.file.Files
import java.nio.file.Paths

/**
 * @deprecated
 * */
@Slf4j
@Singleton
class FileHandler {

    @Value('${files.directory}')
    private String MEDIA_DIRECTORY

    FileConfiguration documentConfig
    FileConfiguration imageConfig
    FileConfiguration videoConfig

    FileHandler(@Named("document") FileConfiguration documentConfig,
                @Named("image") FileConfiguration imageConfig,
                @Named("video") FileConfiguration videoConfig) {
        this.documentConfig = documentConfig
        this.imageConfig = imageConfig
        this.videoConfig = videoConfig
    }

    /**
     *
     * Returns file bytes as Base64 encoded string
     * @return base64 String ({@link Optional})
     * @param path Location of the file
     * */
    Optional<String> read(String path) {
        log.info "Searching image with path: $path"
        // Read file and return encoded base64
        try {
            File file = new File(path)
            return Optional.of(file.bytes.encodeBase64() as String)
        }
        catch (Exception e) {
            log.error "Exception found while trying to find image by path: $path; Reason: ${ e.message }"
            return Optional.empty()
        }
    }

    /**
     *
     * Returns file bytes as Base64 encoded string
     * @return file path as a String ({@link Optional})
     * @param media {@link Media} object for images, videos and files
     * */
    Optional<String> write(Media media) {
        try {
            String formatDir

            // todo external method
            () -> {
                [imageConfig, documentConfig, videoConfig].each {
                    if (it.allowedFormats.any { format -> format.contains(media.extension) }) {
                        formatDir = it.directory
                    }
                }
            }()

            formatDir = formatDir ?: "other-media-type"

            byte[] bytes = Base64.getDecoder().decode(media.getBase64())
            String finalPath = MEDIA_DIRECTORY + "/" + formatDir + "/${ UUID.randomUUID() }.${ media.extension }"

            File newFile = new File(finalPath)
            Files.createDirectories(Paths.get(MEDIA_DIRECTORY + "/" + formatDir))

            FileOutputStream fos = new FileOutputStream(newFile)
            fos.write(bytes)
            fos.close()

            return Optional.of(finalPath)

        } catch (Exception e) {
            log.error "Exception found while attempting to create image locally...; Reason: ${ e.message }"
            return Optional.empty()
        }

    }
}
