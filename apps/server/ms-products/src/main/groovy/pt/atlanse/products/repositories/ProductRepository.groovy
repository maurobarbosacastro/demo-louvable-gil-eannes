package pt.atlanse.products.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.annotation.Query
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.products.domains.BrandEntity
import pt.atlanse.products.domains.CategoryEntity
import pt.atlanse.products.domains.ProductEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ProductRepository extends PageableRepository<ProductEntity, UUID> {
    @Join(value = "rating", type = Join.Type.LEFT_FETCH)
    @Join(value = "reviews", type = Join.Type.LEFT_FETCH)
    @Join(value = "brand", type = Join.Type.LEFT_FETCH)
    @Join(value = "category", type = Join.Type.LEFT_FETCH)
    @Join(value = "ingredients", type = Join.Type.LEFT_FETCH)
    @Join(value = "extras", type = Join.Type.LEFT_FETCH)
    Optional<ProductEntity> findById(UUID id)

    @Join(value = "rating", type = Join.Type.FETCH)
    @Join(value = "reviews", type = Join.Type.LEFT_FETCH)
    @Join(value = "brand", type = Join.Type.FETCH)
    @Join(value = "category", type = Join.Type.FETCH)
    Page<ProductEntity> findAll(Pageable page)

    Long countByBrand(BrandEntity brand)

    Long countByCategory(CategoryEntity category)

    @Join(value = "rating", type = Join.Type.FETCH)
    @Join(value = "reviews", type = Join.Type.LEFT_FETCH)
    @Join(value = "brand", type = Join.Type.FETCH)
    @Join(value = "category", type = Join.Type.FETCH)
    List<ProductEntity> findByPriceBetween(Float minPrice, Float maxPrice);

    List<ProductEntity> findByNameContains(String name);

    @Query('''DELETE FROM product_extras WHERE product_id = :product AND extra_id = :extra;''')
    deleteExtraRelation(UUID product, UUID extra)

}
