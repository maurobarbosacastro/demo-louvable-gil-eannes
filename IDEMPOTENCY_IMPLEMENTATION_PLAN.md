# API Idempotency & De-duplication Implementation Plan

## Overview

This document provides a comprehensive plan to implement an idempotency mechanism that prevents duplicate requests (double-clicks/race conditions) where the source and data are identical.

## System Architecture

### Technology Stack
- **Backend**: Go (Golang) with Echo framework
- **Cache Layer**: Redis for distributed locking and response caching
- **Frontend**: Angular with functional interceptors

### Strategy Summary
1. **Frontend**: Generate SHA-256 hash of request (method + URL + body) as `Idempotency-Key` header
2. **Backend**: Use Redis SETNX for atomic locking and response caching
3. **Behavior**: Block duplicate requests or return cached responses

---

## Part 1: Backend Implementation (Go + Echo)

### File Structure
```
apps/server/ms-tagpeak/
├── internal/
│   └── middleware/
│       └── idempotency.go          # New middleware file
├── cmd/
│   └── server/
│       └── main.go                  # Register middleware here
└── internal/
    └── config/
        └── redis.go                 # Ensure Redis client is configured
```

### Step 1.1: Create Idempotency Middleware

**File**: `apps/server/ms-tagpeak/internal/middleware/idempotency.go`

**Key Components**:
1. **Custom ResponseWriter**: Hijacks and captures response body for caching
2. **CachedResponse**: Struct to store HTTP status and response body
3. **IdempotencyMiddleware**: Main middleware function accepting Redis client

**Implementation Details**:

```go
package middleware

import (
    "bytes"
    "context"
    "encoding/json"
    "io"
    "net/http"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/redis/go-redis/v9"
)

// Configuration
const LockTTL = 30 * time.Second // Adjustment for "double-click" prevention

type CachedResponse struct {
    Status int    `json:"status"`
    Body   string `json:"body"`
}

// Custom writer to capture response body
type bodyDumpResponseWriter struct {
    io.Writer
    http.ResponseWriter
    statusCode int
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
    w.statusCode = code
    w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
    return w.Writer.Write(b)
}

func IdempotencyMiddleware(rdb *redis.Client) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            key := c.Request().Header.Get("Idempotency-Key")
            if key == "" {
                return next(c)
            }

            ctx := context.Background()
            redisKey := "idemp:" + key

            // 1. Try to acquire lock
            isNew, err := rdb.SetNX(ctx, redisKey, "PROCESSING", LockTTL).Result()
            if err != nil {
                return next(c) // Fail open or handle error
            }

            if !isNew {
                val, _ := rdb.Get(ctx, redisKey).Result()
                if val == "PROCESSING" {
                    return echo.NewHTTPError(http.StatusConflict, "Duplicate request processing")
                }

                var cached CachedResponse
                if json.Unmarshal([]byte(val), &cached) == nil {
                    return c.JSONBlob(cached.Status, []byte(cached.Body))
                }
                return echo.NewHTTPError(http.StatusConflict, "Duplicate request detected")
            }

            // 2. Capture Response
            resBody := new(bytes.Buffer)
            mw := io.MultiWriter(c.Response().Writer, resBody)
            writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer, statusCode: http.StatusOK}
            c.Response().Writer = writer

            // 3. Execute Handler
            if err := next(c); err != nil {
                rdb.Del(ctx, redisKey) // On error, release lock immediately
                return err
            }

            // 4. Cache Success (2xx only)
            if writer.statusCode >= 200 && writer.statusCode < 300 {
                data, _ := json.Marshal(CachedResponse{Status: writer.statusCode, Body: resBody.String()})
                rdb.Set(ctx, redisKey, data, LockTTL)
            } else {
                rdb.Del(ctx, redisKey) // Release lock on non-success logic
            }

            return nil
        }
    }
}
```

**Logic Flow**:
1. Extract `Idempotency-Key` header. If missing, skip idempotency check
2. **Atomic Lock**: Use Redis SETNX with 30-second TTL
3. **Cache Hit**:
   - If key exists with "PROCESSING" → return 409 Conflict
   - If key has cached response → return cached JSON
4. **Cache Miss**: Execute handler → capture response → cache successful responses (2xx only)
5. **Error Handling**: Release lock immediately on handler errors

### Step 1.2: Ensure Redis Client Configuration

**File**: `apps/server/ms-tagpeak/internal/config/redis.go` or similar

**Requirements**:
- Redis client must be initialized and available
- Should use go-redis/v9 package
- Client should be passed to middleware during registration

**Example Redis Client Setup**:
```go
import "github.com/redis/go-redis/v9"

func InitRedis() *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // or from environment variable
        Password: "",               // or from environment variable
        DB:       0,
    })
}
```

### Step 1.3: Register Middleware

**File**: `apps/server/ms-tagpeak/cmd/server/main.go` or router initialization file

**Registration**:
```go
import (
    "your-project/internal/middleware"
    // ... other imports
)

func main() {
    // Initialize Echo
    e := echo.New()

    // Initialize Redis client
    rdb := InitRedis() // or get from dependency injection

    // Register idempotency middleware
    e.Use(middleware.IdempotencyMiddleware(rdb))

    // ... register routes and start server
}
```

**Important Notes**:
- Register middleware **before** route registration
- Ensure Redis connection is established before server starts
- Consider using environment variables for Redis configuration
- Add health checks for Redis connectivity

---

## Part 2: Frontend Implementation (Angular)

### File Structure
```
apps/client/backoffice/  # or mobile/shopify as needed
├── src/
│   ├── app/
│   │   ├── interceptors/
│   │   │   └── idempotency.interceptor.ts    # New interceptor
│   │   └── app.config.ts                     # Register interceptor here
```

### Step 2.1: Create Idempotency Interceptor

**File**: `apps/client/backoffice/src/app/interceptors/idempotency.interceptor.ts`

**Key Features**:
- Functional interceptor (Angular 15+ style)
- SHA-256 hash generation using Web Crypto API
- Only applies to mutation methods: POST, PUT, PATCH, DELETE
- Hash payload: `${method}:${url}:${body}`

**Implementation**:

```typescript
import { HttpInterceptorFn, HttpRequest, HttpHandlerFn, HttpEvent } from '@angular/common/http';
import { from, Observable, switchMap } from 'rxjs';

/**
 * Generate SHA-256 hash for idempotency key
 * @param req HTTP request to hash
 * @returns Promise resolving to hex-encoded hash string
 */
async function generateHash(req: HttpRequest<unknown>): Promise<string> {
  const body = req.body ? JSON.stringify(req.body) : '';
  const payload = `${req.method}:${req.url}:${body}`;
  const encoder = new TextEncoder();
  const data = encoder.encode(payload);
  const hashBuffer = await crypto.subtle.digest('SHA-256', data);
  return Array.from(new Uint8Array(hashBuffer))
    .map(b => b.toString(16).padStart(2, '0'))
    .join('');
}

/**
 * Idempotency interceptor for mutation requests
 * Adds Idempotency-Key header with SHA-256 hash of request
 */
export const idempotencyInterceptor: HttpInterceptorFn = (req, next) => {
  const mutationMethods = ['POST', 'PUT', 'PATCH', 'DELETE'];

  // Skip non-mutation requests
  if (!mutationMethods.includes(req.method)) {
    return next(req);
  }

  // Generate hash and add as header
  return from(generateHash(req)).pipe(
    switchMap(key => {
      const cloned = req.clone({
        setHeaders: { 'Idempotency-Key': key }
      });
      return next(cloned);
    })
  );
};
```

**How It Works**:
1. Check if request method is a mutation (POST/PUT/PATCH/DELETE)
2. Generate SHA-256 hash of `method:url:body`
3. Clone request with `Idempotency-Key` header
4. Forward modified request to next handler

### Step 2.2: Register Interceptor

**File**: `apps/client/backoffice/src/app/app.config.ts`

**Registration**:
```typescript
import { ApplicationConfig } from '@angular/core';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { idempotencyInterceptor } from './interceptors/idempotency.interceptor';

export const appConfig: ApplicationConfig = {
  providers: [
    provideHttpClient(
      withInterceptors([
        idempotencyInterceptor,
        // ... other interceptors
      ])
    ),
    // ... other providers
  ]
};
```

**Alternative Registration** (if using module-based app):
```typescript
import { provideHttpClient, withInterceptors } from '@angular/common/http';

@NgModule({
  providers: [
    provideHttpClient(
      withInterceptors([idempotencyInterceptor])
    )
  ]
})
export class AppModule { }
```

---

## Implementation Checklist

### Backend Tasks
- [ ] Create `internal/middleware/idempotency.go` with middleware implementation
- [ ] Verify Redis client initialization and configuration
- [ ] Add Redis dependency to `go.mod` (`github.com/redis/go-redis/v9`)
- [ ] Register middleware in main server file
- [ ] Configure Redis connection from environment variables
- [ ] Add Redis health check endpoint
- [ ] Test Redis SETNX behavior and TTL expiration
- [ ] Add logging for idempotency hits/misses
- [ ] Handle Redis connection failures gracefully (fail open vs fail closed)

### Frontend Tasks (Backoffice)
- [ ] Create `src/app/interceptors/idempotency.interceptor.ts`
- [ ] Register interceptor in `app.config.ts` or app module
- [ ] Test hash generation with sample requests
- [ ] Verify header is added only to mutation methods
- [ ] Test browser compatibility for crypto.subtle API

### Frontend Tasks (Mobile - if needed)
- [ ] Create equivalent interceptor in mobile app
- [ ] Register in mobile app configuration
- [ ] Test on both iOS and Android devices

### Testing Tasks
- [ ] Unit test: Backend middleware with mock Redis
- [ ] Unit test: Frontend interceptor hash generation
- [ ] Integration test: Double-click scenario returns 409 or cached response
- [ ] Integration test: Different requests get different keys
- [ ] Integration test: Identical requests get same key
- [ ] Load test: Concurrent identical requests handled correctly
- [ ] E2E test: User double-clicks submit button
- [ ] Test TTL expiration (wait 30+ seconds and retry)
- [ ] Test error scenarios (Redis down, handler errors)

---

## Configuration Parameters

### Backend Configuration
```go
// Adjustable parameters in idempotency.go
const (
    LockTTL = 30 * time.Second  // Lock duration (adjust based on longest expected request)
)

// Redis key prefix
const redisKeyPrefix = "idemp:"  // Namespace for idempotency keys
```

**Recommended TTL Values**:
- Short requests (< 1s): 10-15 seconds
- Medium requests (1-5s): 30 seconds
- Long requests (5-30s): 60 seconds
- Background jobs: Consider longer TTL or different strategy

### Frontend Configuration
No configuration needed - behavior is deterministic based on request content.

---

## Error Handling Strategy

### Backend Error Scenarios

1. **Redis Unavailable**:
   - Current: Fail open (skip idempotency check)
   - Alternative: Fail closed (return 503 Service Unavailable)
   - **Decision needed**: Choose based on system requirements

2. **Handler Error**:
   - Release Redis lock immediately
   - Allow client to retry

3. **Non-2xx Response**:
   - Release Redis lock
   - Don't cache error responses
   - **Rationale**: Allow retry on transient errors

### Frontend Error Scenarios

1. **Crypto API Unavailable**:
   - Fallback: Skip idempotency header (backend handles gracefully)
   - Log warning to console

2. **Hash Generation Error**:
   - Skip idempotency header
   - Allow request to proceed

---

## Security Considerations

1. **Hash Collision**:
   - SHA-256 provides sufficient collision resistance
   - Probability of collision is negligible

2. **Key Exposure**:
   - Idempotency key is deterministic but not reversible
   - No sensitive data leaked through key

3. **Replay Attacks**:
   - TTL limits replay window to 30 seconds
   - Consider adding timestamp to hash if longer TTL needed

4. **Redis Security**:
   - Use password-protected Redis in production
   - Isolate Redis instance or use ACLs
   - Enable TLS for Redis connections

---

## Monitoring & Observability

### Metrics to Track
- Idempotency hit rate (cache hits vs misses)
- 409 Conflict response rate
- Redis operation latency
- Cache TTL effectiveness

### Logging Points
```go
// Add structured logging in middleware
log.Info("Idempotency cache hit",
    "key", key,
    "cached_status", cached.Status)

log.Info("Idempotency cache miss",
    "key", key,
    "method", c.Request().Method,
    "path", c.Request().URL.Path)
```

---

## Rollout Strategy

### Phase 1: Backend Implementation
1. Implement middleware in separate branch
2. Deploy with feature flag (disabled)
3. Monitor Redis performance
4. Enable for test endpoints only

### Phase 2: Frontend Integration
1. Deploy interceptor to staging environment
2. Test with real user scenarios
3. Monitor 409 responses and cache hits

### Phase 3: Production Rollout
1. Enable for non-critical endpoints
2. Monitor error rates and performance
3. Gradually enable for all mutation endpoints
4. Document any endpoint-specific exclusions

---

## Troubleshooting Guide

### Problem: Too many 409 responses
**Cause**: TTL too long or requests are actually legitimate
**Solution**:
- Reduce TTL
- Add timestamp component to hash
- Review request patterns

### Problem: Redis memory growing
**Cause**: High request volume with 30s TTL
**Solution**:
- Reduce TTL
- Implement Redis eviction policy
- Monitor key count

### Problem: Interceptor not adding header
**Cause**: Interceptor not registered or wrong HTTP client
**Solution**:
- Verify interceptor registration
- Check HttpClient is injected correctly
- Test with browser dev tools (Network tab)

---

## Future Enhancements

1. **Configurable TTL**: Per-endpoint TTL configuration
2. **Response Compression**: Compress cached responses in Redis
3. **Selective Caching**: Configure which endpoints use idempotency
4. **Dashboard**: Real-time monitoring of idempotency metrics
5. **Client-side Retry Logic**: Automatic retry on 409 with backoff

---

## Dependencies

### Backend
```bash
go get github.com/redis/go-redis/v9
```

### Frontend
No additional dependencies - uses native Web Crypto API

---

## Testing Examples

### Backend Unit Test Example
```go
func TestIdempotencyMiddleware_CacheHit(t *testing.T) {
    // Setup mock Redis with existing key
    // Send request with Idempotency-Key
    // Assert 409 or cached response returned
}
```

### Frontend Unit Test Example
```typescript
describe('idempotencyInterceptor', () => {
  it('should add Idempotency-Key header to POST requests', async () => {
    // Create mock POST request
    // Apply interceptor
    // Assert header exists and is valid SHA-256
  });
});
```

---

## Questions for Implementation

1. **Redis Configuration**:
   - Where is Redis currently configured in the project?
   - Should we use existing Redis instance or separate one?

2. **Error Strategy**:
   - Fail open (skip check) or fail closed (return error) when Redis unavailable?

3. **Scope**:
   - Apply to all endpoints or specific ones only?
   - Should GET requests with bodies be included?

4. **TTL Duration**:
   - Is 30 seconds appropriate for all endpoints?
   - Any long-running operations needing longer TTL?

5. **Multiple Clients**:
   - Should this be implemented in backoffice only, or also mobile app?

---

## References

- [Echo Framework Middleware](https://echo.labstack.guide/middleware/)
- [Redis SETNX Documentation](https://redis.io/commands/setnx/)
- [Angular HTTP Interceptors](https://angular.io/guide/http-intercept-requests-and-responses)
- [Web Crypto API - SubtleCrypto](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto)
- [Idempotency Best Practices](https://stripe.com/docs/api/idempotent_requests)

---

## Implementation Contact

For questions or clarifications during implementation, refer to:
- This specification document
- Architecture team for Redis configuration decisions
- Security team for production rollout approval
