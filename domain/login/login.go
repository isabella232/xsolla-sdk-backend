package domainlogin

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"xsolla-sdk-backend/internal/app"
	"xsolla-sdk-backend/internal/store"
)

type LoginDomain struct {
	a *app.Application
}

func NewLoginDomain(a *app.Application) LoginDomain {
	return LoginDomain{
		a: a,
	}
}

func (l *LoginDomain) LoginUser(email fmt.Stringer) (store.UserItem, error) {
	url := fmt.Sprintf("https://api.xsolla.com/merchant/v2/merchants/%d/token", l.a.Config.MerchantID)
	requestBody, err := json.Marshal(map[string]interface{}{
		"user": map[string]interface{}{
			"email": map[string]string{
				"value": email.String(),
			},
			"id": map[string]string{
				"value": "user" + email.String(),
			},
		},
		"settings": map[string]interface{}{
			"project_id": l.a.Config.ProjectID,
		},
	})
	if err != nil {
		return store.UserItem{}, err
	}

	timeout := 5 * time.Second

	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return store.UserItem{}, err
	}

	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d:%s", l.a.Config.MerchantID, l.a.Config.PublisherAPIKey)))

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	req.Header.Add("Content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return store.UserItem{}, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return store.UserItem{}, err
	}
	accessTokenInterface, ok := result["token"]
	if !ok {
		return store.UserItem{}, fmt.Errorf("can not parse result with token. Response body: %s", result)
	}
	accessToken, ok := accessTokenInterface.(string)
	if !ok {
		return store.UserItem{}, fmt.Errorf("failed type assertion. Waited: string, got: %v", accessTokenInterface)
	}

	return store.UserItem{Email: email.String(), AccessToken: accessToken}, nil
}
