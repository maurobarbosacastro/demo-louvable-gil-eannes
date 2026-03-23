package pt.atlanse.blog.models

import com.fasterxml.jackson.annotation.JsonIgnore
import groovy.transform.TupleConstructor

@TupleConstructor
class Notification {
    String creator
    String action
    Article article
    Comment comment
    String createdAt

    @JsonIgnore
    String createdBy
}
