package bookvalidator

import (
	"github.com/eghbalii/libManager/contract/goproto/book"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateAddBookRequest(req book.Book) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Title, validation.Required, validation.Length(3, 100)),
		validation.Field(&req.Author, validation.Required),
		validation.Field(&req.ISBN, validation.Required, validation.By(v.checkISBNUniqueness)),
		validation.Field(&req.PublishDate, validation.Required),
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

func (v Validator) checkISBNUniqueness(value interface{}) error {
	isbn := value.(string)
	isUnique := v.repo.IsISBNUnique(isbn)
	if !isUnique {
		return validation.NewError("ISBN", "ISBN is not unique")
	}
	return nil
}
