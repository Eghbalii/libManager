package bookvalidator

import (
	"github.com/eghbalii/libManager/contract/goproto/book"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateUpdateBookRequest(req book.Book) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.ISBN, validation.Required),
	); err != nil {
		fielErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fielErrors[key] = value.Error()
				}
			}
		}
		return fielErrors, err
	}
	return nil, nil
}
