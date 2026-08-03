package validator

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(v any) error {
	err := validate.Struct(v)
	if err == nil {
		return nil
	}

	var messages []string

	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag(){
			case "required":
				messages = append(messages, 
				e.Field()+" is required")
			case "email":
				messages = append(messages, 
				e.Field()+" must be a valid email")
			case "gt":
				messages = append(messages, 
				e.Field()+ " must be greater than "+e.Param())
			case "gte":
				messages = append(messages, 
				e.Field()+ " must be greater than or equal to " +e.Param())
			case "oneof":
				messages = append(messages, 
				e.Field()+" must be one of: "+e.Param())
			default:
				messages = append(messages, 
				e.Error())
			
		}
	}
	return errors.New(strings.Join(messages, ", "))
}