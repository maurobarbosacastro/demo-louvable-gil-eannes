package pt.atlanse.eduquiz.domain

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Entity
@Table(name = "lesson")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class LessonEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @Column(name = "image_id", nullable = true)
    String image

    @NotNull(message = "Title can not be null")
    @NotBlank(message = "Title can not be blank")
    @Column(name = "title")
    String title

    @NotNull(message = "Subtitle can not be null")
    @NotBlank(message = "Subtitle can not be blank")
    @Column(name = "subtitle")
    String subtitle

    @NotNull(message = "Type can not be null")
    @NotBlank(message = "Type can not be blank")
    @Column(name = "type")
    String type

    @Nullable
    @Column(name = "conclusion")
    String conclusion

    @NotNull(message = "Status can not be null")
    @NotBlank(message = "Status can not be blank")
    @Column(name = "status")
    String status
}
