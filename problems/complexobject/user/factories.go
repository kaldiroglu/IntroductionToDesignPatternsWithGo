package user

// Faithful ports of the Java factory stubs: each is intentionally unimplemented and
// returns nil - the "before" picture the course refactors.

type UserFactory struct{}

// CreateUserFromRequest takes the request as `any`: Go forbids the import cycle that
// would arise from importing the parent complexobject package (which imports this one),
// and the stub ignores the request anyway.
func (UserFactory) CreateUserFromRequest(userRequest any, username, userResourceID,
	remoteAddr, userAgent string) *User {
	var user *User
	// ...
	return user
}

type RoleFactory struct{}

func (RoleFactory) CreateUserRoleFromRequest(roleEntityID, username, userEntityID,
	remoteAddr, userAgent string) *UserRole {
	var userRole *UserRole
	// ...
	return userRole
}

type UserCategoryFactory struct{}

// CreateCatagoryFromRequest keeps the Java original's spelling ("catagory") so the port
// mirrors the source repo exactly; UserService calls it under that name.
func (UserCategoryFactory) CreateCatagoryFromRequest(disciplineResourceID, username,
	userEntityID, remoteAddr, header string) *UserCategory {
	var userCategory *UserCategory
	// ...
	return userCategory
}

type ConfigurationFactory struct{}

func (ConfigurationFactory) CreateConfigurationFromRequest(configurationEntityID, username,
	userEntityID, remoteAddr, header string) *Configuration {
	var configuration *Configuration
	//
	return configuration
}
