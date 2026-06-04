// Package user holds the entity graph the factories assemble. Java JavaBeans become Go
// structs with exported fields. (Grouped into entities.go + factories.go, the idiomatic
// Go equivalent of Java's one-class-per-file package.)
package user

type User struct {
	ID             int64
	EntityID       string
	Deleted        bool
	Username       string
	Password       string
	Email          string
	Status         bool
	Approved       bool
	ListName       string
	Name           string
	LastName       string
	Organization   string
	Roles          []*UserRole
	Categories     []*UserCategory
	Configurations []*Configuration
	Internal       bool
}

type Role struct {
	ResourceID string
	Deleted    bool
	Name       string
	UserRoles  []*UserRole
	Privileges []*RolePrivilege
}

type Category struct {
	ID             int64
	EntityID       string
	Deleted        bool
	Name           string
	UserCategories []*UserCategory
}

type Privilege struct {
	ID             int64
	EntityID       string
	Deleted        bool
	Name           string
	RolePrivileges []*RolePrivilege
}

type Configuration struct {
	ID       int64
	EntityID string
	Deleted  bool
	User     *User
}

type UserRole struct {
	ID       int64
	EntityID string
	Deleted  bool
	Role     *Role
	User     *User
}

type UserCategory struct {
	ID       int64
	EntityID string
	Deleted  bool
	Category *Category
	User     *User
}

type RolePrivilege struct {
	ID        int64
	EntityID  string
	Deleted   bool
	Privilege *Privilege
	Role      *Role
}
