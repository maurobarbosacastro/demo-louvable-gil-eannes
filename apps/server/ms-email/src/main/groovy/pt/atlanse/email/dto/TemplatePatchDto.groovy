package pt.atlanse.email.dto

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable

@Introspected
@TupleConstructor
@ToString(includeFields = true, includePackage = false, includeNames = true)
class TemplatePatchDto {

    @Nullable
    String name

    @Nullable
    String templateHtml

    @Nullable
    String templateJson
}
