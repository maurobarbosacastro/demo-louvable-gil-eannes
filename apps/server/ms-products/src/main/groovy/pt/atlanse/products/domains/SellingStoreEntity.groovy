package pt.atlanse.products.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.NonNull
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank


@Entity
@TupleConstructor
@Table(name = "selling_store")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class SellingStoreEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @Column(name = "name")
    String name

    @Nullable
    @Column(name = "website")
    String website

    @ManyToOne
    @JoinColumn(name = "product_id")
    ProductEntity product
}
