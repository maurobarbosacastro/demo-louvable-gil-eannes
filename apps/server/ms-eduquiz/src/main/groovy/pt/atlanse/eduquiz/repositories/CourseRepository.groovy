package pt.atlanse.eduquiz.repositories

import io.micronaut.data.annotation.Join
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import io.micronaut.data.repository.jpa.JpaSpecificationExecutor
import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.eduquiz.domain.CourseEntity

import java.time.LocalDateTime

@JdbcRepository(dialect = Dialect.POSTGRES)
interface CourseRepository extends PageableRepository<CourseEntity, UUID>, JpaSpecificationExecutor<CourseEntity> {
    Page<CourseEntity> findAll(PredicateSpecification specification, Pageable pageable)

    List<CourseEntity> findByStatusIlike(String status)

    @Join(value = "courseOrder", type = Join.Type.LEFT_FETCH)
    @Join(value = "courseOrder.module", type = Join.Type.LEFT_FETCH)
    Optional<CourseEntity> findTop1ByBeginDateGreaterThan(LocalDateTime endDate)

    @Join(value = "courseOrder", type = Join.Type.LEFT_FETCH)
    @Join(value = "courseOrder.module", type = Join.Type.LEFT_FETCH)
    Optional<CourseEntity> findTop1ByEndDateLessThan(LocalDateTime beginDate)

    @Join(value = "courseOrder", type = Join.Type.LEFT_FETCH)
    @Join(value = "courseOrder.module", type = Join.Type.LEFT_FETCH)
    Optional<CourseEntity> findById(UUID id)
}
