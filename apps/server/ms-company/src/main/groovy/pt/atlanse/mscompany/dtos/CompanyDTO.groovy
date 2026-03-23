package pt.atlanse.mscompany.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable
import pt.atlanse.mscompany.domains.CompanyStatus

import java.time.LocalDateTime

@ToString(includeFields = true, includePackage = false, includeNames = true, excludes = ["description"])
@Introspected
@TupleConstructor
@Serdeable
class CompanyDTO {

    @Nullable
    UUID id

    @Nullable
    String name

    @Nullable
    String emailAddress

    @Nullable
    String phoneNumber

    @Nullable
    CompanyStatus status

    @Nullable
    String category

    @Nullable
    String vatNumber

    @Nullable
    String description

    @Nullable
    String legalName

    @Nullable
    Long deliveryRate

    @Nullable
    Long deliveryQuantity

}
