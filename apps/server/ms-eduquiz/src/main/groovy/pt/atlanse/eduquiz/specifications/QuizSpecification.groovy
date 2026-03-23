package pt.atlanse.eduquiz.specifications

import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.eduquiz.DTO.QuizParams
import pt.atlanse.eduquiz.domain.QuizEntity

class QuizSpecification {
    static PredicateSpecification<QuizEntity> containsTitle(String title) {
        return (root, query) -> query.like(root.get("title"), "%${ title }%")
    }

    static PredicateSpecification<QuizEntity> containsDescription(String description) {
        return (root, query) -> query.like(root.get("description"), "%${ description }%")
    }

    static PredicateSpecification<QuizEntity> equalsParticipant(String participant) {
        return (root, query) -> query.equal(root.get("participant"), participant)
    }

    static PredicateSpecification<QuizEntity> containsTeamId(String teamId) {
        return (root, query) -> query.like(root.get("teamId"), "%${ teamId }%")
    }

    static PredicateSpecification<QuizEntity> createQueryBySpecification(QuizParams params) {
        PredicateSpecification<QuizEntity> query = PredicateSpecification.ALL as PredicateSpecification<QuizEntity>
        if (params) {
            if (params.title) {
                query = query.and(containsTitle(params.title))
            }

            if (params.description) {
                query = query.and(containsDescription(params.description))
            }

            if (params.teamId) {
                query = query.and(containsTeamId(params.teamId))
            }

            if (params.participantId) {
                query = query.and(equalsParticipant(params.participantId))
            }
        }
        return query
    }
}
