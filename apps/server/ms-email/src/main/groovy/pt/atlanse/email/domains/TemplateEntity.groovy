package pt.atlanse.email.domains

import groovy.transform.ToString
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank

@Entity
@Table(name = "templates")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class TemplateEntity extends BaseEntity {

    @Id
    @AutoPopulated
    UUID id

    @NonNull
    @NotBlank
    String code

    @NonNull
    @NotBlank
    String name

    @NonNull
    @NotBlank
    @Column(name = "template_html", columnDefinition = "TEXT")
    String templateHtml

    @NonNull
    @NotBlank
    @Column(name = "template_json", columnDefinition = "TEXT")
    String templateJson
}
