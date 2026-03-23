package pt.atlanse.eduquiz.repositories

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import io.micronaut.data.repository.jpa.JpaSpecificationExecutor
import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.eduquiz.domain.CategoryEntity
import pt.atlanse.eduquiz.domain.QuizEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface QuizRepository extends PageableRepository<QuizEntity, UUID>, JpaSpecificationExecutor<QuizEntity> {
    Long countByCategory(CategoryEntity category)

    Page<QuizEntity> findAll(PredicateSpecification specification, Pageable pageable)
}


