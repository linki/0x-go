package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/linki/0x-go/types"
)

type Client struct {
	url    string
	client *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		url:    url,
		client: http.DefaultClient,
	}
}

func (c *Client) GetTokenPairs(ctx context.Context) ([]types.TokenPair, error) {
	req, err := http.NewRequest(http.MethodGet, c.url+"/token_pairs", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erroneous status code: %s", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tokenPairs := []types.TokenPair{}

	if err := json.Unmarshal(respBody, &tokenPairs); err != nil {
		return nil, fmt.Errorf("error parsing json response: %v", err)
	}

	return tokenPairs, nil
}
