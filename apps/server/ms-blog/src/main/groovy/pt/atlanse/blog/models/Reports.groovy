package pt.atlanse.blog.models

import io.micronaut.data.model.Page

/**
 * @deprecated Use {@link Page}<{@link Report}> instead
 * */
class Reports {
    int pageNumber
    Long totalPages
    Long totalElements
    List<Report> reports
}
