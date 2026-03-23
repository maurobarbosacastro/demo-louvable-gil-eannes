package pt.atlanse.products.domains.compositeIds

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import jakarta.persistence.Column
import jakarta.persistence.Embeddable
import jakarta.persistence.EmbeddedId
import jakarta.persistence.Entity
import jakarta.persistence.FetchType
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.MapsId
import jakarta.persistence.Table
import pt.atlanse.products.domains.IngredientEntity
import pt.atlanse.products.domains.ProductEntity



@Embeddable
class ProductIngredientId implements Serializable {
    @Column(name = "product_id")
    UUID productId

    @Column(name = "ingredient_id")
    UUID ingredientId
}

@Entity
@TupleConstructor
@Table(name = "product_ingredient")
@ToString(includePackage = false, includeNames = true, includeFields = true, excludes = ["product", "ingredient"])
class ProductIngredientRelationEntity {

    @EmbeddedId
    ProductIngredientId id

    @ManyToOne(fetch = FetchType.EAGER)
    @MapsId("productId")
    @JoinColumn(name = "product_id", referencedColumnName = "id")
    ProductEntity product

    @ManyToOne(fetch = FetchType.EAGER)
    @MapsId("ingredientId")
    @JoinColumn(name = "ingredient_id", referencedColumnName = "id")
    IngredientEntity ingredient

    ProductIngredientRelationEntity() {}

    /**
     * @deprecated
     * */
    ProductIngredientRelationEntity(ProductEntity product, IngredientEntity ingredient) {
        this.product = product
        this.ingredient = ingredient
        this.id = new ProductIngredientId(productId: product.id, ingredientId: ingredient.id)
        product.ingredients.add(ingredient)
        ingredient.products.add(product)
    }

}
