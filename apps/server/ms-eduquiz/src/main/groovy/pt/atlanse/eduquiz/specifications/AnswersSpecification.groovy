package pt.atlanse.eduquiz.specifications

import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.eduquiz.DTO.AnswersParams
import pt.atlanse.eduquiz.domain.AnswersEntity

class AnswersSpecification {
    static PredicateSpecification<AnswersEntity> equalIsCorrect(Boolean isCorrect) {
        return (root, query) -> query.equal(root.get("isCorrect"), isCorrect)
    }

    static PredicateSpecification<AnswersEntity> equalQuestion(String question) {
        return (root, query) -> query.equal(root.get("question"), question)
    }

    static PredicateSpecification<AnswersEntity> likeContent(String content) {
        return (root, query) -> query.like(root.get("content"), "%$content%")
    }

    static PredicateSpecification<AnswersEntity> createQueryBySpecification(AnswersParams params) {
        PredicateSpecification<AnswersEntity> query = PredicateSpecification.ALL as PredicateSpecification<AnswersEntity>
       if (params) {
           if (params.isCorrect) {
               query = query.and(equalIsCorrect(params.isCorrect))
           }
           if (params.questionId) {
               query = query.and(equalQuestion(params.questionId))
           }
           if (params.content) {
               query = query.and(likeContent(params.content))
           }
       }
        return query
    }
}
