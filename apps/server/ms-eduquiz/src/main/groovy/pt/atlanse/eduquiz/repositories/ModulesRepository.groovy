package pt.atlanse.eduquiz.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.eduquiz.domain.ModulesEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ModulesRepository extends PageableRepository<ModulesEntity, UUID> {
    @Join(value = "categories", type = Join.Type.LEFT_FETCH)
    Optional<ModulesEntity> findById(UUID id)
}


