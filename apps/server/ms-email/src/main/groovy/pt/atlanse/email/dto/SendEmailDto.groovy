package pt.atlanse.email.dto

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.core.annotation.Nullable
import jakarta.validation.constraints.NotBlank

@Introspected
@TupleConstructor
@ToString(includeFields = true, includePackage = false, includeNames = true)
class SendEmailDto {
    @NonNull
    @NotBlank
    String to

    @NonNull
    @NotBlank
    String subject

    @NonNull
    @NotBlank
    String body

    @Nullable
    String replyTo

    @Nullable
    List<AttachmentDto> attachments
}
