package pt.atlanse.PROJECT_NAME.domains

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.data.annotation.AutoPopulated

import javax.persistence.Entity
import javax.persistence.Id
import javax.persistence.Table

@Entity
@Table
@ToString(includePackage = false, includeNames = false, includeFields = false)
@TupleConstructor
class ENTITY_NAME {
    @Id
    @AutoPopulated
    UUID id
}
