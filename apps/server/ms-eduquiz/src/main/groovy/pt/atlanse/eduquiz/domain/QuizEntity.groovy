package pt.atlanse.eduquiz.domain

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
@Table(name = "quiz")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class QuizEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @Nullable
    @ManyToOne
    @JoinColumn(name = "category_id", referencedColumnName = "id")
    CategoryEntity category

    @NotNull(message = "Title can not be null")
    @NotBlank(message = "Title can not be blank")
    @Column(name = "title")
    String title

    @NotNull(message = "Description can not be null")
    @NotBlank(message = "Description can not be blank")
    @Column(name = "description")
    String description

    @NotNull(message = "Random can not be null")
    @NotBlank(message = "Random can not be blank")
    @Column(name = "random")
    Boolean random

    @Nullable
    String extras
}
