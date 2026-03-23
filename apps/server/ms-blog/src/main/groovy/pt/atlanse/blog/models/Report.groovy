package pt.atlanse.blog.models

import groovy.transform.TupleConstructor

@TupleConstructor
class Report {

    Long id // report id

    String reason
    String status
    Object author
    String articleName
    Comment comment
    Object createdBy
    String createdAt
    String updatedAt
    List<Comment> context

}
