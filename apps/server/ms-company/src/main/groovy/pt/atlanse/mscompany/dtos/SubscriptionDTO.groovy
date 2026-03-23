package pt.atlanse.mscompany.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.serde.annotation.Serdeable
import jakarta.annotation.Nullable
import pt.atlanse.mscompany.domains.SocialType
import pt.atlanse.mscompany.domains.SubscriptionStatus


@ToString(includeFields = true, includePackage = false, includeNames = true)
@Introspected
@TupleConstructor
@Serdeable
class SubscriptionDTO {

    @Nullable
    UUID id

    @Nullable
    SubscriptionStatus status

    @Nullable
    double price

    @Nullable
    String name

    @Nullable
    String description

}
