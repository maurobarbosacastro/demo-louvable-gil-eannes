package pt.atlanse.blog.models

import groovy.transform.TupleConstructor
import io.micronaut.data.model.Page

/**
 * @deprecated Use {@link Page}<{@link Comment}> instead
 * */
@TupleConstructor
class Comments {
    int pageNumber
    Long totalPages
    Long totalElements
    List<Comment> content
}
