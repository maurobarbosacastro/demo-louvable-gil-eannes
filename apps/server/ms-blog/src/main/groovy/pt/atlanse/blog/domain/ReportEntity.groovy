package pt.atlanse.blog.domain

import groovy.transform.ToString
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.GeneratedValue
import jakarta.persistence.ManyToOne

import java.time.LocalDateTime
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.OneToOne
import jakarta.persistence.Table
import jakarta.persistence.Version
import jakarta.validation.constraints.NotBlank

@Entity
@Table(name = "reports")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class ReportEntity {

    @Id
    @GeneratedValue(GeneratedValue.Type.AUTO)
    Long id

    @Version
    Long version

    @NonNull
    @NotBlank
    @Column(name = "reason")
    String reason

    @NonNull
    @NotBlank
    @Column(name = "status")
    String status

    @NonNull
    @NotBlank
    @ManyToOne
    @JoinColumn(name = "article_id", referencedColumnName = "id")
    ArticleEntity article

    @NonNull
    @NotBlank
    @ManyToOne
    @JoinColumn(name = "comment_id", referencedColumnName = "id")
    CommentEntity comment

    @NonNull
    @NotBlank
    @Column(name = "created_by")
    String createdBy

    @GeneratedValue
    @Column(name = "created_at")
    LocalDateTime createdAt

    @NonNull
    @NotBlank
    @Column(name = "updated_by")
    String updatedBy

    @GeneratedValue
    @Column(name = "updated_at")
    LocalDateTime updatedAt
}
