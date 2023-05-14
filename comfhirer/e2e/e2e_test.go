package e2e

import (
	"encoding/json"
	"github.com/rzeAkbari/comfhirer-server/comfhirer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndToEnd(t *testing.T) {

	testCases := []struct {
		name  string
		input map[string]any
		fhir  []byte
	}{
		{
			name: "patient",
			input: map[string]any{
				"Patient.(0).identifier.[0].use":                             "usual",
				"Patient.(0).identifier.[0].type.coding.[0].system":          "http://terminology.hl7.org/CodeSystem/v2-0203",
				"Patient.(0).identifier.[0].type.coding.[0].code":            "MR",
				"Patient.(0).identifier.[0].system":                          "urn:oid:1.2.36.146.595.217.0.1",
				"Patient.(0).identifier.[0].value":                           "12345",
				"Patient.(0).identifier.[0].period.start":                    "2001-05-06",
				"Patient.(0).identifier.[0].assigner.display":                "Acme Healthcare",
				"Patient.(0).identifier.[1].use":                             "official",
				"Patient.(0).identifier.[1].type.coding.[0].system":          "http://terminology.hl7.org/CodeSystem/v2-0206",
				"Patient.(0).identifier.[1].type.coding.[0].code":            "MR",
				"Patient.(0).identifier.[1].type.coding.[0].display":         "Peter",
				"Patient.(0).identifier.[1].system":                          "urn:oid:1.2.36.146.595.217.0.1",
				"Patient.(0).identifier.[1].value":                           "6789",
				"Patient.(0).identifier.[1].period.start":                    "2003-05-06",
				"Patient.(0).identifier.[1].assigner.display":                "Acme Healthcare one",
				"Patient.(0).active":                                         true,
				"Patient.(0).name.[0].use":                                   "official",
				"Patient.(0).name.[0].family":                                "Chalmers",
				"Patient.(0).name.[0].given.{0}":                             "Peter",
				"Patient.(0).name.[0].given.{1}":                             "James",
				"Patient.(0).name.[1].use":                                   "usual",
				"Patient.(0).name.[1].given.{0}":                             "Jim",
				"Patient.(0).name.[2].use":                                   "maiden",
				"Patient.(0).name.[2].family":                                "Windsor",
				"Patient.(0).name.[2].given.{0}":                             "Peter",
				"Patient.(0).name.[2].given.{1}":                             "James",
				"Patient.(0).name.[2].period.end":                            "2002",
				"Patient.(0).telecom.[0].use":                                "home",
				"Patient.(0).telecom.[1].use":                                "work",
				"Patient.(0).telecom.[1].value":                              "(03) 5555 6473",
				"Patient.(0).telecom.[1].system":                             "phone",
				"Patient.(0).telecom.[1].rank":                               1,
				"Patient.(0).telecom.[0].period.end":                         "2014",
				"Patient.(0).telecom.[2].use":                                "mobile",
				"Patient.(0).telecom.[2].value":                              "(03) 3410 5613",
				"Patient.(0).telecom.[2].system":                             "phone",
				"Patient.(0).telecom.[2].rank":                               2,
				"Patient.(0).telecom.[3].use":                                "old",
				"Patient.(0).telecom.[3].value":                              "(03) 5555 8834",
				"Patient.(0).telecom.[3].system":                             "phone",
				"Patient.(0).gender":                                         "male",
				"Patient.(0).birthDate":                                      "1974-12-25",
				"Patient.(0).ext_birthDate.extension.[0].url":                "http://hl7.org/fhir/StructureDefinition/Patient-birthTime",
				"Patient.(0).ext_birthDate.extension.[0].valueDateTime":      "1974-12-25T14:35:45-05:00",
				"Patient.(0).deceasedBoolean":                                false,
				"Patient.(0).address.[0].use":                                "home",
				"Patient.(0).address.[0].type":                               "both",
				"Patient.(0).address.[0].text":                               "534 Erewhon St PeasantVille, Rainbow, Vic  3999",
				"Patient.(0).address.[0].line.{0}":                           "534 Erewhon St",
				"Patient.(0).address.[0].line.{1}":                           "baker street",
				"Patient.(0).address.[0].city":                               "PleasantVille",
				"Patient.(0).address.[0].district":                           "Rainbow",
				"Patient.(0).address.[0].state":                              "Vic",
				"Patient.(0).address.[0].postalCode":                         "3999",
				"Patient.(0).address.[0].period.start":                       "1974-12-25",
				"Patient.(0).contact.[0].relationship.[0].coding.[0].system": "http://terminology.hl7.org/CodeSystem/v2-0131",
				"Patient.(0).contact.[0].relationship.[0].coding.[0].code":   "N",
				"Patient.(0).contact.[0].name.family":                        "du Marché",
				"Patient.(0).contact.[0].name.ext_family.extension.[0].url":  "http://hl7.org/fhir/StructureDefinition/humanname-own-prefix",
				"Patient.(0).contact.[0].name.given.{0}":                     "Bénédicte",
				"Patient.(0).contact.[0].telecom.[0].system":                 "phone",
				"Patient.(0).contact.[0].telecom.[0].value":                  "+33 (237) 998327",
				"Patient.(0).contact.[0].address.use":                        "home",
				"Patient.(0).contact.[0].address.type":                       "both",
				"Patient.(0).contact.[0].address.line.{0}":                   "534 Erewhon St",
				"Patient.(0).contact.[0].address.city":                       "PleasantVille",
				"Patient.(0).contact.[0].address.district":                   "Rainbow",
				"Patient.(0).contact.[0].address.state":                      "Vic",
				"Patient.(0).contact.[0].address.postalCode":                 "3999",
				"Patient.(0).contact.[0].address.period.start":               "1974-12-25",
				"Patient.(0).contact.[0].gender":                             "female",
				"Patient.(0).contact.[0].period.start":                       "2012",
				"Patient.(0).managingOrganization.reference":                 "Organization/1",
			},
			fhir: []byte("{\"resourceType\":\"Bundle\",\"entry\":[{\"resource\":{\"resourceType\":\"Patient\",\"active\":true,\"address\":[{\"city\":\"PleasantVille\",\"district\":\"Rainbow\",\"line\":[\"534 Erewhon St\",\"baker street\"],\"period\":{\"start\":\"1974-12-25\"},\"postalCode\":\"3999\",\"state\":\"Vic\",\"text\":\"534 Erewhon St PeasantVille, Rainbow, Vic  3999\",\"type\":\"both\",\"use\":\"home\"}],\"birthDate\":\"1974-12-25\",\"_birthDate\":{\"extension\":[{\"url\":\"http://hl7.org/fhir/StructureDefinition/Patient-birthTime\",\"valueDateTime\":\"1974-12-25T14:35:45-05:00\"}]},\"contact\":[{\"address\":{\"city\":\"PleasantVille\",\"district\":\"Rainbow\",\"line\":[\"534 Erewhon St\"],\"period\":{\"start\":\"1974-12-25\"},\"postalCode\":\"3999\",\"state\":\"Vic\",\"type\":\"both\",\"use\":\"home\"},\"gender\":\"female\",\"name\":{\"family\":\"du Marché\",\"_family\":{\"extension\":[{\"url\":\"http://hl7.org/fhir/StructureDefinition/humanname-own-prefix\"}]},\"given\":[\"Bénédicte\"]},\"period\":{\"start\":\"2012\"},\"relationship\":[{\"coding\":[{\"code\":\"N\",\"system\":\"http://terminology.hl7.org/CodeSystem/v2-0131\"}]}],\"telecom\":[{\"system\":\"phone\",\"value\":\"+33 (237) 998327\"}]}],\"gender\":\"male\",\"identifier\":[{\"assigner\":{\"display\":\"Acme Healthcare\"},\"period\":{\"start\":\"2001-05-06\"},\"system\":\"urn:oid:1.2.36.146.595.217.0.1\",\"type\":{\"coding\":[{\"code\":\"MR\",\"system\":\"http://terminology.hl7.org/CodeSystem/v2-0203\"}]},\"use\":\"usual\",\"value\":\"12345\"},{\"assigner\":{\"display\":\"Acme Healthcare one\"},\"period\":{\"start\":\"2003-05-06\"},\"system\":\"urn:oid:1.2.36.146.595.217.0.1\",\"type\":{\"coding\":[{\"code\":\"MR\",\"display\":\"Peter\",\"system\":\"http://terminology.hl7.org/CodeSystem/v2-0206\"}]},\"use\":\"official\",\"value\":\"6789\"}],\"managingOrganization\":{\"reference\":\"Organization/1\"},\"name\":[{\"family\":\"Chalmers\",\"given\":[\"Peter\",\"James\"],\"use\":\"official\"},{\"given\":[\"Jim\"],\"use\":\"usual\"},{\"family\":\"Windsor\",\"given\":[\"Peter\",\"James\"],\"period\":{\"end\":\"2002\"},\"use\":\"maiden\"}],\"telecom\":[{\"period\":{\"end\":\"2014\"},\"use\":\"home\"},{\"system\":\"phone\",\"use\":\"work\",\"value\":\"(03) 5555 6473\"},{\"system\":\"phone\",\"use\":\"mobile\",\"value\":\"(03) 3410 5613\"},{\"system\":\"phone\",\"use\":\"old\",\"value\":\"(03) 5555 8834\"}]}}]}"),
		},
		{
			name: "medication",
			input: map[string]any{
				"Medication.(0).identifier.[0].use":                           "usual",
				"Medication.(0).identifier.[0].type.coding.[0].system":        "http://terminology.hl7.org/CodeSystem/v2-0203",
				"Medication.(0).identifier.[0].type.coding.[0].code":          "MR",
				"Medication.(0).identifier.[0].system":                        "urn:oid:1.2.36.146.595.217.0.1",
				"Medication.(0).code.coding.[0].system":                       "http://snomed.info/sct",
				"Medication.(0).code.coding.[0].code":                         "260385009",
				"Medication.(0).code.coding.[1].system":                       "http://system.info/sct",
				"Medication.(0).code.coding.[1].display":                      "Negative",
				"Medication.(0).status":                                       "active",
				"Medication.(0).manufacturer.reference":                       "organization/1",
				"Medication.(0).manufacturer.type":                            "patient",
				"Medication.(0).manufacturer.identifier.system":               "http://snomed.info/manufacturer",
				"Medication.(0).manufacturer.identifier.value":                "manufacturer-code",
				"Medication.(0).manufacturer.identifier.use":                  "manufacturer-code",
				"Medication.(0).manufacturer.identifier.type.coding.[0].code": "manufacturer-code",
				"Medication.(0).manufacturer.identifier.period.start":         "2023-12-01",
				"Medication.(0).manufacturer.identifier.assigner.reference":   "assigner/1",
				"Medication.(0).manufacturer.display":                         "manufacturer",
				"Medication.(0).form.coding.[0].system":                       "http://snomed.info/sct",
				"Medication.(0).form.coding.[0].code":                         "430127000",
				"Medication.(0).form.coding.[0].display":                      "Oral Form Oxycodone (product)",
				"Medication.(0).ingredient.[0].itemReference.reference":       "ingredient/sub03",
				"Medication.(0).ingredient.[0].strength.numerator.value":      5,
				"Medication.(0).ingredient.[0].strength.numerator.system":     "http://unitsofmeasure.org",
				"Medication.(0).ingredient.[0].strength.numerator.code":       "mg",
				"Medication.(0).ingredient.[0].strength.denominator.value":    1,
				"Medication.(0).ingredient.[0].strength.denominator.system":   "http://terminology.hl7.org/CodeSystem/v3-orderableDrugForm",
				"Medication.(0).ingredient.[0].strength.denominator.code":     "TAB",
				"Medication.(0).batch.lotNumber":                              "9494788",
				"Medication.(0).ingredient.[0].isActive":                      true,
				"Medication.(0).batch.expirationDate":                         "2017-05-22",
				"Medication.(0).amount.numerator.value":                       1,
				"Medication.(0).amount.numerator.system":                      "http://unitsofmeasure.org",
				"Medication.(0).amount.numerator.code":                        "mg",
				"Medication.(0).amount.denominator.value":                     2,
				"Medication.(0).amount.denominator.system":                    "http://unitsofmeasure.org",
				"Medication.(0).amount.denominator.code":                      "mg",
			},
			fhir: []byte("{\"resourceType\":\"Bundle\",\"entry\":[{\"resource\":{\"resourceType\":\"Medication\",\"amount\":{\"denominator\":{\"code\":\"mg\",\"system\":\"http://unitsofmeasure.org\",\"value\":2},\"numerator\":{\"code\":\"mg\",\"system\":\"http://unitsofmeasure.org\",\"value\":1}},\"batch\":{\"expirationDate\":\"2017-05-22\",\"lotNumber\":\"9494788\"},\"code\":{\"coding\":[{\"code\":\"260385009\",\"system\":\"http://snomed.info/sct\"},{\"display\":\"Negative\",\"system\":\"http://system.info/sct\"}]},\"form\":{\"coding\":[{\"code\":\"430127000\",\"display\":\"Oral Form Oxycodone (product)\",\"system\":\"http://snomed.info/sct\"}]},\"identifier\":[{\"system\":\"urn:oid:1.2.36.146.595.217.0.1\",\"type\":{\"coding\":[{\"code\":\"MR\",\"system\":\"http://terminology.hl7.org/CodeSystem/v2-0203\"}]},\"use\":\"usual\"}],\"ingredient\":[{\"isActive\":true,\"item\":{},\"itemReference\":{\"reference\":\"ingredient/sub03\"},\"strength\":{\"denominator\":{\"code\":\"TAB\",\"system\":\"http://terminology.hl7.org/CodeSystem/v3-orderableDrugForm\",\"value\":1},\"numerator\":{\"code\":\"mg\",\"system\":\"http://unitsofmeasure.org\",\"value\":5}}}],\"manufacturer\":{\"display\":\"manufacturer\",\"identifier\":{\"assigner\":{\"reference\":\"assigner/1\"},\"period\":{\"start\":\"2023-12-01\"},\"system\":\"http://snomed.info/manufacturer\",\"type\":{\"coding\":[{\"code\":\"manufacturer-code\"}]},\"use\":\"manufacturer-code\",\"value\":\"manufacturer-code\"},\"reference\":\"organization/1\",\"type\":\"patient\"},\"status\":\"active\"}}]}"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := comfhirer.Run(tc.input)
			if len(err) > 0 {
				t.Fatal("got error")
			}
			var gotFhir Bundle
			var wantFhir Bundle

			json.Unmarshal(got, &gotFhir)
			json.Unmarshal(tc.fhir, &wantFhir)

			assert.JSONEq(t, string(tc.fhir), string(got))
		})
	}
}
