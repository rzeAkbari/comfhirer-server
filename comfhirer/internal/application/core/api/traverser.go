package api

import (
	"errors"
	"fmt"
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	"reflect"
	"strconv"
	"unicode"
)

var typeRegistry = map[string]reflect.Type{
	"Patient":    reflect.TypeOf(domain.Patient{}),
	"Medication": reflect.TypeOf(domain.Medication{}),
}

type Traverser struct {
}

func (t Traverser) Travers(ast []domain.ASTNode) (domain.Bundle, []error) {
	var instances = map[string]reflect.Value{}
	traversErr := []error{}
	r := domain.Bundle{
		ResourceType: "Bundle",
	}

	for _, node := range ast {
		fhirResource, err := fhirResourceInstance(node.NodeName, node.NodeIndex, instances)
		if err != nil {
			traversErr = append(traversErr, err)
			continue
		}
		_, err = setField(fhirResource, &node.FhirField, node.NodeValue)
		if err != nil {
			traversErr = append(traversErr, err)
			continue
		}
	}

	for _, fhirResource := range instances {
		r.Entry = append(r.Entry, domain.BundleEntry{Resource: fhirResource.Interface()})
	}

	return r, traversErr
}

func fhirResourceInstance(fhirResourceName string, index string, instances map[string]reflect.Value) (reflect.Value, error) {
	instanceKey := fhirResourceName + "_" + index
	if instances[instanceKey].IsValid() {
		return instances[instanceKey], nil
	}

	fhirResourceType := typeRegistry[fhirResourceName]

	if fhirResourceType == nil {
		return reflect.Value{}, errors.New(fmt.Sprintf("unknown resource %q", fhirResourceName))
	}

	instance := reflect.New(fhirResourceType).Elem()
	setFieldValue(instance, "resourceType", fhirResourceName)

	instances[instanceKey] = instance

	return instance, nil
}

func setField(fhirResource reflect.Value, field *domain.FhirField, value any) (reflect.Value, error) {
	nestedField := field.FhirField

	if field.FieldParsedType == domain.MultipleValueField {
		index, _ := strconv.Atoi(field.Name)
		if index < 0 {
			return reflect.Value{}, errors.New("negative index detected")
		}
		if addNewFhirField(fhirResource, index) {
			fhirResource.Set(reflect.Append(fhirResource, reflect.ValueOf(value)))
		}
		if hasToPopulate(fhirResource, index) {
			for i := 0; i < index; i++ {
				placeHolder := reflect.ValueOf(value)
				fhirResource.Set(reflect.Append(fhirResource, placeHolder))
			}
			fhirResource.Set(reflect.Append(fhirResource, reflect.ValueOf(value)))
		}
		if containsFhirField(fhirResource, index) {
			position := fhirResource.Index(index)
			position.Set(reflect.ValueOf(value))
		}

		return fhirResource, nil
	}
	if nestedField == nil {
		return setFieldValue(fhirResource, field.Name, value)
	}
	fieldName := camelToPascalCase(field.Name)

	if field.FieldParsedType == domain.SingleField {
		var fhirField reflect.Value

		if fieldPointer(fhirResource) {
			fhirField = reflect.Indirect(fhirResource).FieldByName(fieldName)
			if !fhirField.IsValid() && !fhirField.CanSet() {
				return reflect.Value{}, errors.New(fmt.Sprintf("unknown field %q", fieldName))
			}
			if nilFieldPointer(fhirField) {
				fhirField = reflect.New(fhirField.Type().Elem())
			}
		}
		if emptySlice(fhirResource, fieldName) {
			fhirField = reflect.New(fhirResource.FieldByName(fieldName).Type()).Elem() //catch exception when field not exists
		}
		if !fhirField.IsValid() {
			fhirField = fhirResource.FieldByName(fieldName)
			if !fhirField.IsValid() && !fhirField.CanSet() {
				return reflect.Value{}, errors.New(fmt.Sprintf("unknown field %q", fieldName))
			}
			if nilFieldPointer(fhirField) {
				fhirField = reflect.New(fhirField.Type().Elem())
			}
		}

		_, err := setField(fhirField, nestedField, value)
		if err != nil {
			return reflect.Value{}, err
		}
		setFieldValue(fhirResource, field.Name, fhirField.Interface())
	}
	if field.FieldParsedType == domain.MultipleNestedField {
		var fhirField reflect.Value
		fhirFieldIndex, _ := strconv.Atoi(field.Name)
		if fhirFieldIndex < 0 {
			return reflect.Value{}, errors.New("negative position detected")
		}
		if containsFhirField(fhirResource, fhirFieldIndex) {
			fhirField = reflect.ValueOf(fhirResource.Interface()).Index(fhirFieldIndex)
			setField(fhirField, nestedField, value)
		}
		if addNewFhirField(fhirResource, fhirFieldIndex) {
			fhirField = reflect.New(reflect.TypeOf(fhirResource.Interface()).Elem()).Elem()
			_, err := setField(fhirField, nestedField, value)
			if err != nil {
				return reflect.Value{}, err
			}
			fhirResource.Set(reflect.Append(fhirResource, fhirField))
		}
		if hasToPopulate(fhirResource, fhirFieldIndex) {
			for i := 0; i < fhirFieldIndex; i++ {
				placeHolder := reflect.New(reflect.TypeOf(fhirResource.Interface()).Elem()).Elem()
				fhirResource.Set(reflect.Append(fhirResource, placeHolder))
			}
			setField(fhirResource, field, value)
		}
	}

	return fhirResource, nil
}

func setFieldValue(fhirResource reflect.Value, fieldName string, value any) (reflect.Value, error) {
	var field reflect.Value
	fhirResourceField := camelToPascalCase(fieldName)
	if fhirResource.Kind() == reflect.Pointer {
		field = reflect.Indirect(fhirResource).FieldByName(fhirResourceField)
	} else {
		field = fhirResource.FieldByName(fhirResourceField)
	}
	if !field.IsValid() && !field.CanSet() {
		return reflect.Value{}, errors.New(fmt.Sprintf("unknown field %q", fhirResourceField))
	}

	field.Set(reflect.ValueOf(value))

	return fhirResource, nil
}

func camelToPascalCase(fieldName string) string {
	runes := []rune(fieldName)
	if runes[0] != '_' {
		runes[0] = unicode.ToUpper(runes[0])
	} else {
		runes[1] = unicode.ToUpper(runes[1])
	}
	return string(runes)
}

func containsFhirField(fhirResourceSlice reflect.Value, fhirFieldIndex int) bool {
	return fhirResourceSlice.Len() > fhirFieldIndex
}

func addNewFhirField(fhirResourceSlice reflect.Value, fhirFieldIndex int) bool {
	return fhirResourceSlice.Len() == fhirFieldIndex
}

func hasToPopulate(fhirResourceSlice reflect.Value, fhirFieldIndex int) bool {
	return fhirResourceSlice.Len() < fhirFieldIndex
}

func nilFieldPointer(fhirResource reflect.Value) bool {
	return fhirResource.Kind() == reflect.Pointer && fhirResource.IsNil()
}

func fieldPointer(fhirResource reflect.Value) bool {
	return fhirResource.Kind() == reflect.Pointer && !fhirResource.IsNil()
}

func emptySlice(fhirResource reflect.Value, fieldName string) bool {
	return fhirResource.Kind() != reflect.Pointer &&
		fhirResource.FieldByName(fieldName).IsValid() &&
		fhirResource.FieldByName(fieldName).Kind() != reflect.Struct &&
		fhirResource.FieldByName(fieldName).Interface() == nil
}
