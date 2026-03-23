package pt.atlanse.eduquiz.domain

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Entity
@Table(name = "category")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class CategoryEntity extends BaseEntity {
    @Id
    String id

    @NotNull
    @NotBlank(message = "Name can not be blank")
    @Column(name = "name")
    String name
}
