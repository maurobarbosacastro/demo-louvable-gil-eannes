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
class Brand {
    @Nullable
    UUID id

    @Nullable
    String name

    @Nullable
    String description

    @Nullable
    String website

    @Nullable
    String state

    @Nullable
    String numProducts

    @Nullable
    String averageRating

    @Nullable
    String image

    @Nullable
    String banner

}
