package pt.atlanse.email.services

import io.micronaut.context.annotation.ConfigurationProperties

@ConfigurationProperties("email")
class EmailConfig {
    String host
    String port
    String from
    String username
    String password
    String timeout
    boolean sslEnable
    boolean auth
}
