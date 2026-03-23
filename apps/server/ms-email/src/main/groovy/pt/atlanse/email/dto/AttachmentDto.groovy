package pt.atlanse.email.dto

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import jakarta.validation.constraints.NotBlank

@Introspected
@TupleConstructor
@ToString(includeFields = true, includePackage = false, includeNames = true)
class AttachmentDto {
    @NonNull
    @NotBlank
    String filename

    @NonNull
    @NotBlank
    String data

    @NonNull
    @NotBlank
    String mimeType
}
