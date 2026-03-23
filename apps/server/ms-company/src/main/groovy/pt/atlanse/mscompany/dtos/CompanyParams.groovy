package pt.atlanse.mscompany.dtos

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@Introspected
@TupleConstructor
@Serdeable.Deserializable
class CompanyParams {
    @Nullable
    String name

    @Nullable
    String status

}
