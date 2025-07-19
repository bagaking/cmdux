// Package input provides form components.
package input

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bagaking/cmdux/style"
)

// Form represents a collection of input fields.
type Form struct {
	title       string
	fields      []FormField
	titleStyle  *style.Color
	labelStyle  *style.Color
	inputStyle  *style.Color
	errorStyle  *style.Color
	results     map[string]interface{}
}

// FormField represents a single form field.
type FormField struct {
	Name        string
	Label       string
	Type        FieldType
	Required    bool
	Default     interface{}
	Options     []string
	Validator   func(interface{}) error
	Transformer func(string) interface{}
}

// FieldType represents the type of form field.
type FieldType int

const (
	FieldTypeText FieldType = iota
	FieldTypePassword
	FieldTypeNumber
	FieldTypeBoolean
	FieldTypeSelect
	FieldTypeMultiSelect
)

// NewForm creates a new form.
func NewForm(title string) *Form {
	return &Form{
		title:      title,
		fields:     []FormField{},
		titleStyle: style.Primary,
		labelStyle: style.Secondary,
		inputStyle: style.Primary,
		errorStyle: style.Error,
		results:    make(map[string]interface{}),
	}
}

// AddField adds a field to the form.
func (f *Form) AddField(field FormField) *Form {
	f.fields = append(f.fields, field)
	return f
}

// TextField adds a text input field.
func (f *Form) TextField(name, label string, required bool, defaultValue ...string) *Form {
	field := FormField{
		Name:     name,
		Label:    label,
		Type:     FieldTypeText,
		Required: required,
	}
	
	if len(defaultValue) > 0 {
		field.Default = defaultValue[0]
	}
	
	return f.AddField(field)
}

// PasswordField adds a password input field.
func (f *Form) PasswordField(name, label string, required bool) *Form {
	field := FormField{
		Name:     name,
		Label:    label,
		Type:     FieldTypePassword,
		Required: required,
	}
	
	return f.AddField(field)
}

// NumberField adds a number input field.
func (f *Form) NumberField(name, label string, required bool, defaultValue ...int) *Form {
	field := FormField{
		Name:     name,
		Label:    label,
		Type:     FieldTypeNumber,
		Required: required,
	}
	
	if len(defaultValue) > 0 {
		field.Default = defaultValue[0]
	}
	
	return f.AddField(field)
}

// BooleanField adds a boolean (yes/no) field.
func (f *Form) BooleanField(name, label string, defaultValue ...bool) *Form {
	field := FormField{
		Name:     name,
		Label:    label,
		Type:     FieldTypeBoolean,
		Required: false, // Boolean fields are never required
	}
	
	if len(defaultValue) > 0 {
		field.Default = defaultValue[0]
	}
	
	return f.AddField(field)
}

// SelectField adds a single-select field.
func (f *Form) SelectField(name, label string, options []string, required bool) *Form {
	field := FormField{
		Name:     name,
		Label:    label,
		Type:     FieldTypeSelect,
		Required: required,
		Options:  options,
	}
	
	return f.AddField(field)
}

// MultiSelectField adds a multi-select field.
func (f *Form) MultiSelectField(name, label string, options []string) *Form {
	field := FormField{
		Name:     name,
		Label:    label,
		Type:     FieldTypeMultiSelect,
		Required: false,
		Options:  options,
	}
	
	return f.AddField(field)
}

// Run executes the form and collects all input.
func (f *Form) Run() (map[string]interface{}, error) {
	// Display form title
	if f.title != "" {
		fmt.Println(f.titleStyle.Sprint("=== " + f.title + " ==="))
		fmt.Println()
	}
	
	// Process each field
	for _, field := range f.fields {
		value, err := f.processField(field)
		if err != nil {
			return nil, err
		}
		f.results[field.Name] = value
	}
	
	return f.results, nil
}

func (f *Form) processField(field FormField) (interface{}, error) {
	switch field.Type {
	case FieldTypeText:
		return f.processTextField(field)
	case FieldTypePassword:
		return f.processPasswordField(field)
	case FieldTypeNumber:
		return f.processNumberField(field)
	case FieldTypeBoolean:
		return f.processBooleanField(field)
	case FieldTypeSelect:
		return f.processSelectField(field)
	case FieldTypeMultiSelect:
		return f.processMultiSelectField(field)
	default:
		return nil, fmt.Errorf("unknown field type: %v", field.Type)
	}
}

func (f *Form) processTextField(field FormField) (string, error) {
	prompt := NewPrompt(field.Label).
		Required(field.Required)
	
	if field.Default != nil {
		if defaultStr, ok := field.Default.(string); ok {
			prompt.Default(defaultStr)
		}
	}
	
	if field.Validator != nil {
		prompt.Validator(func(input string) error {
			return field.Validator(input)
		})
	}
	
	if field.Transformer != nil {
		prompt.Transformer(func(input string) string {
			if result := field.Transformer(input); result != nil {
				if str, ok := result.(string); ok {
					return str
				}
			}
			return input
		})
	}
	
	return prompt.Run()
}

func (f *Form) processPasswordField(field FormField) (string, error) {
	return Password(field.Label)
}

func (f *Form) processNumberField(field FormField) (int, error) {
	prompt := NewPrompt(field.Label).
		Required(field.Required).
		Validator(func(input string) error {
			if input == "" && !field.Required {
				return nil
			}
			_, err := strconv.Atoi(input)
			return err
		})
	
	if field.Default != nil {
		if defaultInt, ok := field.Default.(int); ok {
			prompt.Default(strconv.Itoa(defaultInt))
		}
	}
	
	input, err := prompt.Run()
	if err != nil {
		return 0, err
	}
	
	if input == "" {
		if field.Default != nil {
			if defaultInt, ok := field.Default.(int); ok {
				return defaultInt, nil
			}
		}
		return 0, nil
	}
	
	return strconv.Atoi(input)
}

func (f *Form) processBooleanField(field FormField) (bool, error) {
	defaultVal := false
	if field.Default != nil {
		if defaultBool, ok := field.Default.(bool); ok {
			defaultVal = defaultBool
		}
	}
	
	return Confirm(field.Label, defaultVal)
}

func (f *Form) processSelectField(field FormField) (string, error) {
	_, selected, err := Select(field.Label, field.Options)
	return selected, err
}

func (f *Form) processMultiSelectField(field FormField) ([]string, error) {
	_, selected, err := MultiSelect(field.Label, field.Options)
	return selected, err
}

// GetResult gets a specific field result by name.
func (f *Form) GetResult(name string) interface{} {
	return f.results[name]
}

// GetString gets a string field result.
func (f *Form) GetString(name string) string {
	if value, ok := f.results[name].(string); ok {
		return value
	}
	return ""
}

// GetInt gets an integer field result.
func (f *Form) GetInt(name string) int {
	if value, ok := f.results[name].(int); ok {
		return value
	}
	return 0
}

// GetBool gets a boolean field result.
func (f *Form) GetBool(name string) bool {
	if value, ok := f.results[name].(bool); ok {
		return value
	}
	return false
}

// GetStringSlice gets a string slice field result.
func (f *Form) GetStringSlice(name string) []string {
	if value, ok := f.results[name].([]string); ok {
		return value
	}
	return []string{}
}

// Bind binds the form results to a struct.
func (f *Form) Bind(target interface{}) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}
	
	v = v.Elem()
	t := v.Type()
	
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		
		if !fieldValue.CanSet() {
			continue
		}
		
		// Look for form tag or use field name
		name := field.Tag.Get("form")
		if name == "" {
			name = strings.ToLower(field.Name)
		}
		
		if result, exists := f.results[name]; exists {
			if err := f.setFieldValue(fieldValue, result); err != nil {
				return fmt.Errorf("error setting field %s: %v", name, err)
			}
		}
	}
	
	return nil
}

func (f *Form) setFieldValue(fieldValue reflect.Value, result interface{}) error {
	resultValue := reflect.ValueOf(result)
	
	if resultValue.Type().AssignableTo(fieldValue.Type()) {
		fieldValue.Set(resultValue)
		return nil
	}
	
	// Try type conversion
	if resultValue.Type().ConvertibleTo(fieldValue.Type()) {
		fieldValue.Set(resultValue.Convert(fieldValue.Type()))
		return nil
	}
	
	return fmt.Errorf("cannot assign %T to %T", result, fieldValue.Interface())
}