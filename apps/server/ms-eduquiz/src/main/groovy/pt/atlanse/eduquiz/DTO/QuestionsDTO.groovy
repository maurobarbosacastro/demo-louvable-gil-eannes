package pt.atlanse.eduquiz.DTO

import com.fasterxml.jackson.annotation.JsonFormat
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable
import pt.atlanse.eduquiz.domain.CategoryEntity

import java.time.LocalDateTime

enum TypeEnum {
    TEXT,
    VIDEO,
    IMAGE
}

@TupleConstructor
@Introspected
@ToString(includePackage = false, includeFields = true, includeNames = true)
@Serdeable.Deserializable
class QuestionsDTO {
    @Nullable
    UUID id

    @Nullable
    TypeEnum type

    @Nullable
    String description

    @Nullable
    ImageDTO content

    @Nullable
    String imageId

    @Nullable
    CategoryEntity category

    @Nullable
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    LocalDateTime beginDate

    @Nullable
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    LocalDateTime endDate

    @Nullable
    String extras

    @Nullable
    Long points

    @Nullable
    String status

    List<AnswersDTO> answers
}
