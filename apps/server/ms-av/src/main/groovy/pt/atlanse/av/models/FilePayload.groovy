package pt.atlanse.av.models

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank


@Introspected
@ToString
@TupleConstructor
@Serdeable.Deserializable
class FilePayload {
    @NonNull
    @NotBlank
    String fileName

    @NonNull
    @NotBlank
    byte[] base64
}
