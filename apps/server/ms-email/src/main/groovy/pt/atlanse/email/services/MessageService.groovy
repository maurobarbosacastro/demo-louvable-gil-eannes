package pt.atlanse.email.services

import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.email.dto.EmailDto
import pt.atlanse.email.dto.SendEmailDto
import pt.atlanse.email.exceptions.CustomException

import javax.activation.DataHandler
import javax.mail.Address
import javax.mail.Message
import javax.mail.Transport
import javax.mail.internet.InternetAddress
import javax.mail.internet.MimeBodyPart
import javax.mail.internet.MimeMessage
import javax.mail.internet.MimeMultipart
import javax.mail.util.ByteArrayDataSource

@Singleton
class MessageService {

    @Inject
    SMTPStartup smtp

    private MimeMessage setMessage(String subject, content, List<String> address, boolean isHtml = false) {
        MimeMessage message = new MimeMessage(smtp.session)
        message.setFrom(new InternetAddress(smtp.emailConfig.from))
        message.setSubject(subject)

        message.addRecipients(Message.RecipientType.TO, address.collect { new InternetAddress(it) } as Address[])

        if (isHtml) {
            MimeBodyPart mimeBodyPart = new MimeBodyPart()
            mimeBodyPart.setContent(content, 'text/html; charset=utf-8')
            MimeMultipart mimeMultipart = new MimeMultipart()
            mimeMultipart.addBodyPart(mimeBodyPart)

            message.setContent(mimeMultipart)
            return message
        }

        message.setText(content as String)
        return message
    }

    void send(String subject, String content, EmailDto payload, boolean isHtml = false) {
        try {
            MimeMessage message = setMessage(subject, content, [payload.to], isHtml)
            Transport.send(message)
        } catch (Exception e) {
            throw new CustomException(
                "Error occurred while trying to send email to address ${ payload.to }",
                e.message,
                HttpStatus.INTERNAL_SERVER_ERROR
            )
        }
    }

    void sendRaw(SendEmailDto payload) {
        try {
            MimeMessage message = new MimeMessage(smtp.session)
            message.setFrom(new InternetAddress(smtp.emailConfig.from))
            message.setSubject(payload.subject)
            message.addRecipients(Message.RecipientType.TO, [new InternetAddress(payload.to)] as Address[])

            if (payload.replyTo) {
                message.setReplyTo([new InternetAddress(payload.replyTo)] as Address[])
            }

            MimeMultipart multipart = new MimeMultipart()

            MimeBodyPart htmlPart = new MimeBodyPart()
            htmlPart.setContent(payload.body, 'text/html; charset=utf-8')
            multipart.addBodyPart(htmlPart)

            if (payload.attachments) {
                payload.attachments.each { attachment ->
                    byte[] fileData = Base64.decoder.decode(attachment.data)
                    ByteArrayDataSource dataSource = new ByteArrayDataSource(fileData, attachment.mimeType)
                    MimeBodyPart attachmentPart = new MimeBodyPart()
                    attachmentPart.setDataHandler(new DataHandler(dataSource))
                    attachmentPart.setFileName(attachment.filename)
                    multipart.addBodyPart(attachmentPart)
                }
            }

            message.setContent(multipart)
            Transport.send(message)
        } catch (Exception e) {
            throw new CustomException(
                "Error occurred while trying to send raw email to address ${ payload.to }",
                e.message,
                HttpStatus.INTERNAL_SERVER_ERROR
            )
        }
    }
}
