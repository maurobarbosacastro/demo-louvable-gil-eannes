package pt.atlanse.blog.config

import io.micronaut.context.annotation.EachProperty
import io.micronaut.context.annotation.Parameter

@EachProperty("articles.roles")
class RoleConfiguration {
	String name

	List<String> roles

	RoleConfiguration(@Parameter String name) {
		this.name = name
	}
}
