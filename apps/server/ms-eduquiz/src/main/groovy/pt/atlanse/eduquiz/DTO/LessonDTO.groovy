package pt.atlanse.eduquiz.DTO

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Introspected
@TupleConstructor
@Serdeable.Deserializable
class LessonDTO {
    @NotNull(message = 'Title can not be null')
    @NotBlank
    String title

    @NotNull(message = 'Subtitle can not be null')
    @NotBlank
    String subtitle

    @Nullable
    ImageDTO content

    @Nullable
    String conclusion

    @NotNull(message = 'Type can not be null')
    @NotBlank
    String type

    @NotNull(message = 'Status can not be null')
    @NotBlank
    String status
}
