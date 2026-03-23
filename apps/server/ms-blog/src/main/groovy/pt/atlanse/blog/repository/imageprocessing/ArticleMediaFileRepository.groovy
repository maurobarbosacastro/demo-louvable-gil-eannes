package pt.atlanse.blog.repository.imageprocessing

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.blog.domain.imageprocessing.ArticleMediaFile

/**
 * @deprecated
 * */
@JdbcRepository(dialect = Dialect.POSTGRES)
interface ArticleMediaFileRepository extends PageableRepository<ArticleMediaFile, Long> {

}
