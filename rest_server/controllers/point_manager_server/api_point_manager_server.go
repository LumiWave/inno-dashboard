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

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), nil)
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

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, buff, nil)
	if err != nil {
		return nil, err
	}

	return data.(*ResSwapInfo), nil
}

func (o *PointManagerServerInfo) PostCoinTransfer(req *ReqCoinTransfer) (*ResCoinTransfer, error) {
	api := ApiList[Api_coin_transfer]
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, api.Uri)

	pbytes, _ := json.Marshal(req)
	buff := bytes.NewBuffer(pbytes)

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, buff, nil)
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransfer), nil
}

func (o *PointManagerServerInfo) GetCoinTransferExistInProgress(auid int64) (*ResCoinTransfer, error) {
	api := ApiList[Api_coin_transfer_exist_inprogress]
	uri := fmt.Sprintf(api.Uri, auid)
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, uri)

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), nil)
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransfer), nil
}

func (o *PointManagerServerInfo) GetCoinTransferNotExistInProgress(auid int64) (*ResCoinTransfer, error) {
	api := ApiList[Api_coin_transfer_exist_inprogress]
	uri := fmt.Sprintf(api.Uri, auid)
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, uri)

	data, err := HttpCall(callUrl, o.ApiKey, api.Method, api.ApiType, bytes.NewBuffer(nil), nil)
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransfer), nil
}
