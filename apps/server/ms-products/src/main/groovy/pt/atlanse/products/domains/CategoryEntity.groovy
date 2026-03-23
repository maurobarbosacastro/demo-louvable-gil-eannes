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
@Table(name = "category")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class CategoryEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @Column(name = "name")
    String name

    @Column(name = "description")
    String description

    @Column(name = "state")
    String state

    @Column(name = "image_uuid")
    UUID image

    CategoryEntity(){}

    public UUID getId(){
        return this.id
    }
}
