package pt.atlanse.blog.services

import groovy.util.logging.Slf4j
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.blog.DTO.ArticleParameters
import pt.atlanse.blog.domain.CommentEntity
import pt.atlanse.blog.domain.LikeEntity
import pt.atlanse.blog.models.Article
import pt.atlanse.blog.models.Comment
import pt.atlanse.blog.models.Notification
import pt.atlanse.blog.repository.CommentRepository
import pt.atlanse.blog.repository.LikeRepository

@Slf4j
@Singleton
class NotificationService {

    @Inject
    ArticleService articles

    @Inject
    CommentService commentsService

    @Inject
    TranslationService translationService

    LikeRepository likes
    CommentRepository comments

    NotificationService(LikeRepository likes, CommentRepository comments) {
        this.likes = likes
        this.comments = comments
    }

    List<Notification> getLikesByUser(String who) {
        List<LikeEntity> likeEntityList = likes.findTop50ByCreatedByOrderByCreatedAtDesc(who)
        List<LikeEntity> likeEntityListComments = likes.findTop50ByCreatedByAndCommentHiddenFalseOrderByCreatedAtDesc(who)
        List<Notification> notificationList = []
        ArticleParameters params = new ArticleParameters(
            image: false
        )
        likeEntityList.each {
            if (it.article != null) {
                Article articleData = articles.findById(it.article.id as String, params, who)
                Article article = new Article(
                    id: articleData.id,
                    title: articleData.title,
                )
                Notification notification = new Notification()
                notification.action = 'likeArticle'
                notification.article = article
                notification.createdAt = it.createdAt
                notificationList.add(notification)
            }
        }

        likeEntityListComments.each {
            if (it.comment != null) {
                CommentEntity commentEntity = commentsService.find(it.comment.id)
                if (commentEntity.createdBy != who) {
                    Article articleData
                    if (commentEntity.article != null) {
                        articleData = articles.findById(commentEntity.article.id as String, params, who)
                    } else {
                        CommentEntity comment = commentsService.find(commentEntity.parent.id)
                        articleData = articles.findById(comment.article.id as String, params, who)
                    }
                    Article article = new Article(
                        id: articleData.id,
                        title: articleData.title,
                    )
                    Notification notification = new Notification()
                    notification.action = 'likeComment'
                    notification.article = article
                    notification.createdAt = it.createdAt
                    notificationList.add(notification)
                }

            }

        }

        return notificationList
    }

    List<Notification> getCommentsByUser(String who) {
        List<CommentEntity> commentEntityList = comments.findTop50ByCreatedByAndHiddenFalseOrderByCreatedAtDesc(who)
        List<Notification> notificationList = []
        ArticleParameters params = new ArticleParameters(
            image: false
        )

        commentEntityList.each {

            Comment comment = new Comment(
                id: it.id,
                content: it.text,
            )

            if (it.article != null) {
                Article articleData = articles.findById(it.article.id as String, params, who)
                Article article = new Article(
                    id: articleData.id,
                    title: articleData.title,
                )

                Notification notification = new Notification()
                notification.action = 'commentArticle'
                notification.article = article
                notification.comment = comment
                notification.createdAt = it.createdAt
                notificationList.add(notification)
            }
            if (it.parent != null) {
                CommentEntity commentEntity = commentsService.find(it.parent.id)
                Article articleData = articles.findById(commentEntity.article.id as String, params, who)
                Article article = new Article(
                    id: articleData.id,
                    title: articleData.title,
                )
                Notification notification = new Notification()
                notification.action = 'replyComment'
                notification.article = article
                notification.comment = comment
                notification.createdAt = it.createdAt
                notificationList.add(notification)

            }

        }

        return notificationList;
    }

    List<Notification> getLikesOnUser(String who) {
        List<LikeEntity> likeEntityList = likes.findTop50ByCommentCreatedByAndCommentHiddenFalseOrderByCreatedAt(who)
        List<Notification> notificationList = [];

        ArticleParameters params = new ArticleParameters(
            image: false
        )
        likeEntityList.each {
            if (it.comment != null) {
                Article articleData
                if (it.comment.article != null) {
                    articleData = articles.findById(it.comment.article.id as String, params, who)
                } else if (it.comment.parent != null) {
                    CommentEntity commentEntity = commentsService.find(it.comment.parent.id)
                    articleData = articles.findById(commentEntity.article.id as String, params, who)
                } else {
                    return
                }

                Article article = new Article(
                    id: articleData.id,
                    title: articleData.title,
                )
                Notification notification = new Notification()
                notification.action = 'likeUserComment'
                notification.article = article
                notification.createdAt = it.createdAt
                notification.createdBy = it.createdBy
                notificationList.add(notification)
            }
        }
        return notificationList;
    }

    List<Notification> getCommentsOnUser(String who) {
        List<CommentEntity> commentEntityList = comments.findTop50ByParentCreatedByAndHiddenFalseOrderByCreatedAt(who)
        List<Notification> notificationList = []
        ArticleParameters params = new ArticleParameters(
            image: false
        )

        commentEntityList.each {
            CommentEntity commentEntity = commentsService.find(it.parent.id)
            if (it.createdBy != who) {
                Article articleData = articles.findById(commentEntity.article.id as String, params, who)
                Article article = new Article(
                    id: articleData.id,
                    title: articleData.title,
                )
                Notification notification = new Notification()
                notification.action = 'commentUserComment'
                notification.article = article
                notification.comment = new Comment(
                    id: it.id,
                    content: it.text
                )
                notification.createdAt = it.createdAt
                notification.createdBy = it.createdBy
                notificationList.add(notification)
            }
        }
        return notificationList;
    }
}
