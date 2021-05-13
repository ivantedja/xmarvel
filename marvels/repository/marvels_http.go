package repository

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/google/uuid"

	entity "github.com/ivantedja/xmarvel/entity"
)

var (
	HttpClient = http.DefaultClient
)

type api struct {
	client     *http.Client
	baseUrl    string
	publicKey  string
	privateKey string
}

func NewAPI(baseUrl, publicKey, privateKey string, timeout time.Duration) *api {
	httpClient := HttpClient
	httpClient.Timeout = timeout
	return &api{
		client:     httpClient,
		baseUrl:    baseUrl,
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (a *api) Search(ctx context.Context, filter map[string]string) (*entity.CharacterCollection, error) {
	uuid := uuid.New().String()
	hash := fmt.Sprintf("%x", md5.Sum([]byte(uuid+a.privateKey+a.publicKey)))

	url, _ := url.Parse(a.baseUrl + "/v1/public/characters")

	q := url.Query()
	q.Set("ts", uuid)
	q.Set("apikey", a.publicKey)
	q.Set("hash", hash)

	for k, v := range filter {
		q.Set(k, v)
	}

	url.RawQuery = q.Encode()

	resp, rerr := a.client.Get(url.String())
	if rerr != nil {
		return &entity.CharacterCollection{}, rerr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		dump, _ := httputil.DumpResponse(resp, true)
		return nil, fmt.Errorf("error response: %q", dump)
	}

	var cc entity.CharacterCollection
	if jerr := json.NewDecoder(resp.Body).Decode(&cc); jerr != nil {
		return &entity.CharacterCollection{}, jerr
	}

	return &cc, nil
}
