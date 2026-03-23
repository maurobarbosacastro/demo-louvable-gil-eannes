package pt.atlanse.blog.domain

import groovy.transform.ToString
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.GeneratedValue
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.Table
import jakarta.persistence.Version
import jakarta.validation.constraints.NotBlank

@Entity
@Table(name = "translations")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class TranslationEntity {

    @Id
    @GeneratedValue(GeneratedValue.Type.AUTO)
    Long id

    @Version
    Long version

    @NonNull
    @NotBlank
    @ManyToOne
    @JoinColumn(name = "article_id")
    ArticleEntity article

    @NonNull
    @NotBlank
    @Column(name = "lang")
    String lang

    @NonNull
    @NotBlank
    @Column(name = "title")
    String title

    @NonNull
    @NotBlank
    @Column(name = "subtitle")
    String subtitle

    @NonNull
    @NotBlank
    @Column(name = "content")
    String content

    @NonNull
    @NotBlank
    @Column(name = "conclusion")
    String conclusion

    @NonNull
    @NotBlank
    @Column(name = "enabled")
    boolean enabled
}
