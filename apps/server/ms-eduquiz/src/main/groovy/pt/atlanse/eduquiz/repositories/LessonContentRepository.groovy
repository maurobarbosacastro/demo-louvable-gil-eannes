package pt.atlanse.eduquiz.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.eduquiz.domain.LessonContentEntity
import pt.atlanse.eduquiz.domain.LessonEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface LessonContentRepository extends PageableRepository<LessonContentEntity, UUID> {

    @Join(value = "content", type = Join.Type.FETCH)
    List<LessonContentEntity> findAllByLesson(LessonEntity lesson)
}


