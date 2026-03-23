package pt.atlanse.products.configs

import io.micronaut.context.annotation.EachProperty
import io.micronaut.context.annotation.Parameter

@EachProperty("files.file-type")
class FileConfiguration {
    String type

    String directory
    String size
    List<String> allowedFormats

    FileConfiguration(@Parameter String type) {
        this.type = type
    }
}
