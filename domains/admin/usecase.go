package admin

type AdminUsecase interface {
	Authenticate(username string, password string) bool
}

type adminUsecase struct {
	adminRepository AdminRepository
}

func NewAdminUsecase(adminRepository AdminRepository) AdminUsecase {
	return &adminUsecase{
		adminRepository: adminRepository,
	}
}

func (a *adminUsecase) Authenticate(username string, password string) bool {
	user := a.adminRepository.FindUserByUsername(username)

	if user == nil {
		return false
	}

	if password != user.Password {
		return false
	}

	return true
}
