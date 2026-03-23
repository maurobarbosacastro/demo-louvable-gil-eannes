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
@Table(name = "subscription")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@Serdeable
class SubscriptionEntity extends BaseEntity {

    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @Column(name = "status")
    SubscriptionStatus status

    @NonNull
    @NotBlank
    @Column(name = "price")
    double price

    @NonNull
    @NotBlank
    @Column(name = "name")
    String name

    @Column(name = "description")
    String description

}

@Introspected
enum SubscriptionStatus {
    DELETED,
    ACTIVE,
}

