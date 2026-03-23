package pt.atlanse.products.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated

import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.ManyToMany
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank

@Entity
@TupleConstructor
@Table(name = "ingredient")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class IngredientEntity extends BaseEntity {

    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @Column(name = "name")
    String name

    @ManyToMany(mappedBy = 'ingredients')
    Set<ProductEntity> products = new HashSet<>()

}
