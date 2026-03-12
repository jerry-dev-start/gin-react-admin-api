package configs

type FileServer struct {
	Type *string `json:"type"`
}

type LocalFileServer struct {
	FilePath *string `mapstructure:"file-path"`
}
