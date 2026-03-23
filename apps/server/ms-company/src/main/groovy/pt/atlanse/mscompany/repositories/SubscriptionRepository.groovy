package pt.atlanse.mscompany.repositories


import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.mscompany.domains.SubscriptionEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface SubscriptionRepository extends PageableRepository<SubscriptionEntity, UUID> {
}
