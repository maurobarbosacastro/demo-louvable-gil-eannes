package pt.atlanse.blog.controllers

import groovy.util.logging.Slf4j
import io.micronaut.core.annotation.NonNull
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
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
import io.swagger.v3.oas.annotations.Operation
import io.swagger.v3.oas.annotations.responses.ApiResponse
import io.swagger.v3.oas.annotations.tags.Tag
import jakarta.inject.Inject
import jakarta.validation.Valid
import jakarta.validation.constraints.NotBlank
import pt.atlanse.blog.DTO.ArticleDTO
import pt.atlanse.blog.DTO.ArticleParameters
import pt.atlanse.blog.domain.ArticleEntity
import pt.atlanse.blog.domain.CommentEntity
import pt.atlanse.blog.domain.LikeEntity
import pt.atlanse.blog.models.Article
import pt.atlanse.blog.models.Articles
import pt.atlanse.blog.models.Comment
import pt.atlanse.blog.models.Comments
import pt.atlanse.blog.models.CustomException
import pt.atlanse.blog.services.ArticleService
import pt.atlanse.blog.services.CommentService
import pt.atlanse.blog.services.LikeService
import java.security.Principal

@Slf4j
@Tag(name = "Articles")
@ExecuteOn(TaskExecutors.IO)
@Controller("/api/articles")
class ArticleController {

    @Inject
    ArticleService articles

    @Inject
    LikeService likes

    @Inject
    CommentService comments

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ ex.toString() }"
        return HttpResponse.status(ex.status).body([
            message: ex.title,
            details: ex.details,
            link   : request.path
        ])
    }

    @Post
    @Operation(summary = "Create new article")
    @ApiResponse(responseCode = "201", description = "Article posted successfully")
    @ApiResponse(responseCode = "400", description = "Payload for article creation is incorrect")
    MutableHttpResponse createArticle(@Body @Valid ArticleDTO article, Principal principal) {
        log.debug("User ${ principal.name } attempting to create new article; $article")
        return HttpResponse.created(articles.create(article, principal.name))
    }

    @Get("{?params*}")
    @Operation(summary = "Gets a list of articles")
    @ApiResponse(responseCode = "200", description = "Retrieves a list of all articles using parameters")
    MutableHttpResponse<Articles> getArticles(ArticleParameters params, @Valid Pageable pageable, @Nullable Principal principal = null) {
        return HttpResponse.ok().body(articles.findAll(params, pageable, principal ? principal.name : null))
    }

    @Get("/{articleId}{?params*}")
    @Operation(summary = "Get article")
    @ApiResponse(responseCode = "200", description = "Retrieves one article using the id and other parameters")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    MutableHttpResponse<Article> getArticle(String articleId, ArticleParameters params, @Nullable Principal principal = null) {
        try {
            return HttpResponse.ok().body(articles.findById(articleId, params, principal ? principal.name : null))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Patch("/{articleId}")
    @Operation(summary = "Patch article")
    @ApiResponse(responseCode = "200", description = "Update article (e.g., changing the status to DRAFT or ARCHIVE)")
    @ApiResponse(responseCode = "400", description = "Payload for article creation is incorrect")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "This endpoint is under maintenance")
    MutableHttpResponse patchArticle(String articleId, @Body @Valid ArticleDTO articlePayload, Principal principal) {
        try {
            return HttpResponse.ok().body(articles.update(articleId, articlePayload, principal.name))
        } catch (Exception ignored) {
            return HttpResponse.notFound()
        }
    }

    @Post("/{articleId}/comments")
    @Operation(summary = "Create comment for the article")
    @ApiResponse(responseCode = "201", description = "Successfully created comment")
    @ApiResponse(responseCode = "400", description = "Payload for comment creation is incorrect")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "This endpoint is under maintenance")
    MutableHttpResponse createComment(Long articleId, @Header(name = "Authorization") String authorization, @NonNull @NotBlank String content, Principal principal) {
        ArticleEntity article = articles.find(articleId)
        CommentEntity comment = comments.add(article, content, principal.name)
        Comment comment1 = comments.getComment(comment.id, comment.createdBy)

        return HttpResponse.ok(comment1)
    }

    @Get("/{articleId}/comments")
    @Operation(summary = "Get comments for the article")
    @ApiResponse(responseCode = "200", description = "Retrieves comments for the article with the specified ID")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "This endpoint is under maintenance")
    MutableHttpResponse getUserComments(Long articleId, @Header(name = "Authorization") String authorization, Pageable pageable, Principal principal) {
        ArticleEntity article = articles.find(articleId)
        Comments commentList = comments.getComments(article, pageable, principal.name)
        Map<String, Object> profiles = new HashMap<>()

        commentList.content.each {
            it.creator = profiles.get(it.author)
        }

        return HttpResponse.ok(commentList)
    }

    @Get("/{articleId}/comments/total")
    @Operation(summary = "Get amount of comments for the article")
    @ApiResponse(responseCode = "200", description = "Retrieves amount of comments for the article with the specified ID")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "This endpoint is under maintenance")
    MutableHttpResponse getCommentCount(Long articleId) {
        ArticleEntity article = articles.find(articleId)
        return HttpResponse.ok([comments: comments.count(article)])
    }

    @Get("/{articleId}/likes")
    @Operation(summary = "Get likes for the article")
    @ApiResponse(responseCode = "200", description = "Retrieves likes for the article with the specified ID")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "This endpoint is under maintenance")
    MutableHttpResponse getLikes(Long articleId) {
        ArticleEntity article = articles.find(articleId)
        return HttpResponse.ok([likes: likes.getLikes(article)])
    }

    @Get("/{articleId}/likes/total")
    @Operation(summary = "Get amount of likes for the article")
    @ApiResponse(responseCode = "200", description = "Retrieves amount of likes for the article with the specified ID")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "This endpoint is under maintenance")
    MutableHttpResponse getLikeCount(Long articleId) {
        ArticleEntity article = articles.find(articleId)
        return HttpResponse.ok([likes: likes.count(article)])
    }

    @Post("/{articleId}/likes")
    @Operation(summary = "Add like on the article")
    @ApiResponse(responseCode = "201", description = "Create Like for the article with the specified ID")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    @ApiResponse(responseCode = "503", description = "This endpoint is under maintenance")
    MutableHttpResponse createLike(Long articleId, Principal principal) {
        ArticleEntity article = articles.find(articleId)
        LikeEntity like = likes.add(article, principal.name)
        return HttpResponse.ok()
    }

    @Delete("/{articleId}/likes")
    MutableHttpResponse removeLike(Long articleId, Principal principal) {
        ArticleEntity article = articles.find(articleId)
        likes.del(article, principal.name)
        return HttpResponse.ok()
    }

    // todo clean this shits
    @Get("/random{?params*}")
    @Operation(summary = "Get random article")
    @ApiResponse(responseCode = "200", description = "Retrieves one article using the id and other parameters")
    @ApiResponse(responseCode = "404", description = "The article using the specified ID was not found")
    MutableHttpResponse randomArticle(ArticleParameters params, @Nullable Principal principal = null, Pageable pageable) {

        if (pageable.size >= 3) {
            List<Article> article = articles.findRandomArticles(params, principal ? principal.name : null)
            return HttpResponse.ok(article)
        }

        Article article = articles.findRandomArticle(params, principal ? principal.name : null)
        return HttpResponse.ok(article)
    }

}
