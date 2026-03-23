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
@Table(name = "discount")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class DiscountEntity  extends BaseEntity{
    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @Column(name = "code")
    String code

    @NonNull
    @NotBlank
    @Column(name = "active")
    Boolean active

    @NonNull
    @NotBlank
    @Column(name = "amount")
    String amount

}
