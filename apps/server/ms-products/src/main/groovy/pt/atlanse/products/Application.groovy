package pt.atlanse.products

import io.micronaut.runtime.Micronaut
import groovy.transform.CompileStatic
import io.swagger.v3.oas.annotations.OpenAPIDefinition
import io.swagger.v3.oas.annotations.info.Info

@OpenAPIDefinition(
    info = @Info(
        title = "ms-product",
        version = "0.0"
    )
)
@CompileStatic
class Application {

    static void main(String[] args) {
        Micronaut.run(Application, args)
    }
}
