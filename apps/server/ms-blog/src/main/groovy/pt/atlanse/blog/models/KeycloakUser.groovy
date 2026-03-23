package pt.atlanse.blog.models

import groovy.transform.ToString
import io.swagger.v3.oas.annotations.media.Schema

@Schema(name = "KeycloakUser", description = "Data obtained from the authentication/authorization tasks")
@ToString
class KeycloakUser {
    private String email
    private String username
    private List<String> roles

    KeycloakUser() {}

    KeycloakUser(String email, String username, List<String> roles) {
        this.email = email
        this.username = username
        this.roles = roles
    }

    void setEmail(String email) {
        this.email = email
    }

    void setUsername(String username) {
        this.username = username
    }

    void setRoles(List<String> roles) {
        this.roles = roles
    }

    String getEmail() {
        return email
    }

    String getUsername() {
        return username
    }

    List<String> getRoles() {
        return roles
    }
}
