package mybinanceapi

import (
	"fmt"
	"time"
)

func (ws *SpotWsStreamClient) SubscribeUserDataStream() (*WsApiResult[SpotWsApiUserDataStreamResult], error) {
	type Timestamp struct {
		Timestamp *int64 `json:"timestamp"`
	}

	type WsSpotSubscribeUserDataStreamReq struct {
		WsApiReqExtend
		Timestamp
	}

	type WsSpotMarginUserListenTokenPostReq struct {
		ListenToken *string `json:"listenToken"`
	}

	switch ws.spotWsType {
	case SPOT_WS_TYPE:
		nowTimestamp := time.Now().UnixMilli() + serverTimeDelta

		timestamp := Timestamp{
			Timestamp: &nowTimestamp,
		}

		req := WsSpotSubscribeUserDataStreamReq{
			WsApiReqExtend: WsApiReqExtend{
				ApiKey:    ws.apiKey,
				Signature: signWsApi(ws.apiKey, ws.apiSecret, &timestamp),
			},
			Timestamp: timestamp,
		}
		return sendApiMsg[WsSpotSubscribeUserDataStreamReq, SpotWsApiUserDataStreamResult](&ws.WsStreamClient, 0, "userDataStream.subscribe.signature", req)
	case SPOT_MARGIN_WS_TYPE:
		res, err := ws.client.NewSpotMarginUserListenTokenPost().Do()
		if err != nil {
			return nil, err
		}
		req := WsSpotMarginUserListenTokenPostReq{
			ListenToken: &res.Token,
		}
		return sendApiMsg[WsSpotMarginUserListenTokenPostReq, SpotWsApiUserDataStreamResult](&ws.WsStreamClient, 0, "userDataStream.subscribe.listenToken", req)
	case SPOT_ISOLATED_MARGIN_WS_TYPE:
		res, err := ws.client.NewSpotMarginUserListenTokenPost().Symbol(ws.isolatedSymbol).IsIsolated(true).Do()
		if err != nil {
			return nil, err
		}
		req := WsSpotMarginUserListenTokenPostReq{
			ListenToken: &res.Token,
		}
		return sendApiMsg[WsSpotMarginUserListenTokenPostReq, SpotWsApiUserDataStreamResult](&ws.WsStreamClient, 0, "userDataStream.subscribe.listenToken", req)
	default:
		return nil, fmt.Errorf("spotWsType is not support")
	}
}
