package pt.atlanse.eduquiz.DTO

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotNull

@Introspected
@TupleConstructor
@ToString(includePackage = false, includeFields = true, includeNames = true)
@Serdeable.Deserializable
class ContentDTO {
    @NotNull(message = 'Content can not be null')
    ImageDTO content
}
