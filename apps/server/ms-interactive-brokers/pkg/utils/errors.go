package utils

import (
	"fmt"
	"github.com/google/uuid"
)

type CustomErrorStruct struct {
	ErrorType    string
	ErrorMessage string
}

func (e CustomErrorStruct) Error() string {
	return e.ErrorMessage
}

func (c CustomErrorStruct) NotFoundError(entity string, uuid uuid.UUID) CustomErrorStruct {
	return CustomErrorStruct{
		ErrorType:    "Not Found",
		ErrorMessage: fmt.Sprintf("%s with UUID %s not found", entity, uuid),
	}
}

func (c CustomErrorStruct) AlreadyExists(entity string, id string) CustomErrorStruct {
	return CustomErrorStruct{
		ErrorType:    "Not Found",
		ErrorMessage: fmt.Sprintf("%s with identity %s already exists", entity, id),
	}
}

func (c CustomErrorStruct) ErrorUploadCSVDuplicatedFK(entity string, code string, store string) CustomErrorStruct {
	return CustomErrorStruct{
		ErrorType:    "Error upload store csv",
		ErrorMessage: fmt.Sprintf("Error saving in %s because already exists code %s with store %s", entity, code, store),
	}
}

func (c CustomErrorStruct) ErrorParsingValue(value string, from string, to string) CustomErrorStruct {
	return CustomErrorStruct{
		ErrorType:    "Error Parsing Value",
		ErrorMessage: fmt.Sprintf("Error parsing value %s from %s to %s", value, from, to),
	}
}

func (c CustomErrorStruct) InternalServerError(message string) CustomErrorStruct {
	return CustomErrorStruct{
		ErrorType:    "Internal Server Error",
		ErrorMessage: message,
	}
}

func (c CustomErrorStruct) BadRequestError(entity string) CustomErrorStruct {
	return CustomErrorStruct{
		ErrorType:    "Bad Request error",
		ErrorMessage: fmt.Sprintf("Error on %s", entity),
	}
}

func (c CustomErrorStruct) ConflictError(message string) CustomErrorStruct {
	return CustomErrorStruct{
		ErrorType:    "Conflict",
		ErrorMessage: fmt.Sprintf(message),
	}
}
