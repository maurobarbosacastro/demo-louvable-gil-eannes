package pt.atlanse.blog.repository

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.blog.domain.ContentEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ContentRepository extends PageableRepository<ContentEntity, Long> {

}


