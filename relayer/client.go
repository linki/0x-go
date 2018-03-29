package relayer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/types"
	"github.com/linki/0x-go/util"
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

type GetTokenPairsOpts struct {
	TokenA common.Address
	TokenB common.Address
}

func (c *Client) GetTokenPairs(ctx context.Context, opts GetTokenPairsOpts) ([]types.TokenPair, error) {
	reqURL, err := url.Parse(c.url + "/token_pairs")
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	if !util.EmptyAddress(opts.TokenA) {
		query["tokenA"] = []string{strings.ToLower(opts.TokenA.Hex())}
	}
	if !util.EmptyAddress(opts.TokenB) {
		query["tokenB"] = []string{strings.ToLower(opts.TokenB.Hex())}
	}
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
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

type GetOrdersOpts struct {
	ExchangeContractAddress common.Address
	TokenAddress            common.Address
	MakerTokenAddress       common.Address
	TakerTokenAddress       common.Address
	Maker                   common.Address
	Taker                   common.Address
	Trader                  common.Address
	FeeRecipient            common.Address
}

func (c *Client) GetOrders(ctx context.Context, opts GetOrdersOpts) ([]types.Order, error) {
	reqURL, err := url.Parse(c.url + "/orders")
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	if !util.EmptyAddress(opts.ExchangeContractAddress) {
		query["exchangeContractAddress"] = []string{strings.ToLower(opts.ExchangeContractAddress.Hex())}
	}
	if !util.EmptyAddress(opts.TokenAddress) {
		query["tokenAddress"] = []string{strings.ToLower(opts.TokenAddress.Hex())}
	}
	if !util.EmptyAddress(opts.MakerTokenAddress) {
		query["makerTokenAddress"] = []string{strings.ToLower(opts.MakerTokenAddress.Hex())}
	}
	if !util.EmptyAddress(opts.TakerTokenAddress) {
		query["takerTokenAddress"] = []string{strings.ToLower(opts.TakerTokenAddress.Hex())}
	}
	if !util.EmptyAddress(opts.Maker) {
		query["maker"] = []string{strings.ToLower(opts.Maker.Hex())}
	}
	if !util.EmptyAddress(opts.Taker) {
		query["taker"] = []string{strings.ToLower(opts.Taker.Hex())}
	}
	if !util.EmptyAddress(opts.Trader) {
		query["trader"] = []string{strings.ToLower(opts.Trader.Hex())}
	}
	if !util.EmptyAddress(opts.FeeRecipient) {
		query["feeRecipient"] = []string{strings.ToLower(opts.FeeRecipient.Hex())}
	}
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
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

	orders := []types.Order{}

	if err := json.Unmarshal(respBody, &orders); err != nil {
		return nil, fmt.Errorf("error parsing json response: %v", err)
	}

	return orders, nil
}

func (c *Client) GetOrder(ctx context.Context, orderHash common.Hash) (types.Order, error) {
	req, err := http.NewRequest(http.MethodGet, c.url+"/order/"+orderHash.Hex(), nil)
	if err != nil {
		return types.Order{}, err
	}

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return types.Order{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return types.Order{}, fmt.Errorf("order not found: %s", orderHash.Hex())
	}

	if resp.StatusCode != http.StatusOK {
		return types.Order{}, fmt.Errorf("erroneous status code: %s", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.Order{}, err
	}

	order := types.Order{}

	if err := json.Unmarshal(respBody, &order); err != nil {
		return types.Order{}, fmt.Errorf("error parsing json response: %v", err)
	}

	return order, nil
}

func (c *Client) CreateOrder(ctx context.Context, order types.Order) error {
	reqBody, err := order.MarshalJSON()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.url+"/order", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("erroneous status code: %s", resp.Status)
	}

	return nil
}

func (c *Client) GetFees(ctx context.Context, order types.UnsignedOrder) (types.Fees, error) {
	reqBody, err := order.MarshalJSON()
	if err != nil {
		return types.Fees{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.url+"/fees", bytes.NewReader(reqBody))
	if err != nil {
		return types.Fees{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return types.Fees{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.Fees{}, fmt.Errorf("erroneous status code: %s", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.Fees{}, err
	}

	fees := types.Fees{}

	if err := json.Unmarshal(respBody, &fees); err != nil {
		return types.Fees{}, fmt.Errorf("error parsing json response: %v", err)
	}

	return fees, nil
}
