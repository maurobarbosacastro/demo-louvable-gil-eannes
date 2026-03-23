package pt.atlanse.mscompany.domains

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
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank

@Introspected
@Entity
@TupleConstructor
@Table(name = "socials")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@Serdeable
class SocialsEntity extends BaseEntity {

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
    @Column(name = "type")
    SocialType type

    @NonNull
    @NotBlank
    @Column(name = "link")
    String link

    @NonNull
    @NotBlank
    @Column(name = "image_id")
    String image
}



@Introspected
enum SocialType {
    INSTAGRAM,
    FACEBOOK,
    X,
    TIKTOK
}
