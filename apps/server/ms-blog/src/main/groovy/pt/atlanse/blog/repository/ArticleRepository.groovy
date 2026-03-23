package pt.atlanse.blog.repository

import io.micronaut.data.annotation.Query
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.blog.domain.ArticleEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ArticleRepository extends PageableRepository<ArticleEntity, Long> {
    Page<ArticleEntity> findByStatusIlikeAndViewWebTrue(String Status, Pageable pageable)

    Page<ArticleEntity> findByStatusIlikeAndViewMobileTrue(String Status, Pageable pageable)

    Page<ArticleEntity> findByStatusIlike(String Status, Pageable pageable)

    Page<ArticleEntity> findByViewMobileTrue(Pageable pageable)

    Page<ArticleEntity> findByViewWebTrue(Pageable pageable)

    @Query(value = "SELECT * FROM articles WHERE STATUS = 'PUBLISHED' AND view_mobile = 'true' ORDER BY RANDOM() LIMIT 1", nativeQuery = true)
    ArticleEntity getRandomArticle()

    @Query(value = "SELECT * FROM articles WHERE STATUS = 'PUBLISHED' AND view_web = 'true' ORDER BY RANDOM() LIMIT 3", nativeQuery = true)
    List<ArticleEntity> getRandomArticles()
}
