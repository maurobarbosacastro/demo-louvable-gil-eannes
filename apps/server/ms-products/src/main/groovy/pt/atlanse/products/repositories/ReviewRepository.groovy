package pt.atlanse.products.repositories

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.Sort
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.products.domains.ProductEntity
import pt.atlanse.products.domains.ReviewEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ReviewRepository extends PageableRepository<ReviewEntity, UUID> {
    Page<ReviewEntity> findByProduct(ProductEntity product, Pageable pageable)

    Page<ReviewEntity> findAll(Pageable pageable, Sort sort)
}
