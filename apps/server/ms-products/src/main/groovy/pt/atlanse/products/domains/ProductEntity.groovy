package pt.atlanse.products.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.CascadeType
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.FetchType
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.JoinTable
import jakarta.persistence.ManyToMany
import jakarta.persistence.ManyToOne
import jakarta.persistence.OneToMany
import jakarta.persistence.OneToOne
import jakarta.persistence.Table

@Entity
@Table(name = "product")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class ProductEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @Column(name = "name", nullable = false)
    String name

    @Column(name = "description")
    String description

    @Column(name = "state")
    String state

    @Column(name = "image_uuid")
    UUID image

    @Column(name = "production_time")
    Long productionTime

    @ManyToOne
    @JoinColumn(name = "brand_id", referencedColumnName = "id")
    BrandEntity brand

    @ManyToOne
    @JoinColumn(name = "discount_id", referencedColumnName = "id", nullable = true)
    DiscountEntity discount

    @OneToOne
    @JoinColumn(name = "category_id", referencedColumnName = "id")
    CategoryEntity category

    @OneToOne
    @JoinColumn(name = "rating_id", referencedColumnName = "id", nullable = true)
    RatingEntity rating

    @Column(name = "price")
    Float price

    @ManyToMany(fetch = FetchType.EAGER)
    @JoinTable(name = "product_ingredient",
        joinColumns = @JoinColumn(name = "product_id"),
        inverseJoinColumns = @JoinColumn(name = "ingredient_id"))
    Set<IngredientEntity> ingredients = new HashSet<>()

    @OneToMany(mappedBy = "product", cascade = [CascadeType.ALL])
    List<SellingStoreEntity> sellingStores = new ArrayList<>()

    @OneToMany(mappedBy = "product", cascade = CascadeType.ALL)
    List<ReviewEntity> reviews = new ArrayList<>()

    @ManyToMany(fetch = FetchType.EAGER, cascade = CascadeType.ALL )
    @JoinTable(name = "product_extras",
        joinColumns = @JoinColumn(name = "product_id"),
        inverseJoinColumns = @JoinColumn(name = "extra_id"))
    Set<ExtrasEntity> extras = new HashSet<>()

}
