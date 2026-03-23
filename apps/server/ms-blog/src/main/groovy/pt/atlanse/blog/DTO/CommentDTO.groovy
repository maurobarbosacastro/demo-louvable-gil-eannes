package pt.atlanse.blog.DTO

import groovy.transform.ToString
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank


@Introspected
@ToString(includePackage = false, includeFields = true, includeNames = true)
@Serdeable.Deserializable
class CommentDTO {

	@NonNull
	@NotBlank
	String text

	@Nullable
	String parentId

	CommentDTO(){}

	@NonNull
	String getText() {
		return text
	}

	@Nullable
	String getParentId() {
		return parentId
	}

	void setText(@NonNull String text) {
		this.text = text
	}

	void setParentId(@Nullable String parentId) {
		this.parentId = parentId
	}
}
