package pt.atlanse.blog.DTO

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank


@Introspected
@TupleConstructor
@Serdeable.Deserializable
class ImageDTO implements MediaDTO {

    @NonNull
    @NotBlank
    private String fileName

    @NonNull
    @NotBlank
    private String extension

    @NonNull
    @NotBlank
    private String base64

    @Override
    String getFileName() {
        return null
    }

    @Override
    String getExtension() {
        return null
    }

    @Override
    String getBase64() {
        return null
    }

    @Override
    void setFileName(String fileName) {

    }

    @Override
    void setExtension(String extension) {

    }

    @Override
    void setBase64(String base64) {
        int y = base64.endsWith("==") ? 2 : base64.endsWith("=") ? 1 : 0
        long bytes = ((base64.length() * (3 / 4)) - y).longValue()

        // 5e+6 bytes ====== 5MB
        // todo get these bytes from the application.yml -> articles.file-type.image using configuration object
        if (bytes > 5e+6) {
            throw new Exception("Image too long; Max 5MB")
        }

        this.base64 = base64
    }
}
