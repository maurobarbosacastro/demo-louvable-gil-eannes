package pt.atlanse.eduquiz.domain

import com.fasterxml.jackson.annotation.JsonFormat
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.CascadeType
import jakarta.persistence.FetchType
import jakarta.persistence.OneToMany
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull
import java.time.LocalDateTime

@Entity
@Table(name = "course")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class CourseEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @Column(name = "image_id")
    String image

    @NotNull
    @NotBlank(message = "Title must not be blank")
    @Column(name = "title")
    String title

    @Nullable
    @Column(name = "description")
    String description

    @NotNull
    @NotBlank(message = "Status must not be blank")
    @Column(name = "status")
    String status

    @Nullable
    @Column(name = "begin_date")
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    LocalDateTime beginDate

    @Nullable
    @Column(name = "end_date")
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    LocalDateTime endDate

    @Nullable
    @Column(name = "extras")
    String extras

    @Nullable
    @OneToMany(mappedBy = "course", cascade = [CascadeType.ALL], fetch = FetchType.EAGER)
    List<CourseOrderEntity> courseOrder = new ArrayList<>()
}
