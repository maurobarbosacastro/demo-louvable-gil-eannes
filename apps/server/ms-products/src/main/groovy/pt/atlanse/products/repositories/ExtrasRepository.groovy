package pt.atlanse.products.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.annotation.Query
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.products.domains.ExtrasEntity
import pt.atlanse.products.domains.ProductEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ExtrasRepository extends PageableRepository<ExtrasEntity, UUID> {

    @Join(value = "products", type = Join.Type.LEFT_FETCH)
    Optional<ExtrasEntity> findById(UUID id)

    @Query('''DELETE FROM product_extras WHERE extra_id = :extra;''')
    deleteProductRelation(UUID extra);
}
