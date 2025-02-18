package internal

import (
	"testing"

	"github.com/qkbyte/go-zero/core/stringx"
	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	endpoints := []string{
		"192.168.0.2:2379",
		"192.168.0.3:2379",
		"192.168.0.4:2379",
	}
	username := "foo" + stringx.Rand()
	password := "bar"
	anotherPassword := "any"

	_, ok := GetAccount(endpoints)
	assert.False(t, ok)

	AddAccount(endpoints, username, password)
	account, ok := GetAccount(endpoints)
	assert.True(t, ok)
	assert.Equal(t, username, account.User)
	assert.Equal(t, password, account.Pass)

	AddAccount(endpoints, username, anotherPassword)
	account, ok = GetAccount(endpoints)
	assert.True(t, ok)
	assert.Equal(t, username, account.User)
	assert.Equal(t, anotherPassword, account.Pass)
}
