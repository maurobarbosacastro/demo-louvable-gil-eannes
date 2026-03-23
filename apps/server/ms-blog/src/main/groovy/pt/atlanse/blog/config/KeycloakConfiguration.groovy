package pt.atlanse.blog.config

import groovy.transform.ToString
import io.micronaut.context.annotation.EachProperty
import io.micronaut.context.annotation.Parameter
import io.micronaut.core.annotation.Nullable

@EachProperty("keycloak-configuration")
@ToString(includePackage = false, includeFields = true, includeNames = true)
class KeycloakConfiguration {
	String name
	String clientId
	String clientSecret
	String realm
	String url
	String tokenUrlIntrospect

	@Nullable
	String adminUsername
	@Nullable
	String adminPassword

	KeycloakConfiguration(@Parameter String name) {
		this.name = name
	}
}
