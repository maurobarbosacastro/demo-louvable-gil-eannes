package pt.atlanse.blog.repository

import io.micronaut.core.annotation.Nullable
import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.blog.domain.TranslationEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface TranslationRepository extends PageableRepository<TranslationEntity, Long> {
    Optional<TranslationEntity> findByArticleIdAndLang(Long articleId, String lang)

    @Join(value = "article", type = Join.Type.FETCH)
    Page<TranslationEntity> findByTitleContainsOrContentContains(String title, String content, Pageable pageable)

    @Join(value = "article", type = Join.Type.FETCH)
    Page<TranslationEntity> findByTitleContainsOrContentContainsAndArticleStatusIlike(@Nullable String title, @Nullable String content, @Nullable String status, Pageable pageable)

    @Join(value = "article", type = Join.Type.FETCH)
    Page<TranslationEntity> findByTitleContainsOrContentContainsAndArticleStatusIlikeAndArticleViewWebTrue(@Nullable String title, @Nullable String content, @Nullable String status, Pageable pageable)

    @Join(value = "article", type = Join.Type.FETCH)
    Page<TranslationEntity> findByTitleContainsOrContentContainsAndArticleStatusIlikeAndArticleViewMobileTrue(@Nullable String title, @Nullable String content, @Nullable String status, Pageable pageable)

    @Join(value = "article", type = Join.Type.FETCH)
    Page<TranslationEntity> findByTitleContainsOrContentContainsAndArticleViewMobileTrue(@Nullable String title, @Nullable String content, Pageable pageable)

    @Join(value = "article", type = Join.Type.FETCH)
    Page<TranslationEntity> findByTitleContainsOrContentContainsAndArticleViewWebTrue(@Nullable String title, @Nullable String content, Pageable pageable)
}
