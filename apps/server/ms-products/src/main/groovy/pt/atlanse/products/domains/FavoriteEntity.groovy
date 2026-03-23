package pt.atlanse.products.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.NonNull
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
@Table(name = "favorite")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class FavoriteEntity  extends BaseEntity{
    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @ManyToOne
    @JoinColumn(name = "product_id", referencedColumnName = "id")
    ProductEntity productId

    @NonNull
    @NotBlank
    @Column(name = "user_id")
    String userId

}
