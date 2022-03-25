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

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), nil, &MePointInfo{})
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

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, buff, nil, &ResSwapInfo{})
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

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, buff, nil, &ResCoinTransferFromParentWallet{})
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

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, buff, nil, &ResCoinTransferFromUserWallet{})
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransferFromUserWallet), nil
}

func (o *PointManagerServerInfo) GetCoinTransferExistInProgress(auid int64) (*ResCoinTransferFromUserWallet, error) {
	api := ApiList[Api_coin_transfer_exist_inprogress]
	uri := fmt.Sprintf(api.Uri, auid)
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, uri)

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), nil, &ResCoinTransferFromUserWallet{})
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransferFromUserWallet), nil
}

func (o *PointManagerServerInfo) GetCoinTransferNotExistInProgress(auid int64) (*ResCoinTransferFromUserWallet, error) {
	api := ApiList[Api_coin_transfer_exist_inprogress]
	uri := fmt.Sprintf(api.Uri, auid)
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, uri)

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), nil, &ResCoinTransferFromUserWallet{})
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransferFromUserWallet), nil
}

func (o *PointManagerServerInfo) GetCoinFee(req *ReqCoinFee) (*ResCoinFeeInfo, error) {
	urlInfo := ApiList[Api_get_coin_fee]
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, urlInfo.Uri)

	data, err := HttpCall(callUrl, o.ApiKey, urlInfo.Method, urlInfo.ApiType, bytes.NewBuffer(nil), req, &ResCoinFeeInfo{})
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinFeeInfo), nil
}
