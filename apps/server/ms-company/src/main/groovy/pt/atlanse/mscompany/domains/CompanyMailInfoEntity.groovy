package pt.atlanse.mscompany.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated
import io.micronaut.serde.annotation.Serdeable
import jakarta.annotation.Nullable
import jakarta.persistence.*
import jakarta.validation.constraints.NotBlank

@Introspected
@Entity
@TupleConstructor
@Table(name = "company_mail_info")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@Serdeable
class CompanyMailInfoEntity extends BaseEntity {

    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @OneToOne
    @JoinColumn(name = "company_id", referencedColumnName = "id")
    CompanyEntity company

    @NonNull
    @NotBlank
    @Column(name = "address_1")
    String address1


    @NonNull
    @NotBlank
    @Column(name = "address_2")
    String address2


    @NonNull
    @NotBlank
    @Column(name = "postal_code")
    String postalCode

    @NonNull
    @NotBlank
    @Column(name = "locality")
    String locality

    @NonNull
    @NotBlank
    @Column(name = "country_name")
    String countryName

    @Nullable
    @Column(name = "longitude")
    Double longitude

    @Nullable
    @Column(name = "latitude")
    Double latitude

}

