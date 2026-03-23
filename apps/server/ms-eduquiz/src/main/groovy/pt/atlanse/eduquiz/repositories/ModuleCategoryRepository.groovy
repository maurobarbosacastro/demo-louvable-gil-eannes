package pt.atlanse.eduquiz.repositories

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.eduquiz.domain.CategoryEntity
import pt.atlanse.eduquiz.domain.ModulesEntity
import pt.atlanse.eduquiz.domain.compositeIds.ModuleCategoryEntity
import pt.atlanse.eduquiz.domain.compositeIds.ModuleCategoryId

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ModuleCategoryRepository extends PageableRepository<ModuleCategoryEntity, ModuleCategoryId> {
    ModuleCategoryEntity findByModuleAndCategory(ModulesEntity module, CategoryEntity category)

    List<ModuleCategoryEntity> findAllByModule(ModulesEntity module)
}
