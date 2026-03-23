package pt.atlanse.blog.services.imageprocessing

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Named
import jakarta.inject.Singleton
import pt.atlanse.blog.DTO.ImageDTO
import pt.atlanse.blog.DTO.MediaDTO
import pt.atlanse.blog.config.ArticleConfiguration
import pt.atlanse.blog.domain.imageprocessing.ArticleMediaFile
import pt.atlanse.blog.domain.imageprocessing.MediaFileDoc
import pt.atlanse.blog.models.CustomException
import pt.atlanse.blog.repository.imageprocessing.ArticleMediaFileRepository
import pt.atlanse.blog.repository.imageprocessing.MediaFileDocRepository

/**
 * @deprecated
 * */
@Slf4j
@Singleton
class ArticleMediaFileService {

    @Inject
    MediaService mediaService

    ArticleConfiguration imageConfig

    ArticleMediaFileRepository articleMediaFileRepository
    MediaFileDocRepository mediaFileDocRepository

    ArticleMediaFileService(
        @Named("image") ArticleConfiguration imageConfig,
        ArticleMediaFileRepository articleMediaFileRepository,
        MediaFileDocRepository mediaFileDocRepository) {
        this.imageConfig = imageConfig
        this.articleMediaFileRepository = articleMediaFileRepository
        this.mediaFileDocRepository = mediaFileDocRepository
    }

    /**
     * Create a document and return the object of class {@link pt.atlanse.blog.domain.imageprocessing.ArticleMediaFile}
     * @param media of the {@link pt.atlanse.blog.DTO.MediaDTO} object
     * @param code type of media (e.g., image)
     * @param userIdentifier The user that created the media
     * @return object of type {@link pt.atlanse.blog.domain.imageprocessing.ArticleMediaFile}
     * */
    // todo review Usefulness of the MediaFileDoc table
    ArticleMediaFile create(MediaDTO media, String code, String userIdentifier) {
        try {
            // CODE has to be one of the values on the yml file (e.g., image, video or document)
            MediaFileDoc mediaFileDoc = mediaService.createMediaFile(media, imageConfig)

            ArticleMediaFile articleMediaFile = new ArticleMediaFile()
            articleMediaFile.mediaFile = mediaFileDoc
            articleMediaFile.code = code
            articleMediaFile.createdBy = userIdentifier
            articleMediaFile.updatedBy = userIdentifier

            articleMediaFileRepository.save(articleMediaFile)
        }
        catch (Exception e) {
            log.error "Exception thrown while trying to access configuration w/ code $code; Reason: ${ e.message }"
            return null
        }
    }

    /**
     * Retrieve object of type {@link ArticleMediaFile}
     * @param id of the {@link ArticleMediaFile} object
     * @return Optional of type {@link MediaFileDoc}
     * */
    Optional<ArticleMediaFile> findById(Long id) {
        articleMediaFileRepository.findById(id)
    }

    /**
     * Retrieve object of type {@link MediaFileDoc}
     * @param id of the {@link MediaFileDoc} object
     * @return object of type {@link MediaFileDoc}
     * */
    MediaFileDoc loadDoc(long id) {
        mediaFileDocRepository.findById(id).get()
    }

    /**
     * Returns image bytes as Base64 encoded string
     * @return base64 String
     * @param articleMediaFile {@link ArticleMediaFile}
     * @deprecated
     * */
    String loadImage(ArticleMediaFile articleMediaFile) {
        try {
            // Without lazy loading, the articleMediaFile contains only the ID
            // It is therefore required to load the object using the repository
            Optional<ArticleMediaFile> mediaFile = findById(articleMediaFile.id)

            // Previous object only has a reference to the MediaFileDoc object which in turn, contains the path of the file
            MediaFileDoc doc = this.mediaFileDocRepository.findById(mediaFile.get().mediaFile.id).get()

            // Read file and return encoded base64
            File file = new File(doc.path)
            return file.bytes.encodeBase64()
        } catch (Exception e) {
            log.error "Error while trying to upload image; Reason ${ e.message }"
            throw new CustomException(
                "Error fetching image",
                "This image may be temporarily unavailable or was removed from the server; We apologize for the inconvenience :(",
                HttpStatus.UNPROCESSABLE_ENTITY
            )
        }
    }

    /**
     * Retrieve the media
     * @param articleMediaFile The mediafile of type {@ArticleMediaFile}
     * @return Object of class {@link MediaDTO}
     * */
    MediaDTO loadMedia(ArticleMediaFile articleMediaFile) {
        try {// Without lazy loading, the articleMediaFile contains only the ID
            // It is therefore required to load the object using the repository
            Optional<ArticleMediaFile> mediaFile = findById(articleMediaFile.id)

            // Previous object only has a reference to the MediaFileDoc object which in turn, contains the path of the file
            MediaFileDoc doc = this.mediaFileDocRepository.findById(mediaFile.get().mediaFile.id).get()

            // Get file
            File file = new File(doc.path)

            MediaDTO media = new ImageDTO()
            media.fileName = file.name
            media.extension = file.name.substring(file.name.lastIndexOf(".") + 1)
            media.base64 = file.bytes.encodeBase64() as String

            return media
        } catch (Exception e) {
            throw new CustomException(
                "Error creating media file",
                "This may happen because the image is temporarily unavailable or was removed from the server; We apologize for the inconvenience :(; Reason: ${ e.message }",
                HttpStatus.BAD_REQUEST
            )
        }

    }

}
