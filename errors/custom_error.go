package errors

import "link-converter/models"

func NewCustomErr(text string) *models.CustomError {
	return &models.CustomError{Message: text}
}
