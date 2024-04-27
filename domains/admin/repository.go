package admin

import "github.com/larb26656/assessment-tax/config"

type AdminRepository interface {
	FindUserByUsername(username string) *User
}

type adminRepository struct {
	appConfig *config.AppConfig
	users     map[string]*User
}

func NewAdminRepository(appConfig *config.AppConfig) AdminRepository {
	// append users by config
	users := make(map[string]*User)

	users[appConfig.AdminUsername] = &User{
		Username: appConfig.AdminUsername,
		Password: appConfig.AdminPassword,
	}

	return &adminRepository{
		appConfig: appConfig,
		users:     users,
	}
}

func (r *adminRepository) FindUserByUsername(username string) *User {
	user := r.users[username]

	if user == nil {
		return nil
	}

	return user
}
