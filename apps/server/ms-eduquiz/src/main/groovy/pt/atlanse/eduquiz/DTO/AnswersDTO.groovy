package pt.atlanse.eduquiz.DTO

import com.fasterxml.jackson.annotation.JsonFormat
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

import java.time.LocalDateTime

@Introspected
@ToString(includePackage = false, includeFields = false, includeNames = true)
@Serdeable.Deserializable
@TupleConstructor
class AnswersDTO {
    @Nullable
    UUID id

    @NotBlank(message = 'Content can not be null')
    @NotNull
    String content

    @NotBlank(message = 'QuestionId can not be null')
    @NotNull
    String questionId

    @NotBlank(message = 'isCorrect can not be null')
    @NotNull
    Boolean isCorrect

    @NotBlank(message = 'Points can not be null')
    @NotNull
    Long points

    @Nullable
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    LocalDateTime createdAt
}
