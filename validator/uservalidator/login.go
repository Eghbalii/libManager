package uservalidator

import (
	"regexp"

	"github.com/eghbalii/libManager/param"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateLoginRequest(req param.LoginRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,

		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error("phone number is not valid"),
			validation.By(v.doesPhoneNumberExist)),

		validation.Field(&req.Password, validation.Required),
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

func (v Validator) doesPhoneNumberExist(value interface{}) error {
	phoneNumber := value.(string)
	_, err := v.repo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return err
	}
	return nil
}
