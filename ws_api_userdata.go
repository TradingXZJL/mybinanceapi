package mybinanceapi

import "time"

func (ws *SpotWsStreamClient) SubscribeUserDataStream() (*WsApiResult[SpotWsApiUserDataStreamResult], error) {
	type Timestamp struct {
		Timestamp *int64 `json:"timestamp"`
	}

	type WsSpotSubscribeUserDataStreamReq struct {
		WsApiReqExtend
		Timestamp
	}

	nowTimestamp := time.Now().UnixMilli()

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
}
