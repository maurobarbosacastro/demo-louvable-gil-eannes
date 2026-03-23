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
class LessonParams {
    @Nullable
    String title

    @Nullable
    String subtitle

    @Nullable
    String state

    @Nullable
    String conclusion

    @Nullable
    String type
}
