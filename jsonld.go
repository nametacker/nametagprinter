package nametagprinter

type JsonLDTypedModel struct {
	JsonLDContext string `json:"@context,omitempty"`
	JsonLDId      string `json:"@id,omitempty"`
	JsonLDType    string `json:"@type,omitempty"`
}
