package pt.atlanse.blog.repository.imageprocessing

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.blog.domain.imageprocessing.MediaFileDoc

/**
 * @deprecated
 * */
@JdbcRepository(dialect = Dialect.POSTGRES)
interface MediaFileDocRepository extends PageableRepository<MediaFileDoc, Long> {
    Optional<MediaFileDoc> findByHash(String hash)
}
