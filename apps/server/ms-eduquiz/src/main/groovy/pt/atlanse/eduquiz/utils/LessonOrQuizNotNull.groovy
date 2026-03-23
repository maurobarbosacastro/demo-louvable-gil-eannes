package pt.atlanse.eduquiz.utils

import jakarta.validation.Constraint
import jakarta.validation.Payload

import java.lang.annotation.Documented
import java.lang.annotation.ElementType
import java.lang.annotation.Retention
import java.lang.annotation.RetentionPolicy
import java.lang.annotation.Target

@Documented
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = CourseOrQuizNotNullValidator.class)
@Target([ElementType.TYPE, ElementType.METHOD, ElementType.TYPE_USE])
@interface LessonOrQuizNotNull {

    String message() default "Either lesson or quiz must not be null"

    Class<?>[] groups() default []

    Class<? extends Payload>[] payload() default [];
}
