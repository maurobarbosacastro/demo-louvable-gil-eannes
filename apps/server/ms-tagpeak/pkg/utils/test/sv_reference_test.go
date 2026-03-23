package test

import (
	"ms-tagpeak/pkg/utils"
	"strings"
	"testing"
)

func TestEncodeAwinStoreVisitRef(t *testing.T) {
	t.Run("EncodeStoreVisitRef", func(t *testing.T) {
		ref := "TP-37"
		userUuid := "fdc6c557-9e67-41c3-9042-4e4618538ae9"

		result := utils.EncodeAwinStoreVisitRef(ref, userUuid)
		if result != "VFAtMzdfZmRjNmM1NTc=" {
			t.Errorf("Expected result to be 'VFAtMzdfZmRjNmM1NTc=', got '%s'", result)
		}
	})
}

func TestDecodeStoreVisitRef(t *testing.T) {
	t.Run("DecodeStoreVisitRef", func(t *testing.T) {
		ref := "TP-37"
		userUuid := "fdc6c557-9e67-41c3-9042-4e4618538ae9"
		encodedRef := utils.EncodeAwinStoreVisitRef(ref, userUuid)
		userUuidSplit := strings.Split(userUuid, "-")

		decodedRef, sectionCheck := utils.DecodeStoreVisitRef(encodedRef)
		if decodedRef != ref {
			t.Errorf("Expected decodedRef to be '%s', got '%s'", ref, decodedRef)
		}
		if sectionCheck != userUuidSplit[0] {
			t.Errorf("Expected sectionCheck to be '%s', got '%s'", sectionCheck, userUuidSplit[0])
		}

	})
}

func TestValidateAwinDecodedStoreVisitRef(t *testing.T) {
	t.Run("ValidateDecodedStoreVisitRef", func(t *testing.T) {
		ref := "TP-37"
		userUuid := "fdc6c557-9e67-41c3-9042-4e4618538ae9"
		encodedRef := utils.EncodeAwinStoreVisitRef(ref, userUuid)
		_, sectionCheck := utils.DecodeStoreVisitRef(encodedRef)

		if !utils.ValidateAwinDecodedStoreVisitRef(sectionCheck, userUuid) {
			t.Errorf("Expected ValidateDecodedStoreVisitRef to be true, got false")
		}
	})
}
