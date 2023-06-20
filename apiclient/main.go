package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthResponse struct {
	Issuer       string `json:"issuer"`
	RefreshToken string `json:"refreshToken"`
	User         string `json:"user"`
	Email        string `json:"mail"`
	// ReplacementRT string `json:"x-cpcli-replacementrefreshtoken"`
}

type APIClient struct {
	BaseURL      string
	RefreshToken string
	Subdomain    string
	// ReplacementRT string
}

func main() {
	apiClient := APIClient{
		BaseURL:   "https://cpcli.cf.eu10.hana.ondemand.com", // Replace with your API base URL
		Subdomain: "27eadf16trial-ga",
	}

	err := apiClient.Login("P2006255217", "Carlos123*")
	if err != nil {
		panic(err)
	}

	err = apiClient.SendCommand()
	if err != nil {
		fmt.Println(string(err.Error()))
		panic(err)
	}
}

func (c *APIClient) Login(username, password string) error {
	loginURL := c.BaseURL + "/login/v2.29.0"
	loginBody := map[string]string{
		"userName":  username,
		"password":  password,
		"subdomain": c.Subdomain,
	}
	loginJSON, _ := json.Marshal(loginBody)

	// fmt.Println(string(loginJSON))
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
	// c.ReplacementRT = authResponse.ReplacementRT
	// fmt.Println(string(c.RefreshToken))

	return nil
}

func (c *APIClient) SendCommand() error {
	commandURL := c.BaseURL + "/command/v2.29.0/accounts/subaccount?list"
	commandBody := map[string]interface{}{
		"paramValues": map[string]string{
			"globalAccount": "27eadf16trial-ga",
		},
	}
	commandJSON, _ := json.Marshal(commandBody)

	req, err := http.NewRequest("POST", commandURL, bytes.NewBuffer(commandJSON))
	if err != nil {
		fmt.Println(string(err.Error()))
		return err
	}
	// fmt.Println(req)
	// fmt.Println(string(commandURL))
	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-cpcli-refreshtoken", c.RefreshToken)
	req.Header.Add("x-cpcli-subdomain", c.Subdomain)
	req.Header.Add("x-cpcli-format", "json")
	// fmt.Println(string(commandJSON))
	// fmt.Println(string(c.Subdomain))
	// fmt.Println(string(c.RefreshToken))
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(string(err.Error()))
		return err
	}
	// fmt.Println(resp.Header)
	// fmt.Println(resp.StatusCode)
	defer resp.Body.Close()

	// c.ReplacementRT = resp.Header.Get("x-cpcli-replacementrefreshtoken")

	// Process the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(err.Error()))
		return err
	}

	fmt.Println(string(responseBody))

	return nil
}
