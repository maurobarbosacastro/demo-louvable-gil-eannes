package pt.atlanse.products.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank


@Entity
@TupleConstructor
@Table(name = "rating")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class RatingEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @Column(name = "amount", nullable = false)
    Long amount
}
