package pt.atlanse.blog.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.blog.domain.ArticleEntity
import pt.atlanse.blog.domain.CommentEntity
import pt.atlanse.blog.models.Comment
import pt.atlanse.blog.models.Comments
import pt.atlanse.blog.models.CustomException
import pt.atlanse.blog.models.Likeable
import pt.atlanse.blog.repository.CommentRepository

import java.time.format.DateTimeFormatter

@Slf4j
@Singleton
class CommentService {

    CommentRepository comments

    @Inject
    LikeService likes

    CommentService(CommentRepository comments) {
        this.comments = comments
    }

    /**
     * This method searches for a comment using it's ID and throws an exception if the comment was not found
     *
     * @param id The identifier of the comment
     * @return object of type {@link CommentEntity}
     * @throw Exception of type {@link pt.atlanse.blog.models.CustomException} if the comment was not found
     * */
    CommentEntity find(Long id, String onError = null) {
        return comments.findById(id).orElseThrow {
            new CustomException(
                "comment not found",
                onError ?: "Article with id $id was not found",
                HttpStatus.NOT_FOUND
            )
        }
    }

    Comment getComment(Long id, String onError = null) {
        CommentEntity commentEntity = comments.findById(id).orElseThrow {
            new CustomException(
                "comment not found",
                onError ?: "Article with id $id was not found",
                HttpStatus.NOT_FOUND
            )
        }

        return new Comment(
            id: commentEntity.id,
            author: commentEntity.createdBy,
            content: commentEntity.text,
            createdAt: commentEntity.createdAt.format(DateTimeFormatter.ISO_LOCAL_DATE_TIME),
            updatedAt: commentEntity.updatedAt.format(DateTimeFormatter.ISO_LOCAL_DATE_TIME),
            likes: likes.count(commentEntity),
            comments: count(commentEntity),
        )
    }

    Comments getComments(Likeable liked, Pageable pageable = null, String user = null) {
        List<Comment> commentList = new ArrayList<>()
        Page<CommentEntity> page = null
        long commentCount = 0

        if (liked instanceof CommentEntity) {
            // It's a comment
            log.info "Getting comments for parent comment"
            page = comments.findAllByParentAndHiddenFalseOrderByCreatedAtDesc(liked, pageable)
            commentCount = comments.countByParent(liked)

        } else if (liked instanceof ArticleEntity) {
            // It's an article
            log.info "Getting comments for parent comment"
            page = comments.findAllByArticleAndHiddenFalseOrderByCreatedAtDesc(liked, pageable)
            commentCount = comments.countByArticle(liked)
        }

        page.toList().each { comment ->
            commentList.add(new Comment(
                id: comment.id,
                author: comment.createdBy,
                content: comment.text,
                createdAt: comment.createdAt.format(DateTimeFormatter.ISO_LOCAL_DATE_TIME),
                updatedAt: comment.updatedAt.format(DateTimeFormatter.ISO_LOCAL_DATE_TIME),
                likes: likes.count(comment),
                comments: count(comment),
                liked: user ? likes.userLikedComment(comment, user) : null
            ))
        }

        return new Comments(
            pageNumber: pageable.getNumber() + 1,
            totalPages: page.totalPages,//	Math.ceil(commentCount / pageable.size),
            totalElements: page.totalSize,
            content: commentList
        )
    }

    /**
     *
     * @return number of comments for the articles or comments
     * */
    int count(Likeable liked) {
        if (liked instanceof CommentEntity) {
            // It's a comment
            log.info "Counting likes for comment"
            return comments.countByParent(liked)

        } else if (liked instanceof ArticleEntity) {
            // It's an article
            log.info "Counting likes for article"
            return comments.countByArticle(liked)
        }
    }

    CommentEntity add(Likeable liked, String content, String who) {
        log.info "Creating comment for Likeable object: ${ liked.toString() }"

        try {
            CommentEntity comment = new CommentEntity()
            comment.setCreatedBy(who)

            if (liked instanceof CommentEntity) {
                comment.setParent(liked)
            } else if (liked instanceof ArticleEntity) {
                comment.setArticle(liked)
            }

            comment.setHidden(false)
            comment.setText(content)
            comment.setUpdatedBy(who)

            return comments.save(comment)
        } catch (Exception e) {
            log.error "Error while trying to create Comment entity; Reason: ${ e.message }"
            throw new CustomException(
                "Error creating Comment",
                "An error occured while creating the Comment entity; More details about the error: ${ e.message }",
                HttpStatus.INTERNAL_SERVER_ERROR
            )
        }

    }

    CommentEntity setCommentHidden(long commentId) {
        CommentEntity commentEntity = comments.findById(commentId).orElseThrow {
            new CustomException(
                "comment not found",
                "Comment with id $commentId was not found",
                HttpStatus.NOT_FOUND
            )
        }
        commentEntity.setHidden(true);
        return comments.update(commentEntity)
    }
}
