package pt.atlanse.mscompany.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.serde.annotation.Serdeable
import jakarta.annotation.Nullable
import pt.atlanse.mscompany.domains.ScheduleType

import java.time.DayOfWeek
import java.time.LocalDateTime

@ToString(includeFields = true, includePackage = false, includeNames = true)
@Introspected
@TupleConstructor
@Serdeable
class ScheduleParams {

    @Nullable
    ScheduleType type

    @Nullable
    DayOfWeek weekDay

    @Nullable
    String company

}
