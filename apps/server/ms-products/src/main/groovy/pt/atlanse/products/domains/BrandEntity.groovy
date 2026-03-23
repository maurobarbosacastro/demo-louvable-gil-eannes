package pt.atlanse.products.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank


@Introspected
@Entity
@TupleConstructor
@Table(name = "brand")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class BrandEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @Column(name = "name")
    String name

    @Column(name = "description")
    String description

    @Column(name = "website")
    String website

    @Column(name = "state")
    String state

    @Column(name = "image_uuid")
    UUID image

    @Column(name = "banner_uuid")
    UUID banner
}
