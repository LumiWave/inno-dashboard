package point_manager_server

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (o *PointManagerServerInfo) GetPointAppList(muid, databaseid int64) (*MePointInfo, error) {
	uri := fmt.Sprintf(ApiList[Api_get_point_list].Uri, muid, databaseid)
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, uri)

	data, err := HttpCall(callUrl, o.ApiKey, "GET", Api_get_point_list, bytes.NewBuffer(nil), nil)
	if err != nil {
		return nil, err
	}

	return data.(*MePointInfo), nil
}

func (o *PointManagerServerInfo) PostPointCoinSwap(swapInfo *ReqSwapInfo) (*ResSwapInfo, error) {
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, ApiList[Api_post_swap].Uri)

	pbytes, _ := json.Marshal(swapInfo)
	buff := bytes.NewBuffer(pbytes)

	data, err := HttpCall(callUrl, o.ApiKey, "POST", Api_post_swap, buff, nil)
	if err != nil {
		return nil, err
	}

	return data.(*ResSwapInfo), nil
}

func (o *PointManagerServerInfo) PostCoinTransfer(req *ReqCoinTransfer) (*ResCoinTransfer, error) {
	callUrl := fmt.Sprintf("%s%s%s", o.IntHostUri, o.IntVer, ApiList[Api_coin_transfer].Uri)

	pbytes, _ := json.Marshal(req)
	buff := bytes.NewBuffer(pbytes)

	data, err := HttpCall(callUrl, o.ApiKey, "POST", Api_coin_transfer, buff, nil)
	if err != nil {
		return nil, err
	}

	return data.(*ResCoinTransfer), nil
}
