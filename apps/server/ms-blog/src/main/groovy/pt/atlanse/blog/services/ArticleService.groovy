package pt.atlanse.blog.services

import groovy.util.logging.Slf4j
import io.micronaut.context.annotation.Property
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.blog.DTO.ArticleDTO
import pt.atlanse.blog.DTO.ArticleParameters
import pt.atlanse.blog.domain.ArticleEntity
import pt.atlanse.blog.domain.TranslationEntity
import pt.atlanse.blog.domain.imageprocessing.MediaFileDoc
import pt.atlanse.blog.models.Article
import pt.atlanse.blog.models.Articles
import pt.atlanse.blog.models.CustomException
import pt.atlanse.blog.repository.ArticleRepository
import pt.atlanse.blog.repository.imageprocessing.MediaFileDocRepository
import pt.atlanse.blog.services.imageprocessing.ArticleMediaFileService

import java.time.format.DateTimeFormatter

@Slf4j
@Singleton
class ArticleService {

    /**@deprecated */
    private final String DEFAULT_LANG = "en"

    // Todo attempt to inject during introspection
    @Property(name = "articles.filter-word")
    private int FILTER_MIN_SIZE

    @Inject
    ArticleMediaFileService mediaService

    @Inject
    TranslationService translationService

    /**
     * @deprecated
     * */
    @Inject
    KeycloakService keycloakService

    @Inject
    LikeService likes

    @Inject
    CommentService comments

    ArticleRepository articles

    MediaFileDocRepository mediaFileDocRepository

    ArticleService(ArticleRepository articles, MediaFileDocRepository mediaFileDocRepository) {
        this.articles = articles
        this.mediaFileDocRepository = mediaFileDocRepository
    }

    /**
     * Build the holy bible of pages of type {@link TranslationEntity} or even {@link ArticleEntity}
     * @param params Other parameters. Class type {@link ArticleParameters}. Allows usage of filters for image, status and text content
     * @param pageable ({@link Pageable}) obtained from the parameters and used for the pagination (e.g., size and page)
     * @return An entire {@link Page} book made of type {@link TranslationEntity} or {@link ArticleEntity}
     * */
    private Page<?> applyFilters(ArticleParameters params, Pageable pageable) {
        if (params.status) {
            if (params.searchText && params.target) {
                return translationService.findPatternStatusTarget(params.searchText, params.status, params.target, pageable)
            }

            if (params.searchText) {
                return translationService.findPatternStatus(params.searchText, params.status, pageable)
            }

            if (params.target) {
                return (params.target.toLowerCase() == "mobile") ?
                    articles.findByStatusIlikeAndViewMobileTrue(params.status, pageable) :
                    articles.findByStatusIlikeAndViewWebTrue(params.status, pageable)
            }

            // search by status only
            return articles.findByStatusIlike(params.status, pageable)
        }

        if (params.searchText && params.target) {
            return translationService.findPatternTarget(params.searchText, params.target, pageable)
        }

        if (params.searchText) {
            return translationService.findByPattern(params.searchText, pageable)
        }

        if (params.target) {
            return (params.target.toLowerCase() == "mobile") ?
                articles.findByViewMobileTrue(pageable) :
                articles.findByViewWebTrue(pageable)
        }

        return articles.findAll(pageable)

    }

    /**
     * Finds all {@link ArticleEntity}
     * @return List of {@link ArticleEntity}
     * */
    Articles findAll(ArticleParameters params, Pageable pageable, String user = null) {
        log.info "Using pageable arguments: Page_number: ${ pageable.offset }; Amount_of_articles: ${ pageable.size }"
        List<Article> articleList = new ArrayList<>()

        Page<?> pages = applyFilters(params, pageable)

        pages.getContent().each { page ->

            // Verify if its a Translation object
            if (page instanceof TranslationEntity) {
                articleList.add(buildArticleResponse(page.article, page, params, user))
                return
            }

            // Get translation if its actually an ArticleEntity
            Optional<TranslationEntity> translation = translationService.findByArticleIdAndLang((page as ArticleEntity).id, DEFAULT_LANG)
            articleList.add(
                buildArticleResponse((page as ArticleEntity), translation.orElseThrow {
                    new CustomException(
                        "Error fetching translation",
                        "Error fetching translation for lang: $DEFAULT_LANG; Not found",
                        HttpStatus.BAD_REQUEST
                    )
                }, params, user)
            )
        }

        return new Articles(
            pageNumber: pageable.getNumber() + 1,
            totalPages: pages.totalPages,
            totalElements: pages.totalSize,
            content: articleList
        )

    }

    /**
     * This method searches for an article using it's ID and throws an exception if the article was not found
     *
     * @param id The identifier of the article
     * @return object of type {@link ArticleEntity}
     * @throw Exception of type {@link CustomException} if the article was not found
     * */
    ArticleEntity find(Long id, String onError = null) {
        return articles.findById(id).orElseThrow {
            new CustomException(
                "Article not found",
                onError ?: "Article with id $id was not found",
                HttpStatus.NOT_FOUND
            )
        }
    }

    /**
     * Runs the private method {@link ArticleEntity} and converts the object to a representation
     * @param id String with the {@link ArticleEntity} ID
     * @return {@link ArticleEntity}
     * */
    Article findById(String id, ArticleParameters params, String user = null) {
        // 1. Article by ID
        ArticleEntity entity = find(Long.parseLong(id))

        // 2. Get translations for this Article
        // todo Add multi lang by replacing the DEFAULT_LANG var
        Optional<TranslationEntity> translation = translationService.findByArticleIdAndLang(entity.id, DEFAULT_LANG)

        // 3. Build Pojo
        Article pojo = buildArticleResponse(entity, translation.orElseThrow {
            new CustomException(
                "Error fetching translation",
                "Error fetching translation for lang: $DEFAULT_LANG; Not found",
                HttpStatus.BAD_REQUEST
            )
        }, params, user)

        return pojo
    }

    private Article buildArticleResponse(ArticleEntity entity, TranslationEntity translation, ArticleParameters params, String user = null) {
        // Run the tuple constructor
        Article article = new Article(
            // Article main object
            id: entity.id,
            status: entity.status,
            viewMobile: entity.viewMobile,
            viewWeb: entity.viewWeb,
            createdAt: entity.createdAt.format(DateTimeFormatter.ISO_LOCAL_DATE_TIME),
            updatedAt: entity.updatedAt.format(DateTimeFormatter.ISO_LOCAL_DATE_TIME),

            // Returns "User was not found" if the user was not found
            // This is useful for specific situations such as User bans or when the user deletes the account
            createdBy: keycloakService.getUserEmail(entity.createdBy), // Creation
            updatedBy: keycloakService.getUserEmail(entity.updatedBy),
            firstName: keycloakService.findUser(entity.createdBy).firstName,
            lastName: keycloakService.findUser(entity.createdBy).lastName,

            // translations
            title: translation.title,
            subtitle: translation.subtitle,
            content: translation.content,
            conclusion: translation.conclusion,
            enabled: translation.enabled,

            likes: likes.count(entity),
            comments: comments.count(entity),

            liked: user ? likes.userLikedArticle(entity, user) : null
        )

        // 4. Check if we need to add the image and if it exists
        if (params.image) {
            // This catch is used because we dont want to break the cycle;
            // If the article doesn't have an image (for some reason) the article should still
            // be available for editing
            try {
                article.setImage(mediaService.loadMedia(entity.articleMediaFile))
            } catch (CustomException e) {
                log.error "Error while fetching this image; ${ e.message }"
            }
        }

        return article
    }

    /**
     * Create object. Recommended for administration, otherwise the api should be register
     *
     * @param {@link ArticleDTO}
     * @param authorEmail The author of the article
     *
     * @return void
     * */
    void create(ArticleDTO details, String author) {
        // 1. Build and Save an article
        ArticleEntity article = this.articles.save(fillArticleEntity(new ArticleEntity(), details, author))

        // 2. Create and Save translations for the article above
        this.translationService.createTranslation(article, details.translations.first())
    }

    /**
     * This method selects a User with the ID received as a param
     * and updates the information using a body of type {@link ArticleDTO}.
     *
     * @param id String with the user ID
     * @param R {@link ArticleDTO} object with data about user
     *
     * @return response object of type
     * */
    void update(String id, ArticleDTO details, String updatedBy) {
        // Search article
        ArticleEntity article = find(Long.parseLong(id))

        // Update entity
        article = fillArticleEntity(article, details, updatedBy, true)

        this.articles.update(article)
        this.translationService.createTranslation(article, details.translations.first(), true)
    }

    /**
     * @return ArticleEntity object
     * */
    private ArticleEntity fillArticleEntity(ArticleEntity article, ArticleDTO details, String author, boolean isUpdate = false) {

        Closure buildMediaFile = { return mediaService.create(details.image, "image", author) }
        Closure searchMediaFile = { return mediaService.findById(article.articleMediaFile.id) }
        Closure getDecodedImage = { return Base64.decoder.decode(details.image.base64).sha256() }

        // Default to change
        article.setStatus(details.status)
        article.setViewMobile(details.isMobile)
        article.setViewWeb(details.isWeb)
        article.setUpdatedBy(author)

        if (!isUpdate && !article.articleMediaFile) {
            // If the isUpdate flag is Off, this means the article is being created and not updated.
            // Add the author as the created
            article.setCreatedBy(author)
            // todo Verify if the image already exists
            article.setArticleMediaFile(buildMediaFile())

            return article
        }

        searchMediaFile().ifPresent { file ->
            MediaFileDoc doc = mediaService.loadDoc(file.mediaFile.id)

            // Run this if the image is different; UPDATE
            if (getDecodedImage() != doc.hash) {
                log.debug "The hash is new"

                // todo Delete replaced image
                article.setArticleMediaFile(buildMediaFile())
            }
        }

        return article

    }

    /**
     * This method deletes a users with the id received as a param
     *
     * @param id String with the {@link ArticleDTO}  ID
     * */
    void delete(String id) {
        throw new CustomException(
            "Endpoint unavailable",
            "This endpoint is currently under maintenance",
            HttpStatus.SERVICE_UNAVAILABLE
        )
    }

    /**
     * Runs the private method {@link ArticleEntity} and converts the object to a representation
     * @return {@link ArticleEntity}
     * */
    Article findRandomArticle(ArticleParameters params, String user) {
        // 1. Article by ID
        ArticleEntity entity = articles.getRandomArticle()

        // 2. Get translations for this Article
        // todo Add multi lang by replacing the DEFAULT_LANG var
        Optional<TranslationEntity> translation = translationService.findByArticleIdAndLang(entity.id, DEFAULT_LANG)

        // 3. Build Pojo
        Article pojo = buildArticleResponse(entity, translation.orElseThrow {
            new CustomException(
                "Error fetching translation",
                "Error fetching translation for lang: $DEFAULT_LANG; Not found",
                HttpStatus.BAD_REQUEST
            )
        }, params, user)

        return pojo
    }

    List<Article> findRandomArticles(ArticleParameters params, String user = null) {

        // 1. Article by ID
        List<ArticleEntity> entities = articles.getRandomArticles()
        List<Article> articleList = new ArrayList<>()

        entities.each { entity ->

            // Get translation if its actually an ArticleEntity
            Optional<TranslationEntity> translation = translationService.findByArticleIdAndLang((entity as ArticleEntity).id, DEFAULT_LANG)
            articleList.add(
                buildArticleResponse((entity as ArticleEntity), translation.orElseThrow {
                    new CustomException(
                        "Error fetching translation",
                        "Error fetching translation for lang: $DEFAULT_LANG; Not found",
                        HttpStatus.BAD_REQUEST
                    )
                }, params, user)
            )
        }

        return articleList
    }
}
