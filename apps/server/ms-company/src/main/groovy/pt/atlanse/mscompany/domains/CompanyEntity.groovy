package pt.atlanse.mscompany.domains

import com.fasterxml.jackson.annotation.JsonFormat
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated
import io.micronaut.serde.annotation.Serdeable
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank

import java.time.LocalDateTime

@Introspected
@Entity
@TupleConstructor
@Table(name = "company")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@Serdeable
class CompanyEntity extends BaseEntity {

    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @Column(name = "name")
    String name

    @NonNull
    @NotBlank
    @Column(name = "email_address")
    String emailAddress

    @NonNull
    @NotBlank
    @Column(name = "phone_number")
    String phoneNumber

    @Column(name = "status")
    CompanyStatus status

    @Column(name = "category")
    String category

    @NonNull
    @NotBlank
    @Column(name = "vat_number")
    String vatNumber

    @Column(name = "description")
    String description

    @NonNull
    @NotBlank
    @Column(name = "legal_name")
    String legalName

    @Column(name = "delivery_rate")
    Long deliveryRate

    @Column(name = "delivery_quantity")
    Long deliveryQuantity

}


@Introspected
enum CompanyStatus {
    PENDING,
    SUSPENDED,
    DELETED,
    ACTIVE,
}
