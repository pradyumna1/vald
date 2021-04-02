package config

type Firestore struct {
	ProjectID           string `json:"projectID" yaml:"projectID"`
	CredentialsFilePath string `json:"credentials_file_path" yaml:"credentials_file_path"`
	CredentialsJSON     string `json:"credentials_json" yaml:"credentials_json"`
	CollectionPath      string `json:"collection_path" yaml:"collection_path"`
}

func (f *Firestore) Bind() *Firestore {
	return f
}
