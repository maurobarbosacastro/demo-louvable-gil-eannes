package pt.atlanse.products.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@ToString(includeFields = true, includePackage = false, includeNames = true, excludes = ["description"])
@Introspected
@TupleConstructor
@Serdeable.Deserializable
class Extras {
    @Nullable
    UUID id

    @Nullable
    String name

    @Nullable
    String description

    @Nullable
    Float price

    @Nullable
    String state

    @Nullable
    ImageDTO image

    @Nullable
    Long productsCount

    @Nullable
    String imageId

    @Nullable
    Set<GetProduct> products
}
