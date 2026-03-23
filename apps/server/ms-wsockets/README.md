# GenericWebSocket - Micronaut WebSocket Server

The `GenericWebSocket` class represents a WebSocket server for handling WebSocket connections and messages in a Micronaut application. This class is based on the Micronaut WebSocket guide and includes basic security implementation.

## Usage

1. **Class Overview:**
    - `ms-wsockets` is a Micronaut `ServerWebSocket` annotated class.
    - It handles WebSocket connections and messages.
    - Security rules for this WebSocket are anonymous, meaning no authentication is required.

2. **Dependencies:**
    - This class relies on Micronaut's WebSocket functionality and basic security.
    - Micronaut WebSocketBroadcaster is used to send messages to connected clients.

3. **WebSocket Events:**
    - The following WebSocket events are defined in this class:

    - `@OnOpen`: Invoked when a WebSocket connection is established. It broadcasts a message to all connected clients in the same topic.

    - `@OnMessage`: Invoked when a WebSocket message is received. It broadcasts the received message to all clients in the same topic.

    - `@OnClose`: Invoked when a WebSocket connection is closed. It broadcasts a message indicating that the user has left the topic.

4. **Logger:**
    - The class uses a logger to log WebSocket events.

5. **Helper Methods:**
    - `isValid(String topic)`: A helper method to check if a WebSocket session is valid for a given topic.

## Example Usage

```java
// Example of how to use GenericWebSocket in your Micronaut application
@Secured(SecurityRule.IS_ANONYMOUS)
@ServerWebSocket("/ws/{topic}/{username}")
public class GenericWebSocket {

    // Constructor with WebSocketBroadcaster injection

    // onOpen, onMessage, onClose methods

    // Helper methods and logger

}
```

## Handler

This microservice also have a handler to manager diferent groups of users
The usage is simple:
Inject the handle and in the @onOpen meoth add the user to the group
To send a message use the handle broacast metho

# Testing the Chat Application Using Postman WebSocket

This section provides instructions on how to set up a simple chat application using Postman's WebSocket feature. By following these steps, you can simulate two users communicating through a WebSocket connection.

## Prerequisites

- [Postman](https://www.postman.com/) installed on your computer.
- A running WebSocket server (e.g., `ms-wsockets`) on `ws://localhost:8094/ws/{topic}/{username}`. Ensure your WebSocket server is configured to handle incoming WebSocket connections.

## Steps to Set Up the Chat

1. **Open Postman:**
   Launch the Postman application on your computer.

2. **Create a WebSocket Request (User 1):**

    - Click on the "New" button in Postman.
    - Select "WebSocket Request" from the request types dropdown.

3. **Configure WebSocket Connection (User 1):**

    - In the "Connect to WebSocket" dialog, enter the following WebSocket URL:
      ```
      ws://localhost:8094/ws/{topic}/{username}
      ```
      Replace `{topic}` and `{username}` with appropriate values or placeholders as per your WebSocket server configuration.

    - Click the "Connect" button to initiate the WebSocket connection for User 1.

4. **Create Another WebSocket Request (User 2):**

    - Click on the "New" button in Postman.
    - Select "WebSocket Request" from the request types dropdown.

5. **Configure WebSocket Connection (User 2):**

    - In the "Connect to WebSocket" dialog, enter the same WebSocket URL as User 1:
      ```
      ws://localhost:8094/ws/{topic}/{username}
      ```

      Replace `{topic}` with the same topic used in the first user and use a different `{username}` to represent User 2.

    - Click the "Connect" button to initiate the WebSocket connection for User 2.

6. **Start Chatting:**
    - User 1 and User 2 are now connected to the WebSocket server.
    - You can send messages from User 1 by entering a message in the "User 1" WebSocket request tab and clicking the "Send" button.
    - Similarly, send messages from User 2 using the "User 2" WebSocket request tab.

7. **Observe Chat Messages:**
    - You will observe that messages sent from User 1 are received by User 2 and vice versa.
    - This simulates a basic chat interaction between two WebSocket clients.

8. **Close WebSocket Connections:**
    - When you're done testing the chat, you can close the WebSocket connections by clicking the "Disconnect" button in the WebSocket request tabs for both User 1 and User 2.
