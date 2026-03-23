package pt.atlanse.mscompany.domains

import com.fasterxml.jackson.annotation.JsonFormat
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated
import io.micronaut.data.annotation.DateCreated
import io.micronaut.serde.annotation.Serdeable
import jakarta.persistence.*
import jakarta.validation.constraints.NotBlank

import java.time.DayOfWeek
import java.time.LocalDateTime
import java.time.LocalTime

@Introspected
@Entity
@TupleConstructor
@Table(name = "company_history")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@Serdeable
class CompanyHistoryEntity extends BaseEntity {

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
    @Column(name = "change_type")
    CompanyHistoryType changeType

    @NonNull
    @NotBlank
    @Column(name = "description")
    String description
}

@Introspected
enum CompanyHistoryType {
    UPDATE,
    DELETE
}

