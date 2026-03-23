package test

import (
	"encoding/json"
	"ms-tagpeak/external/notifications"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	testKeycloakURL          = os.Getenv("K_URL")
	testKeycloakClientID     = os.Getenv("IN_CLIENT_ID")
	testKeycloakClientSecret = os.Getenv("IN_CLIENT_SECRET")
	testKeycloakRealm        = os.Getenv("K_REALM")
)

func TestSaveToken(t *testing.T) {
	if testKeycloakURL == "" {
		t.Skip("Skipping test: testKeycloakURL not set")
	}

	t.Run("should save token successfully", func(t *testing.T) {
		expectedToken := notifications.UserToken{
			UUID:     uuid.New(),
			UserUUID: "user-123",
			Token:    "fcm-token-456",
			BaseEntity: notifications.BaseEntity{
				CreatedAt: time.Now(),
				CreatedBy: "system",
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/token" {
				t.Errorf("Expected path '/token', got %s", r.URL.Path)
			}
			if r.Method != http.MethodPost {
				t.Errorf("Expected POST method, got %s", r.Method)
			}

			var body notifications.TokenSave
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("Failed to decode request body: %v", err)
			}

			if body.UserUUID != "user-123" || body.Token != "fcm-token-456" {
				t.Errorf("Unexpected body: %+v", body)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(expectedToken)
		}))
		defer server.Close()

		os.Setenv("MS_FIREBASE_GO_URL", server.URL+"/")
		os.Setenv("KEYCLOAK_URL", testKeycloakURL)
		os.Setenv("INTERNAL_KEYCLOAK_CLIENT_ID", testKeycloakClientID)
		os.Setenv("INTERNAL_KEYCLOAK_HOST_SECRET", testKeycloakClientSecret)
		os.Setenv("KEYCLOAK_REALM", testKeycloakRealm)

		result, err := notifications.SaveToken(notifications.TokenSave{
			UserUUID: "user-123",
			Token:    "fcm-token-456",
		})
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.UserUUID != expectedToken.UserUUID {
			t.Errorf("Expected UserUUID %s, got %s", expectedToken.UserUUID, result.UserUUID)
		}
		if result.Token != expectedToken.Token {
			t.Errorf("Expected Token %s, got %s", expectedToken.Token, result.Token)
		}
	})
}

func TestDeleteToken(t *testing.T) {
	if testKeycloakURL == "" {
		t.Skip("Skipping test: testKeycloakURL not set")
	}

	t.Run("should delete token successfully", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/token/fcm-token-456" {
				t.Errorf("Expected path '/token/fcm-token-456', got %s", r.URL.Path)
			}
			if r.Method != http.MethodDelete {
				t.Errorf("Expected DELETE method, got %s", r.Method)
			}

			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		os.Setenv("MS_FIREBASE_GO_URL", server.URL+"/")
		os.Setenv("KEYCLOAK_URL", testKeycloakURL)
		os.Setenv("INTERNAL_KEYCLOAK_CLIENT_ID", testKeycloakClientID)
		os.Setenv("INTERNAL_KEYCLOAK_HOST_SECRET", testKeycloakClientSecret)
		os.Setenv("KEYCLOAK_REALM", testKeycloakRealm)

		err := notifications.DeleteToken("fcm-token-456")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestDeleteTokensByUser(t *testing.T) {
	if testKeycloakURL == "" {
		t.Skip("Skipping test: testKeycloakURL not set")
	}

	t.Run("should delete tokens by user successfully", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/token/user/user-123" {
				t.Errorf("Expected path '/token/user/user-123', got %s", r.URL.Path)
			}
			if r.Method != http.MethodDelete {
				t.Errorf("Expected DELETE method, got %s", r.Method)
			}

			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		os.Setenv("MS_FIREBASE_GO_URL", server.URL+"/")
		os.Setenv("KEYCLOAK_URL", testKeycloakURL)
		os.Setenv("INTERNAL_KEYCLOAK_CLIENT_ID", testKeycloakClientID)
		os.Setenv("INTERNAL_KEYCLOAK_HOST_SECRET", testKeycloakClientSecret)
		os.Setenv("KEYCLOAK_REALM", testKeycloakRealm)

		err := notifications.DeleteTokensByUser("user-123")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestCreateNotification(t *testing.T) {
	if testKeycloakURL == "" {
		t.Skip("Skipping test: testKeycloakURL not set")
	}

	t.Run("should create notification successfully", func(t *testing.T) {
		notifUUID := uuid.New()
		notifDate := time.Now()
		expectedNotification := notifications.Notification{
			UUID:  notifUUID,
			Title: "Test Notification",
			Body:  "This is a test",
			Date:  notifDate,
			State: notifications.NotificationStateDraft,
			BaseEntity: notifications.BaseEntity{
				CreatedAt: time.Now(),
				CreatedBy: "system",
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/notification" {
				t.Errorf("Expected path '/notification', got %s", r.URL.Path)
			}
			if r.Method != http.MethodPost {
				t.Errorf("Expected POST method, got %s", r.Method)
			}

			var body notifications.CreateNotificationDto
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("Failed to decode request body: %v", err)
			}

			if body.Title != "Test Notification" || body.Content != "This is a test" {
				t.Errorf("Unexpected body: %+v", body)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(expectedNotification)
		}))
		defer server.Close()

		os.Setenv("MS_FIREBASE_GO_URL", server.URL+"/")
		os.Setenv("KEYCLOAK_URL", testKeycloakURL)
		os.Setenv("INTERNAL_KEYCLOAK_CLIENT_ID", testKeycloakClientID)
		os.Setenv("INTERNAL_KEYCLOAK_HOST_SECRET", testKeycloakClientSecret)
		os.Setenv("KEYCLOAK_REALM", testKeycloakRealm)

		result, err := notifications.CreateNotification(notifications.CreateNotificationDto{
			Title:   "Test Notification",
			Content: "This is a test",
			Date:    notifDate,
		})
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.Title != expectedNotification.Title {
			t.Errorf("Expected Title %s, got %s", expectedNotification.Title, result.Title)
		}
		if result.State != notifications.NotificationStateDraft {
			t.Errorf("Expected State %s, got %s", notifications.NotificationStateDraft, result.State)
		}
	})
}

func TestSendNotification(t *testing.T) {
	if testKeycloakURL == "" {
		t.Skip("Skipping test: testKeycloakURL not set")
	}

	t.Run("should send notification successfully", func(t *testing.T) {
		notifUUID := uuid.New().String()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectedPath := "/notification/" + notifUUID + "/send"
			if r.URL.Path != expectedPath {
				t.Errorf("Expected path '%s', got %s", expectedPath, r.URL.Path)
			}
			if r.Method != http.MethodPost {
				t.Errorf("Expected POST method, got %s", r.Method)
			}

			// Expect empty body for the new unified endpoint
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		os.Setenv("MS_FIREBASE_GO_URL", server.URL+"/")
		os.Setenv("KEYCLOAK_URL", testKeycloakURL)
		os.Setenv("INTERNAL_KEYCLOAK_CLIENT_ID", testKeycloakClientID)
		os.Setenv("INTERNAL_KEYCLOAK_HOST_SECRET", testKeycloakClientSecret)
		os.Setenv("KEYCLOAK_REALM", testKeycloakRealm)

		err := notifications.SendNotification(notifUUID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestGetAllTopics(t *testing.T) {
	if testKeycloakURL == "" {
		t.Skip("Skipping test: testKeycloakURL not set")
	}

	t.Run("should get all topics successfully", func(t *testing.T) {
		expectedTopics := []notifications.Topic{
			{
				UUID: uuid.New(),
				Name: "announcements",
				BaseEntity: notifications.BaseEntity{
					CreatedAt: time.Now(),
					CreatedBy: "system",
				},
			},
			{
				UUID: uuid.New(),
				Name: "updates",
				BaseEntity: notifications.BaseEntity{
					CreatedAt: time.Now(),
					CreatedBy: "system",
				},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/topic" {
				t.Errorf("Expected path '/topic', got %s", r.URL.Path)
			}
			if r.Method != http.MethodGet {
				t.Errorf("Expected GET method, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(expectedTopics)
		}))
		defer server.Close()

		os.Setenv("MS_FIREBASE_GO_URL", server.URL+"/")
		os.Setenv("KEYCLOAK_URL", testKeycloakURL)
		os.Setenv("INTERNAL_KEYCLOAK_CLIENT_ID", testKeycloakClientID)
		os.Setenv("INTERNAL_KEYCLOAK_HOST_SECRET", testKeycloakClientSecret)
		os.Setenv("KEYCLOAK_REALM", testKeycloakRealm)

		result, err := notifications.GetAllTopics()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(result) != len(expectedTopics) {
			t.Errorf("Expected %d topics, got %d", len(expectedTopics), len(result))
		}
		if result[0].Name != expectedTopics[0].Name {
			t.Errorf("Expected first topic name %s, got %s", expectedTopics[0].Name, result[0].Name)
		}
		if result[1].Name != expectedTopics[1].Name {
			t.Errorf("Expected second topic name %s, got %s", expectedTopics[1].Name, result[1].Name)
		}
	})
}

func TestAddTokenToTopic(t *testing.T) {
	if testKeycloakURL == "" {
		t.Skip("Skipping test: testKeycloakURL not set")
	}

	t.Run("should add user to topic successfully", func(t *testing.T) {
		topicUUID := uuid.New().String()
		userUUID := "user-123"

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectedPath := "/topic/" + topicUUID + "/user"
			if r.URL.Path != expectedPath {
				t.Errorf("Expected path '%s', got %s", expectedPath, r.URL.Path)
			}
			if r.Method != http.MethodPost {
				t.Errorf("Expected POST method, got %s", r.Method)
			}

			var body notifications.AddRemoveUserToTopicDto
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("Failed to decode request body: %v", err)
			}

			if body.UserUuid != userUUID {
				t.Errorf("Unexpected body: %+v", body)
			}

			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		os.Setenv("MS_FIREBASE_GO_URL", server.URL+"/")
		os.Setenv("KEYCLOAK_URL", testKeycloakURL)
		os.Setenv("INTERNAL_KEYCLOAK_CLIENT_ID", testKeycloakClientID)
		os.Setenv("INTERNAL_KEYCLOAK_HOST_SECRET", testKeycloakClientSecret)
		os.Setenv("KEYCLOAK_REALM", testKeycloakRealm)

		err := notifications.AddUserToTopic(topicUUID, userUUID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestRemoveTokenFromTopic(t *testing.T) {
	if testKeycloakURL == "" {
		t.Skip("Skipping test: testKeycloakURL not set")
	}

	t.Run("should remove user from topic successfully", func(t *testing.T) {
		topicUUID := uuid.New().String()
		userUUID := "user-123"

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectedPath := "/topic/" + topicUUID + "/user/" + userUUID
			if r.URL.Path != expectedPath {
				t.Errorf("Expected path '%s', got %s", expectedPath, r.URL.Path)
			}
			if r.Method != http.MethodDelete {
				t.Errorf("Expected DELETE method, got %s", r.Method)
			}

			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		os.Setenv("MS_FIREBASE_GO_URL", server.URL+"/")
		os.Setenv("KEYCLOAK_URL", testKeycloakURL)
		os.Setenv("INTERNAL_KEYCLOAK_CLIENT_ID", testKeycloakClientID)
		os.Setenv("INTERNAL_KEYCLOAK_HOST_SECRET", testKeycloakClientSecret)
		os.Setenv("KEYCLOAK_REALM", testKeycloakRealm)

		err := notifications.RemoveUserFromTopic(topicUUID, userUUID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
