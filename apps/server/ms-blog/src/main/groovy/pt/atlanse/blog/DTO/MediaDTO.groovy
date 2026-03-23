package pt.atlanse.blog.DTO

import io.micronaut.serde.annotation.Serdeable

@Serdeable.Deserializable
interface MediaDTO {
	String getFileName()

	String getExtension()

	String getBase64()

	void setFileName(String fileName)

	void setExtension(String extension)

	void setBase64(String base64)

}
