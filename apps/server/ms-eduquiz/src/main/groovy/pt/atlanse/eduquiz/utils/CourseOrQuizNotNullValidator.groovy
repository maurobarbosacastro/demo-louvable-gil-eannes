package pt.atlanse.eduquiz.utils

import io.micronaut.core.annotation.AnnotationValue
import io.micronaut.core.annotation.NonNull
import io.micronaut.core.annotation.Nullable
import io.micronaut.validation.validator.constraints.ConstraintValidator
import io.micronaut.validation.validator.constraints.ConstraintValidatorContext
import jakarta.inject.Singleton

interface ICourseOrQuiz {
    def getCourse()

    def getQuiz()
}

@Singleton
class CourseOrQuizNotNullValidator implements ConstraintValidator<CourseOrQuizNotNull, ICourseOrQuiz> {

    @Override
    boolean isValid(@Nullable ICourseOrQuiz value, @NonNull AnnotationValue<CourseOrQuizNotNull> annotationMetadata, @NonNull ConstraintValidatorContext context) {
        return (value.course != null || value.quiz != null)
    }
}
