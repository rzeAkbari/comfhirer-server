package api

import (
	"github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/core/domain"
	fhir_r4 "github.com/rzeAkbari/comfhirer-server/comfhirer/internal/application/fhir/r4"
	"reflect"
	"unicode"
)

var typeRegistry = map[string]reflect.Type{
	"Patient": reflect.TypeOf(fhir_r4.Patient{}),
}

func Travers(ast []domain.ASTNode) fhir_r4.Bundle {
	r := fhir_r4.Bundle{
		ResourceType: "Bundle",
	}

	fhirResourceName := ast[0].NodeName
	fhirResource := fhirResourceInstance(fhirResourceName)

	setFieldValue(fhirResource, "resourceType", ast[0].NodeName)

	setField(fhirResource, &ast[0].FhirField, ast[0].NodeValue)

	r.Entry = append(r.Entry, fhir_r4.BundleEntry{Resource: fhirResource.Interface()})

	return r
}

func fhirResourceInstance(fhirResourceName string) reflect.Value {
	fhirResourceType := typeRegistry[fhirResourceName]

	return reflect.New(fhirResourceType).Elem()
}

func setField(fhirResource reflect.Value, field *domain.FhirField, value any) reflect.Value {
	nestedField := field.FhirField
	if nestedField == nil {
		return setFieldValue(fhirResource, field.Name, value)
	}
	fieldName := camelToPascalCase(field.Name)

	if field.FieldParsedType == domain.SingleField {
		fhirField := reflect.New(fhirResource.FieldByName(fieldName).Type()).Elem() //catch exception when field not exists
		setField(fhirField, nestedField, value)
		setFieldValue(fhirResource, field.Name, fhirField.Interface())
	}
	if field.FieldParsedType == domain.MultipleNestedField {
		fhirField := reflect.New(reflect.TypeOf(fhirResource.Interface()).Elem()).Elem()
		setField(fhirField, nestedField, value)
		fhirResource.Set(reflect.Append(fhirResource, fhirField))
	}

	return fhirResource
}

func setFieldValue(fhirResource reflect.Value, fieldName string, value any) reflect.Value {
	fhirResourceField := camelToPascalCase(fieldName)
	field := fhirResource.FieldByName(fhirResourceField)
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
