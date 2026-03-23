package pt.atlanse.blog.domain

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.GeneratedValue
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.persistence.Version
import jakarta.validation.constraints.NotBlank
import java.time.LocalDateTime

@Entity
@Table(name = "category")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class Category {

    @Id
    @GeneratedValue(GeneratedValue.Type.AUTO)
    Long id

    @Version
    Long version

    @NonNull
    @NotBlank
    @Column(name = "created_by")
    String createdBy

    @GeneratedValue
    @Column(name = "created_at")
    LocalDateTime createdAt

    @NonNull
    @NotBlank
    @Column(name = "updated_by")
    String updatedBy

    @GeneratedValue
    @Column(name = "updated_at")
    LocalDateTime updatedAt

    @NonNull
    @NotBlank
    @Column(name = "name")
    String name

    @NonNull
    @NotBlank
    @Column(name = "description")
    String description

    @NonNull
    @NotBlank
    @Column(name = "status")
    String status

}
