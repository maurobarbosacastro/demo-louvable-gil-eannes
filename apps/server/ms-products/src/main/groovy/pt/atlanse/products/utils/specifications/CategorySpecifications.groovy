package pt.atlanse.products.utils.specifications

import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.products.domains.CategoryEntity
import pt.atlanse.products.dtos.CategoryParams

/**
 * @deprecated
 * */
class CategorySpecifications {

    static PredicateSpecification<CategoryEntity> containsName(String name) {

        return (root, query) -> query.like(root.get("name"), "%${ name }%")
    }

    static PredicateSpecification<CategoryEntity> stateEquals(String state) {
        return (root, query) -> query.equal(root.get("state"), state)
    }

    static PredicateSpecification<CategoryEntity> createQueryBySpecifications(CategoryParams params) {

        PredicateSpecification<CategoryEntity> query = PredicateSpecification.ALL as PredicateSpecification<CategoryEntity>

        if (params.name) {
            query = query.and(containsName(params.name))
        }

        if (params.state) {
            query = query.and(stateEquals(params.state))
        }

        return query
    }
}
