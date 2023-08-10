package point_manager_server

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (o *PointManagerServerInfo) GetPointAppList(muid, databaseid int64) (*MePointInfo, error) {
	api := ApiList[Api_get_point_list]
	uri := fmt.Sprintf(api.Uri, muid, databaseid)
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, uri)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), nil, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*MePointInfo), nil
}

func (o *PointManagerServerInfo) PostPointCoinSwap(swapInfo *ReqSwapInfo) (*ResSwapInfo, error) {
	api := ApiList[Api_post_swap]
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, api.Uri)

	pbytes, _ := json.Marshal(swapInfo)
	buff := bytes.NewBuffer(pbytes)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, buff, nil, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*ResSwapInfo), nil
}

func (o *PointManagerServerInfo) PostCoinTransferFromParentWallet(req *ReqCoinTransferFromParentWallet) (*ResCoinTransferFromParentWallet, error) {
	api := ApiList[Api_coin_transfer_from_parentwallet]
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, api.Uri)

	pbytes, _ := json.Marshal(req)
	buff := bytes.NewBuffer(pbytes)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, buff, nil, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransferFromParentWallet), nil
}

func (o *PointManagerServerInfo) PostCoinTransferFromUserWallet(req *ReqCoinTransferFromUserWallet) (*ResCoinTransferFromUserWallet, error) {
	api := ApiList[Api_coin_transfer_from_userwallet]
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, api.Uri)

	pbytes, _ := json.Marshal(req)
	buff := bytes.NewBuffer(pbytes)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, buff, nil, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransferFromUserWallet), nil
}

func (o *PointManagerServerInfo) GetCoinTransferExistInProgress(auid int64) (*ResCoinTransferFromUserWallet, error) {
	api := ApiList[Api_coin_transfer_exist_inprogress]
	uri := fmt.Sprintf(api.Uri, auid)
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, uri)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), nil, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransferFromUserWallet), nil
}

func (o *PointManagerServerInfo) GetCoinTransferNotExistInProgress(auid int64) (*ResCoinTransferFromUserWallet, error) {
	api := ApiList[Api_coin_transfer_exist_inprogress]
	uri := fmt.Sprintf(api.Uri, auid)
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, uri)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), nil, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransferFromUserWallet), nil
}

func (o *PointManagerServerInfo) GetCoinFee(req *ReqCoinFee) (*ResCoinFeeInfo, error) {
	api := ApiList[Api_get_coin_fee]
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, api.Uri)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), req, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinFeeInfo), nil
}
func (o *PointManagerServerInfo) PostCoinReload(req *CoinReload) (*ResCoinReload, error) {
	api := ApiList[Api_post_coin_reload]
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, api.Uri)

	pbytes, _ := json.Marshal(req)
	buff := bytes.NewBuffer(pbytes)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, buff, req, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinReload), nil
}

func (o *PointManagerServerInfo) GetBalance(req *ReqBalance) (*ResBalance, error) {
	api := ApiList[Api_get_balance]
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, api.Uri)

	pbytes, _ := json.Marshal(req)
	buff := bytes.NewBuffer(pbytes)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, buff, req, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*ResBalance), nil
}

func (o *PointManagerServerInfo) GetBalanceAll(req *ReqBalanceAll) (*ResBalanceAll, error) {
	api := ApiList[Api_get_balances]
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, api.Uri)

	pbytes, _ := json.Marshal(req)
	buff := bytes.NewBuffer(pbytes)

	data, err := HttpCall(api.client, callUrl, o.ApiKey, api.Method, api.ApiType, buff, req, api.ResponseFuncType())
	if err != nil {
		return nil, err
	}

	return data.(*ResBalanceAll), nil
}
