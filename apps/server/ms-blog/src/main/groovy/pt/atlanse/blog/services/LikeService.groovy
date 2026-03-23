package pt.atlanse.blog.services

import groovy.util.logging.Slf4j
import io.micronaut.http.HttpStatus
import jakarta.inject.Singleton
import pt.atlanse.blog.domain.ArticleEntity
import pt.atlanse.blog.domain.CommentEntity
import pt.atlanse.blog.domain.LikeEntity
import pt.atlanse.blog.models.CustomException
import pt.atlanse.blog.models.Likeable
import pt.atlanse.blog.repository.LikeRepository

@Slf4j
@Singleton
class LikeService {
    LikeRepository likes

    LikeService(LikeRepository likes) {
        this.likes = likes
    }

    /**
     * Check if user likes the article
     * @param user The user used to verify if the like exists
     * @param The article entity of type {@link ArticleEntity}
     * @return boolean
     * */
    boolean userLikedArticle(ArticleEntity article, String user) {
        likes.findByArticleAndCreatedBy(article, user).isPresent()
    }

    boolean userLikedComment(CommentEntity comment, String user) {
        likes.findByCommentAndCreatedBy(comment, user).isPresent()
    }

    /**
     * Collects all the likes for the {@link Likeable} entity
     * @param liked {@link Likeable} object
     * @return list of likes for the {@link Likeable} entity
     * */
    List<LikeEntity> getLikes(Likeable liked) {
        if (liked instanceof CommentEntity) {
            // It's a comment
            log.info "Getting likes for comment"
            return likes.findAllByComment(liked)
        }

        // It's an article
        log.info "Getting likes for article"
        return likes.findAllByArticle(liked as ArticleEntity)
    }

    /**
     * Count all the likes for the {@link Likeable} entity
     * @param liked {@link Likeable} object
     * @return number of likes for the {@link Likeable} entity
     * */
    int count(Likeable liked) {
        if (liked instanceof CommentEntity) {
            // It's a comment
            log.info "Counting likes for comment"
            return likes.countByComment(liked)
        }

        // It's an article
        log.info "Counting likes for article"
        return likes.countByArticle(liked as ArticleEntity)
    }

    /**
     * Builds a {@link LikeEntity} object for comments or articles
     * @param liked {@link Likeable} object (e.g., object of type {@link CommentEntity} or {@link ArticleEntity}
     * @param who The author of this like
     * @return A fully built object of the {@link pt.atlanse.blog.domain.LikeEntity} class
     * */
    private LikeEntity fillLikeEntity(Likeable liked, String who) {
        LikeEntity like = new LikeEntity()
        like.setCreatedBy(who)

        if (liked instanceof CommentEntity) {
            like.setComment(liked)
            return like
        }

        if (likes.findByArticleAndCreatedBy(liked as ArticleEntity, who)) {
            throw new CustomException(
                "User already like this article",
                "This article has already 1 like from this user",
                HttpStatus.CONFLICT
            )
        }

        like.setArticle(liked as ArticleEntity)

        return like
    }

    /**
     * This method creates the {@link LikeEntity}
     * @return the created {@link LikeEntity} entity
     * */
    LikeEntity add(Likeable liked, String who) {
        log.info "Creating like for Likeable object: ${ liked.toString() }"

        LikeEntity like = fillLikeEntity(liked, who)

        try {
            return likes.save(like)
        } catch (Exception e) {
            log.error "Error while trying to create Like entity; Reason: ${ e.message }"
            throw new CustomException(
                "Error creating Like",
                "An error occured while creating the Like entity; More details about the error: ${ e.message }",
                HttpStatus.INTERNAL_SERVER_ERROR
            )
        }
    }

    /**
     * This method removes a like {@link LikeEntity}
     * */
    void del(Likeable liked, String user) {
        if (liked instanceof CommentEntity) {
            log.info "Removing like from comment ${ liked.id }..."
            likes.findByCommentAndCreatedBy(liked, user).ifPresent {
                log.info "Found like for user $user, comment: ${ liked.id }..."
                likes.delete(it)
            }
        }
        if (liked instanceof ArticleEntity) {
            log.info "Removing like from article ${ liked.id }..."
            likes.findByArticleAndCreatedBy(liked, user).ifPresent {
                log.info "Found like for user $user, article: ${ liked.id }..."
                likes.delete(it)
            }
        }
    }
}
