package pt.atlanse.eduquiz.DTO

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@Introspected
@TupleConstructor
@Serdeable.Deserializable
class RulesParams {
    @Nullable
    String code

    @Nullable
    String value

    @Nullable
    String title

    @Nullable
    String description
}
