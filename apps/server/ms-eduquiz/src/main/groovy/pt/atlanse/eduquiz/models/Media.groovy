package pt.atlanse.eduquiz.models

interface Media {
	String getFileName()

	String getExtension()

	String getBase64()

	void setFileName(String fileName)

	void setExtension(String extension)

	void setBase64(String base64)

}
