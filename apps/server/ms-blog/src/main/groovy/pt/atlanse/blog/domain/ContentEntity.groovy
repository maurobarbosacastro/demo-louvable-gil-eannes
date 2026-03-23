package pt.atlanse.blog.domain

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.NonNull
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.annotation.GeneratedValue
import jakarta.validation.constraints.NotNull
import pt.atlanse.blog.DTO.ImageDTO
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import java.time.LocalDateTime

@Entity
@Table(name = "content")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class ContentEntity {
    @Id
    @GeneratedValue(GeneratedValue.Type.AUTO)
    Long id

    @NotNull
    @Column(name = "content_path")
    String contentPath

    @Nullable
    @Column(name = "file_name")
    String fileName

    @Nullable
    @Column(name = "extension")
    String extension

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

    ImageDTO toImageDTO() {
        return new ImageDTO(
            fileName: fileName,
            base64: contentPath,
            extension: extension
        )
    }
}
