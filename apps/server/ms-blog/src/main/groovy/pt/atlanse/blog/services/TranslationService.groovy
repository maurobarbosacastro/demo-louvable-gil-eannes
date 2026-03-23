package pt.atlanse.blog.services

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import jakarta.inject.Singleton
import jakarta.transaction.Transactional
import pt.atlanse.blog.DTO.TranslationDTO
import pt.atlanse.blog.domain.ArticleEntity
import pt.atlanse.blog.domain.TranslationEntity
import pt.atlanse.blog.repository.TranslationRepository


@Slf4j
@Singleton
@Transactional
class TranslationService {

    private final String DEFAULT_LANG = "en"

    TranslationRepository translations

    TranslationService(TranslationRepository translations) {
        this.translations = translations
    }

    /**
     * Used internally to fill a {@link pt.atlanse.blog.domain.TranslationEntity} object
     * @param article The article entity of type {@link pt.atlanse.blog.domain.ArticleEntity}
     * @param details The translation details of class {@link pt.atlanse.blog.DTO.TranslationDTO}
     * @param isUpdated The boolean that serves as a flag to identify an update (to be removed)
     * @return Object of type {@link pt.atlanse.blog.domain.TranslationEntity}
     * */
    TranslationEntity fillTranslationEntity(ArticleEntity article, TranslationDTO details, boolean isUpdate = false) {
        Optional<TranslationEntity> translationOpt = translations.findByArticleIdAndLang(article.id, DEFAULT_LANG) // Search only english

        // Fill translation data like it was a form
        Closure setTranslation = (TranslationEntity t) -> {
            t.lang = details.lang
            t.title = details.title
            t.subtitle = details.subtitle
            t.content = details.content
            t.conclusion = details.conclusion
            t.enabled = details.enabled

            return t
        }
        // If the optional has a translation:
        if (!translationOpt.isEmpty()) {
            return setTranslation(translationOpt.get())
        }

        // Generate new translation
        TranslationEntity translation = setTranslation(new TranslationEntity())
        translation.setArticle(article)
        return translation

    }

    /**
     * Gets a translation for given language on a given article
     * @param id A long for the id of the article
     * @param lang language using 2 chars (e.g., EN, FR or IT)
     * @return Optional of {@link TranslationEntity}
     * */
    Optional<TranslationEntity> findByArticleIdAndLang(Long id, String lang) {
        log.info "Searching translations for article #$id"
        lang = DEFAULT_LANG //todo remove this
        return translations.findByArticleIdAndLang(id, lang)
    }

    /**
     * Retrieves a list of {@link TranslationEntity} using the article id
     * @param id The id of the article
     * @deprecated Using default language EN
     * */
    List<TranslationEntity> findByArticleId(Long id) {
        List<TranslationEntity> translations = new ArrayList<>()

        log.info "Searching translations for article #$id"
        this.findByArticleIdAndLang(id, DEFAULT_LANG).ifPresent {
            translations.add(it)
        }

        return translations
    }

    /**
     * This method creates a translation according the the payload parameter
     * @param payload An object of class {@link TranslationDTO}
     * @param article An {@link ArticleEntity} object
     * @param isUpdate A boolean that serves as a <i>flag</i> to identify if the user is updating
     * or creating a new translation
     * @return void
     * */
    void createTranslation(ArticleEntity article, TranslationDTO payload, boolean isUpdate = false) {
        log.info "Creating translation: ${ payload.toString() }"

        TranslationEntity translation = this.fillTranslationEntity(article, payload)

        if (isUpdate) {
            translations.update(translation)
            return
        }

        translations.save(translation)
    }

    /**
     * Temporarily not being used <br>
     * This method creates multiple translations according the the translationList parameter
     * @param translationList An array of objects of class {@link TranslationDTO}
     * @param article An {@link ArticleEntity} object
     * @return void
     * */
    void createTranslations(ArticleEntity article, List<TranslationDTO> translationList) {
        translationList.each {
            this.createTranslation(article, it)
        }
    }

    /**
     * Finds translations using a pattern <br>
     * @param pattern String pattern to match
     * @param pageable Micronaut object {@link Pageable}
     * */
    Page<TranslationEntity> findByPattern(String pattern, Pageable pageable) {
        translations.findByTitleContainsOrContentContains(pattern, pattern, pageable)
    }

    Page<TranslationEntity> findPatternStatus(@Nullable String pattern, @Nullable String status, Pageable pageable) {
        translations.findByTitleContainsOrContentContainsAndArticleStatusIlike(pattern, pattern, status, pageable)
    }

    Page<TranslationEntity> findPatternStatusTarget(@Nullable String pattern, @Nullable String status, String target, Pageable pageable) {

        // is Mobile view
        if (target.toLowerCase() == "mobile") {
            return translations.findByTitleContainsOrContentContainsAndArticleStatusIlikeAndArticleViewMobileTrue(pattern, pattern, status, pageable)
        }

        // is Web view
        return translations.findByTitleContainsOrContentContainsAndArticleStatusIlikeAndArticleViewWebTrue(pattern, pattern, status, pageable)
    }

    Page<TranslationEntity> findPatternTarget(@Nullable String pattern, String target, Pageable pageable) {

        // is Mobile view
        if (target.toLowerCase() == "mobile") {
            return translations.findByTitleContainsOrContentContainsAndArticleViewMobileTrue(pattern, pattern, pageable)
        }

        // is Web view
        return translations.findByTitleContainsOrContentContainsAndArticleViewWebTrue(pattern, pattern, pageable)
    }

    //TODO : SEE
    TranslationEntity findByArticle(Long id) {

        return translations.findByArticleId(id)
    }
}
