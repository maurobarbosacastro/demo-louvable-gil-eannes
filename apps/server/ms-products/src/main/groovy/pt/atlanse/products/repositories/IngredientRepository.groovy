package pt.atlanse.products.repositories

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.products.domains.IngredientEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface IngredientRepository extends PageableRepository<IngredientEntity, UUID> {

    Page<IngredientEntity> findAllByNameContains(String name, Pageable pageable)

}

