package pt.atlanse.eduquiz.utils

import io.micronaut.core.annotation.AnnotationValue
import io.micronaut.core.annotation.NonNull
import io.micronaut.core.annotation.Nullable
import io.micronaut.validation.validator.constraints.ConstraintValidator
import io.micronaut.validation.validator.constraints.ConstraintValidatorContext
import jakarta.inject.Singleton

interface ILessonOrQuiz {
    def getLesson()

    def getQuiz()
}

@Singleton
class LessonOrQuizNotNullValidator implements ConstraintValidator<LessonOrQuizNotNull, ILessonOrQuiz> {
    @Override
    boolean isValid(@Nullable ILessonOrQuiz value, @NonNull AnnotationValue<LessonOrQuizNotNull> annotationMetadata, @NonNull ConstraintValidatorContext context) {
        return (value.lesson != null || value.quiz != null)
    }
}
