package pt.atlanse.mscompany.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.serde.annotation.Serdeable
import jakarta.annotation.Nullable
import pt.atlanse.mscompany.domains.CompanySubscriptionStatus

import java.time.LocalDateTime

@ToString(includeFields = true, includePackage = false, includeNames = true)
@Introspected
@TupleConstructor
@Serdeable
class CompanySubscriptionDTO {

    @Nullable
    UUID id

    @Nullable
    UUID company

    @Nullable
    UUID subscription

    @Nullable
    LocalDateTime expireDate

    @Nullable
    CompanySubscriptionStatus status

    @Nullable
    double price

    @Nullable
    LocalDateTime startDate

    @Nullable
    LocalDateTime buyDate

}
