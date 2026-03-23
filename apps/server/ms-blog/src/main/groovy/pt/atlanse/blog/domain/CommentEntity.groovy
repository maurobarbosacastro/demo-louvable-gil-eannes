package pt.atlanse.blog.domain

import groovy.transform.ToString
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.GeneratedValue
import jakarta.persistence.ManyToOne
import pt.atlanse.blog.models.Likeable
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.Table
import jakarta.persistence.Version
import jakarta.validation.constraints.NotBlank
import java.time.LocalDateTime

@Entity
@Table(name = "comments")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class CommentEntity implements Likeable {

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
    @JoinColumn(name = "parent_id", referencedColumnName = "id")
    private CommentEntity parent

    @NonNull
    @NotBlank
    @Column(name = "hidden")
    private boolean hidden

    @NonNull
    @NotBlank
    @Column(name = "text")
    private String text

    @NonNull
    @NotBlank
    @Column(name = "created_by")
    private String createdBy

    @GeneratedValue
    @Column(name = "created_at")
    private LocalDateTime createdAt

    @NonNull
    @NotBlank
    @Column(name = "updated_by")
    private String updatedBy

    @GeneratedValue
    @Column(name = "updated_at")
    private LocalDateTime updatedAt

    CommentEntity() {}

    void setId(Long id) {
        this.id = id
    }

    void setVersion(Long version) {
        this.version = version
    }

    void setArticle(@NonNull ArticleEntity article) {
        this.article = article
    }

    void setParent(@NonNull CommentEntity parent) {
        this.parent = parent
    }

    void setHidden(boolean hidden) {
        this.hidden = hidden
    }

    void setText(@NonNull String text) {
        this.text = text
    }

    void setCreatedBy(@NonNull String createdBy) {
        this.createdBy = createdBy
    }

    void setCreatedAt(LocalDateTime createdAt) {
        this.createdAt = createdAt
    }

    void setUpdatedBy(@NonNull String updatedBy) {
        this.updatedBy = updatedBy
    }

    void setUpdatedAt(LocalDateTime updatedAt) {
        this.updatedAt = updatedAt
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
    CommentEntity getParent() {
        return parent
    }

    boolean getHidden() {
        return hidden
    }

    @NonNull
    String getText() {
        return text
    }

    @NonNull
    String getCreatedBy() {
        return createdBy
    }

    LocalDateTime getCreatedAt() {
        return createdAt
    }

    @NonNull
    String getUpdatedBy() {
        return updatedBy
    }

    LocalDateTime getUpdatedAt() {
        return updatedAt
    }
}
