package pt.atlanse.eduquiz.domain.compositeIds

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import jakarta.persistence.Column
import jakarta.persistence.Embeddable
import jakarta.persistence.EmbeddedId
import jakarta.persistence.Entity
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.MapsId
import jakarta.persistence.Table
import pt.atlanse.eduquiz.domain.CategoryEntity
import pt.atlanse.eduquiz.domain.ModulesEntity


@Embeddable
class ModuleCategoryId implements Serializable {
    @Column(name = "module_id")
    UUID moduleId

    @Column(name = "category_id")
    String categoryId
}

@Entity
@TupleConstructor
@Table(name = "module_category")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class ModuleCategoryEntity {
    @EmbeddedId
    ModuleCategoryId id

    @ManyToOne
    @MapsId("moduleId")
    @JoinColumn(name = "module_id", referencedColumnName = "id")
    ModulesEntity module

    @ManyToOne
    @MapsId("categoryId")
    @JoinColumn(name = "category_id", referencedColumnName = "id")
    CategoryEntity category
}
