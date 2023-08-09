package zogo

import (
	"errors"
	"fmt"
)

// FieldValidator defines the signature of a field validation function.
type FieldValidator func(value interface{}) error

// RuleValidator is a struct to manage validation rules for fields.
type RuleValidator struct {
	Rules map[string][]FieldValidator
}

// NewRuleValidator creates a new instance of RuleValidator.
func NewRuleValidator() *RuleValidator {
	return &RuleValidator{
		Rules: make(map[string][]FieldValidator),
	}
}

// AddRule adds a validation rule for a specific field.
func (v *RuleValidator) AddRule(field string, validators ...FieldValidator) {
	v.Rules[field] = append(v.Rules[field], validators...)
}

// Validate validates the provided data against the defined rules.
func (v *RuleValidator) Validate(data map[string]interface{}) error {
	for field, rules := range v.Rules {
		value, exists := data[field]
		if !exists {
			return fmt.Errorf("field '%s' not found", field)
		}

		for _, rule := range rules {
			if err := rule(value); err != nil {
				return fmt.Errorf("field '%s': %v", field, err)
			}
		}
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

func (v *StringValidator) Min(minLength int) *StringValidator {
	v.validators = append(v.validators, func(value string) error {
		if len(value) < minLength {
			return fmt.Errorf("must have a minimum length of %d", minLength)
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







