package pt.atlanse.mscompany.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.serde.annotation.Serdeable
import jakarta.annotation.Nullable
import pt.atlanse.mscompany.domains.ScheduleType
import pt.atlanse.mscompany.domains.SocialType

import java.time.DayOfWeek
import java.time.LocalDateTime
import java.time.LocalTime


@ToString(includeFields = true, includePackage = false, includeNames = true)
@Introspected
@TupleConstructor
@Serdeable
class ScheduleDTO {

    @Nullable
    UUID id

    @Nullable
    ScheduleType type

    @Nullable
    LocalTime startDate

    @Nullable
    LocalTime endDate

    @Nullable
    DayOfWeek weekDay

    @Nullable
    String company

}
