package pt.atlanse.blog.config

import io.micronaut.context.annotation.EachProperty
import io.micronaut.context.annotation.Parameter

@EachProperty("articles.file-type")
class ArticleConfiguration {
	String type

	String directory
	String size
	List<String> allowedFormats

	ArticleConfiguration(@Parameter String type) {
		this.type = type
	}
}
