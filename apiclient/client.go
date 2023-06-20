package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SubaccountInfo struct {
	SubacccountID     string `json:"guid"`
	TechnicalName     string `json:"technicalName"`
	DisplayName       string `json:"displayName"`
	GlobalAccountGUID string `json:"globalAccountGUID"`
	ParentGUID        string `json:"parentGUID"`
	ParentType        string `json:"parentType"`
	Region            string `json:"region"`
	Subdomain         string `json:"subdomain"`
	BetaEnabled       bool   `json:"betaEnabled"`
	UsedForProduction string `json:"usedForProduction"`
	State             string `json:"state"`
	StateMessage      string `json:"stateMessage"`
	CreatedDate       string `json:"createdDate"`
	ModifiedDate      string `json:"modifiedDate"`
}

type AuthResponse struct {
	Issuer       string `json:"issuer"`
	RefreshToken string `json:"refreshToken"`
	User         string `json:"user"`
	Email        string `json:"mail"`
}

type APIClient struct {
	BaseURL       string
	RefreshToken  string
	Username      string
	Password      string
	GlobalAccount string
	ReplacementRT string
}

func (c *APIClient) Login(s *SubaccountInfo) error {
	loginURL := c.BaseURL + "/login/v2.29.0"
	loginBody := map[string]string{
		"userName":  c.Username,
		"password":  c.Password,
		"subdomain": c.GlobalAccount,
	}
	loginJSON, _ := json.Marshal(loginBody)

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	authResponse := AuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return err
	}

	c.RefreshToken = authResponse.RefreshToken

	return nil
}

func (c *APIClient) ListSubaccountCommand(s *SubaccountInfo) (string, error) {
	commandURL := c.BaseURL + "/command/v2.29.0/accounts/subaccount?list"
	commandBody := map[string]interface{}{
		"paramValues": map[string]string{
			"globalAccount": c.GlobalAccount,
		},
	}
	commandJSON, _ := json.Marshal(commandBody)
	if c.ReplacementRT != "" {
		c.RefreshToken = c.ReplacementRT
	}
	req, err := http.NewRequest("POST", commandURL, bytes.NewBuffer(commandJSON))
	if err != nil {
		fmt.Println(string(err.Error()))
		return "", err
	}
	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-cpcli-refreshtoken", c.RefreshToken)
	req.Header.Add("x-cpcli-subdomain", c.GlobalAccount)
	req.Header.Add("x-cpcli-format", "json")
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(string(err.Error()))
		return "", err
	}
	defer resp.Body.Close()

	c.ReplacementRT = resp.Header.Get("x-cpcli-replacementrefreshtoken")
	if c.ReplacementRT == "" {
		c.Login(s)
	}
	fmt.Println(c.ReplacementRT)
	// Process the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(err.Error()))
		return "", err
	}

	fmt.Println(string(responseBody))

	return string(responseBody), nil
}

func (c *APIClient) GetSubaccountCommand(s *SubaccountInfo) (*SubaccountInfo, error) {
	commandURL := c.BaseURL + "/command/v2.29.0/accounts/subaccount?get"
	commandBody := map[string]interface{}{
		"paramValues": map[string]string{
			"globalAccount": c.GlobalAccount,
			"subaccount":    s.SubacccountID,
		},
	}
	commandJSON, _ := json.Marshal(commandBody)
	if c.ReplacementRT != "" {
		c.RefreshToken = c.ReplacementRT
	}
	req, err := http.NewRequest("POST", commandURL, bytes.NewBuffer(commandJSON))
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-cpcli-refreshtoken", c.RefreshToken)
	req.Header.Add("x-cpcli-subdomain", c.GlobalAccount)
	req.Header.Add("x-cpcli-format", "json")
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	c.ReplacementRT = resp.Header.Get("x-cpcli-replacementrefreshtoken")
	if c.ReplacementRT == "" {
		c.Login(s)
	}
	fmt.Println(c.ReplacementRT)
	// Process the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	f := SubaccountInfo{}
	fmt.Println(f)
	err = json.Unmarshal([]byte(responseBody), &f)
	if err != nil {
		return nil, err
	}
	fmt.Println("hey2")
	fmt.Println(f)
	fmt.Println(string(responseBody))
	fmt.Println("Region:", f.Region)
	return &f, nil
}
func (c *APIClient) CreateSubaccountCommand(s *SubaccountInfo) (*SubaccountInfo, error) {
	commandURL := c.BaseURL + "/command/v2.29.0/accounts/subaccount?create"
	commandBody := map[string]interface{}{
		"paramValues": map[string]string{
			"displayName": s.DisplayName,
			"region":      s.Region,
			"subdomain":   s.Subdomain,
		},
	}
	commandJSON, _ := json.Marshal(commandBody)
	if c.ReplacementRT != "" {
		c.RefreshToken = c.ReplacementRT
	}
	req, err := http.NewRequest("POST", commandURL, bytes.NewBuffer(commandJSON))
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-cpcli-refreshtoken", c.RefreshToken)
	req.Header.Add("x-cpcli-subdomain", c.GlobalAccount)
	req.Header.Add("x-cpcli-format", "json")
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	c.ReplacementRT = resp.Header.Get("x-cpcli-replacementrefreshtoken")
	if c.ReplacementRT == "" {
		c.Login(s)
	}
	fmt.Println(c.ReplacementRT)
	// Process the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	f := SubaccountInfo{}
	fmt.Println(f)
	err = json.Unmarshal([]byte(responseBody), &f)
	if err != nil {
		return nil, err
	}
	// fmt.Println("hey2")
	// fmt.Println(f)
	fmt.Println(string(responseBody))
	// fmt.Println("Region:", f.Region)
	return &f, nil
}
func (c *APIClient) UpdateSubaccountCommand(s *SubaccountInfo) (*SubaccountInfo, error) {
	commandURL := c.BaseURL + "/command/v2.29.0/accounts/subaccount?update"
	commandBody := map[string]interface{}{
		"paramValues": map[string]string{
			"displayName":   s.DisplayName,
			"subaccount":    s.SubacccountID,
			"globalAccount": c.GlobalAccount,
		},
	}
	commandJSON, _ := json.Marshal(commandBody)
	if c.ReplacementRT != "" {
		c.RefreshToken = c.ReplacementRT
	}
	req, err := http.NewRequest("POST", commandURL, bytes.NewBuffer(commandJSON))
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-cpcli-refreshtoken", c.RefreshToken)
	req.Header.Add("x-cpcli-subdomain", c.GlobalAccount)
	req.Header.Add("x-cpcli-format", "json")
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	c.ReplacementRT = resp.Header.Get("x-cpcli-replacementrefreshtoken")
	if c.ReplacementRT == "" {
		c.Login(s)
	}
	fmt.Println(c.ReplacementRT)
	// Process the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	f := SubaccountInfo{}
	fmt.Println(f)
	err = json.Unmarshal([]byte(responseBody), &f)
	if err != nil {
		return nil, err
	}
	// fmt.Println("hey2")
	// fmt.Println(f)
	fmt.Println(string(responseBody))
	// fmt.Println("Region:", f.Region)
	return &f, nil
}
func (c *APIClient) DeleteSubaccountCommand(s *SubaccountInfo) (*SubaccountInfo, error) {
	commandURL := c.BaseURL + "/command/v2.29.0/accounts/subaccount?delete"
	commandBody := map[string]interface{}{
		"paramValues": map[string]string{
			"globalAccount": c.GlobalAccount,
			"subaccount":    s.SubacccountID,
		},
	}
	commandJSON, _ := json.Marshal(commandBody)
	if c.ReplacementRT != "" {
		c.RefreshToken = c.ReplacementRT
	}
	req, err := http.NewRequest("POST", commandURL, bytes.NewBuffer(commandJSON))
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-cpcli-refreshtoken", c.RefreshToken)
	req.Header.Add("x-cpcli-subdomain", c.GlobalAccount)
	req.Header.Add("x-cpcli-format", "json")
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	c.ReplacementRT = resp.Header.Get("x-cpcli-replacementrefreshtoken")
	if c.ReplacementRT == "" {
		c.Login(s)
	}
	fmt.Println(c.ReplacementRT)
	// Process the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	f := SubaccountInfo{}
	fmt.Println(f)
	err = json.Unmarshal([]byte(responseBody), &f)
	if err != nil {
		return nil, err
	}
	// fmt.Println("hey2")
	// fmt.Println(f)
	fmt.Println(string(responseBody))
	// fmt.Println("Region:", f.Region)
	return &f, nil
}
