package pt.atlanse.eduquiz.domain

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.data.annotation.AutoPopulated
import pt.atlanse.eduquiz.utils.LessonOrQuizNotNull

import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Entity
@Table(name = "modules_order")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
@LessonOrQuizNotNull
class ModulesOrderEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @NotNull(message = "ModuleId can not be blank")
    @NotBlank(message = "ModuleId can not be blank")
    @ManyToOne
    @JoinColumn(name = "module_id", referencedColumnName = "id")
    ModulesEntity module

    @ManyToOne
    @JoinColumn(name = "lesson_id", referencedColumnName = "id")
    LessonEntity lesson

    @ManyToOne
    @JoinColumn(name = "quiz_id", referencedColumnName = "id")
    QuizEntity quiz

    @NotNull(message = "Position can not be null")
    @NotBlank(message = "Position can not be blank")
    @Column(name = "position")
    Long position
}
