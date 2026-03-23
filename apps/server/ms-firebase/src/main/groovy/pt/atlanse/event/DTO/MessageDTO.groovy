package pt.atlanse.event.DTO

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull


@Introspected
@ToString(includePackage = false, includeFields = false, includeNames = true)
@TupleConstructor
@Serdeable.Deserializable

class MessageDTO {

    @NotNull
    @NotBlank
    String title

    @NotNull
    @NotBlank
    String body

}
