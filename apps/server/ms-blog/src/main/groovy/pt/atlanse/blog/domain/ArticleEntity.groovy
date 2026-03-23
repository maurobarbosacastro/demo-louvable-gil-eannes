package pt.atlanse.blog.domain

import groovy.transform.ToString
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.GeneratedValue
import pt.atlanse.blog.domain.imageprocessing.ArticleMediaFile
import pt.atlanse.blog.models.Likeable
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.OneToOne
import jakarta.persistence.Table
import jakarta.persistence.Version
import jakarta.validation.constraints.NotBlank
import java.time.LocalDateTime

@Entity
@Table(name = "articles")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class ArticleEntity implements Likeable {

    @Id
    @GeneratedValue(GeneratedValue.Type.AUTO)
    private Long id

    @Version
    private Long version

    @NonNull
    @NotBlank
    @OneToOne
    @JoinColumn(name = "article_media_file", referencedColumnName = "id")
    ArticleMediaFile articleMediaFile

    @NonNull
    @NotBlank
    @Column(name = "status")
    private String status

    @NonNull
    @NotBlank
    @Column(name = "view_mobile")
    private boolean viewMobile

    @NonNull
    @NotBlank
    @Column(name = "view_web")
    private boolean viewWeb

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

    @NonNull
    @NotBlank
    @OneToOne
    @JoinColumn(name = "category", referencedColumnName = "id")
    Category category

    ArticleEntity() {}

    void setCategory(@NonNull Category category) {
        this.category = category
    }

    @NonNull
    Category getCategory() {
        return category
    }

    void setId(Long id) {
        this.id = id
    }

    void setVersion(Long version) {
        this.version = version
    }

    void setArticleMediaFile(@NonNull ArticleMediaFile articleMediaFile) {
        this.articleMediaFile = articleMediaFile
    }

    void setStatus(@NonNull String status) {
        this.status = status
    }

    void setViewMobile(boolean viewMobile) {
        this.viewMobile = viewMobile
    }

    void setViewWeb(boolean viewWeb) {
        this.viewWeb = viewWeb
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
    ArticleMediaFile getArticleMediaFile() {
        return articleMediaFile
    }

    @NonNull
    String getStatus() {
        return status
    }

    boolean getViewMobile() {
        return viewMobile
    }

    boolean getViewWeb() {
        return viewWeb
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
