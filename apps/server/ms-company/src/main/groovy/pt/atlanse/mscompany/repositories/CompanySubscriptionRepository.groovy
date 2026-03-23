package pt.atlanse.mscompany.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.mscompany.domains.CompanyEntity
import pt.atlanse.mscompany.domains.CompanySubscriptionEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface CompanySubscriptionRepository extends PageableRepository<CompanySubscriptionEntity, UUID> {

    @Join(value = "company", type = Join.Type.LEFT_FETCH)
    @Join(value = "subscription", type = Join.Type.LEFT_FETCH)
    Optional<CompanySubscriptionEntity> findById(UUID id)

    Optional<CompanySubscriptionEntity> findByIdAndCompany(UUID id, CompanyEntity company)
}
