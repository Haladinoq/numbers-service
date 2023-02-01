package rest

// Validator interface.
type Validator interface {
	// IsValid method is used as pre validations in order to enable defaults.
	IsValid() error
}
