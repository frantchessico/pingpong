package zogo

import (
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


// func NewObjectSchema(fields map[string]FieldValidator) FieldValidator {
// 	return func(value interface{}) error {
// 		data, ok := value.(map[string]interface{})
// 		if !ok {
// 			return errors.New("must be an object")
// 		}

// 		for fieldName, validator := range fields {
// 			fieldValue, exists := data[fieldName]
// 			if !exists {
// 				return fmt.Errorf("field '%s' not found in object", fieldName)
// 			}

// 			if err := validator(fieldValue); err != nil {
// 				return fmt.Errorf("field '%s': %v", fieldName, err)
// 			}
// 		}

// 		return nil
// 	}
// }


func NewObjectSchema(validators map[string]FieldValidator) FieldValidator {
	return func(value interface{}) error {
		if data, ok := value.(map[string]interface{}); ok {
			for fieldName, validator := range validators {
				fieldValue, exists := data[fieldName]
				if !exists {
					return fmt.Errorf("field '%s' not found", fieldName)
				}
				if err := validator(fieldValue); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
