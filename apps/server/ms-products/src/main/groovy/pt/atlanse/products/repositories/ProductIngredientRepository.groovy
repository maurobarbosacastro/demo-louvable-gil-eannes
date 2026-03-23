package pt.atlanse.products.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.products.domains.IngredientEntity
import pt.atlanse.products.domains.ProductEntity
import pt.atlanse.products.domains.compositeIds.ProductIngredientId
import pt.atlanse.products.domains.compositeIds.ProductIngredientRelationEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ProductIngredientRepository extends PageableRepository<ProductIngredientRelationEntity, ProductIngredientId> {

    @Join(value = "ingredient", type = Join.Type.FETCH)
    Page<ProductIngredientRelationEntity> findByProduct(ProductEntity product, Pageable pageable)

    @Join(value = "product.rating", type = Join.Type.FETCH)
    @Join(value = "product.reviews", type = Join.Type.LEFT_FETCH)
    @Join(value = "product.brand", type = Join.Type.FETCH)
    @Join(value = "product.category", type = Join.Type.FETCH)
    Page<ProductIngredientRelationEntity> findByIngredient(IngredientEntity ingredient, Pageable pageable)

    Page<ProductIngredientRelationEntity> findByProductAndIngredient(ProductEntity product, IngredientEntity ingredient, Pageable pageable)

}

