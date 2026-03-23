package pt.atlanse.eduquiz.domain

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Entity
@Table(name = "quiz_questions")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class QuizQuestionsEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @NotNull
    @NotBlank(message = "QuizId can not be blank")
    @ManyToOne
    @JoinColumn(name = "quiz_id", referencedColumnName = "id")
    QuizEntity quiz

    @NotNull
    @NotBlank(message = "QuestionsId can not be blank")
    @ManyToOne
    @JoinColumn(name = "questions_id", referencedColumnName = "id")
    QuestionsEntity questions
}
