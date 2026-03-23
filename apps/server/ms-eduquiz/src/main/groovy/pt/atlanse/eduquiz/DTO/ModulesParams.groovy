package pt.atlanse.eduquiz.DTO

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@Introspected
@TupleConstructor
@ToString(includeNames = true, includePackage = false, includeFields = true)
@Serdeable.Deserializable
class ModulesParams {
    @Nullable
    String title

    @Nullable
    String description

    @Nullable
    String status

    @Nullable
    String category
}
