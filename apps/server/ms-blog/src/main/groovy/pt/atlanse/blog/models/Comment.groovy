package pt.atlanse.blog.models

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Nullable

@TupleConstructor
class Comment {
	Long id // comment id


	String author

	@Nullable
	String creator

	String content
	String text
	String createdAt
	String updatedAt
	Long likes
	Long comments
	boolean liked
}
