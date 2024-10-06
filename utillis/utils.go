package utillis

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	UncorrectCerdentials = errors.New("Uncorrect cerdentials")
	AuthFailed           = errors.New("Authentication Failed")
)

func BasicAuthString(user, pass string) string {
	bstr := []byte(user + ":" + pass)
	return fmt.Sprint("Basic ", base64.StdEncoding.EncodeToString(bstr))
}

func CorrectCerdentials(user, pass string) error {
	if strings.Contains(user, " ") || strings.Contains(pass, " ") {
		return UncorrectCerdentials
	}
	return nil
}

func GetRequest(host *url.URL, auth, path string) ([]byte, error) {
	endpoint := host
	endpoint = endpoint.JoinPath(path)

	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		panic(err)
	}

	if req.Header == nil {
		req.Header = make(http.Header)
	}
	req.Header.Add("Authorization", auth)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, AuthFailed
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
