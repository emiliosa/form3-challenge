package account

import (
	"fmt"
	"github.com/google/uuid"
	"go/types"
	internalClient "interview-accountapi/internal/client"
	"net/http"
	"strconv"
)

type Account struct {
	Attributes     *Attributes `json:"attributes,omitempty"`
	ID             string      `json:"id,omitempty"`
	OrganisationID string      `json:"organisation_id,omitempty"`
	Type           string      `json:"type,omitempty"`
	Version        *int64      `json:"version,omitempty"`
}

type Attributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

type Response struct {
	Data  Account     `json:"data"`
	Links types.Slice `json:"links"`
}

type ResponseArray struct {
	Data  []Account   `json:"data"`
	Links types.Slice `json:"links"`
}

type Request struct {
	Data Account `json:"data"`
}

func Fetch(uuid string) (*Account, error) {
	var client = internalClient.New("organisation/accounts")
	var accountResponse Response
	request, err := client.Request(http.MethodGet, "/"+uuid, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response, err := client.Do(request, &accountResponse)

	if err != nil {
		err = fmt.Errorf("error on Fetch Account [Response: %v, Error: %v]", response, err)
		return nil, err
	}

	return &accountResponse.Data, nil
}

// Deprecated: not required, just for fun
func List() ([]Account, error) {
	var client = internalClient.New("organisation/accounts")
	var accounts ResponseArray
	request, err := client.Request(http.MethodGet, "", nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = client.Do(request, &accounts)
	return accounts.Data, nil
}

func Create(data Account) (*Account, error) {
	var client = internalClient.New("organisation/accounts")
	var accountResponse Response
	var accountRequest = Request{Data: data}
	request, err := client.Request(http.MethodPost, "", accountRequest)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response, err := client.Do(request, &accountResponse)

	if err != nil {
		err = fmt.Errorf("error on Create Account [Response: %v, Error: %v]", response, err)
		return nil, err
	}

	return &accountResponse.Data, nil
}

func Delete(uuid string, version int64) (bool, error) {
	var client = internalClient.New("organisation/accounts")
	versionStr := strconv.FormatInt(version, 10)
	request, err := client.Request(http.MethodDelete, "/"+uuid+"?version="+versionStr, nil)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	response, err := client.Do(request, nil)

	if err != nil {
		err = fmt.Errorf("error on Delete Account [Response: %v, Error: %v]", response, err)
		return false, err
	}

	return true, nil
}

func MakeAccount() Account {
	country := "GB"
	accountClassification := "Personal"
	version := int64(0)
	accountAttributes := Attributes{
		AccountClassification: &accountClassification,
		BankID:                "TEST",
		BankIDCode:            "TEST",
		BaseCurrency:          "GBP",
		Bic:                   "NWBKGB42",
		Country:               &country,
		Name:                  []string{"Emiliano Abarca"},
	}
	return Account{
		Attributes:     &accountAttributes,
		ID:             uuid.New().String(),
		OrganisationID: uuid.New().String(),
		Type:           "accounts",
		Version:        &version,
	}
}
