package pt.atlanse.products.repositories

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.products.domains.DiscountEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface DiscountRepository extends PageableRepository<DiscountEntity, UUID> {

}
