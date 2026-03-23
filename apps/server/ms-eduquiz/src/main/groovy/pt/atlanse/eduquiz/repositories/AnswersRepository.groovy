package pt.atlanse.eduquiz.repositories

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import io.micronaut.data.repository.jpa.JpaSpecificationExecutor
import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.eduquiz.domain.AnswersEntity
import pt.atlanse.eduquiz.domain.QuestionsEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface AnswersRepository extends PageableRepository<AnswersEntity, UUID>, JpaSpecificationExecutor<AnswersEntity> {
    Page<AnswersEntity> findAll(PredicateSpecification predicateSpecification, Pageable pageable)

    Page<AnswersEntity> findAllByQuestion(QuestionsEntity question, Pageable pageable)
}
