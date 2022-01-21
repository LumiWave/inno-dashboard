package point_manager_server

import (
	"bytes"
	"fmt"
)

func (o *PointManagerServerInfo) GetPointAppList(muid, databaseid int64) (*MePointInfo, error) {
	uri := fmt.Sprintf(ApiList[Api_get_point_list].Uri, muid, databaseid)
	callUrl := fmt.Sprintf("%s%s%s", o.HostUri, o.Ver, uri)

	data, err := HttpCall(callUrl, o.ApiKey, "GET", Api_get_point_list, bytes.NewBuffer(nil), nil)
	if err != nil {
		return nil, err
	}

	return data.(*MePointInfo), nil
}

// func GetCandleMinutes(upbitInfo *PointManagerInfo, interval int, symbol string, count int, to string) (*[]CandleMinute, error) {
// 	uri := fmt.Sprintf(ApiList[Api_upbit_candle_minutes].Uri, interval, symbol, count, to)
// 	callUrl := fmt.Sprintf("%s%s%s", upbitInfo.HostUri, upbitInfo.Ver, uri)

// 	data, err := HttpCall(callUrl, upbitInfo.ApiKey, "GET", Api_upbit_candle_minutes, bytes.NewBuffer(nil), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data.(*[]CandleMinute), nil
// }

// func GetCandleMinutesByStruct(upbitInfo *UpbitInfo, interval int, req *ReqCandleMinute) (*[]CandleMinute, error) {
// 	uri := fmt.Sprintf(ApiList[Api_upbit_candle_minutes_bystruct].Uri, interval)
// 	callUrl := fmt.Sprintf("%s%s%s", upbitInfo.HostUri, upbitInfo.Ver, uri)

// 	data, err := HttpCall(callUrl, upbitInfo.ApiKey, "GET", Api_upbit_candle_minutes_bystruct, bytes.NewBuffer(nil), req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data.(*[]CandleMinute), nil
// }

// func GetCandleDays(upbitInfo *UpbitInfo, symbol string, count int, to, convertingPriceUnit string) (*[]CandleDay, error) {
// 	uri := fmt.Sprintf(ApiList[Api_upbit_candle_day].Uri, symbol, count, to, convertingPriceUnit)
// 	callUrl := fmt.Sprintf("%s%s%s", upbitInfo.HostUri, upbitInfo.Ver, uri)

// 	data, err := HttpCall(callUrl, upbitInfo.ApiKey, "GET", Api_upbit_candle_day, bytes.NewBuffer(nil), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data.(*[]CandleDay), nil
// }

// func GetCandleDaysByStruct(upbitInfo *UpbitInfo, req *ReqCandleDay) (*[]CandleDay, error) {
// 	uri := fmt.Sprintf(ApiList[Api_upbit_candle_day_bystruct].Uri)
// 	callUrl := fmt.Sprintf("%s%s%s", upbitInfo.HostUri, upbitInfo.Ver, uri)

// 	data, err := HttpCall(callUrl, upbitInfo.ApiKey, "GET", Api_upbit_candle_day_bystruct, bytes.NewBuffer(nil), req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data.(*[]CandleDay), nil
// }

// func GetCandleWeeks(upbitInfo *UpbitInfo, symbol string, count int, to string) (*[]CandleWeek, error) {
// 	uri := fmt.Sprintf(ApiList[Api_upbit_candle_week].Uri, symbol, count, to)
// 	callUrl := fmt.Sprintf("%s%s%s", upbitInfo.HostUri, upbitInfo.Ver, uri)

// 	data, err := HttpCall(callUrl, upbitInfo.ApiKey, "GET", Api_upbit_candle_week, bytes.NewBuffer(nil), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data.(*[]CandleWeek), nil
// }

// func GetCandleWeeksByStruct(upbitInfo *UpbitInfo, req *ReqCandleWeek) (*[]CandleWeek, error) {
// 	uri := fmt.Sprintf(ApiList[Api_upbit_candle_week_bystruct].Uri)
// 	callUrl := fmt.Sprintf("%s%s%s", upbitInfo.HostUri, upbitInfo.Ver, uri)

// 	data, err := HttpCall(callUrl, upbitInfo.ApiKey, "GET", Api_upbit_candle_week_bystruct, bytes.NewBuffer(nil), req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data.(*[]CandleWeek), nil
// }

// func GetCandleMonths(upbitInfo *UpbitInfo, symbol string, count int, to string) (*[]CandleMonth, error) {
// 	uri := fmt.Sprintf(ApiList[Api_upbit_candle_month].Uri, symbol, count, to)
// 	callUrl := fmt.Sprintf("%s%s%s", upbitInfo.HostUri, upbitInfo.Ver, uri)

// 	data, err := HttpCall(callUrl, upbitInfo.ApiKey, "GET", Api_upbit_candle_month, bytes.NewBuffer(nil), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data.(*[]CandleMonth), nil
// }

// func GetCandleMonthsByStruct(upbitInfo *UpbitInfo, req *ReqCandleMonth) (*[]CandleMonth, error) {
// 	uri := fmt.Sprintf(ApiList[Api_upbit_candle_month_bystruct].Uri)
// 	callUrl := fmt.Sprintf("%s%s%s", upbitInfo.HostUri, upbitInfo.Ver, uri)

// 	data, err := HttpCall(callUrl, upbitInfo.ApiKey, "GET", Api_upbit_candle_month_bystruct, bytes.NewBuffer(nil), req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data.(*[]CandleMonth), nil
// }
