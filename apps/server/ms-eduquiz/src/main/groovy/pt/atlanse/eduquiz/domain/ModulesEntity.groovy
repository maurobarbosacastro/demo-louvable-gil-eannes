package pt.atlanse.eduquiz.domain

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.CascadeType
import jakarta.persistence.FetchType
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.JoinTable
import jakarta.persistence.ManyToMany
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Entity
@TupleConstructor
@Table(name = "modules")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class ModulesEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @Column(name = "image_id", nullable = true)
    String image

    @NotNull(message = "Title can not be null")
    @NotBlank(message = "Title can not be blank")
    @Column(name = "title")
    String title

    @Nullable
    @Column(name = "description")
    String description

    @NotNull(message = "Status can not be null")
    @NotBlank(message = "Status can not be blank")
    @Column(name = "status")
    String status

    @Nullable
    @ManyToMany(fetch = FetchType.EAGER, cascade = CascadeType.ALL)
    @JoinTable(name = "module_category",
        joinColumns = @JoinColumn(name = "module_id"),
        inverseJoinColumns = @JoinColumn(name = "category_id"))
    List<CategoryEntity> categories = new ArrayList<>()

    @Nullable
    @Column(name = "extras")
    String extras
}
