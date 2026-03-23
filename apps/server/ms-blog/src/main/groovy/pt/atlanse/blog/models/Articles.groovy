package pt.atlanse.blog.models

import groovy.transform.TupleConstructor
import io.micronaut.data.model.Page

/**
 * @deprecated Use {@link Page}<{@link Article}> instead
 * */
@TupleConstructor
class Articles {
    int pageNumber
    Long totalPages
    Long totalElements
    List<Article> content
}
