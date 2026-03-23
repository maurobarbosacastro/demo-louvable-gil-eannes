package pt.atlanse.blog.domain.imageprocessing

import groovy.transform.ToString
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.GeneratedValue
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.JoinColumn
import jakarta.persistence.OneToOne
import jakarta.persistence.Table
import jakarta.persistence.Version
import jakarta.validation.constraints.NotBlank
import java.time.LocalDateTime

/**
 * @deprecated
 * */
@Entity
@Table(name = "article_media_file")
@ToString(includePackage = false, includeNames = true, includeFields = true)
class ArticleMediaFile {

	@Id
	@GeneratedValue(GeneratedValue.Type.AUTO)
	Long id

	@Version
	private Long version

	@NonNull
	@NotBlank
	@OneToOne
	@JoinColumn(name = "media_file_id", referencedColumnName = "id")
	MediaFileDoc mediaFile

	@NonNull
	@NotBlank
	@Column(name = "code")
	String code

	@NonNull
	@NotBlank
	@Column(name = "created_by")
	String createdBy

	@GeneratedValue
	@Column(name = "created_at")
	LocalDateTime createdAt

	@NonNull
	@NotBlank
	@Column(name = "updated_by")
	String updatedBy

	@GeneratedValue
	@Column(name = "updated_at")
	LocalDateTime updatedAt

	ArticleMediaFile() {

	}

	void setCreatedAt(LocalDateTime createdAt) {
		this.createdAt = createdAt
	}

	void setUpdatedAt(LocalDateTime updatedAt) {
		this.updatedAt = updatedAt
	}

	LocalDateTime getCreatedAt() {
		return createdAt
	}

	LocalDateTime getUpdatedAt() {
		return updatedAt
	}

	void setId(Long id) {
		this.id = id
	}

	void setVersion(Long version) {
		this.version = version
	}

	void setMediaFile(@NonNull MediaFileDoc mediaFile) {
		this.mediaFile = mediaFile
	}

	void setCode(@NonNull String code) {
		this.code = code
	}

	void setCreatedBy(@NonNull String createdBy) {
		this.createdBy = createdBy
	}


	void setUpdatedBy(@NonNull String updatedBy) {
		this.updatedBy = updatedBy
	}


	Long getId() {
		return id
	}

	Long getVersion() {
		return version
	}

	@NonNull
	MediaFileDoc getMediaFile() {
		return mediaFile
	}

	@NonNull
	String getCode() {
		return code
	}

	@NonNull
	String getCreatedBy() {
		return createdBy
	}

	@NonNull
	String getUpdatedBy() {
		return updatedBy
	}
}
