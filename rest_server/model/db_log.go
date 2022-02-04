package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/controllers/context"
)

// 특정 개수만큼 포인트 유동량 history 정보 업데이트
func (o *DB) LoadFullPointLiquidity(interval int64) {
	start := time.Now()
	// 포인트 별 hour, day, week, month 값을 가져와서 redis에 저장
	// point loop
	for appId, app := range o.AppPointsMap {
		for _, point := range app.Points {
			// candle type loop
			req := &context.ReqPointLiquidity{}
			for candleType, candleProcedure := range context.CandleTypeOfPoint {
				req.AppID = appId
				req.PointID = point.PointId
				req.Candle = candleType
				req.Interval = interval
				req.BaseDate = ""

				key := MakeLogKeyOfPoint(appId, point.PointId, candleType)

				loopCnt := 0
				for {
					loopCnt++
					if pointLiqs, err := o.GetListPointLiquidity(candleProcedure, req); err != nil {
						log.Errorf("GetListPointLiquidity error : %v", err)
						break
					} else {
						if len(pointLiqs) == 0 {
							//log.Debugf("appID : %v, pointid : %v, candleType : %v, loopCnt:%v", appId, point.PointId, candleType, loopCnt)
							break
						} else {
							// redis에 저장하고 다음 데이터 수집한다.
							for _, pointLiq := range pointLiqs {
								// if strings.EqualFold(candleType, "hour") {
								// 	log.Debugf("appID : %v, pointid : %v, candleType : %v, loopCnt:%v, baseDate : %v %v",
								// 		appId, point.PointId, candleType, loopCnt, pointLiq.BaseDate, pointLiq.BaseDateToNumber)
								// } else {
								// 	log.Debugf("appID : %v, pointid : %v, candleType : %v, loopCnt:%v, baseDate : %v %v",
								// 		appId, point.PointId, candleType, loopCnt, pointLiq.BaseDate, pointLiq.BaseDateToNumber)
								// }
								o.ZRemRangeByScorePoint(key, strconv.FormatInt(pointLiq.BaseDateToNumber, 10), strconv.FormatInt(pointLiq.BaseDateToNumber, 10))
								if err := o.ZADDLogOfPoint(key, pointLiq.BaseDateToNumber, pointLiq); err != nil {
									log.Errorf("ZADDLogOfPoint error : %v", err)
								}
							}
							req.BaseDate = pointLiqs[len(pointLiqs)-1].BaseDate.String()
						}
					}
				}
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("LoadFullPointLiquidity took %s \n", elapsed)
}

// 특정 개수만큼 코인 유동량 history 정보 업데이트
func (o *DB) LoadFullCoinLiquidity(interval int64) {
	start := time.Now()
	// 코인 별 hour, day, week, month 값을 가져와서 redis에 저장
	// coin loop
	for coinID, _ := range o.CoinsMap {
		// candle type loop
		req := &context.ReqCoinLiquidity{}
		for candleType, candleProcedure := range context.CandleTypeOfCoin {
			req.CoinID = coinID
			req.Candle = candleType
			req.Interval = interval
			req.BaseDate = ""

			key := MakeLogKeyOfCoin(coinID, candleType)

			loopCnt := 0
			for {
				loopCnt++
				if coinLiqs, err := o.GetListCoinLiquidity(candleProcedure, req); err != nil {
					log.Errorf("GetListCoinLiquidity error : %v", err)
					break
				} else {
					if len(coinLiqs) == 0 {
						//log.Debugf("coinID : %v, candleType : %v, loopCnt:%v", coinID, candleType, loopCnt)
						break
					} else {
						// redis에 저장하고 다음 데이터 수집한다.
						for _, coinLiq := range coinLiqs {
							// if strings.EqualFold(candleType, "hour") {
							// 	log.Debugf("coinID : %v, candleType : %v, loopCnt:%v, baseDate : %v %v",
							// 		coinID, candleType, loopCnt, coinLiq.BaseDate, coinLiq.BaseDateToNumber)
							// } else {
							// 	log.Debugf("appID : %v, candleType : %v, loopCnt:%v, baseDate : %v %v",
							// 		coinID, candleType, loopCnt, coinLiq.BaseDate, coinLiq.BaseDateToNumber)
							// }
							o.ZRemRangeByScoreOfCoin(key, strconv.FormatInt(coinLiq.BaseDateToNumber, 10), strconv.FormatInt(coinLiq.BaseDateToNumber, 10))
							if err := o.ZADDLogOfCoin(key, coinLiq.BaseDateToNumber, coinLiq); err != nil {
								log.Errorf("ZADDLogOfPoint error : %v", err)
							}
						}
						req.BaseDate = coinLiqs[len(coinLiqs)-1].BaseDate.String()
					}
				}
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("LoadFullCoinLiquidity took %s \n", elapsed)
}

func (o *DB) UpdateLiquidity() {
	go func() {
		for {
			timer := time.NewTimer(1 * time.Minute)
			//timer := time.NewTimer(5 * time.Second)
			<-timer.C
			timer.Stop()

			o.LoadFullPointLiquidity(1)
			o.LoadFullCoinLiquidity(1)
		}
	}()
}

func ChangeStringDayTime(strTime string) *time.Time {
	if len(strTime) == 0 {
		return nil
	}
	var baseDate *time.Time
	t, err := time.Parse("2006-01-02", strTime)
	if err != nil {
		log.Errorf("time.Parse [err%v]", err)
		return nil
	} else {
		baseDate = &t
	}
	if t.IsZero() {
		baseDate = nil
	}
	return baseDate
}

func ChangeStringHourTime(t *time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
