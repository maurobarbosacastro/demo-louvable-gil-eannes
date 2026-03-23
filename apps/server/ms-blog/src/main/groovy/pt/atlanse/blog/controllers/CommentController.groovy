package pt.atlanse.blog.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.HttpStatus
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.http.annotation.Error
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Header
import io.micronaut.http.annotation.Patch
import io.micronaut.http.annotation.Post
import io.micronaut.scheduling.TaskExecutors
import io.micronaut.scheduling.annotation.ExecuteOn
import io.micronaut.security.annotation.Secured
import io.micronaut.security.rules.SecurityRule
import io.swagger.v3.oas.annotations.Operation
import io.swagger.v3.oas.annotations.responses.ApiResponse
import io.swagger.v3.oas.annotations.tags.Tag
import jakarta.inject.Inject
import jakarta.validation.Valid
import jakarta.validation.constraints.NotBlank
import pt.atlanse.blog.domain.CommentEntity
import pt.atlanse.blog.models.Comment
import pt.atlanse.blog.models.Comments
import pt.atlanse.blog.models.CustomException
import pt.atlanse.blog.models.Report
import pt.atlanse.blog.models.Reports
import pt.atlanse.blog.services.CommentService
import pt.atlanse.blog.services.LikeService
import pt.atlanse.blog.services.ReportService

import java.security.Principal

@ExecuteOn(TaskExecutors.IO)
@Slf4j
@Tag(name = "Comments")
@Controller("/api/comments")
class CommentController {

    @Inject
    CommentService comments

    @Inject
    LikeService likes

    @Inject
    ReportService reports

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ex.toString()}"
        return HttpResponse.status(ex.status).body([
                message: ex.title,
                details: ex.details,
                link   : request.path
        ])
    }

    @Secured("create-comment")
    @Post("/{commentId}/comments")
    @Operation(summary = "Create comment for the article")
    @ApiResponse(responseCode = "201", description = "Successfully created comment")
    @ApiResponse(responseCode = "400", description = "Payload for comment creation is incorrect")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "This endpoint is under maintenance")
    MutableHttpResponse createComment(Long commentId, @NonNull @NotBlank String content, Principal principal, @Header(name = "Authorization") String authorization) {
        CommentEntity comment = comments.find(commentId)
        CommentEntity reply = comments.add(comment, content, principal.name)
        Comment comment1 = comments.getComment(reply.id, reply.createdBy)
        comment1.creator = principal.name
        return HttpResponse.ok(comment1)
    }

    @Secured("read-comment")
    @Get("/{commentId}/comments")
    @Operation(summary = "Get comments for the article")
    @ApiResponse(responseCode = "200", description = "Retrieves comments for the article with the specified ID")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "This endpoint is under maintenance")
    HttpResponse getUserComments(Long commentId, @Header(name = "Authorization") String authorization, Pageable pageable, Principal principal) {
        CommentEntity comment = comments.find(commentId)
        Comments commentList = comments.getComments(comment, pageable, principal.name)
        Map<String, Object> profiles = new HashMap<>()

        // Key: Keycloak user ID; Value: User profile
        commentList.content*.author.each {
            profiles.put(it, principal.name)
        }

        commentList.content.each {
            it.creator = profiles.get(it.author)
        }

        return HttpResponse.ok(commentList)
    }

    @Secured("read-comment")
    @Get("/{commentId}/likes")
    @Operation(summary = "Get list of likes for the comment", description = "Collects a list of users that liked the comment")
    @ApiResponse(responseCode = "200", description = "Retrieves amount of likes for the article with the specified ID")
    @ApiResponse(responseCode = "404", description = "The comment using the specified ID was not found")
    MutableHttpResponse getLikes(Long commentId) {
        CommentEntity comment = comments.find(commentId)
        return HttpResponse.ok([likes: likes.getLikes(comment)])
    }

    @Secured("read-comment")
    @Get("/{commentId}/likes/total")
    @Operation(summary = "Get amount of likes for the comment")
    @ApiResponse(responseCode = "200", description = "Retrieves amount of likes for the comment with the specified ID")
    @ApiResponse(responseCode = "404", description = "The comment using the specified ID was not found")
    MutableHttpResponse getLikeCount(Long commentId) {
        CommentEntity comment = comments.find(commentId)
        return HttpResponse.ok([likes: likes.count(comment)])
    }

    @Patch("/{commentId}")
    @Operation(summary = "Update the HIDDEN field of the comment")
    @ApiResponse(responseCode = "200", description = "Update comment (e.g., changing the HIDDEN field to true)")
    @ApiResponse(responseCode = "400", description = "Payload for comment manipulation is incorrect")
    @ApiResponse(responseCode = "404", description = "The comment using the specified ID was not found")
    MutableHttpResponse patchArticle(String commentId, @Body @Valid Object articlePayload) {
        return HttpResponse.status(HttpStatus.SERVICE_UNAVAILABLE)
    }

    @Secured("read-comment")
    @Post("/{commentId}/likes")
    @Operation(summary = "Add like on the comment")
    @ApiResponse(responseCode = "201", description = "Create Like for the comment with the specified ID")
    @ApiResponse(responseCode = "404", description = "The comment using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "The service is currently unavailable")
    MutableHttpResponse createLike(Long commentId, Principal principal) {
        CommentEntity comment = comments.find(commentId)
        likes.add(comment, principal.name)
        return HttpResponse.ok()
    }

    @Secured("read-comment")
    @Delete("/{commentId}/likes")
    @Operation(summary = "Delete like on the comment")
    @ApiResponse(responseCode = "201", description = "Delete Like for the comment with the specified ID")
    @ApiResponse(responseCode = "404", description = "The comment using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "The service is currently unavailable")
    MutableHttpResponse deleteLike(Long commentId, Principal principal) {
        CommentEntity comment = comments.find(commentId)
        likes.del(comment, principal.name)
        return HttpResponse.ok()
    }



    //TODO change role
    @Secured("report-comment")
    @Get("/reports")
    @Operation(summary = "Gets reported comment with a few more for context")
    @ApiResponse(
            responseCode = "200",
            description = "Retrieves list of reports with the users and comment information and some comments as context"
    )
    MutableHttpResponse getReported(@Header(name = "Authorization") String authorization, Pageable page, Principal principal) {

        try {
            Reports reportList = reports.getReports(page, 'Pending', principal)

            if (reportList.reports.isEmpty()) return HttpResponse.ok("{}")

            return HttpResponse.ok(reportList)

        } catch (Exception e) {
            return HttpResponse.notFound()
        }
    }

    @Secured(SecurityRule.IS_ANONYMOUS)
    @Post("/{commentId}/reported")
    MutableHttpResponse reportComment(Long commentId, @NonNull @NotBlank String reason, Principal principal) {

        try {
            CommentEntity comment = comments.find(commentId)
            Report report = reports.add(comment, reason, 'Pending', principal.name)
            return HttpResponse.ok(report)

        } catch (Exception e) {
            return HttpResponse.notFound()
        }
    }

    @Secured(SecurityRule.IS_ANONYMOUS)
    @Patch("/reported/{reportId}")
    MutableHttpResponse updateReportStatus(Long reportId, @NonNull @NotBlank String status) {
        try {
            //Update report status
            reports.updateReportStatus(reportId, status);
            return HttpResponse.ok()

        } catch (Exception e) {
            return HttpResponse.notFound()
        }
    }
}
