package comfhirer_server

type Bundle struct {
	ResourceType string  `default:"bundle", json:"resourceType"`
	Entry        []Entry `json:"entry"`
}
type Entry struct {
	Resource FhirResource `json:"resource"`
}
type FhirResource struct {
	ResourceType string `json:"resourceType"`
}
