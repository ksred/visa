package visa

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_PROD = "https://api.visa.com"
const API_SANDBOX = "https://sandbox.api.visa.com"
const API_PORT = 443

var USER_ID = ""
var USER_PASSWORD = ""

var API_URL = API_SANDBOX // Default to sandbox

// FUND TRANSFER API
// - Pull funds
// - Push funds
// - Refund funds
// mVISA API
// Watch list screening API

func setVariables(user_id string, user_password string) {
	USER_ID = user_id
	USER_PASSWORD = user_password
}

func getApiUrl(production bool) {
	switch production {
	case false:
		API_URL = API_SANDBOX
		break
	case true:
		API_URL = API_PROD
		break
	}
}

func Client(userId string, userPassword string, url string, reqType string, production bool, body []byte, transactionID string) (response []byte, err error) {
	setVariables(userId, userPassword)
	authHeader := createAuthHeader()

	//fmt.Printf("Body: %v\n", bytes.NewBuffer(body))
	req, err := http.NewRequest(reqType, url, bytes.NewBuffer(body))
	req.Header.Set("X-Client-Transaction-ID", transactionID)
	req.Header.Set("Authorization:Basic ", authHeader)
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Accept", "application/json,application/octet-stream")
	req.Header.Set("Accept", "application/json")

	// Load client cert
	cert, err := tls.LoadX509KeyPair(SSL_PUBLIC_KEY_PATH, SSL_PRIVATE_KEY_PATH)
	if err != nil {
		return nil, fmt.Errorf("Could not load key pair: %v", err)
	}
	// Load CA cert
	caCert, err := ioutil.ReadFile(SSL_CAPRIVATE_KEY_PATH)
	if err != nil {
		return nil, fmt.Errorf("Could not load CA key: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true, //@FIXME: This call *must* be secure
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not load HTTPS client: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		response, _ = ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP error: %s, %v", resp.Status, string(response))
	}

	response, _ = ioutil.ReadAll(resp.Body)
	fmt.Printf("Response: %v\n", string(response))
	return response, nil
}

func createAuthHeader() (authHeader string) {
	// Auth header = \ase64(userid:user_password)
	authHeader = base64.StdEncoding.EncodeToString([]byte(USER_ID + ":" + USER_PASSWORD))
	return
}

// PushFundsTransactions or MultiPushFundsTransactions POST operation credits pushes) funds to a recipient's Visa account or multiple Visa accounts, respectively.
func PushFundsTransactionsPost() {
}

// PushFundsTransactions or MultiPushFundsTransactions POST operation credits pushes) funds to a recipient's Visa account or multiple Visa accounts, respectively.
func MultiPushFundsTransactionsPost() {
}

// ReverseFundsTransactions POST operation credits (pushes back) funds to the sender's Visa account
// (for example, when authorization for the PushFundsTransactions POST operation is declined)
func ReverseFundsTransactionsPost() {
}
