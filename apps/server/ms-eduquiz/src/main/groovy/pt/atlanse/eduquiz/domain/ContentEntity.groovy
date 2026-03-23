package pt.atlanse.eduquiz.domain

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Entity
@Table(name = "content")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class ContentEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @NotNull(message = "ContentPath can not be null")
    @NotBlank(message = "ContentPath can not be blank")
    @Column(name = "content_path")
    String contentPath

    @NotNull(message = "Filename can not be null")
    @NotBlank(message = "Filename can not be blank")
    @Column(name = "file_name")
    String fileName

    @NotNull(message = "Extension can not be null")
    @NotBlank(message = "Extension can not be blank")
    @Column(name = "extension")
    String extension
}
