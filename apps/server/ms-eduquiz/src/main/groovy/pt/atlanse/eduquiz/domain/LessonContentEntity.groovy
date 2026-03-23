package pt.atlanse.eduquiz.domain

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.Table

@Entity
@TupleConstructor
@Table(name = "lessons_content")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class LessonContentEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @Nullable
    @ManyToOne
    @JoinColumn(name = "lesson_id", referencedColumnName = "id")
    LessonEntity lesson

    @Nullable
    @ManyToOne
    @JoinColumn(name = "content_id", referencedColumnName = "id")
    ContentEntity content
}
