package pt.atlanse.blog.domain.imageprocessing

import groovy.transform.ToString
import groovy.transform.TupleConstructor
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.annotation.GeneratedValue
import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import jakarta.persistence.Version
import jakarta.validation.constraints.NotBlank

import java.time.LocalDateTime


/**
 * @deprecated
 * */
@Entity
@Table(name = "media_file_doc")
@ToString(includePackage = false, includeNames = true, includeFields = true)
@TupleConstructor
class MediaFileDoc {

	@Id
	@GeneratedValue(GeneratedValue.Type.AUTO)
	Long id

	@Version
	private Long version

	@NonNull
	@NotBlank
	@Column(name = "hash")
	String hash

	@NonNull
	@NotBlank
	@Column(name = "name")
	String name

	@NonNull
	@NotBlank
	@Column(name = "path")
	String path

	// 3 UUIDs appended that are used as the unique key to access this resource
	@NonNull
	@NotBlank
	@Column(name = "public_code")
	String publicCode

	@GeneratedValue
	@Column(name = "created_at")
	LocalDateTime createdAt

	@GeneratedValue
	@Column(name = "updated_at")
	LocalDateTime updatedAt

	MediaFileDoc() {}

	void invalidatePublicCode() {
		publicCode = UUID.randomUUID().toString() + UUID.randomUUID().toString() + UUID.randomUUID().toString()
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

	void setHash(@NonNull String hash) {
		this.hash = hash
	}

	void setName(@NonNull String name) {
		this.name = name
	}

	void setPath(@NonNull String path) {
		this.path = path
	}

	void setPublicCode(@NonNull String publicCode) {
		this.publicCode = publicCode
	}


	Long getId() {
		return id
	}

	Long getVersion() {
		return version
	}

	@NonNull
	String getHash() {
		return hash
	}

	@NonNull
	String getName() {
		return name
	}

	@NonNull
	String getPath() {
		return path
	}

	@NonNull
	String getPublicCode() {
		return publicCode
	}

}
