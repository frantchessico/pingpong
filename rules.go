package zogo

import (
	"errors"
	"fmt"
	"regexp"
)

func MinLengthValidator(minLength int) FieldValidator {
	return func(value interface{}) error {
		strValue, ok := value.(string)
		if !ok {
			return errors.New("must be a string")
		}
		if len(strValue) < minLength {
			return fmt.Errorf("must have a minimum length of %d", minLength)
		}
		return nil
	}
}

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

func MaxLengthValidator(maxLength int) FieldValidator {
	return func(value interface{}) error {
		strValue, ok := value.(string)
		if !ok {
			return errors.New("must be a string")
		}
		if len(strValue) > maxLength {
			return fmt.Errorf("must have a maximum length of %d", maxLength)
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


func CombineValidators(validators ...FieldValidator) FieldValidator {
	return func(value interface{}) error {
		for _, validator := range validators {
			if err := validator(value); err != nil {
				return err
			}
		}
		return nil
	}
}


func EmailMinMaxLengthValidator(min, max int) FieldValidator {
	return CombineValidators(
		StringSchema,
		EmailSchema,
		MinLengthValidator(min),
		MaxLengthValidator(max),
	)
}


func NestedFieldValidator() FieldValidator {
	return CombineValidators(
		NewObjectSchema(map[string]FieldValidator{
			"here": StringSchema,
		}),
	)
}


func NestedObjectValidator() FieldValidator {
	return CombineValidators(
		NewObjectSchema(map[string]FieldValidator{
			"somewhere": NewObjectSchema(map[string]FieldValidator{
				"here": StringSchema,
			}),
		}),
	)
}


func SecondLevelFieldValidator(fieldName string, fieldValidator FieldValidator) FieldValidator {
	return func(value interface{}) error {
		if data, ok := value.(map[string]interface{}); ok {
			if fieldValue, exists := data[fieldName]; exists {
				return fieldValidator(fieldValue)
			}
		}
		return nil
	}
}
