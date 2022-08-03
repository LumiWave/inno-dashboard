package model

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/ONBUFF-IP-TOKEN/basedb"
	"github.com/ONBUFF-IP-TOKEN/baseutil/datetime"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
)

func (o *DB) PublishEvent(channel string, val interface{}) error {
	msg, _ := json.Marshal(val)
	return o.Cache.GetDB().Publish(MakePubSubKey(channel), string(msg))
}

func (o *DB) ListenSubscribeEvent(termSec int64) error {
	receiveCh := make(chan basedb.PubSubMessageV8)
	defer close(receiveCh)

	channel := MakePubSubKey(InternalCmd)
	rch, err := o.Cache.GetDB().Subscribe(receiveCh, channel)
	if err != nil {
		log.Errorf("pubsub Subscribe err:%v", err)
		return err
	}

	log.Info("ListenSubscribeEvent() has been started")

	conf := config.GetInstance()
	bKeepAlive := func() bool {
		if conf.App.LiquidityUpdate {
			return true
		} else {
			return false
		}
	}

	defer func(ch string, termSec int64) {
		o.Cache.GetDB().Unsubscribe(ch)
		o.Cache.GetDB().ClosePubSub()
		if recver := recover(); recver != nil {
			log.Error("Recoverd in listenPubSubEvent()", recver)
		}
		go o.ListenSubscribeEvent(termSec)
	}(channel, termSec)

	if bKeepAlive() {
		go func() {
			ticker := time.NewTicker(time.Duration(termSec) * time.Second)

			for {
				msg := &PSHealthCheck{
					PSHeader: PSHeader{
						Type: PubSub_cmd_healthcheck,
					},
				}
				msg.Value.Timestamp = datetime.GetTS2MilliSec()

				if err := o.PublishEvent(InternalCmd, msg); err != nil {
					log.Errorf("pubsub health check err : %v", err)
				}
				<-ticker.C
			}

		}()
	}

	for {
		msg, ok := <-rch
		if msg == nil || !ok {
			continue
		}

		if strings.Contains(msg.Channel, MakePubSubKey(InternalCmd)) {
			o.PubSubCmdByInternal(msg)
		}

		log.Debugf("subscribe channel: %v, val: %v", msg.Channel, msg.Payload)
	}

	return nil
}

func (o *DB) PubSubCmdByInternal(msg basedb.PubSubMessageV8) error {

	header := &PSHeader{}
	json.Unmarshal([]byte(msg.Payload), header)

	if strings.EqualFold(header.Type, PubSub_type_maintenance) {
		psPacket := &PSMaintenance{}
		json.Unmarshal([]byte(msg.Payload), psPacket)
		SetMaintenance(psPacket.Value.Enable)
	} else if strings.EqualFold(header.Type, PubSub_type_Swap) {
		psPacket := &PSSwap{}
		json.Unmarshal([]byte(msg.Payload), psPacket)
		SetSwapEnable(psPacket.Value.ToCoinEnable, psPacket.Value.ToPointEnable)
	} else if strings.EqualFold(header.Type, PubSub_type_CoinTransferExternal) {
		psPacket := &PSCoinTransferExternal{}
		json.Unmarshal([]byte(msg.Payload), psPacket)
		SetExternalTransferEnable(psPacket.Value.Enable)
	} else if strings.EqualFold(header.Type, PubSub_type_meta_refresh) {
		// db meta refresh
		gDB.GetPointList()
		gDB.GetBaseCoins()
		gDB.GetAppCoins()
		gDB.GetCoins()
		gDB.GetApps()
		gDB.GetAppPoints()
		gDB.GetScanExchangeGoods()
		log.Infof("pubsub cmd : %v", PubSub_type_meta_refresh)
	}
	return nil
}
