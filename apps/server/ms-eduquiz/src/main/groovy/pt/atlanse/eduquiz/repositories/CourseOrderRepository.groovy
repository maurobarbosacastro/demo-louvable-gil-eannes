package pt.atlanse.eduquiz.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.eduquiz.domain.CourseEntity
import pt.atlanse.eduquiz.domain.CourseOrderEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface CourseOrderRepository extends PageableRepository<CourseOrderEntity, UUID> {
    Long countByCourse(CourseEntity course)

    @Join(value = "module", type = Join.Type.FETCH)
    Page<CourseOrderEntity> findAllByCourse(CourseEntity course, Pageable page)

    List<CourseOrderEntity> findAllByCourse(CourseEntity course)
}


