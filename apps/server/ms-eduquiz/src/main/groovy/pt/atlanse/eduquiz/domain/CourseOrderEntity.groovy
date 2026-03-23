package pt.atlanse.eduquiz.domain

import com.fasterxml.jackson.annotation.JsonIgnore
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Entity
@Table(name = "course_order")
@ToString(includePackage = false, includeNames = true, includeFields = true, excludes = ["course"])
@TupleConstructor
class CourseOrderEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @NotNull(message = "CourseId can not be null")
    @NotBlank(message = "CourseId can not be blank")
    @ManyToOne
    @JsonIgnore
    @JoinColumn(name = "course_id", referencedColumnName = "id")
    CourseEntity course

    @NotNull(message = "ModuleId can not be null")
    @NotBlank(message = "ModuleId can not be blank")
    @ManyToOne
    @JoinColumn(name = "module_id", referencedColumnName = "id")
    ModulesEntity module

    @NotNull(message = "Position can not be null")
    @NotBlank(message = "Position can not be blank")
    @Column(name = "position")
    Long position
}
