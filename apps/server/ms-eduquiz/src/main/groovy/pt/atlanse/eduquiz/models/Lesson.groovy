package pt.atlanse.eduquiz.models

import groovy.transform.ToString
import groovy.transform.TupleConstructor

@TupleConstructor
@ToString(includeFields = true, includePackage = false, includeNames = true)
class Lesson {
    UUID id
    String image
    String title
    String subtitle
    String type
    Media[] contents
    String conclusion
    String status
    String createdBy
    String createdAt
    String updatedBy
    String updatedAt
}
