package pt.atlanse.blog.repository

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.blog.domain.ArticleEntity
import pt.atlanse.blog.domain.CommentEntity
import pt.atlanse.blog.domain.LikeEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface LikeRepository extends PageableRepository<LikeEntity, Long> {
    List<LikeEntity> findAllByArticle(ArticleEntity article)

    List<LikeEntity> findAllByComment(CommentEntity comment)

    long countByArticle(ArticleEntity article)

    long countByComment(CommentEntity comment)

    Optional<LikeEntity> findByArticleAndCreatedBy(ArticleEntity article, String author)

    Optional<LikeEntity> findByCommentAndCreatedBy(CommentEntity comment, String author)

    List<LikeEntity> findTop50ByCreatedByOrderByCreatedAtDesc(String author)

    @Join(value = "comment", type = Join.Type.FETCH)
    List<LikeEntity> findTop50ByCreatedByAndCommentHiddenFalseOrderByCreatedAtDesc(String author)

    @Join(value = "comment", type = Join.Type.FETCH)
    List<LikeEntity> findTop50ByCommentCreatedByAndCommentHiddenFalseOrderByCreatedAt(String author)

}
