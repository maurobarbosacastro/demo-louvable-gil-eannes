package pt.atlanse.eduquiz.DTO

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@TupleConstructor
@Introspected
@ToString(includePackage = false, includeFields = true, includeNames = true)
@Serdeable.Deserializable
class ModulesDTO {
    @Nullable
    ImageDTO image

    @NotNull(message = 'Title can not be null')
    @NotBlank
    String title

    @Nullable
    String description

    @Nullable
    List<CategoryDTO> categories

    @Nullable
    String extras

    @Nullable
    String status
}
