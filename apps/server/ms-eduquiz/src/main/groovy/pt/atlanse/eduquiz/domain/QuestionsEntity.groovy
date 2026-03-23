package pt.atlanse.eduquiz.domain

import com.fasterxml.jackson.annotation.JsonFormat
import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Nullable
import io.micronaut.data.annotation.AutoPopulated
import jakarta.persistence.CascadeType
import jakarta.persistence.FetchType
import jakarta.persistence.OneToMany
import java.time.LocalDateTime
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.ManyToOne
import jakarta.persistence.Table
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Entity
@Table(name = "questions")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class QuestionsEntity extends BaseEntity {
    @Id
    @AutoPopulated
    UUID id

    @Nullable
    @Column(name = "type")
    String type

    @NotNull(message = "Description can not be null")
    @NotBlank(message = "Description can not be blank")
    @Column(name = "description")
    String description

    @Column(name = "image_id", nullable = true)
    String image

    @Nullable
    @ManyToOne
    @JoinColumn(name = "category_id", referencedColumnName = "id")
    CategoryEntity category

    @Nullable
    @Column(name = "points")
    Long points

    @Nullable
    @Column(name = "begin_date")
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    LocalDateTime beginDate

    @Nullable
    @Column(name = "end_date")
    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss")
    LocalDateTime endDate

    @Nullable
    @Column(name = "extras")
    String extras

    @OneToMany(mappedBy = "question", cascade = [CascadeType.ALL], fetch = FetchType.EAGER)
    List<AnswersEntity> answers = new ArrayList<>()

    @Nullable
    String status
}
