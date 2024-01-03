package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	usr, err := NewUser("John", "test@test.com", "test123")
	assert.Nil(t, err)
	assert.NotNil(t, usr)
	assert.NotEmpty(t, usr.ID)
	assert.NotEmpty(t, usr.Name)
	assert.NotEmpty(t, usr.Email)
	assert.NotEmpty(t, usr.Password)
	assert.Equal(t, usr.Name, "John")
	assert.Equal(t, usr.Email, "test@test.com")

}

func TestNewUserPassword(t *testing.T) {
	usr, err := NewUser("John", "test@test.com", "test123")
	assert.Nil(t, err)
	assert.NotNil(t, usr)
	assert.NotEqual(t, usr.Password, "test1234")
	assert.True(t, usr.ValidatePassword("test123"))
	assert.False(t, usr.ValidatePassword("test1234"))
}
