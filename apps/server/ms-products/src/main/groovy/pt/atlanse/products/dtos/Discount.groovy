package pt.atlanse.products.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank


@ToString(includeFields = true, includePackage = false, includeNames = true)
@Introspected
@TupleConstructor
@Serdeable.Deserializable
class Discount {
    @NonNull
    @NotBlank
    String code

    @NonNull
    @NotBlank
    boolean active

    @NonNull
    @NotBlank
    String amount

}
