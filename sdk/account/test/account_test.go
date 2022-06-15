package test

import (
	"github.com/stretchr/testify/assert"
	accountCRUD "interview-accountapi/account"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("BEGIN e2e testing")
	result := m.Run()
	log.Println("END e2e testing")
	os.Exit(result)
}

func TestListAccountsEmpty(t *testing.T) {
	t.Skip("not required, just for fun")
	result, _ := accountCRUD.List()
	assert.Empty(t, result)
}

func TestListAccounts(t *testing.T) {
	t.Skip("not required, just for fun")
	result, _ := accountCRUD.List()
	assert.NotEmpty(t, result)
}

func TestCreateAccount(t *testing.T) {
	expectedAccount := accountCRUD.MakeAccount()
	actualAccount, err := accountCRUD.Create(expectedAccount)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, *actualAccount)
}

func TestCreateAccountFail(t *testing.T) {
	expectedAccount := accountCRUD.MakeAccount()
	expectedAccount.Attributes = nil
	actualAccount, err := accountCRUD.Create(expectedAccount)

	assert.Error(t, err)
	assert.Nil(t, actualAccount)
}

func TestDeleteAccount(t *testing.T) {
	account, err := accountCRUD.Create(accountCRUD.MakeAccount())

	assert.NoError(t, err)

	result, err := accountCRUD.Delete(account.ID, *account.Version)

	assert.NoError(t, err)
	assert.True(t, result)
}

func TestDeleteAccountFail(t *testing.T) {
	account := accountCRUD.MakeAccount()
	result, err := accountCRUD.Delete(account.ID, *account.Version)

	assert.Error(t, err)
	assert.False(t, result)
}

func TestFetchAccount(t *testing.T) {
	expectedAccount, err := accountCRUD.Create(accountCRUD.MakeAccount())

	assert.NoError(t, err)

	actualAccount, err := accountCRUD.Fetch(expectedAccount.ID)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, actualAccount)
}

func TestFetchAccountFail(t *testing.T) {
	account := accountCRUD.MakeAccount()
	result, err := accountCRUD.Fetch(account.ID)

	assert.Error(t, err)
	assert.Nil(t, result)
}
