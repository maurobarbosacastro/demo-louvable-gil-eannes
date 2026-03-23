package pt.atlanse.eduquiz.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.eduquiz.domain.QuizEntity
import pt.atlanse.eduquiz.domain.QuizQuestionsEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface QuizQuestionsRepository extends PageableRepository<QuizQuestionsEntity, UUID> {
    @Join(value = "questions", type = Join.Type.FETCH)
    Page<QuizQuestionsEntity> findAllByQuiz(QuizEntity quiz, Pageable pageable)
}


