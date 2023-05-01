# Comfhirer

Comfhirer compiles flat, high level json key values to Fhir json.
Being a compiler it has its own syntactical rules and regulations.<br/>
## Use Case
Assume we have some information about a patient extracted from a 
service which needs to be converted to a <a href="https://hl7.org/fhir/R4/patient.html">Fhir Patient Resource</a>.
Having extracted data assign to Comfhirer syntax, it will generated the Fhir Patient json.

```json
{
  "Patient.identifier.[0].use":"usual",
  "Patient.identifier.[0].type.coding.[0].system":"http://terminology.hl7.org/CodeSystem/v2-0203",
  "Patient.identifier.[0].type.coding.[0].code":"MR",
  "Patient.identifier.[0].system":"urn:oid:1.2.36.146.595.217.0.1",
  "Patient.identifier.[0].value":"12345",
  "Patient.identifier.[0].period.start":"2001-05-06",
  "Patient.identifier.[0].assigner.display":"Acme Healthcare",
  "Patient.name.[0].use":"official",
  "Patient.name.[0].family":"Chalmers",
  "Patient.name.[0].given.{0}":"Peter"
}
```
```json
{
  "resourceType": "Patient",
  "identifier": [
    {
      "use": "usual",
      "type": {
        "coding": [
          {
            "system": "http://terminology.hl7.org/CodeSystem/v2-0203",
            "code": "MR"
          }
        ]
      },
      "system": "urn:oid:1.2.36.146.595.217.0.1",
      "value": "12345",
      "period": {
        "start": "2001-05-06"
      },
      "assigner": {
        "display": "Acme Healthcare"
      }
    }
  ],
  "name": [
    {
      "use": "official",
      "family": "Chalmers",
      "given": ["Peter"]
    }
  ]
}

```
## Syntactical Regulations
- Each Key starts with associated Fhir resource name having the first letter in upper case.
- Special Characters:
  - [] : square brackets are used when an associated Fhir field accepts json array.
    ```json
    {"Patient.identifier.[0].use":"usual"}
    ```
  - {} : graph brackets are used when an associated Fhir field accepts multiple values.
    ```json
    {"Patient.name.[0].given.{0}":"Peter"}
    ```
  - () : parenthesis are used for indexing Fhir resources of the same type.
    ```json
    {
      "Patient.(0).active":"false",
      "Patient.(1).active":"true"
    }
    ```
- Before Each special character a pointer must be put. 