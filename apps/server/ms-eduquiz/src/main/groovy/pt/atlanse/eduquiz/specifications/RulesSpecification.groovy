package pt.atlanse.eduquiz.specifications

import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.eduquiz.DTO.RulesParams
import pt.atlanse.eduquiz.domain.RulesEntity

class RulesSpecification {

    static PredicateSpecification<RulesEntity> containsDescription(String description) {
        return (root, query) -> query.like(root.get("description"), "%${ description }%")
    }

    static PredicateSpecification<RulesEntity> containsValue(String value) {
        return (root, query) -> query.like(root.get("value"), "%${ value }%")
    }

    static PredicateSpecification<RulesEntity> containsTitle(String title) {
        return (root, query) -> query.like(root.get("title"), "%${ title }%")
    }

    static PredicateSpecification<RulesEntity> containsCode(String code) {
        return (root, query) -> query.like(root.get("code"), "%${ code }%")
    }

    static PredicateSpecification<RulesEntity> createQueryBySpecification(RulesParams params) {
        PredicateSpecification<RulesEntity> query = PredicateSpecification.ALL as PredicateSpecification<RulesEntity>
        if (params) {
            if (params.description) {
                query = query.and(containsDescription(params.description))
            }

            if (params.value) {
                query = query.and(containsValue(params.value))
            }

            if (params.title) {
                query = query.and(containsTitle(params.title))
            }

            if (params.code) {
                query = query.and(containsCode(params.code))
            }
        }
        return query
    }
}
