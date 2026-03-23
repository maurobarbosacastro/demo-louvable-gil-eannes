package pt.atlanse.blog.repository

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.blog.domain.ArticleEntity
import pt.atlanse.blog.domain.CommentEntity

import java.time.LocalDateTime

@JdbcRepository(dialect = Dialect.POSTGRES)
interface CommentRepository extends PageableRepository<CommentEntity, Long> {
    Page<CommentEntity> findAllByArticle(ArticleEntity article, Pageable pageable)

    Page<CommentEntity> findAllByArticleAndHiddenFalseOrderByCreatedAtDesc(ArticleEntity article, Pageable pageable)

    Page<CommentEntity> findAllByParent(CommentEntity comment, Pageable pageable)

    Page<CommentEntity> findAllByParentAndHiddenFalseOrderByCreatedAtDesc(CommentEntity comment, Pageable pageable)

    List<CommentEntity> findTop3ByParentAndHiddenFalseAndCreatedAtLessThanOrderByCreatedAtDesc(CommentEntity comment, LocalDateTime date)

    List<CommentEntity> findTop4ByParentIsNullAndCreatedAtLessThanOrderByCreatedAtDesc(LocalDateTime date)

    List<CommentEntity> findTop3ByParentAndHiddenFalseAndCreatedAtLessThanOrderByCreatedAtAsc(CommentEntity parent, LocalDateTime date)

    List<CommentEntity> findTop4ByParentIsNullAndCreatedAtLessThanOrderByCreatedAtAsc(LocalDateTime date)

    Optional<CommentEntity> findById(Long id)

    long countByArticle(ArticleEntity article)

    long countByParent(CommentEntity comment)

    List<CommentEntity> findTop50ByCreatedByAndHiddenFalseOrderByCreatedAtDesc(String author)

    List<CommentEntity> findTop50ByParentCreatedByAndHiddenFalseOrderByCreatedAt(String author)
}
