package domain

type Bundle struct {
	ResourceType string  `default:"bundle"`
	Entry        []Entry `json:"entry"`
}
type Entry struct {
	Resource FhirResource `json:"resource"`
}
type FhirResource struct {
	ResourceType string `json:"resourceType"`
}
