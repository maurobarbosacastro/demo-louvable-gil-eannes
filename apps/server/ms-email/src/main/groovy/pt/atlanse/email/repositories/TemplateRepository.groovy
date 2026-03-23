package pt.atlanse.email.repositories

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.email.domains.TemplateEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface TemplateRepository extends PageableRepository<TemplateEntity, UUID> {

    Optional<TemplateEntity> findByCode(String code)

}
