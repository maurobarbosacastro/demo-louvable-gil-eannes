package pt.atlanse.products.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.OneToOne
import jakarta.persistence.Table


@Entity
@TupleConstructor
@Table(name = "review")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class ReviewEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @Column(name = "content", nullable = false)
    String content

    @OneToOne
    @JoinColumn(name = "rating_id", referencedColumnName = "id", nullable = false)
    RatingEntity rating

    @ManyToOne
    @JoinColumn(name = "product_id", referencedColumnName = "id", nullable = false)
    ProductEntity product
}
