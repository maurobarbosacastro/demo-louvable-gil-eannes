package pt.atlanse.products.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.CascadeType
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.FetchType
import jakarta.persistence.Id
import jakarta.persistence.ManyToMany
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank


@Introspected
@Entity
@TupleConstructor
@Table(name = "extras")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class ExtrasEntity {

    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @Column(name = "name")
    String name

    @NonNull
    @NotBlank
    @Column(name = "price")
    Float price

    @Column(name = "description")
    String description

    @NonNull
    @NotBlank
    @Column(name = "state")
    String state

    @Column(name = "image_uuid")
    UUID image

    @ManyToMany(mappedBy = 'extras', fetch = FetchType.EAGER, cascade = CascadeType.ALL)
    Set<ProductEntity> products = new HashSet<>()

}

