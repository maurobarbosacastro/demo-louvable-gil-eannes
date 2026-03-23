package pt.atlanse.blog.models

import groovy.transform.TupleConstructor
import pt.atlanse.blog.DTO.MediaDTO

@TupleConstructor
class Article {
    Long id
    MediaDTO image
    String status
    boolean viewMobile
    boolean viewWeb
    String createdBy
    String createdAt
    String updatedBy
    String updatedAt

    // Translations object
    String lang
    String title
    String subtitle
    String content
    String conclusion
    boolean enabled

    Long likes
    Long comments
    boolean liked

    String firstName
    String LastName
}
