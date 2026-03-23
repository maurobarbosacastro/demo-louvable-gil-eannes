package pt.atlanse.eduquiz.specifications

import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.eduquiz.DTO.QuestionsParams
import pt.atlanse.eduquiz.domain.QuestionsEntity

class QuestionSpecification {
    static PredicateSpecification<QuestionsEntity> containsType(String type) {
        return (root, query) -> query.like(root.get("type"), "%${ type }%")
    }

    static PredicateSpecification<QuestionsEntity> equalsCategory(String category) {
        return (root, query) -> query.equal(root.get("category"), category)
    }

    static PredicateSpecification<QuestionsEntity> createQueryBySpecification(QuestionsParams params) {
        PredicateSpecification<QuestionsEntity> query = PredicateSpecification.ALL as PredicateSpecification<QuestionsEntity>
        if (params) {
            if (params.category && params.category !== null) {
                query = query.and(equalsCategory(params.category))
            }
            if (params.type && params.type !== null) {
                query = query.and(containsType(params.type))
            }
        }
        return query
    }
}
