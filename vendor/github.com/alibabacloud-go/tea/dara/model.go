package dara

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Model interface {
	Validate() error
	ToMap() map[string]interface{}
	copyWithouStream() Model
}

func Validate(params interface{}) error {
	if params == nil {
		return nil
	}
	requestValue := reflect.ValueOf(params)
	if requestValue.IsNil() {
		return nil
	}
	err := validate(requestValue.Elem())
	return err
}

// Verify whether the parameters meet the requirements
func validate(dataValue reflect.Value) error {
	if strings.HasPrefix(dataValue.Type().String(), "*") { // Determines whether the input is a structure object or a pointer object
		if dataValue.IsNil() {
			return nil
		}
		dataValue = dataValue.Elem()
	}
	dataType := dataValue.Type()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		valueField := dataValue.Field(i)
		for _, value := range validateParams {
			err := validateParam(field, valueField, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateParam(field reflect.StructField, valueField reflect.Value, tagName string) error {
	tag, containsTag := field.Tag.Lookup(tagName) // Take out the checked regular expression
	if containsTag && tagName == "require" {
		err := checkRequire(field, valueField)
		if err != nil {
			return err
		}
	}
	if strings.HasPrefix(field.Type.String(), "[]") { // Verify the parameters of the array type
		err := validateSlice(field, valueField, containsTag, tag, tagName)
		if err != nil {
			return err
		}
	} else if valueField.Kind() == reflect.Ptr { // Determines whether it is a pointer object
		err := validatePtr(field, valueField, containsTag, tag, tagName)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateSlice(field reflect.StructField, valueField reflect.Value, containsregexpTag bool, tag, tagName string) error {
	if valueField.IsValid() && !valueField.IsNil() { // Determines whether the parameter has a value
		if containsregexpTag {
			if tagName == "maxItems" {
				err := checkMaxItems(field, valueField, tag)
				if err != nil {
					return err
				}
			}

			if tagName == "minItems" {
				err := checkMinItems(field, valueField, tag)
				if err != nil {
					return err
				}
			}
		}

		for m := 0; m < valueField.Len(); m++ {
			elementValue := valueField.Index(m)
			if elementValue.Type().Kind() == reflect.Ptr { // Determines whether the child elements of an array are of a basic type
				err := validatePtr(field, elementValue, containsregexpTag, tag, tagName)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func validatePtr(field reflect.StructField, elementValue reflect.Value, containsregexpTag bool, tag, tagName string) error {
	if elementValue.IsNil() {
		return nil
	}
	if isFilterType(elementValue.Elem().Type().String(), basicTypes) {
		if containsregexpTag {
			if tagName == "pattern" {
				err := checkPattern(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}

			if tagName == "maxLength" {
				err := checkMaxLength(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}

			if tagName == "minLength" {
				err := checkMinLength(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}

			if tagName == "maximum" {
				err := checkMaximum(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}

			if tagName == "minimum" {
				err := checkMinimum(field, elementValue.Elem(), tag)
				if err != nil {
					return err
				}
			}
		}
	} else {
		err := validate(elementValue)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkRequire(field reflect.StructField, valueField reflect.Value) error {
	name, _ := field.Tag.Lookup("json")
	strs := strings.Split(name, ",")
	name = strs[0]
	if !valueField.IsNil() && valueField.IsValid() {
		return nil
	}
	return errors.New(name + " should be setted")
}

func checkPattern(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		value := valueField.String()
		r, _ := regexp.Compile("^" + tag + "$")
		if match := r.MatchString(value); !match { // Determines whether the parameter value satisfies the regular expression or not, and throws an error
			return errors.New(value + " is not matched " + tag)
		}
	}
	return nil
}

func checkMaxItems(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		maxItems, err := strconv.Atoi(tag)
		if err != nil {
			return err
		}
		length := valueField.Len()
		if maxItems < length {
			errMsg := fmt.Sprintf("The length of %s is %d which is more than %d", field.Name, length, maxItems)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMinItems(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() {
		minItems, err := strconv.Atoi(tag)
		if err != nil {
			return err
		}
		length := valueField.Len()
		if minItems > length {
			errMsg := fmt.Sprintf("The length of %s is %d which is less than %d", field.Name, length, minItems)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMaxLength(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		maxLength, err := strconv.Atoi(tag)
		if err != nil {
			return err
		}
		length := valueField.Len()
		if valueField.Kind().String() == "string" {
			length = strings.Count(valueField.String(), "") - 1
		}
		if maxLength < length {
			errMsg := fmt.Sprintf("The length of %s is %d which is more than %d", field.Name, length, maxLength)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMinLength(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() {
		minLength, err := strconv.Atoi(tag)
		if err != nil {
			return err
		}
		length := valueField.Len()
		if valueField.Kind().String() == "string" {
			length = strings.Count(valueField.String(), "") - 1
		}
		if minLength > length {
			errMsg := fmt.Sprintf("The length of %s is %d which is less than %d", field.Name, length, minLength)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMaximum(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		maximum, err := strconv.ParseFloat(tag, 64)
		if err != nil {
			return err
		}
		byt, _ := json.Marshal(valueField.Interface())
		num, err := strconv.ParseFloat(string(byt), 64)
		if err != nil {
			return err
		}
		if maximum < num {
			errMsg := fmt.Sprintf("The size of %s is %f which is greater than %f", field.Name, num, maximum)
			return errors.New(errMsg)
		}
	}
	return nil
}

func checkMinimum(field reflect.StructField, valueField reflect.Value, tag string) error {
	if valueField.IsValid() && valueField.String() != "" {
		minimum, err := strconv.ParseFloat(tag, 64)
		if err != nil {
			return err
		}

		byt, _ := json.Marshal(valueField.Interface())
		num, err := strconv.ParseFloat(string(byt), 64)
		if err != nil {
			return err
		}
		if minimum > num {
			errMsg := fmt.Sprintf("The size of %s is %f which is less than %f", field.Name, num, minimum)
			return errors.New(errMsg)
		}
	}
	return nil
}
