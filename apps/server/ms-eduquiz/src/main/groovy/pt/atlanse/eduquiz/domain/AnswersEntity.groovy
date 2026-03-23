package pt.atlanse.eduquiz.domain

import com.fasterxml.jackson.annotation.JsonIgnore
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Nullable
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
@Table(name = "answers")
@ToString(includePackage = false, includeNames = true, includeFields = true, excludes = ["question"])
@TupleConstructor
class AnswersEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @JsonIgnore
    @NotNull(message = "QuestionId can not be null")
    @NotBlank(message = "QuestionId can not be blank")
    @ManyToOne
    @JoinColumn(name = "question_id", referencedColumnName = "id")
    QuestionsEntity question

    @Nullable
    @Column(name = "content")
    String content

    @NotNull(message = "isCorrect can not be null")
    @NotBlank(message = "isCorrect can not be blank")
    @Column(name = "is_correct")
    Boolean isCorrect

    @NotNull(message = "Points can not be null")
    @NotBlank(message = "Points can not be blank")
    @Column(name = "points")
    Long points
}
