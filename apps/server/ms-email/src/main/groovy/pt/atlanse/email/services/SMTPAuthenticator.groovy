package pt.atlanse.email.services

import javax.mail.PasswordAuthentication

class SMTPAuthenticator extends javax.mail.Authenticator {

    String username
    String password

    SMTPAuthenticator(String username, String password) {
        this.username = username
        this.password = password
    }

    PasswordAuthentication getPasswordAuthentication() {
        return new PasswordAuthentication(username, password)
    }
}
