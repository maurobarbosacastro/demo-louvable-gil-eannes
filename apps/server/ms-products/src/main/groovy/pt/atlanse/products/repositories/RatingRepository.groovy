package pt.atlanse.products.repositories

import io.micronaut.data.annotation.Query
import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.products.domains.RatingEntity

@JdbcRepository(dialect = Dialect.POSTGRES)
interface RatingRepository extends PageableRepository<RatingEntity, UUID> {
    @Query("select sum(amount) from products.rating r inner join products.product p on r.id = p.rating_id where p.brand_id =:brandId")
    Optional<Long> sumRatingByBrand(UUID brandId);
}
