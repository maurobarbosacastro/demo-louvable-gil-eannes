package pt.atlanse.blog.config

import io.micronaut.context.annotation.ConfigurationProperties

@ConfigurationProperties("articles.default-roles")
class DefaultRolesConfiguration {
	List<String> userRoles
	List<String> adminRoles
}
