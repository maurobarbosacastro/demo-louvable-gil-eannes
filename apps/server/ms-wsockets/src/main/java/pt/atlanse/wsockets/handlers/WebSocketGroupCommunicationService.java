package pt.atlanse.wsockets.handlers;
import io.micronaut.websocket.WebSocketSession;
import jakarta.inject.Singleton;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.http.WebSocket;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ConcurrentHashMap;

/**
 * Manages WebSocket sessions organized into groups for efficient communication.
 */
@Singleton // Assumes this class is a singleton.
public class WebSocketGroupCommunicationService {
    private static final Logger LOG = LoggerFactory.getLogger(WebSocket.class);
    private final Map<String, List<WebSocketSession>> sessionsByGroup = new ConcurrentHashMap<>();

    /**
     * Adds a WebSocket session to a specified group.
     *
     * @param group   The group identifier.
     * @param session The WebSocket session to add.
     */
    public void addSessionToGroup(String group, WebSocketSession session) {
        LOG.info("Add session: " + session.getId() + " to: " + group);
        sessionsByGroup.computeIfAbsent(group, k -> new ArrayList<>()).add(session);
    }

    /**
     * Removes a WebSocket session from a specified group.
     *
     * @param group   The group identifier.
     * @param session The WebSocket session to remove.
     */
    public void removeSessionFromGroup(String group, WebSocketSession session) {
        List<WebSocketSession> sessions = sessionsByGroup.get(group);
        if (sessions != null) {
            sessions.remove(session);
        }
    }

    /**
     * Broadcasts a message to all WebSocket sessions in a specified group.
     *
     * @param group   The group identifier.
     * @param message The message to broadcast.
     */
    public void broadcastToGroup(String group, String message) {
        List<WebSocketSession> sessions = sessionsByGroup.get(group);
        if (sessions != null) {
            List<WebSocketSession> sessionsToRemove = new ArrayList<>();

            for (WebSocketSession session : sessions) {
                if (session.isOpen()) {
                    try {
                        CompletableFuture<String> sendFuture = session.sendAsync(message);
                        sendFuture.thenAccept(result -> {
                            LOG.info("Message sent successfully to: " + session.getId());
                        }).exceptionally(ex -> {
                            LOG.error("Failed to send message to: " + session.getId(), ex);
                            // Handle the exception, e.g., update session status or remove from map
                            return null;
                        });
                    } catch (Exception e) {
                        LOG.error("Error sending message to: " + session.getId(), e);
                    }
                } else {
                    LOG.warn("Skipping closed session: " + session.getId());
                    LOG.warn("Marking closed session for removal: " + session.getId());
                    sessionsToRemove.add(session);
                }
            }
            // Remove closed sessions outside the loop
            sessions.removeAll(sessionsToRemove);
        }
    }
}
