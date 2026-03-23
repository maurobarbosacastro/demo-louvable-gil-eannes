package pt.atlanse.event.service

import com.google.auth.oauth2.GoogleCredentials
import com.google.common.util.concurrent.ListenableFuture
import com.google.firebase.FirebaseApp
import com.google.firebase.FirebaseOptions
import com.google.firebase.messaging.FirebaseMessaging
import com.google.firebase.messaging.Message
import com.google.firebase.messaging.Notification
import groovy.util.logging.Slf4j
import io.micronaut.context.annotation.Value
import pt.atlanse.event.DTO.MessageDTO
import jakarta.inject.Singleton

import java.util.concurrent.CompletableFuture

@Slf4j
@Singleton
class FirebaseService {
    private final FirebaseMessaging firebaseMessaging

    FirebaseService(@Value('${firebase.credentials.file}') String credentialsFile) throws IOException {
        log.info "Loading FirebaseService Singleton"
        FileInputStream serviceAccount = new FileInputStream(credentialsFile)

        FirebaseOptions options = FirebaseOptions.builder()
            .setCredentials(GoogleCredentials.fromStream(serviceAccount))
            .build()

        FirebaseApp.initializeApp(options)
        firebaseMessaging = FirebaseMessaging.getInstance()
    }

    CompletableFuture<String> sendGroupMessage(String group, MessageDTO message) {
        log.info "Sending topic message"
        CompletableFuture<String> completableFuture = new CompletableFuture<>()
        Message firebaseMessage = Message.builder()
            .setTopic('/topics/' + group)
            .setNotification(
                Notification.builder()
                    .setTitle(message.title)
                    .setBody(message.body)
                    .build()
            )
            .putData("message", message.body)
            .build()

        ListenableFuture<String> listenableFuture = firebaseMessaging.sendAsync(firebaseMessage)
        CompletableFuture<String> future = CompletableFuture.supplyAsync {
            try {
                return listenableFuture.get()
            } catch (Exception e) {
                throw new RuntimeException(e)
            }
        }

        future.whenComplete { result, error ->
            if (error != null) {
                completableFuture.completeExceptionally(error)
            } else {
                completableFuture.complete(result)
            }
        }

        return completableFuture
    }


    CompletableFuture<String> sendIndividualMessage(String token, MessageDTO message) {
        log.info "Sending token message"
        CompletableFuture<String> completableFuture = new CompletableFuture<>()
        Message firebaseMessage = Message.builder()
            .setToken(token)
            .setNotification(
                Notification.builder()
                    .setTitle(message.title)
                    .setBody(message.body)
                    .build()
            )
            .putData("message", message.body)
            .build()

        ListenableFuture<String> listenableFuture = firebaseMessaging.sendAsync(firebaseMessage)
        CompletableFuture<String> future = CompletableFuture.supplyAsync {
            try {
                return listenableFuture.get()
            } catch (Exception e) {
                throw new RuntimeException(e)
            }
        }

        future.whenComplete { result, error ->
            if (error != null) {
                completableFuture.completeExceptionally(error)
            } else {
                completableFuture.complete(result)
            }
        }

        return completableFuture
    }
}
