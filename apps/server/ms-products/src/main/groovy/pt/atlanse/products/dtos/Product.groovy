package pt.atlanse.products.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@ToString
@Introspected
@TupleConstructor
@Serdeable.Deserializable
class Product {
    @Nullable
    String name

    @Nullable
    String description

    @Nullable
    UUID brand

    @Nullable
    UUID discount

    @Nullable
    UUID category

    @Nullable
    Float price

    @Nullable
    Long productionTime

    @Nullable
    ImageDTO image

    @Nullable
    String state

    // todo Associate with product during creation
    @Nullable
    List<Ingredient> ingredients

    @Nullable
    List<SellingStore> sellingStores
}
