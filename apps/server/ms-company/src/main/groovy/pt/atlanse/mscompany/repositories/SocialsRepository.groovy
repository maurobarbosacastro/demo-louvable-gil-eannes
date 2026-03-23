package pt.atlanse.mscompany.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.mscompany.domains.CompanyEntity
import pt.atlanse.mscompany.domains.CompanyHistoryEntity
import pt.atlanse.mscompany.domains.SocialsEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface SocialsRepository extends PageableRepository<SocialsEntity, UUID> {

    Optional<SocialsEntity> findByIdAndCompany(UUID id, CompanyEntity company)

}
