package pt.atlanse.blog.DTO

import groovy.transform.ToString
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank


// todo SIZE ANNOTATION NOT WORKING
@Introspected
@ToString(includePackage = false, includeFields = true, includeNames = true)
@Serdeable.Deserializable
class TranslationDTO {

	// Default values
	private final int TITLE_MAX_SIZE = 100
	private final int SUBTITLE_MAX_SIZE = 120
	private final int CONTENT_MAX_SIZE = 2000

	@NonNull
	@NotBlank
	private String lang

	@NonNull
	@NotBlank
	private String title

	@NonNull
	@NotBlank
	private String subtitle

	@NonNull
	@NotBlank
	private String content

	@NonNull
	@NotBlank
	private String conclusion

	@NonNull
	@NotBlank
	private boolean enabled

	void setLang(@NonNull String lang) {
		this.lang = lang
	}

	void setConclusion(@NonNull String conclusion) {
		this.conclusion = conclusion
	}

	void setTitle(@NonNull String title) {
		if (title.length() > TITLE_MAX_SIZE) {
			throw new Exception("Title too long; Max $TITLE_MAX_SIZE chars")
		}
		this.title = title
	}

	void setSubtitle(@NonNull String subtitle) {
		if (subtitle.length() > SUBTITLE_MAX_SIZE) {
			throw new Exception("Subtitle too long; Max $SUBTITLE_MAX_SIZE chars")
		}
		this.subtitle = subtitle
	}

	void setContent(@NonNull String content) {
//		if (content.length() > CONTENT_MAX_SIZE) {
//			throw new Exception("Description too long; Max $CONTENT_MAX_SIZE chars")
//		}
		this.content = content
	}

	void isEnabled(boolean enabled) {
		this.enabled = enabled
	}

	@NonNull
	String getLang() {
		return lang
	}

	@NonNull
	String getTitle() {
		return title
	}

	@NonNull
	String getSubtitle() {
		return subtitle
	}

	@NonNull
	String getContent() {
		return content
	}

	@NonNull
	String getConclusion() {
		return conclusion
	}

	boolean isEnabled() {
		return enabled
	}
}
