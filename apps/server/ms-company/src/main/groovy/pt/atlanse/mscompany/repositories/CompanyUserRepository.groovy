package pt.atlanse.mscompany.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.mscompany.domains.CompanyEntity
import pt.atlanse.mscompany.domains.CompanyHistoryEntity
import pt.atlanse.mscompany.domains.CompanyUserEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface CompanyUserRepository extends PageableRepository<CompanyUserEntity, UUID> {

    Optional<CompanyUserEntity> findByIdAndCompany(UUID id, CompanyEntity company)

}
