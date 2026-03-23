package pt.atlanse.av.domains

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.GeneratedValue
import io.micronaut.data.annotation.Id
import io.micronaut.data.annotation.MappedEntity
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank


@MappedEntity
@TupleConstructor
@Serdeable.Deserializable
class FileScan {

    @Id
    @GeneratedValue
    String id

    @NonNull
    @NotBlank
    String name

    @NonNull
    @NotBlank
    String md5

    @NonNull
    Integer bytes

    @NonNull
    String clamAv
}
