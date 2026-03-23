package pt.atlanse.products.dtos

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@Introspected
@TupleConstructor
@Serdeable.Deserializable
class ExtrasParams {
    @Nullable
    String name

    @Nullable
    String state
}
