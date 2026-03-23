package pt.atlanse.products.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@Introspected
@TupleConstructor
@ToString(includeNames = false, includeFields = false, includePackage = false)
@Serdeable.Deserializable
class ProductSearchDTO {

    @Nullable
    String productName

    @Nullable
    List<UUID> popularIngredients

    @Nullable
    Float minPrice

    @Nullable
    Float maxPrice

    @Nullable
    Boolean promo

    @Nullable
    String state

}
