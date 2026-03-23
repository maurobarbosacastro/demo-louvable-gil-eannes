package pt.atlanse.event

import io.micronaut.context.annotation.Value
import io.micronaut.runtime.Micronaut
import groovy.transform.CompileStatic

@CompileStatic
class Application {

    @Value('${firebase.credentials.file}')
    String credentialsFile

    static void main(String[] args) {
        Application application = new Application()

        if (application.credentialsFile == null || !new File(application.credentialsFile).exists()) {
            // Throw an error or exception to inform the user that the firebase-credentials file is not in the file system
            throw new RuntimeException("The required resource 'firebase-credentials.json' does not exist.")
        }

        Micronaut.run(Application, args)
    }
}
