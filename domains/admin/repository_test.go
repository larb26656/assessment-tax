package admin

import (
	"testing"

	"github.com/larb26656/assessment-tax/mock"
	"github.com/stretchr/testify/assert"
)

var mockConfig = mock.NewMockAppConfig()

// FindUserByUsername
func TestFindUserByUsername_ShouldReturnNil_WhenUsernameInvalid(t *testing.T) {
	// Arrange
	repository := NewAdminRepository(&mockConfig)

	// Act
	user := repository.FindUserByUsername("User01")

	// Assert
	assert.Nil(t, user)
}

// FindUserByUsername
func TestFindUserByUsername_ShouldReturnUser_WhenUsernameValid(t *testing.T) {
	// Arrange
	repository := NewAdminRepository(&mockConfig)

	// Act
	user := repository.FindUserByUsername("adminTax")

	// Assert
	assert.NotNil(t, user)
	assert.Equal(t, user.Username, mockConfig.AdminUsername)
	assert.Equal(t, user.Password, mockConfig.AdminPassword)
}
