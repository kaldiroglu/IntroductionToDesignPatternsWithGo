package complexobject

import "dev.kaldiroglu/dp/intro/problems/complexobject/user"

// UserService coordinates several factories to build a User graph from a request. Because
// the factories are deliberately left as stubs (they return nil), this builds a nil graph -
// the "complex object construction" smell that motivates the Factory / Builder patterns.
type UserService struct{}

func (UserService) CreateUser(newUserRequest *NewUserRequest, username, userEntityID string) *user.User {
	userFactory := user.UserFactory{}
	var userRoleList []*user.UserRole
	var userCategories []*user.UserCategory
	var configurations []*user.Configuration
	u := userFactory.CreateUserFromRequest(newUserRequest, username, userEntityID,
		newUserRequest.RemoteAddr, newUserRequest.GetHeader("User-Agent"))

	roleFactory := user.RoleFactory{}
	if len(newUserRequest.RoleREntityIds) > 0 {
		for _, roleEntityID := range newUserRequest.RoleREntityIds {
			userRole := roleFactory.CreateUserRoleFromRequest(roleEntityID, username, userEntityID,
				newUserRequest.RemoteAddr, newUserRequest.GetHeader("User-Agent"))
			userRole.User = u
			userRoleList = append(userRoleList, userRole)
		}
		u.Roles = userRoleList
	}

	userCategoryFactory := user.UserCategoryFactory{}
	if len(newUserRequest.CategoryEntityIds) > 0 {
		for _, disciplineResourceID := range newUserRequest.CategoryEntityIds {
			userCategory := userCategoryFactory.CreateCatagoryFromRequest(disciplineResourceID, username, userEntityID,
				newUserRequest.RemoteAddr, newUserRequest.GetHeader("User-Agent"))
			userCategory.User = u
			userCategories = append(userCategories, userCategory)
		}
		u.Categories = userCategories
	}

	configurationFactory := user.ConfigurationFactory{}
	if len(newUserRequest.ConfigurationEntityIds) > 0 {
		for _, configurationEntityID := range newUserRequest.ConfigurationEntityIds {
			configuration := configurationFactory.CreateConfigurationFromRequest(configurationEntityID, username, userEntityID,
				newUserRequest.RemoteAddr, newUserRequest.GetHeader("User-Agent"))
			configuration.User = u
			configurations = append(configurations, configuration)
		}
		u.Configurations = configurations
	}

	return u
}
