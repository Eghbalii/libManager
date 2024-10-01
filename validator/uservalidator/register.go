package uservalidator

import (
	"fmt"
	"regexp"

	"github.com/eghbalii/libManager/param"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name,
			validation.Required,
			validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required),
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error("phone number is not valid"),
			validation.By(v.checkPhoneNumberUniqueness)),
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

func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)
	isUnique, err := v.repo.IsPhoneNumberUnique(phoneNumber)
	if err != nil {
		return err
	}
	if !isUnique {
		return fmt.Errorf("phone number is already exists")
	}
	return nil
}
