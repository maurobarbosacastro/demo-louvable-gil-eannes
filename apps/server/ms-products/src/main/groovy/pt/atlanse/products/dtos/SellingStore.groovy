package pt.atlanse.products.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank


@Introspected
@TupleConstructor
@ToString(includeFields = true, includePackage = false, includeNames = true)
@Serdeable.Deserializable
class SellingStore {
    @NonNull
    @NotBlank
    String name

    @Nullable
    String website
}
