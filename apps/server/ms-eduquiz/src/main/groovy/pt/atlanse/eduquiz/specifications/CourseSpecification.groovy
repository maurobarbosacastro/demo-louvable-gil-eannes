package pt.atlanse.eduquiz.specifications

import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.eduquiz.DTO.CourseParams
import pt.atlanse.eduquiz.domain.CourseEntity

class CourseSpecification {
    static PredicateSpecification<CourseEntity> likeTitle(String searchText) {
        return (root, query) -> query.like(root.get("title"), "%${ searchText }%")
    }

    static PredicateSpecification<CourseEntity> likeDescription(String searchText) {
        return (root, query) -> query.like(root.get("description"), "%${ searchText }%")
    }

    static PredicateSpecification<CourseEntity> likeState(String state) {
        return (root, query) -> query.like(root.get("state"), "%${ state }%")
    }

    static PredicateSpecification<CourseEntity> createQueryBySpecification(CourseParams params) {
        PredicateSpecification<CourseEntity> query = PredicateSpecification.ALL as PredicateSpecification<CourseEntity>
        if (params) {
            if (params.searchText) {
                query = query.or(likeTitle(params.searchText))
                query = query.or(likeDescription((params.searchText)))
            }
            if (params.state) {
                query = query.and(likeState(params.state))
            }
        }
        return query
    }
}
