package pt.atlanse.products.dtos

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank


@ToString
@Introspected
@TupleConstructor
@Serdeable.Deserializable
class GetProduct {
    @NonNull
    @NotBlank
    String name

    @NonNull
    @NotBlank
    UUID id

    @NonNull
    @NotBlank(message = "Products must have a description")
    String description

    @NonNull
    @NotBlank
    Map brand = [id: '', name: '']

    @Nullable
    UUID discount

    @Nullable
    Map category = [id: '', name: '']

    @NonNull
    @NotBlank
    Float price

    @Nullable
    String image

    @Nullable
    String state

    @Nullable
    String rating

    @Nullable
    Long productionTime

    @Nullable
    String numReviews

    @Nullable
    Object ingredients

    @Nullable
    Object sellingStores
}
