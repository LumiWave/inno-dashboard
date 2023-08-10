package point_manager_server

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

type api_kind int

const (
	Api_get_point_list                  = 0 // 포인트 리스트 조회 : GetPointAppList
	Api_post_swap                       = 1 // swap 요청 : PostPointCoinSwap
	Api_coin_transfer_from_parentwallet = 2 // 외부 지갑 coin 전송  : PostCoinTransferFromParentWallet
	Api_coin_transfer_from_userwallet   = 3 // 외부 지갑 coin 전송  : PostCoinTransferFromUserWallet
	Api_coin_transfer_exist_inprogress  = 4 // 외부 지갑 coin 전송 중 체크 : GetCoinTransferExistInProgress
	Api_get_coin_fee                    = 5 // 코인 가스비 정보 요청 : GetCoinFee
	Api_post_coin_reload                = 6 // 코인 mainnet reload : PostCoinReload
	Api_get_balance                     = 7 // 단일 코인 balance 조회 : GetBalance
	Api_get_balances                    = 8 // 유저의 전체 코인 balance 조회 : GetBalanceAll
)

type ApiInfo struct {
	ApiType          api_kind
	Uri              string
	Method           string
	ResponseFuncType func() interface{}
	client           *http.Client
}

var ApiList = map[api_kind]ApiInfo{
	Api_get_point_list: ApiInfo{ApiType: Api_get_point_list, Uri: "/point/app?mu_id=%d&database_id=%d", Method: "GET",
		ResponseFuncType: func() interface{} { return new(MePointInfo) }, client: NewClient()},
	Api_post_swap: ApiInfo{ApiType: Api_post_swap, Uri: "/swap", Method: "POST",
		ResponseFuncType: func() interface{} { return new(ResSwapInfo) }, client: NewClient()},
	Api_coin_transfer_from_parentwallet: ApiInfo{ApiType: Api_coin_transfer_from_parentwallet, Uri: "/transfer/parent", Method: "POST",
		ResponseFuncType: func() interface{} { return new(ResCoinTransferFromParentWallet) }, client: NewClient()},
	Api_coin_transfer_from_userwallet: ApiInfo{ApiType: Api_coin_transfer_from_userwallet, Uri: "/transfer/user", Method: "POST",
		ResponseFuncType: func() interface{} { return new(ResCoinTransferFromUserWallet) }, client: NewClient()},
	Api_coin_transfer_exist_inprogress: ApiInfo{ApiType: Api_coin_transfer_exist_inprogress, Uri: "/transfer/existinprogress?au_id=%d", Method: "GET",
		ResponseFuncType: func() interface{} { return new(ResCoinTransferFromUserWallet) }, client: NewClient()},
	Api_get_coin_fee: ApiInfo{ApiType: Api_get_coin_fee, Uri: "/coin/fee", Method: "GET",
		ResponseFuncType: func() interface{} { return new(ResCoinFeeInfo) }, client: NewClient()},
	Api_post_coin_reload: ApiInfo{ApiType: Api_post_coin_reload, Uri: "/coin/reload", Method: "POST",
		ResponseFuncType: func() interface{} { return new(ResCoinReload) }, client: NewClient()},
	Api_get_balance: ApiInfo{ApiType: Api_get_balance, Uri: "/coin/address/balance", Method: "GET",
		ResponseFuncType: func() interface{} { return new(ResBalance) }, client: NewClient()},
	Api_get_balances: ApiInfo{ApiType: Api_get_balances, Uri: "/coin/address/balance/all", Method: "GET",
		ResponseFuncType: func() interface{} { return new(ResBalanceAll) }, client: NewClient()},
}

func NewClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxIdleConnsPerHost = 100
	t.IdleConnTimeout = 30 * time.Second
	t.DisableKeepAlives = false
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{
		Timeout:   60 * time.Second,
		Transport: t,
	}
	return client
}

func MakeHttpClient(callUrl string, auth string, method string, body *bytes.Buffer, queryStr string) *http.Request {
	req, err := http.NewRequest(method, callUrl, body)
	if err != nil {
		return nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if len(auth) > 0 {
		req.Header.Add("Authorization", "Bearer "+auth)
	}
	if len(queryStr) > 0 {
		req.URL.RawQuery = queryStr
	}

	return req
}

func HttpCall(client *http.Client, callUrl string, auth string, method string, kind api_kind, body *bytes.Buffer, queryStruct interface{}, response interface{}) (interface{}, error) {

	var v url.Values
	var queryStr string
	if queryStruct != nil {
		v, _ = query.Values(queryStruct)
		queryStr = v.Encode()
	}

	req := MakeHttpClient(callUrl, auth, method, body, queryStr)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ParseResponse(resp, kind, response)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ParseResponse(resp *http.Response, kind api_kind, response interface{}) (interface{}, error) {
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, errors.New(resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(response)
	if err != nil {
		return nil, err
	}
	return response, err
}
