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
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.OneToOne
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank

import java.time.LocalDateTime

@Introspected
@Entity
@TupleConstructor
@Table(name = "company_subscription")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@Serdeable
class CompanySubscriptionEntity extends BaseEntity {

    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    @OneToOne()
    @JoinColumn(name = "company_id", referencedColumnName = "id", nullable = false)
    CompanyEntity company

    @NonNull
    @NotBlank
    @ManyToOne
    @JoinColumn(name = "subscription_id", referencedColumnName = "id", nullable = false)
    SubscriptionEntity subscription

    @NonNull
    @NotBlank
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    @Column(name = "expire_date")
    LocalDateTime expireDate

    @Column(name = "status")
    CompanySubscriptionStatus status

    @Column(name = "price")
    double price

    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    @Column(name = "start_date")
    LocalDateTime startDate

    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    @Column(name = "buy_date")
    LocalDateTime buyDate

}


@Introspected
enum CompanySubscriptionStatus {
    ACTIVE,
    EXPIRING,
    SUSPENDED,
    CANCELLED
}
