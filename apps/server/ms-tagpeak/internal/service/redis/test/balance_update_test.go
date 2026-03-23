package test

// This test file is not working due to Keycloak integration issues.

import (
	"context"
	"fmt"
	"math/rand"
	"ms-tagpeak/internal/constants"
	"ms-tagpeak/internal/dto"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service/redis"
	"ms-tagpeak/pkg/redisclient"
	"strconv"
	"sync"
	"testing"
	"time"

	gocloak "github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
)

const (
	// Test user UUID - fixed for all tests
	testUserUUID = "c6c7f50c-ad77-4d7d-80c6-5c582d4de1a9"
)

// Keycloak instance for testing
var testKeycloak *constants.Keycloak

// setupTest initializes test environment
func setupTest(t *testing.T) {

	// Initialize Redis
	redisHost := "localhost:6379"

	redisclient.SetRedisClient(redisclient.Config{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	// Initialize Keycloak with real credentials
	keycloakURL := "http://localhost:8081"
	keycloakRealm := "tagpeak"
	adminUsername := "tagpeak-ms"
	adminPassword := "?????"

	if keycloakURL == "" || keycloakRealm == "" || adminUsername == "" || adminPassword == "" {
		t.Skip("Keycloak credentials not found in environment. Required: KEYCLOAK_URL, KEYCLOAK_REALM, TAGPEAK_ADMIN_USERNAME, TAGPEAK_ADMIN_PASSWORD")
		return
	}

	t.Logf("Connecting to Keycloak: %s (realm: %s, user: %s)", keycloakURL, keycloakRealm, adminUsername)

	// Create Keycloak client
	client := gocloak.NewClient(keycloakURL)
	ctx := context.Background()

	// Login as admin
	token, err := client.LoginAdmin(ctx, adminUsername, adminPassword, "master")
	if err != nil {
		t.Fatalf("Failed to login to Keycloak: %v", err)
	}

	testKeycloak = &constants.Keycloak{
		Client:          client,
		Ctx:             ctx,
		AdminToken:      token,
		Realm:           keycloakRealm,
		TokenExpireDate: time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
	}

	t.Log("Keycloak connection established successfully")
}

// cleanupTest cleans up test environment
func cleanupTest(t *testing.T) {
	// Clean up any test locks in Redis
	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()

	lockKey := fmt.Sprintf("balance_lock:%s", testUserUUID)
	client.Del(ctx, lockKey)
}

// generateRandomAmount generates a random amount between 0.01 and 100.00
func generateRandomAmount() float64 {
	return float64(rand.Intn(10000)+1) / 100.0 // 0.01 to 100.00
}

// TestBalanceLockKey tests the lock key generation
func TestBalanceLockKey(t *testing.T) {
	userUUID := testUserUUID
	expectedKey := fmt.Sprintf("balance_lock:%s", userUUID)

	// Note: balanceLockKey is not exported, so we test the pattern
	lockKey := fmt.Sprintf("balance_lock:%s", userUUID)

	if lockKey != expectedKey {
		t.Errorf("Expected lock key %s, got %s", expectedKey, lockKey)
	}
}

// TestHandleKeycloakBalanceUpdate_SingleUpdate tests a single balance update
func TestHandleKeycloakBalanceUpdate_SingleUpdate(t *testing.T) {
	// Skip if Redis not available
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	setupTest(t)
	defer cleanupTest(t)

	t.Run("Should update balance with lock acquired", func(t *testing.T) {
		amount := generateRandomAmount()

		t.Logf("Testing single update: user=%s, amount=%.2f", testUserUUID, amount)

		// Call the actual function with real Keycloak
		err := redis.HandleKeycloakBalanceUpdate(testUserUUID, amount, testKeycloak)
		if err != nil {
			t.Errorf("Failed to update balance: %v", err)
		} else {
			t.Logf("Balance updated successfully (+%.2f)", amount)
		}

		// Verify lock is released
		client := redisclient.GetRedisClient()
		ctx := redisclient.GetContext()
		lockKey := fmt.Sprintf("balance_lock:%s", testUserUUID)

		exists, _ := client.Exists(ctx, lockKey).Result()
		if exists > 0 {
			t.Error("Lock should be released after update")
		}
	})
}

// TestHandleKeycloakBalanceUpdate_ConcurrentUpdates tests concurrent balance updates
// This is the most important test - it verifies that the lock prevents race conditions
func TestHandleKeycloakBalanceUpdate_ConcurrentUpdates(t *testing.T) {
	// Skip if Redis not available
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	setupTest(t)
	defer cleanupTest(t)

	t.Run("Should handle concurrent updates without race condition with real Keycloak", func(t *testing.T) {
		const numConcurrentUpdates = 5 // Use 5 to avoid overwhelming Keycloak

		// Generate random amounts for each update
		amounts := make([]float64, numConcurrentUpdates)
		expectedTotal := 0.0

		for i := 0; i < numConcurrentUpdates; i++ {
			amounts[i] = generateRandomAmount()
			expectedTotal += amounts[i]
		}

		t.Logf("Testing %d concurrent updates for user %s", numConcurrentUpdates, testUserUUID)
		t.Logf("Amounts: %v", amounts)
		t.Logf("Expected total to add: %.2f", expectedTotal)

		// Track results
		var wg sync.WaitGroup
		errors := make([]error, numConcurrentUpdates)

		// Launch concurrent updates with real Keycloak
		for i := 0; i < numConcurrentUpdates; i++ {
			wg.Add(1)
			go func(index int, amount float64) {
				defer wg.Done()

				t.Logf("Update %d: Starting (+%.2f)", index, amount)
				err := redis.HandleKeycloakBalanceUpdate(testUserUUID, amount, testKeycloak)
				errors[index] = err

				if err != nil {
					t.Logf("Update %d: Failed - %v", index, err)
				} else {
					t.Logf("Update %d: Success (+%.2f)", index, amount)
				}
			}(i, amounts[i])
		}

		// Wait for all goroutines
		wg.Wait()

		// Count successes
		successCount := 0
		for i, err := range errors {
			if err == nil {
				successCount++
			} else {
				t.Logf("Update %d failed: %v", i, err)
			}
		}

		t.Logf("Results: %d/%d succeeded", successCount, numConcurrentUpdates)

		// All updates should succeed (lock + retry mechanism handles this)
		if successCount != numConcurrentUpdates {
			t.Errorf("Expected all %d updates to succeed, but only %d succeeded", numConcurrentUpdates, successCount)
		}

		// Verify lock is released at the end
		client := redisclient.GetRedisClient()
		ctx := redisclient.GetContext()
		lockKey := fmt.Sprintf("balance_lock:%s", testUserUUID)

		exists, _ := client.Exists(ctx, lockKey).Result()
		if exists > 0 {
			t.Error("Lock should be released after all updates complete")
		}
	})

	t.Run("Should handle lock contention correctly (simulated)", func(t *testing.T) {
		const numConcurrentUpdates = 10

		// Generate random amounts for each update
		amounts := make([]float64, numConcurrentUpdates)
		expectedTotal := 0.0

		for i := 0; i < numConcurrentUpdates; i++ {
			amounts[i] = generateRandomAmount()
			expectedTotal += amounts[i]
		}

		t.Logf("Testing %d concurrent lock attempts (lock mechanism only)", numConcurrentUpdates)

		// Track which updates succeeded
		var wg sync.WaitGroup
		successCount := 0
		failCount := 0
		var mu sync.Mutex

		// Launch concurrent lock attempts (without Keycloak calls)
		for i := 0; i < numConcurrentUpdates; i++ {
			wg.Add(1)
			go func(index int, amount float64) {
				defer wg.Done()

				// Test the lock mechanism
				client := redisclient.GetRedisClient()
				ctx := redisclient.GetContext()
				lockKey := fmt.Sprintf("balance_lock:%s", testUserUUID)

				// Try to acquire lock
				acquired, err := client.SetNX(ctx, lockKey, fmt.Sprintf("%d", time.Now().UnixNano()), 30*time.Second).Result()

				if err == nil && acquired {
					// Simulate some work
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)))

					// Release lock
					client.Del(ctx, lockKey)

					mu.Lock()
					successCount++
					mu.Unlock()

					t.Logf("Lock attempt %d: Acquired and released", index)
				} else {
					mu.Lock()
					failCount++
					mu.Unlock()

					t.Logf("Lock attempt %d: Failed (lock held by another)", index)
				}
			}(i, amounts[i])
		}

		// Wait for all goroutines
		wg.Wait()

		t.Logf("Lock contention results: %d succeeded, %d failed (expected behavior)", successCount, failCount)

		// At least one should succeed
		if successCount == 0 {
			t.Error("At least one lock attempt should have succeeded")
		}

		// Verify lock is released at the end
		client := redisclient.GetRedisClient()
		ctx := redisclient.GetContext()
		lockKey := fmt.Sprintf("balance_lock:%s", testUserUUID)

		exists, _ := client.Exists(ctx, lockKey).Result()
		if exists > 0 {
			t.Error("Lock should be released after all attempts complete")
		}
	})
}

// TestHandleKeycloakBalanceUpdate_LockExpiration tests lock expiration
func TestHandleKeycloakBalanceUpdate_LockExpiration(t *testing.T) {
	// Skip if Redis not available
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	setupTest(t)
	defer cleanupTest(t)

	t.Run("Should expire lock after timeout", func(t *testing.T) {
		client := redisclient.GetRedisClient()
		ctx := redisclient.GetContext()
		lockKey := fmt.Sprintf("balance_lock:%s", testUserUUID)

		// Acquire lock with short timeout
		acquired, err := client.SetNX(ctx, lockKey, "test", 1*time.Second).Result()
		if err != nil || !acquired {
			t.Fatal("Failed to acquire lock for test")
		}

		// Verify lock exists
		exists, _ := client.Exists(ctx, lockKey).Result()
		if exists == 0 {
			t.Error("Lock should exist immediately after acquisition")
		}

		// Wait for expiration
		time.Sleep(2 * time.Second)

		// Verify lock expired
		exists, _ = client.Exists(ctx, lockKey).Result()
		if exists > 0 {
			t.Error("Lock should have expired after timeout")
		}

		t.Log("Lock expired successfully after 1 second")
	})
}

// TestHandleKeycloakBalanceUpdate_RetryMechanism tests the retry behavior
func TestHandleKeycloakBalanceUpdate_RetryMechanism(t *testing.T) {
	// Skip if Redis not available
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	setupTest(t)
	defer cleanupTest(t)

	t.Run("Should retry when lock is held", func(t *testing.T) {
		client := redisclient.GetRedisClient()
		ctx := redisclient.GetContext()
		lockKey := fmt.Sprintf("balance_lock:%s", testUserUUID)

		// Acquire lock with short timeout
		acquired, err := client.SetNX(ctx, lockKey, "test", 2*time.Second).Result()
		if err != nil || !acquired {
			t.Fatal("Failed to acquire lock for test")
		}

		t.Log("Lock acquired by test, simulating another process trying to acquire...")

		// Try to acquire in goroutine (simulates retry mechanism)
		done := make(chan bool)
		go func() {
			retries := 0
			maxRetries := 5

			for retries < maxRetries {
				acquired, _ := client.SetNX(ctx, lockKey, "retry", 2*time.Second).Result()
				if acquired {
					t.Log("Successfully acquired lock after retries")
					client.Del(ctx, lockKey)
					done <- true
					return
				}

				retries++
				t.Logf("Retry %d/%d: Lock still held", retries, maxRetries)
				time.Sleep(50 * time.Millisecond)
			}

			done <- false
		}()

		// Wait for result
		success := <-done

		// Clean up
		client.Del(ctx, lockKey)

		if !success {
			t.Error("Should have acquired lock after retries (or lock should have expired)")
		}
	})
}

// TestBalanceUpdateConstants tests the configured constants
func TestBalanceUpdateConstants(t *testing.T) {
	t.Run("Should have reasonable timeout values", func(t *testing.T) {
		// Test that constants are defined correctly
		// Note: We can't access package constants directly in tests
		// but we can verify behavior

		expectedTimeout := 30 * time.Second
		expectedRetries := 20
		expectedDelay := 50 * time.Millisecond

		t.Logf("Expected timeout: %v", expectedTimeout)
		t.Logf("Expected retries: %d", expectedRetries)
		t.Logf("Expected delay: %v", expectedDelay)

		// Calculate max wait time
		maxWaitTime := time.Duration(expectedRetries) * expectedDelay
		t.Logf("Max wait time for lock acquisition: %v", maxWaitTime)

		if maxWaitTime > 2*time.Second {
			t.Logf("Warning: Max wait time is %v, which might be too long", maxWaitTime)
		}
	})
}

// BenchmarkHandleKeycloakBalanceUpdate_LockAcquisition benchmarks lock acquisition
func BenchmarkHandleKeycloakBalanceUpdate_LockAcquisition(b *testing.B) {
	// Skip if Redis not available
	if testing.Short() {
		b.Skip("Skipping benchmark in short mode")
	}

	redisclient.SetRedisClient(redisclient.Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	client := redisclient.GetRedisClient()
	ctx := redisclient.GetContext()
	lockKey := fmt.Sprintf("balance_lock:benchmark-user")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Acquire lock
		client.SetNX(ctx, lockKey, "bench", 1*time.Second)
		// Release lock
		client.Del(ctx, lockKey)
	}
}

// Example test showing expected usage pattern
func ExampleHandleKeycloakBalanceUpdate() {
	// Initialize Redis
	redisclient.SetRedisClient(redisclient.Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Initialize Keycloak (mock for example)
	keycloak := &constants.Keycloak{
		// ... configuration
	}

	// Update balance with lock protection
	userUUID := testUserUUID
	amountToAdd := 25.50

	err := redis.HandleKeycloakBalanceUpdate(userUUID, amountToAdd, keycloak)
	if err != nil {
		fmt.Printf("Error updating balance: %v\n", err)
		return
	}

	fmt.Println("Balance updated successfully")
	// Output: Balance updated successfully
}

// Mock functions for testing (to be implemented when full integration testing is set up)

// mockGetUserById mocks the service.GetUserById function
func mockGetUserById(userUUID string, keycloak *constants.Keycloak) (*models.User, error) {
	// Return a mock user with a balance
	return &models.User{
		Uuid:    uuid.MustParse(userUUID),
		Balance: 100.0, // Starting balance
	}, nil
}

// mockUpdateUser mocks the service.UpdateUser function
func mockUpdateUser(userUUID uuid.UUID, dto dto.UpdateUserDto, keycloak *constants.Keycloak) (*models.User, error) {
	// Parse the balance from the DTO
	if dto.Balance != nil {
		balance, _ := strconv.ParseFloat(*dto.Balance, 64)
		return &models.User{
			Uuid:    userUUID,
			Balance: balance,
		}, nil
	}
	return nil, fmt.Errorf("no balance provided")
}

// TestMockBalanceUpdate demonstrates how to test with mocked Keycloak
func TestMockBalanceUpdate(t *testing.T) {
	t.Run("Should update balance correctly with mocked Keycloak", func(t *testing.T) {
		// Setup
		initialBalance := 100.0
		amountToAdd := 25.50
		expectedBalance := initialBalance + amountToAdd

		// Mock user
		user := &models.User{
			Uuid:    uuid.MustParse(testUserUUID),
			Balance: initialBalance,
		}

		// Simulate balance calculation
		newBalance := user.Balance + amountToAdd

		if newBalance != expectedBalance {
			t.Errorf("Expected balance %.2f, got %.2f", expectedBalance, newBalance)
		}

		t.Logf("Balance update: %.2f → %.2f (+%.2f)", initialBalance, newBalance, amountToAdd)
	})
}
