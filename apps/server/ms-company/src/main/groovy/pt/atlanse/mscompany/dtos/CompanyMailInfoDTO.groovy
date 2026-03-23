package pt.atlanse.mscompany.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.serde.annotation.Serdeable
import jakarta.annotation.Nullable
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@ToString(includeFields = true, includePackage = false, includeNames = true)
@Introspected
@TupleConstructor
@Serdeable
class CompanyMailInfoDTO {

    @Nullable
    String id

    @Nullable
    String company

    @Nullable
    String countryName

    @Nullable
    String address1

    @Nullable
    String address2

    @Nullable
    String postalCode

    @Nullable
    CoordinatesDTO coordinates

    @Nullable
    String locality

}

@Introspected
@Serdeable
class CoordinatesDTO {

    @NotBlank
    @NotNull
    double latitude

    @NotBlank
    @NotNull
    double longitude
}
