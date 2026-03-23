package pt.atlanse.products.dtos

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.serde.annotation.Serdeable
import jakarta.annotation.Nullable
import jakarta.validation.constraints.NotBlank


@Introspected
@TupleConstructor
@Serdeable.Deserializable
class ImageDTO implements Media {

    @NonNull
    @NotBlank
    String fileName

    @NotBlank
    @Nullable
    String extension

    @NonNull
    @NotBlank
    String base64

    @Override
    String getFileName() {
        return fileName
    }

    @Override
    String getExtension() {
        return extension
    }

    @Override
    String getBase64() {
        return base64
    }

    @Override
    void setFileName(String fileName) {
        this.fileName = fileName
    }

    @Override
    void setExtension(String extension) {
        this.extension = extension
    }

    @Override
    void setBase64(String base64) {
        int y = base64.endsWith("==") ? 2 : base64.endsWith("=") ? 1 : 0
        long bytes = ((base64.length() * (3 / 4)) - y).longValue()

        // 5e+6 bytes ====== 5MB
        if (bytes > 5e+6) {
            throw new Exception("Image too long; Max 5MB")
        }

        this.base64 = base64
    }
}
