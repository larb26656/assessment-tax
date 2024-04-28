package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAdminRepositoryCaseUserNotFound struct {
}

func (*mockAdminRepositoryCaseUserNotFound) FindUserByUsername(username string) *User {
	return nil
}

// Authenticate
func TestAuthenticate_ShouldReturnFalse_WhenUserNotFound(t *testing.T) {
	// Arrange
	usecase := NewAdminUsecase(&mockAdminRepositoryCaseUserNotFound{})

	// Act
	result := usecase.Authenticate("Admin", "Pass")

	// Assert
	assert.False(t, result)
}

type mockAdminRepositoryCasestruct struct {
}

func (*mockAdminRepositoryCasestruct) FindUserByUsername(username string) *User {
	return &User{
		Username: "adminTax",
		Password: "admin!",
	}
}

func TestAuthenticate_ShouldReturnFalse_WhenPasswordInvalid(t *testing.T) {
	// Arrange
	usecase := NewAdminUsecase(&mockAdminRepositoryCasestruct{})

	// Act
	result := usecase.Authenticate("adminTax", "Pass")

	// Assert
	assert.False(t, result)
}

func TestAuthenticate_ShouldReturnTrue_WhenInputCorrect(t *testing.T) {
	// Arrange
	usecase := NewAdminUsecase(&mockAdminRepositoryCasestruct{})

	// Act
	result := usecase.Authenticate("adminTax", "admin!")

	// Assert
	assert.True(t, result)
}
