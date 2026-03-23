package pt.atlanse.products.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@Introspected
@TupleConstructor
@ToString(includeFields = true, includePackage = false, includeNames = true, excludes = ["description"])
@Serdeable.Deserializable
class Category {
    @Nullable
    UUID id

    @Nullable
    String name

    @Nullable
    String description

    @Nullable
    ImageDTO image

    @Nullable
    String state

    @Nullable
    String numProducts
}
