package pt.atlanse.blog.DTO

import groovy.transform.ToString
import io.micronaut.core.annotation.Introspected
import io.micronaut.core.annotation.NonNull
import io.micronaut.serde.annotation.Serdeable
import jakarta.validation.constraints.NotBlank
import jakarta.validation.constraints.NotNull

@Introspected
@ToString(includePackage = false, includeFields = true, includeNames = true)
@Serdeable.Deserializable
class ArticleDTO {

	@NotNull
	boolean isMobile

	@NonNull
	boolean isWeb

	@NonNull
	@NotBlank
	ImageDTO image

	@NonNull
	@NotBlank
	String status

	@NonNull
	@NotBlank
	TranslationDTO[] translations

	// todo replace by enum maybe?
	private final List<String> DEFAULT_STATUS = ["DRAFT", "PUBLISHED", "ARCHIVED"]

	// todo The translation cannot have repeated languages e.g., lang=en and another object with lang=en
	ArticleDTO() {

	}

	void isMobile(boolean isMobile) {
		this.isMobile = isMobile
	}

	void isWeb(boolean isWeb) {
		this.isWeb = isWeb
	}

	void setImage(@NonNull ImageDTO image) {
		this.image = image
	}

	void setStatus(@NonNull String status) {

		if (!DEFAULT_STATUS.contains(status.toUpperCase())) {
			throw new Exception("Wrong status; Default values are: DRAFT, PUBLISHED and ARCHIVED")
		}

		this.status = status.toUpperCase()
	}

	void setTranslations(@NonNull TranslationDTO[] translations) {
		this.translations = translations
	}

	boolean isMobile() {
		return isMobile
	}

	boolean isWeb() {
		return isWeb
	}

	@NonNull
	ImageDTO getImage() {
		return image
	}

	@NonNull
	String getStatus() {
		return status
	}

	@NonNull
	TranslationDTO[] getTranslations() {
		return translations
	}
}
