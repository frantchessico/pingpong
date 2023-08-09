package zogo

import (
	"errors"
	"fmt"
	"regexp"
)

func MinValueValidator(minValue int) FieldValidator {
	return func(value interface{}) error {
		num, ok := value.(int)
		if !ok {
			return errors.New("must be an integer")
		}
		if num < minValue {
			return fmt.Errorf("must be greater than or equal to %d", minValue)
		}
		return nil
	}
}

func MaxValueValidator(maxValue float64) FieldValidator {
	return func(value interface{}) error {
		num, ok := value.(float64)
		if !ok {
			return errors.New("must be a number")
		}
		if num > maxValue {
			return fmt.Errorf("must be less than or equal to %f", maxValue)
		}
		return nil
	}
}

func StringNotEmptyValidator(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("must be a string")
	}
	if str == "" {
		return errors.New("cannot be empty")
	}
	return nil
}



// EmailRegexValidator is a field validation function for email using RegEx.
var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func EmailSchema(value interface{}) error {
	if value == nil || value == "" {
		return nil
	}

	email, ok := value.(string)
	if !ok {
		return fmt.Errorf("must be a string or nil")
	}

	if !EmailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}

	return nil
}


func BooleanSchema(value interface{}) error {
	if value == nil {
		return nil // Field is undefined or null
	}

	_, ok := value.(bool)
	if !ok {
		return fmt.Errorf("must be a boolean or nil")
	}

	return nil
}


// StringSchema is a field validation function for string values.
func StringSchema(value interface{}) error {
	if value == nil {
		return nil // Field is undefined or null
	}

	_, ok := value.(string)
	if !ok {
		return fmt.Errorf("must be a string or nil")
	}

	return nil
}


func NumberSchema(value interface{}) error {
	if value == nil {
		return nil // Field is undefined or null
	}

	_, ok := value.(float64)
	if !ok {
		return fmt.Errorf("must be a number or nil")
	}

	return nil
}





type StringValidator struct {
	validators []func(value string) error
}

func NewStringValidator() *StringValidator {
	return &StringValidator{}
}

func (v *StringValidator) Email() *StringValidator {
	v.validators = append(v.validators, func(value string) error {
		if !EmailRegex.MatchString(value) {
			return errors.New("invalid email format")
		}
		return nil
	})
	return v
}

func (v *StringValidator) NonEmpty() *StringValidator {
	v.validators = append(v.validators, func(value string) error {
		if value == "" {
			return errors.New("cannot be empty")
		}
		return nil
	})
	return v
}

func (v *StringValidator) Validate(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("must be a string")
	}

	for _, validator := range v.validators {
		if err := validator(strValue); err != nil {
			return err
		}
	}

	return nil
}



