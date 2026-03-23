package utils

import (
	"encoding/base64"
	"strings"
)

// EncodeStoreVisitRef Takes a ref (TP-XXX) and the 1st section of the user uuid to create a unique ref
func EncodeAwinStoreVisitRef(ref string, userUuid string) string {
	usersUuidSplit := strings.Split(userUuid, "-")
	return base64.StdEncoding.EncodeToString([]byte(ref + "_" + usersUuidSplit[0]))
}

// DecodeStoreVisitRef Takes and encoded ref and returns both parts of it (TP ref and user uuid 1st section)
func DecodeStoreVisitRef(encodedRef string) (string, string) {
	decodedRef, err := base64.StdEncoding.DecodeString(encodedRef)
	if err != nil {
		// To support legacy store visits, we return the ref if failed to decode.
		return encodedRef, ""
	}
	splitDecodedRef := strings.Split(string(decodedRef), "_")
	return splitDecodedRef[0], splitDecodedRef[1]
}

// ValidateDecodedStoreVisitRef Validates the 1st section of the user uuid against the sectionCheck from the decode ref.
func ValidateAwinDecodedStoreVisitRef(sectionCheck string, userUuid string) bool {
	usersUuidSplit := strings.Split(userUuid, "-")

	if sectionCheck != usersUuidSplit[0] {
		return false
	}

	return true
}
