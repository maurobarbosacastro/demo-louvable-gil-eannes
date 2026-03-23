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
@Table(name = "rules")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class RulesEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @NotNull(message = 'Code can not be null')
    @NotBlank(message = "Code can not be blank")
    @Column(name = "code")
    String code

    @NotNull(message = 'Value can not be null')
    @NotBlank(message = "Value can not be blank")
    @Column(name = "value")
    String value

    @NotNull(message = 'Title can not be null')
    @NotBlank(message = "Title can not be blank")
    @Column(name = "title")
    String title

    @NotNull(message = 'Description can not be null')
    @NotBlank(message = "Description can not be blank")
    @Column(name = "description")
    String description
}
