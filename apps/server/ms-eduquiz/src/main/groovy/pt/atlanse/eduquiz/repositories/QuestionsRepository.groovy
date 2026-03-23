package pt.atlanse.eduquiz.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import io.micronaut.data.repository.jpa.JpaSpecificationExecutor
import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.eduquiz.domain.QuestionsEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface QuestionsRepository extends PageableRepository<QuestionsEntity, UUID>, JpaSpecificationExecutor<QuestionsEntity> {
    @Join(value = "category", type = Join.Type.LEFT_FETCH)
    @Join(value = "answers", type = Join.Type.LEFT_FETCH)
    Page<QuestionsEntity> findAll(PredicateSpecification specification, Pageable pageable)

    @Join(value = "category", type = Join.Type.LEFT_FETCH)
    @Join(value = "answers", type = Join.Type.LEFT_FETCH)
    Optional<QuestionsEntity> findById(UUID id)
}
