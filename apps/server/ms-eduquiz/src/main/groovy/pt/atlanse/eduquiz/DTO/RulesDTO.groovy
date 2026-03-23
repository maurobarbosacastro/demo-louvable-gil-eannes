package pt.atlanse.eduquiz.DTO

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@TupleConstructor
@Introspected
@ToString(includePackage = false, includeFields = false, includeNames = true)
@Serdeable.Deserializable
class RulesDTO {
    @Nullable
    String code

    @Nullable
    String value

    @Nullable
    String title

    @Nullable
    String description
}
