// Package complexobject ports the "complex object construction" smell: a UserService
// that assembles a User graph from a NewUserRequest through several stub factories.
package complexobject

type NewUserRequest struct {
	Username               string
	Status                 bool
	Email                  string
	Approved               bool
	ListName               string
	Name                   string
	LastName               string
	Organization           string
	RoleREntityIds         []string
	CategoryEntityIds      []string
	ConfigurationEntityIds []string
	Internal               bool
	RemoteAddr             string
}

// GetHeader mirrors the Java original, which ignored its field and returned the argument.
func (NewUserRequest) GetHeader(header string) string {
	return header
}

// SetHeader is a stub, as in the Java source.
func (NewUserRequest) SetHeader(header, value string) {
	// ...
}
