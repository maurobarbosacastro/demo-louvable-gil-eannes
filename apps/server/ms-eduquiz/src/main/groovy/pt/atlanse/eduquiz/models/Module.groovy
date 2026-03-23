package pt.atlanse.eduquiz.models

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import pt.atlanse.eduquiz.DTO.ImageDTO

@TupleConstructor
@ToString(includeNames = true, includePackage = false, includeFields = true)
class Module {
    UUID id
    ImageDTO image
    String title
    String description
    String createdBy
    String createdAt
    String updatedBy
    String updatedAt
}
