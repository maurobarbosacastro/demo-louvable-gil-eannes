/**
 * Represents a generic WebSocket server for handling WebSocket connections and messages.
 * This class is based on the Micronaut WebSocket guide
 * (https://guides.micronaut.io/latest/micronaut-websocket-maven-java.html)
 * and includes basic security implementation.
 */
package pt.atlanse.wsockets.controllers;

import io.micronaut.security.annotation.Secured;
import io.micronaut.security.rules.SecurityRule;
import io.micronaut.websocket.WebSocketBroadcaster;
import io.micronaut.websocket.WebSocketSession;
import io.micronaut.websocket.annotation.OnClose;
import io.micronaut.websocket.annotation.OnMessage;
import io.micronaut.websocket.annotation.OnOpen;
import io.micronaut.websocket.annotation.ServerWebSocket;
import org.reactivestreams.Publisher;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.function.Predicate;

/**
 * This class represents a WebSocket server that handles WebSocket connections and messages.
 * - The class is annotated with @ServerWebSocket, indicating that it's a Micronaut WebSocket server.
 * - The @Secured annotation specifies that security rules for this WebSocket are anonymous, meaning no authentication is required.
 * - WebSocket events, such as onOpen, onMessage, and onClose, are defined along with corresponding actions.
 * - A WebSocketBroadcaster is injected to send messages to connected clients.
 */
@Secured(SecurityRule.IS_ANONYMOUS)
@ServerWebSocket("/ws/{topic}/{username}")
public class GenericWebSocket {

    // Initialize a logger for logging WebSocket events
    private static final Logger LOG = LoggerFactory.getLogger(GenericWebSocket.class);

    // Inject a WebSocketBroadcaster to send messages to connected clients
    private final WebSocketBroadcaster broadcaster;

    public GenericWebSocket(WebSocketBroadcaster broadcaster) {
        this.broadcaster = broadcaster;
    }

    /**
     * Invoked when a WebSocket connection is established.
     * Broadcasts a message to all connected clients in the same topic.
     *
     * @param topic    The topic associated with the WebSocket connection.
     * @param username The username associated with the WebSocket connection.
     * @param session  The WebSocket session.
     * @return A Publisher for broadcasting messages.
     */
    @OnOpen
    public Publisher<String> onOpen(String topic, String username, WebSocketSession session) {
        log("onOpen", session, username, topic);

        // Broadcast a message to all connected clients in the same topic
        return broadcaster.broadcast(String.format("[%s] Joined %s!", username, topic), isValid(topic));
    }

    /**
     * Invoked when a WebSocket message is received.
     * Broadcasts the received message to all clients in the same topic.
     *
     * @param topic    The topic associated with the WebSocket connection.
     * @param username The username associated with the WebSocket connection.
     * @param message  The received message.
     * @param session  The WebSocket session.
     * @return A Publisher for broadcasting messages.
     */
    @OnMessage
    public Publisher<String> onMessage(
            String topic,
            String username,
            String message,
            WebSocketSession session) {

        log("onMessage", session, username, topic);

        // Broadcast the received message to all clients in the same topic
        return broadcaster.broadcast(String.format("[%s] %s", username, message), isValid(topic));
    }

    /**
     * Invoked when a WebSocket connection is closed.
     * Broadcasts a message indicating that the user has left the topic.
     *
     * @param topic    The topic associated with the WebSocket connection.
     * @param username The username associated with the WebSocket connection.
     * @param session  The WebSocket session.
     * @return A Publisher for broadcasting messages.
     */
    @OnClose
    public Publisher<String> onClose(
            String topic,
            String username,
            WebSocketSession session) {

        log("onClose", session, username, topic);

        // Broadcast a message indicating the user has left the topic
        return broadcaster.broadcast(String.format("[%s] Leaving %s!", username, topic), isValid(topic));
    }

    /**
     * Helper method to log WebSocket events.
     *
     * @param event    The WebSocket event (e.g., onOpen, onMessage, onClose).
     * @param session  The WebSocket session.
     * @param username The username associated with the WebSocket connection.
     * @param topic    The topic associated with the WebSocket connection.
     */
    private void log(String event, WebSocketSession session, String username, String topic) {
        LOG.info("* WebSocket: {} received for session {} from '{}' regarding '{}'",
                event, session.getId(), username, topic);
    }

    /**
     * Helper method to check if a WebSocket session is valid for a given topic.
     *
     * @param topic The topic to validate against.
     * @return A Predicate that checks if the WebSocket session is valid for the specified topic.
     */
    private Predicate<WebSocketSession> isValid(String topic) {
        return s -> topic.equalsIgnoreCase(s.getUriVariables().get("topic", String.class, null));
    }
}
