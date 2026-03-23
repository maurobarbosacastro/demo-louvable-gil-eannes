package pt.atlanse.products.utils.specifications

import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.products.domains.BrandEntity
import pt.atlanse.products.dtos.BrandParams

/**
 * @deprecated
 * */
class BrandSpecifications {

    static PredicateSpecification<BrandEntity> containsName(String name) {

        return (root, query) -> query.like(root.get("name"), "%${ name }%")
    }

    static PredicateSpecification<BrandEntity> stateEquals(String state) {
        return (root, query) -> query.equal(root.get("state"), state)
    }

    static PredicateSpecification<BrandEntity> createQueryBySpecifications(BrandParams params) {

        PredicateSpecification<BrandEntity> query = PredicateSpecification.ALL as PredicateSpecification<BrandEntity>

        if (params.name) {
            query = query.and(containsName(params.name))
        }

        if (params.state) {
            query = query.and(stateEquals(params.state))
        }

        return query
    }
}
