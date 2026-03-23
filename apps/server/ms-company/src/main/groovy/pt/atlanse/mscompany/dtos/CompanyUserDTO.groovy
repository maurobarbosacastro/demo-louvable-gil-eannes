package pt.atlanse.mscompany.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.serde.annotation.Serdeable
import jakarta.annotation.Nullable
import pt.atlanse.mscompany.domains.CompanyHistoryType

import java.time.LocalDateTime

@ToString(includeFields = true, includePackage = false, includeNames = true)
@Introspected
@TupleConstructor
@Serdeable
class CompanyUserDTO {

    @Nullable
    UUID id

    @Nullable
    UUID keycloakUserId

    @Nullable
    UUID company


}
