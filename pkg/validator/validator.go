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
	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors){
		return err
	}

	for _, e := range validationErrors {
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
			case "min":
				messages = append(messages, 
				e.Field()+" must be at least"+e.Param()+" characters")
			case "max":
				messages = append(messages, 
				e.Field()+" must not exceed"+e.Param()+" characters")
				case "uiid":
					messages = append(messages, 
					e.Field()+" must be a valid UUID")
			default:
				messages = append(messages, 
				e.Error())
			
		}
	}
	return errors.New(strings.Join(messages, ", "))
}