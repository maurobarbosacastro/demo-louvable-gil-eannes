package pt.atlanse.blog.DTO

import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.Nullable
import io.micronaut.serde.annotation.Serdeable

@Introspected
@TupleConstructor
@Serdeable.Deserializable
class ArticleParameters {

	@Nullable
	String lang = "en"

	@Nullable
	boolean image = false

	@Nullable
	String searchText

	@Nullable
	String status

	@Nullable
	String target

	@Nullable
	private boolean viewMobile

	@Nullable
	private boolean viewWeb

	// Target it's either web or mobile
	void setTarget(@Nullable String target) {
		this.target = target

		if (target == "web") {
			this.viewWeb = true
			return
		}

		this.viewMobile = true
	}

	boolean isViewMobile() {
		return viewMobile
	}

	boolean isViewWeb() {
		return viewWeb
	}

//	@Nullable
//	Optional<Boolean> getImage() {
//		return Optional.ofNullable(image)
//	}
//
//	@Nullable
//	Optional<String> getLang() {
//		return Optional.ofNullable(lang)
//	}
//
//	@Nullable
//	Optional<String> getFilter() {
//		return Optional.ofNullable(filter)
//	}

}
