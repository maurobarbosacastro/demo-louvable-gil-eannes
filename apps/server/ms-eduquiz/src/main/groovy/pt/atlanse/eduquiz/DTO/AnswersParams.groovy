package pt.atlanse.eduquiz.DTO

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@Introspected
@ToString(includePackage = false, includeFields = false, includeNames = true)
@TupleConstructor
@Serdeable.Deserializable
class AnswersParams {
    @Nullable
    Boolean isCorrect

    @Nullable
    String questionId

    @Nullable
    String content
}
