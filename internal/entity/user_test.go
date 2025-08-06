package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	name := "usuario"
	email := "email@email.com"
	password := "senha"
	user, err := NewUser(name, email, password)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
}

func TestValidatePassword(t *testing.T) {
	name := "usuario"
	email := "email@email.com"
	password := "senha"
	user, err := NewUser(name, email, password)
	assert.Nil(t, err)
	assert.NotEqual(t, user.Password, password)
	assert.True(t, user.ValidadePassword(password))
	assert.False(t, user.ValidadePassword("password"))
}
