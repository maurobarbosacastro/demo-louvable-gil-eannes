package pt.atlanse.eduquiz.DTO

import com.fasterxml.jackson.annotation.JsonFormat
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable
import pt.atlanse.eduquiz.domain.CourseOrderEntity

import java.time.LocalDateTime

@Introspected
@TupleConstructor
@Serdeable.Deserializable
class CourseDTO {
    @Nullable
    String id

    @Nullable
    ImageDTO image

    @Nullable
    String imageId

    @Nullable
    String title

    @Nullable
    String description

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
    List<String> moduleIds

    @Nullable
    List<CourseOrderEntity> courseOrder = new ArrayList<>()
}
