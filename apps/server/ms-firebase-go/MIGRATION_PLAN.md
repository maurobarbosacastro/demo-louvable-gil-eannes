# Firebase Service Migration Plan - Micronaut to Go

This document outlines the missing features and differences between the original Micronaut Firebase service (`ms-firebase`) and the Go port (`ms-firebase-go`).

## Current Status

The Go port implements basic individual message sending but is missing several critical features from the original Micronaut service.

## Missing Features in Go Service

### 1. **Topic-based Messaging (Group Messages)**
- **Micronaut**: Has `sendGroupMessage()` that sends to Firebase topics using `/topics/{group}` format
- **Go**: Missing this functionality entirely - only supports individual token-based messages
- **Impact**: Cannot send messages to groups/topics of users

### 2. **API Endpoints Structure**
- **Micronaut**: Two distinct endpoints:
  - `POST /api/messages/topic/{group}` - for group/topic messages
  - `POST /api/messages/token/{group}` - for individual messages  
- **Go**: Only has:
  - `POST /api/message/send` - for individual messages
  - `POST /api/message/group` - defined but not implemented (nil handler)
- **Impact**: API structure doesn't match original service

### 3. **Security/Authentication**
- **Micronaut**: Full Keycloak integration with JWT authentication:
  - `@Secured("keycloak-administrator")` annotations on endpoints
  - Complete OAuth2 configuration in `application.yml`
  - JWT signature validation from Keycloak
- **Go**: **No authentication/security implemented at all**
- **Impact**: Service is completely unsecured

### 4. **Configuration Management**
- **Micronaut**: Comprehensive configuration via `application.yml`:
  - Security settings
  - CORS configuration  
  - Keycloak integration
  - Port and environment variables
- **Go**: Basic environment variable handling but no structured configuration file
- **Impact**: Less maintainable configuration

### 5. **DTO Structure Differences**
- **Micronaut**: `MessageDTO` has `title` and `body` fields only
- **Go**: `MessageDTO` has `token`, `title`, and `body` fields (token should be a path parameter, not in DTO)
- **Impact**: API contract doesn't match original

### 6. **Error Handling & Response Structure**
- **Micronaut**: Returns simple responses, relies on framework error handling
- **Go**: Has structured `NotificationResponse` but inconsistent error handling
- **Impact**: Different error response formats

### 7. **Data in Messages**
- **Micronaut**: Adds `.putData("message", message.body)` to Firebase messages
- **Go**: Missing this additional data field
- **Impact**: Mobile apps may expect this data field

### 8. **Async Processing**
- **Micronaut**: Uses `CompletableFuture` for async message sending
- **Go**: Synchronous processing only
- **Impact**: Potential performance issues under load

### 9. **Infrastructure Files**
- **Micronaut**: Has `Dockerfile`, `Dockerfile-Conf`, and `README.md` with documentation
- **Go**: Missing these infrastructure files
- **Impact**: Cannot be deployed consistently

### 10. **Import Path Issues**  
The Go service has an incorrect import in `firebase.service.go`:
```go
"ms-firebase-go/internal/firebase" // This path doesn't exist
```
Should be:
```go  
"ms-firebase-go/external/firebase"
```

## Implementation Plan

### Phase 1: Critical Fixes
1. **Fix Import Path**
   - Correct the import in `internal/services/firebase.service.go`
   - Update from `internal/firebase` to `external/firebase`

2. **Implement Security & Authentication**
   - Add Keycloak JWT authentication middleware
   - Implement `keycloak-administrator` role checking  
   - Add CORS configuration
   - Create middleware for JWT token validation

### Phase 2: Core Functionality
3. **Topic-Based Messaging**
   - Implement `SendGroupMessage()` function in service
   - Add proper endpoint handler for `/api/message/group`
   - Support Firebase topics format `/topics/{group}`

4. **API Structure Fixes**
   - Update endpoint paths to match Micronaut version:
     - `/api/message/send` → `/api/messages/token/{group}`
     - `/api/message/group` → `/api/messages/topic/{group}`
   - Fix DTO structure (remove token field, use path parameter)
   - Add missing data field to Firebase messages

### Phase 3: Enhancement & Infrastructure
5. **Async Processing**
   - Implement goroutines for async message processing
   - Add proper context handling

6. **Configuration Management**
   - Create structured configuration file support
   - Add validation for required environment variables

7. **Infrastructure Files**
   - Create Dockerfile
   - Add README.md with deployment instructions
   - Add any necessary configuration files

### Phase 4: Testing & Documentation
8. **Testing**
   - Add unit tests for service functions
   - Add integration tests for endpoints
   - Test security middleware

9. **Documentation**
   - Update API documentation
   - Add deployment guide
   - Document configuration options

## Files to Create/Modify

### New Files
- `internal/middleware/auth.go` - JWT authentication middleware
- `internal/config/config.go` - Configuration management
- `Dockerfile` - Container configuration
- `README.md` - Service documentation

### Files to Modify
- `internal/services/firebase.service.go` - Fix imports, add group messaging
- `internal/controllers/firebase.controller.go` - Update endpoints, add security
- `internal/dto/message.dto.go` - Fix DTO structure
- `internal/config/server.go` - Add middleware, update routes
- `go.mod` - Add security dependencies

## Dependencies to Add
- JWT handling library (e.g., `github.com/golang-jwt/jwt/v5`)
- Keycloak integration library
- Configuration management library (e.g., `github.com/spf13/viper`)

## Success Criteria
- [ ] All endpoints match original Micronaut service
- [ ] Security/authentication works with Keycloak
- [ ] Both individual and group messaging work
- [ ] Service can be deployed via Docker
- [ ] API responses match original format
- [ ] Performance is comparable or better than Micronaut version