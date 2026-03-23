package pt.atlanse.mscompany.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated
import io.micronaut.serde.annotation.Serdeable
import jakarta.persistence.*
import jakarta.validation.constraints.NotBlank

@Introspected
@Entity
@TupleConstructor
@Table(name = "company_user")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@Serdeable
class CompanyUserEntity extends BaseEntity {

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
    @Column(name = "keycloak_user_id")
    UUID keycloakUserId

}

