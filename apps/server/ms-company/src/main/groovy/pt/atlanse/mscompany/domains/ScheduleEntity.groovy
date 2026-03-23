package pt.atlanse.mscompany.domains

import com.fasterxml.jackson.annotation.JsonFormat
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated
import io.micronaut.serde.annotation.Serdeable
import jakarta.persistence.*
import jakarta.validation.constraints.NotBlank

import java.time.DayOfWeek
import java.time.LocalDateTime
import java.time.LocalTime

@Introspected
@Entity
@TupleConstructor
@Table(name = "schedule")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@Serdeable
class ScheduleEntity extends BaseEntity {

    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @ManyToOne
    @JoinColumn(name = "company_id", referencedColumnName = "id")
    CompanyEntity company

    @NonNull
    @NotBlank
    @JsonFormat(pattern = "HH:mm:ss")
    @Column(name = "start_date")
    LocalTime startDate

    @NonNull
    @NotBlank
    @JsonFormat(pattern = "HH:mm:ss")
    @Column(name = "end_date")
    LocalTime endDate

    @NonNull
    @NotBlank
    @Column(name = "week_day")
    DayOfWeek weekDay

    @NonNull
    @NotBlank
    @Column(name = "type")
    ScheduleType type  = ScheduleType.DELIVERY
}

@Introspected
enum ScheduleType {
    COMPANY,
    DELIVERY
}
