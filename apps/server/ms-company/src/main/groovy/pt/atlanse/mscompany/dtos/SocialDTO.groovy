package pt.atlanse.mscompany.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.serde.annotation.Serdeable
import jakarta.annotation.Nullable
import pt.atlanse.mscompany.domains.SocialType


@ToString(includeFields = true, includePackage = false, includeNames = true)
@Introspected
@TupleConstructor
@Serdeable
class SocialDTO {

    @Nullable
    String id

    @Nullable
    SocialType type

    @Nullable
    String link

    @Nullable
    String imageId

    @Nullable
    String company

    @Nullable
    ImageDTO image

}
