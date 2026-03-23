package pt.atlanse.email.services

import io.micronaut.context.event.ApplicationEventListener
import io.micronaut.runtime.server.event.ServerStartupEvent
import jakarta.inject.Singleton
import pt.atlanse.email.services.EmailConfig
import pt.atlanse.email.services.SMTPAuthenticator
import io.micronaut.context.annotation.Value

import javax.mail.Authenticator
import javax.mail.Session

@Singleton
class SMTPStartup implements ApplicationEventListener<ServerStartupEvent> {
    EmailConfig emailConfig

    // SMTP data
    Properties properties
    Authenticator auth
    Session session

    SMTPStartup(EmailConfig emailConfig) {
        this.emailConfig = emailConfig
    }

    /**
     * Handle an application event.
     *
     * @param event the event to respond to
     */
    @Override
    void onApplicationEvent(ServerStartupEvent event) {
        setProperties()
        setAuth()
        setSession()
    }



    void setProperties() {


        properties = new Properties()

        String host = emailConfig.host
        String port = emailConfig.port
        String auth = emailConfig.auth

        String username = "your-email@gmail.com"
        String password = "your-password"

        properties.put("mail.smtp.host", System.getenv("EMAIL_SMTP_HOST"))
        properties.put("mail.smtp.port", System.getenv("EMAIL_SMTP_PORT"))
        properties.put("mail.smtp.starttls.enable", System.getenv("EMAIL_SMTP_AUTH"))  // Enable TLS
        properties.put("mail.smtp.auth", System.getenv("EMAIL_SMTP_AUTH"))
        properties.put("mail.smtp.connectiontimeout", "5000")  // Timeout in milliseconds
        properties.put("mail.smtp.timeout", "5000")

        System.printf("properties - Host: %s\n", properties."mail.smtp.host")
        System.printf("properties - Port: %s\n", properties."mail.smtp.port")
        System.printf("properties - Auth: %s\n", properties."mail.smtp.auth")

        System.printf("variable - Host: %s\n", host)
        System.printf("variable - Port: %s\n", port)
        System.printf("variable - Auth: %s\n", auth)

        System.printf("emailConfig - Host: %s\n", emailConfig.host)
        System.printf("emailConfig - Port: %s\n", emailConfig.port)
        System.printf("emailConfig - Auth: %s\n", emailConfig.auth)

        System.printf("env - EMAIL_SMTP_PORT: %s\n", System.getenv("EMAIL_SMTP_PORT"))

        /*
        properties."mail.smtp.port" = emailConfig.port
        properties."mail.smtp.host" = emailConfig.host
        properties."mail.smtp.auth" = emailConfig.auth
        properties."mail.smtp.connectiontimeout" = emailConfig.timeout
        properties."mail.smtp.timeout" = emailConfig.timeout
        properties."mail.transport.protocol" = "smtp"

        if (emailConfig.getSslEnable()) {
            properties."mail.smtp.starttls.enable" = "true"
            properties."mail.smtp.socketFactory.port" = emailConfig.port
            properties."mail.smtp.socketFactory.class" = "javax.net.ssl.SSLSocketFactory"
        }
        */
    }

    void setAuth() {
        auth = new SMTPAuthenticator(emailConfig.username, emailConfig.password)
    }

    void setSession() {
        session = Session.getInstance(properties, auth)
    }
}

