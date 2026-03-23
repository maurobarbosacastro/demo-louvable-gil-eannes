package pt.atlanse.products.repositories

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.products.domains.ProductEntity
import pt.atlanse.products.domains.SellingStoreEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface SellingStoreRepository extends PageableRepository<SellingStoreEntity, UUID> {
    List<SellingStoreEntity> findByProduct(ProductEntity product)
}
