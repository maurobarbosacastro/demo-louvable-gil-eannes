package pt.atlanse.eduquiz.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.eduquiz.domain.LessonEntity
import pt.atlanse.eduquiz.domain.ModulesEntity
import pt.atlanse.eduquiz.domain.ModulesOrderEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ModulesOrderRepository extends PageableRepository<ModulesOrderEntity, UUID> {
    Long countByLessonAndModule(LessonEntity lesson, ModulesEntity module)

    @Join(value = "lesson", type = Join.Type.FETCH)
    Page<ModulesOrderEntity> findAllByModule(ModulesEntity module, Pageable pageable)
}


