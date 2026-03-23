package pt.atlanse.eduquiz.DTO

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@Introspected
@TupleConstructor
@Serdeable.Deserializable
class QuizParams {
    @Nullable
    String categoryId

    @Nullable
    String title

    @Nullable
    String description

    @Nullable
    String participantId

    @Nullable
    String teamId
}
