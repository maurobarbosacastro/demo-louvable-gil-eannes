package pt.atlanse.blog.domain

import groovy.transform.ToString
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.GeneratedValue
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.Table
import jakarta.persistence.Version
import jakarta.validation.constraints.NotBlank
import java.time.LocalDateTime

@Entity
@Table(name = "likes")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class LikeEntity {

    @Id
    @GeneratedValue(GeneratedValue.Type.AUTO)
    private Long id

    @Version
    private Long version

    @NonNull
    @NotBlank
    @ManyToOne
    @JoinColumn(name = "article_id", referencedColumnName = "id")
    private ArticleEntity article

    @NonNull
    @NotBlank
    @ManyToOne
    @JoinColumn(name = "comment_id", referencedColumnName = "id")
    private CommentEntity comment

    @NonNull
    @NotBlank
    @Column(name = "created_by")
    private String createdBy

    @GeneratedValue
    @Column(name = "created_at")
    private LocalDateTime createdAt

    LikeEntity() {}

    void setId(Long id) {
        this.id = id
    }

    void setVersion(Long version) {
        this.version = version
    }

    void setArticle(@NonNull ArticleEntity article) {
        this.article = article
    }

    void setComment(@NonNull CommentEntity comment) {
        this.comment = comment
    }

    void setCreatedBy(@NonNull String createdBy) {
        this.createdBy = createdBy
    }

    void setCreatedAt(LocalDateTime createdAt) {
        this.createdAt = createdAt
    }

    Long getId() {
        return id
    }

    Long getVersion() {
        return version
    }

    @NonNull
    ArticleEntity getArticle() {
        return article
    }

    @NonNull
    CommentEntity getComment() {
        return comment
    }

    @NonNull
    String getCreatedBy() {
        return createdBy
    }

    LocalDateTime getCreatedAt() {
        return createdAt
    }
}
