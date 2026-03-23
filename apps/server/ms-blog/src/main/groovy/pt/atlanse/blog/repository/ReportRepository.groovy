package pt.atlanse.blog.repository

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.blog.domain.ArticleEntity
import pt.atlanse.blog.domain.CommentEntity
import pt.atlanse.blog.domain.ReportEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ReportRepository extends PageableRepository<ReportEntity, Long> {

    List<ReportEntity> findAllByArticle(ArticleEntity article)

    List<ReportEntity> findAllByComment(CommentEntity comment)

    @Join(value = "comment", type = Join.Type.FETCH)
    Page<ReportEntity> findAllByStatus(String status, Pageable pages)

    Optional<ReportEntity> findByArticleAndCreatedBy(ArticleEntity article, String author)

    Optional<ReportEntity> findByCommentAndCreatedBy(CommentEntity comment, String author)

}

