package api

import (
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

func (t Traverser) Travers(ast []domain.ASTNode) domain.Bundle {
	var instances = map[string]reflect.Value{}

	r := domain.Bundle{
		ResourceType: "Bundle",
	}

	for _, node := range ast {
		fhirResource := fhirResourceInstance(node.NodeName, instances)
		setField(fhirResource, &node.FhirField, node.NodeValue)
	}

	for _, fhirResource := range instances {
		r.Entry = append(r.Entry, domain.BundleEntry{Resource: fhirResource.Interface()})
	}

	return r
}

func fhirResourceInstance(fhirResourceName string, instances map[string]reflect.Value) reflect.Value {
	if instances[fhirResourceName].IsValid() {
		return instances[fhirResourceName]
	}

	fhirResourceType := typeRegistry[fhirResourceName]
	instance := reflect.New(fhirResourceType).Elem()
	setFieldValue(instance, "resourceType", fhirResourceName)

	instances[fhirResourceName] = instance

	return instance
}

func setField(fhirResource reflect.Value, field *domain.FhirField, value any) reflect.Value {
	nestedField := field.FhirField

	if field.FieldParsedType == domain.MultipleValueField {
		fhirResource.Set(reflect.Append(fhirResource, reflect.ValueOf(value)))
		return fhirResource
	}
	if nestedField == nil {
		return setFieldValue(fhirResource, field.Name, value)
	}
	fieldName := camelToPascalCase(field.Name)

	if field.FieldParsedType == domain.SingleField {
		var fhirField reflect.Value

		if fieldPointer(fhirResource) {
			fhirField = reflect.Indirect(fhirResource).FieldByName(fieldName)
		}
		if emptySlice(fhirResource, fieldName) {
			fhirField = reflect.New(fhirResource.FieldByName(fieldName).Type()).Elem() //catch exception when field not exists
		}
		if !fhirField.IsValid() {
			fhirField = fhirResource.FieldByName(fieldName)
			if nilFieldPointer(fhirField) {
				fhirField = reflect.New(fhirField.Type().Elem())
			}
		}

		setField(fhirField, nestedField, value)
		setFieldValue(fhirResource, field.Name, fhirField.Interface())
	}
	if field.FieldParsedType == domain.MultipleNestedField {
		var fhirField reflect.Value
		fhirFieldIndex, _ := strconv.Atoi(field.Name)

		if hasFhirField(fhirResource, fhirFieldIndex) {
			fhirField = reflect.ValueOf(fhirResource.Interface()).Index(fhirFieldIndex)
			setField(fhirField, nestedField, value)
		} else {
			fhirField = reflect.New(reflect.TypeOf(fhirResource.Interface()).Elem()).Elem()
			setField(fhirField, nestedField, value)
			fhirResource.Set(reflect.Append(fhirResource, fhirField))
		}
	}

	return fhirResource
}

func setFieldValue(fhirResource reflect.Value, fieldName string, value any) reflect.Value {
	var field reflect.Value
	fhirResourceField := camelToPascalCase(fieldName)
	if fhirResource.Kind() == reflect.Pointer {
		field = reflect.Indirect(fhirResource).FieldByName(fhirResourceField)
	} else {
		field = fhirResource.FieldByName(fhirResourceField)
	}
	if field.IsValid() && field.CanSet() {
		field.Set(reflect.ValueOf(value))
	}

	return fhirResource
}

func camelToPascalCase(fieldName string) string {
	runes := []rune(fieldName)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func hasFhirField(fhirResourceSlice reflect.Value, fhirFieldIndex int) bool {
	return fhirResourceSlice.Len()-1 == fhirFieldIndex
}

func nilFieldPointer(fhirResource reflect.Value) bool {
	return fhirResource.Kind() == reflect.Pointer && fhirResource.IsNil()
}

func fieldPointer(fhirResource reflect.Value) bool {
	return fhirResource.Kind() == reflect.Pointer && !fhirResource.IsNil()
}

func emptySlice(fhirResource reflect.Value, fieldName string) bool {
	return fhirResource.Kind() != reflect.Pointer &&
		fhirResource.FieldByName(fieldName).Kind() != reflect.Struct && //catch error on fieldByName
		fhirResource.FieldByName(fieldName).Interface() == nil
}
