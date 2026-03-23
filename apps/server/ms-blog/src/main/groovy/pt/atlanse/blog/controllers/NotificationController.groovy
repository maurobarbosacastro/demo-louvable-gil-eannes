package pt.atlanse.blog.controllers
import groovy.util.logging.Slf4j
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpResponse
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Header
import io.micronaut.scheduling.TaskExecutors
import io.micronaut.scheduling.annotation.ExecuteOn
import io.micronaut.security.annotation.Secured
import io.micronaut.security.rules.SecurityRule
import io.swagger.v3.oas.annotations.tags.Tag
import jakarta.inject.Inject
import pt.atlanse.blog.models.Notification
import pt.atlanse.blog.services.CommentService
import pt.atlanse.blog.services.LikeService
import pt.atlanse.blog.services.NotificationService

import java.security.Principal

@Slf4j
@Tag(name = "Comments")
@ExecuteOn(TaskExecutors.IO)
@Controller("/api/notifications")
class NotificationController {

    @Inject
    CommentService comments

    @Inject
    LikeService likes

    @Inject
    NotificationService notificationService

    @Secured(SecurityRule.IS_AUTHENTICATED)
    @Get("/")
    HttpResponse getUserComments(@Header(name = "Authorization") String authorization, Pageable pageable, Principal principal) {

        /*---------------------------------------------------------------------------------------*/
        //User likes and comments

        //Get user likes on comments and articles
        List<Notification> listLikes = notificationService.getLikesByUser(principal.name);

        //Get user comments and replies
        List<Notification> listComments = notificationService.getCommentsByUser(principal.name)
        Object userProfile = principal.name

        List<Notification> responseList = []
        responseList.addAll(listLikes)
        responseList.addAll(listComments)

        responseList.each {
            it.creator = userProfile
        }


        /*---------------------------------------------------------------------------------------*/


        /*---------------------------------------------------------------------------------------*/
        //Interactions with user comments
        //Find likes on the user comment
        List<Notification> likesOnComments = notificationService.getLikesOnUser(principal.name)
        likesOnComments.each {
            it.creator = principal.name
        }

        //Find reply on user comments
        List<Notification> commentsOnComments = notificationService.getCommentsOnUser(principal.name)
        commentsOnComments.each {
            it.creator = principal.name
        }

        responseList.addAll(likesOnComments)
        responseList.addAll(commentsOnComments)

        responseList.sort{a,b-> b.createdAt<=>a.createdAt}
        int topIndex = 50;
        if (responseList.size() > topIndex) {
            responseList.subList(0, topIndex)
        }

        Map<String, Object> response = new HashMap<>()
        response.put('content', responseList)

        return HttpResponse.ok(response)
    }

}
