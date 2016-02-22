package visa

import (
	"encoding/base64"
	"testing"
)

func TestSetConstants(t *testing.T) {
	cases := []struct {
		user_id, user_password string
	}{
		{"test_user_id", "dfjLRHzl4Mujn3aihy318OpySANPvO"},
	}

	for _, c := range cases {
		setVariables(c.user_id, c.user_password)
		if c.user_id != USER_ID && c.user_password != USER_PASSWORD {
			t.Errorf("Setting of constants does not pass. Looking for %s, got %s", c.user_id, USER_ID)
		}
	}
}

func TestURL(t *testing.T) {
	cases := []struct {
		prodEnv bool
	}{
		{true},
		{false},
	}

	for _, c := range cases {
		testUrl := ""
		getApiUrl(c.prodEnv)

		switch c.prodEnv {
		case true:
			testUrl = "https://api.visa.com"
			break
		case false:
			testUrl = "https://sandbox.api.visa.com"
			break
		}

		if API_URL != testUrl {
			t.Errorf("API url does not pass. Looking for %s, got %s", testUrl, API_URL)
		}
	}
}

func TestAuthHeader(t *testing.T) {
	cases := []struct {
		user_id, user_password string
	}{
		{"test_user_id", "dfjLRHzl4Mujn3aihy318OpySANPvO"},
	}

	for _, c := range cases {
		testBase64 := base64.StdEncoding.EncodeToString([]byte(c.user_id + ":" + c.user_password))

		setVariables(c.user_id, c.user_password)
		authHeaderTest := createAuthHeader()
		if authHeaderTest != testBase64 {
			t.Errorf("Auth header does not pass. Looking for %s, got %s", testBase64, authHeaderTest)
		}
	}
}
