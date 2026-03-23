package pt.atlanse.eduquiz.DTO

import com.fasterxml.jackson.annotation.JsonFormat
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

import java.time.LocalDateTime

@Introspected
@ToString(includePackage = false, includeFields = false, includeNames = true)
@TupleConstructor
@Serdeable.Deserializable
class CourseParams {
    @Nullable
    String title

    @Nullable
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    LocalDateTime beginDate

    @Nullable
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    LocalDateTime endDate

    @Nullable
    String extras

    @Nullable
    String status

    @Nullable
    boolean completeInformation
}
