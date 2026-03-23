package pt.atlanse.blog.services

import groovy.util.logging.Slf4j
import jakarta.inject.Named
import jakarta.inject.Singleton
import org.keycloak.admin.client.Keycloak
import org.keycloak.admin.client.KeycloakBuilder
import org.keycloak.admin.client.resource.UsersResource
import org.keycloak.representations.idm.UserRepresentation
import pt.atlanse.blog.config.KeycloakConfiguration

/**
 * @deprecated
 * */
@Slf4j
@Singleton
class KeycloakService {

    public Keycloak keycloak

    KeycloakConfiguration keycloakConfigAdmin
    private KeycloakConfiguration keycloakConfigClient

    KeycloakService(
        @Named("master") KeycloakConfiguration keycloakConfigAdmin,
        @Named("client") KeycloakConfiguration keycloakConfigClient) {
        this.keycloakConfigAdmin = keycloakConfigAdmin
        this.keycloakConfigClient = keycloakConfigClient
    }

    Keycloak getKeycloak() {
        return this.keycloak ?: setKeycloak()
    }

    /**
     * Start keycloak instance using the config file
     * @return Object of class {@link Keycloak}
     * */
    Keycloak setKeycloak() {
        try {
            log.info "Creating new bean for Keycloak"

            this.keycloak = KeycloakBuilder.builder()
                .serverUrl("${ this.keycloakConfigAdmin.url }/auth")
                .realm(this.keycloakConfigAdmin.realm)
                .clientId(this.keycloakConfigAdmin.clientId)
                .clientSecret(this.keycloakConfigAdmin.clientSecret)
                .username(this.keycloakConfigAdmin.adminUsername)
                .password(this.keycloakConfigAdmin.adminPassword)
                .build()

            return this.keycloak

        } catch (Exception e) {
            log.error "Error attempting to config Keycloak object; Reason: $e.message"
        }
    }

    /**
     * Get Keycloaks configuration object
     * @return Object of class {@link KeycloakConfiguration}
     * */
    KeycloakConfiguration getClientConfiguration() {
        return this.keycloakConfigClient
    }

    /**
     * Find users inside client realm
     * @return The users object of class {@link UsersResource}
     * */
    UsersResource getUsers() {
        return getKeycloak().realm(this.keycloakConfigClient.realm).users()
    }

    /**
     * Find user representation object and return
     * @param id The user id that matches a keycloak's user UUID
     * @return The user object of class {@link UserRepresentation}
     * */
    UserRepresentation findUser(String id) {
        try {
            return getUsers().get(id).toRepresentation()
        } catch (Exception e) {
            log.error "Failed to retrieve user for id $id; Reason: ${ e.message }"
        }
    }

    /**
     * Find user representation object and return the email address
     * @param id The user id that matches a keycloak's user UUID
     * @return User email or 'User was not found' as String
     * */
    String getUserEmail(String id) {
        return findUser(id)?.email ?: "User was not found"
    }

}
