package pt.atlanse.products.utils.specifications

import io.micronaut.data.repository.jpa.criteria.PredicateSpecification
import pt.atlanse.products.domains.ProductEntity
import pt.atlanse.products.dtos.ProductParams

/**
 * @deprecated
 * */
class ProductSpecifications {

    static PredicateSpecification<ProductEntity> containsName(String name) {

        return (root, query) -> query.like(root.get("name"), "%${ name }%")
    }

    static PredicateSpecification<ProductEntity> brandEquals(String brand) {
        return (root, query) -> query.equal(root.get("brand"), brand)
    }

    static PredicateSpecification<ProductEntity> departmentEquals(String department) {
        return (root, query) -> query.equal(root.get("category"), department)
    }

    static PredicateSpecification<ProductEntity> stateEquals(String state) {
        return (root, query) -> query.equal(root.get("state"), state)
    }

    static PredicateSpecification<ProductEntity> createQueryBySpecifications(ProductParams params) {

        PredicateSpecification<ProductEntity> query = PredicateSpecification.ALL as PredicateSpecification<ProductEntity>

        if (params.name) {
            query = query.and(containsName(params.name))
        }

        if (params.brand) {
            query = query.and(brandEquals(params.brand))
        }

        if (params.department) {
            query = query.and(departmentEquals(params.department))
        }

        if (params.state) {
            query = query.and(stateEquals(params.state))
        }

        return query
    }
}
