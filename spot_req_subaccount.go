package mybinanceapi

import "github.com/shopspring/decimal"

type SpotSubAccountListReq struct {
	Email      *string `json:"email"`
	Isfreeze   *string `json:"isFreeze"`
	Page       *int    `json:"page"`
	Limit      *int    `json:"limit"`
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SpotSubAccountListApi struct {
	client *SpotRestClient
	req    *SpotSubAccountListReq
}

func (api *SpotSubAccountListApi) Email(Email string) *SpotSubAccountListApi {
	api.req.Email = GetPointer(Email)
	return api
}
func (api *SpotSubAccountListApi) Isfreeze(Isfreeze bool) *SpotSubAccountListApi {
	if Isfreeze {
		api.req.Isfreeze = GetPointer("true")
	} else {
		api.req.Isfreeze = GetPointer("false")
	}
	return api
}
func (api *SpotSubAccountListApi) Page(Page int) *SpotSubAccountListApi {
	api.req.Page = GetPointer(Page)
	return api
}
func (api *SpotSubAccountListApi) Limit(Limit int) *SpotSubAccountListApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotSubAccountListApi) RecvWindow(RecvWindow int64) *SpotSubAccountListApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountListApi) Timestamp(Timestamp int64) *SpotSubAccountListApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountUniversalTransferHistoryReq struct {
	FromEmail    *string `json:"fromEmail"`    //NO
	ToEmail      *string `json:"toEmail"`      //NO
	ClientTranId *string `json:"clientTranId"` //NO
	StartTime    *int64  `json:"startTime"`    //NO
	EndTime      *int64  `json:"endTime"`      //NO
	Page         *int    `json:"page"`         //NO	默认 1
	Limit        *int    `json:"limit"`        //NO	默认 500, 最大 500
	RecvWindow   *int64  `json:"recvWindow"`
	Timestamp    *int64  `json:"timestamp"`
}
type SpotSubAccountUniversalTransferHistoryApi struct {
	client *SpotRestClient
	req    *SpotSubAccountUniversalTransferHistoryReq
}

func (api *SpotSubAccountUniversalTransferHistoryApi) FromEmail(FromEmail string) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.FromEmail = GetPointer(FromEmail)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) ToEmail(ToEmail string) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.ToEmail = GetPointer(ToEmail)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) ClientTranId(ClientTranId string) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.ClientTranId = GetPointer(ClientTranId)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) StartTime(StartTime int64) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) EndTime(EndTime int64) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) Page(Page int) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.Page = GetPointer(Page)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) Limit(Limit int) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) RecvWindow(RecvWindow int64) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) Timestamp(Timestamp int64) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountAssetsReq struct {
	Email      *string `json:"email"` //YES	子账户邮箱 备注
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SpotSubAccountAssetsApi struct {
	client *SpotRestClient
	req    *SpotSubAccountAssetsReq
}

func (api *SpotSubAccountAssetsApi) Email(Email string) *SpotSubAccountAssetsApi {
	api.req.Email = GetPointer(Email)
	return api
}
func (api *SpotSubAccountAssetsApi) RecvWindow(RecvWindow int64) *SpotSubAccountAssetsApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountAssetsApi) Timestamp(Timestamp int64) *SpotSubAccountAssetsApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountFuturesAccountReq struct {
	Email       *string `json:"email"`       //YES	子账户邮箱 备注
	FuturesType *int    `json:"futuresType"` //YES	1:USDT Margined Futures, 2:COIN Margined Futures
	RecvWindow  *int64  `json:"recvWindow"`  //NO
	Timestamp   *int64  `json:"timestamp"`   //YES
}
type SpotSubAccountFuturesAccountApi struct {
	client *SpotRestClient
	req    *SpotSubAccountFuturesAccountReq
}

func (api *SpotSubAccountFuturesAccountApi) Email(Email string) *SpotSubAccountFuturesAccountApi {
	api.req.Email = GetPointer(Email)
	return api
}
func (api *SpotSubAccountFuturesAccountApi) FuturesType(FuturesType int) *SpotSubAccountFuturesAccountApi {
	api.req.FuturesType = GetPointer(FuturesType)
	return api
}
func (api *SpotSubAccountFuturesAccountApi) RecvWindow(RecvWindow int64) *SpotSubAccountFuturesAccountApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountFuturesAccountApi) Timestamp(Timestamp int64) *SpotSubAccountFuturesAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountApiIpRestrictionReq struct {
	Email            *string `json:"email"`            //YES	Sub-account email
	SubAccountApiKey *string `json:"subAccountApiKey"` //YES
	RecvWindow       *int64  `json:"recvWindow"`
	Timestamp        *int64  `json:"timestamp"`
}
type SpotSubAccountApiIpRestrictionApi struct {
	client *SpotRestClient
	req    *SpotSubAccountApiIpRestrictionReq
}

func (api *SpotSubAccountApiIpRestrictionApi) Email(Email string) *SpotSubAccountApiIpRestrictionApi {
	api.req.Email = GetPointer(Email)
	return api
}
func (api *SpotSubAccountApiIpRestrictionApi) SubAccountApiKey(SubAccountApiKey string) *SpotSubAccountApiIpRestrictionApi {
	api.req.SubAccountApiKey = GetPointer(SubAccountApiKey)
	return api
}
func (api *SpotSubAccountApiIpRestrictionApi) RecvWindow(RecvWindow int64) *SpotSubAccountApiIpRestrictionApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountApiIpRestrictionApi) Timestamp(Timestamp int64) *SpotSubAccountApiIpRestrictionApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountTransferSubUserHistoryReq struct {
	Asset             *string `json:"asset"`             //NO	如不提供，返回所有asset 划转记录
	Type              *int64  `json:"type"`              //NO	1: transfer in, 2: transfer out; 如不提供，返回transfer out方向划转记录
	StartTime         *int64  `json:"startTime"`         //NO
	EndTime           *int64  `json:"endTime"`           //NO
	Limit             *int64  `json:"limit"`             //NO	默认值: 500
	ReturnFailHistory *bool   `json:"returnFailHistory"` //NO	默认False，返回PROCESS和SUCCESS状态的数据；如果传True返回PROCESS、SUCCESS、FAILURE状态的数据
	RecvWindow        *int64  `json:"recvWindow"`        //NO
	Timestamp         *int64  `json:"timestamp"`         //YES
}
type SpotSubAccountTransferSubUserHistoryApi struct {
	client *SpotRestClient
	req    *SpotSubAccountTransferSubUserHistoryReq
}

func (api *SpotSubAccountTransferSubUserHistoryApi) Asset(Asset string) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) Type(Type int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) StartTime(StartTime int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) EndTime(EndTime int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) Limit(Limit int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) ReturnFailHistory(ReturnFailHistory bool) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.ReturnFailHistory = GetPointer(ReturnFailHistory)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) RecvWindow(RecvWindow int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) Timestamp(Timestamp int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotManagedSubAccountQueryTransLogReq struct {
	StartTime                   *int64  `json:"startTime"`                   //YES	开始时间
	EndTime                     *int64  `json:"endTime"`                     //YES	结束时间(开始时间结束时间间隔不能超过半年)
	Page                        *int    `json:"page"`                        //YES	页数
	Limit                       *int    `json:"limit"`                       //YES	每页数量 (最大值: 500)
	Transfers                   *string `json:"transfers"`                   //NO	划转方向 (FROM/TO)
	TransferFunctionAccountType *string `json:"transferFunctionAccountType"` //NO	划转账户类型 (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE)
	RecvWindow                  *int64  `json:"recvWindow"`                  //NO
	Timestamp                   *int64  `json:"timestamp"`                   //YES
}
type SpotManagedSubAccountQueryTransLogApi struct {
	client *SpotRestClient
	req    *SpotManagedSubAccountQueryTransLogReq
}

func (api *SpotManagedSubAccountQueryTransLogApi) StartTime(StartTime int64) *SpotManagedSubAccountQueryTransLogApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotManagedSubAccountQueryTransLogApi) EndTime(EndTime int64) *SpotManagedSubAccountQueryTransLogApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotManagedSubAccountQueryTransLogApi) Page(Page int) *SpotManagedSubAccountQueryTransLogApi {
	api.req.Page = GetPointer(Page)
	return api
}
func (api *SpotManagedSubAccountQueryTransLogApi) Limit(Limit int) *SpotManagedSubAccountQueryTransLogApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotManagedSubAccountQueryTransLogApi) Transfers(Transfers string) *SpotManagedSubAccountQueryTransLogApi {
	api.req.Transfers = GetPointer(Transfers)
	return api
}
func (api *SpotManagedSubAccountQueryTransLogApi) TransferFunctionAccountType(TransferFunctionAccountType string) *SpotManagedSubAccountQueryTransLogApi {
	api.req.TransferFunctionAccountType = GetPointer(TransferFunctionAccountType)
	return api
}
func (api *SpotManagedSubAccountQueryTransLogApi) RecvWindow(RecvWindow int64) *SpotManagedSubAccountQueryTransLogApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotManagedSubAccountQueryTransLogApi) Timestamp(Timestamp int64) *SpotManagedSubAccountQueryTransLogApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountVirtualSubAccountReq struct {
	SubAccountString *string `json:"subAccountString"`
	RecvWindow       *int64  `json:"recvWindow"`
	Timestamp        *int64  `json:"timestamp"`
}
type SpotSubAccountVirtualSubAccountApi struct {
	client *SpotRestClient
	req    *SpotSubAccountVirtualSubAccountReq
}

func (api *SpotSubAccountVirtualSubAccountApi) SubAccountString(SubAccountString string) *SpotSubAccountVirtualSubAccountApi {
	api.req.SubAccountString = GetPointer(SubAccountString)
	return api
}
func (api *SpotSubAccountVirtualSubAccountApi) RecvWindow(RecvWindow int64) *SpotSubAccountVirtualSubAccountApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountVirtualSubAccountApi) Timestamp(Timestamp int64) *SpotSubAccountVirtualSubAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountUniversalTransferReq struct {
	FromEmail       *string          `json:"fromEmail"`       //NO
	ToEmail         *string          `json:"toEmail"`         //NO
	FromAccountType *string          `json:"fromAccountType"` //YES	"SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	ToAccountType   *string          `json:"toAccountType"`   //YES	"SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	ClientTranId    *string          `json:"clientTranId"`    //NO	不可重复
	Symbol          *string          `json:"symbol"`          //NO	仅在ISOLATED_MARGIN类型下使用
	Asset           *string          `json:"asset"`           //YES
	Amount          *decimal.Decimal `json:"amount"`          //YES
	RecvWindow      *int64           `json:"recvWindow"`
	Timestamp       *int64           `json:"timestamp"`
}
type SpotSubAccountUniversalTransferApi struct {
	client *SpotRestClient
	req    *SpotSubAccountUniversalTransferReq
}

func (api *SpotSubAccountUniversalTransferApi) FromEmail(FromEmail string) *SpotSubAccountUniversalTransferApi {
	api.req.FromEmail = GetPointer(FromEmail)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) ToEmail(ToEmail string) *SpotSubAccountUniversalTransferApi {
	api.req.ToEmail = GetPointer(ToEmail)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) FromAccountType(FromAccountType string) *SpotSubAccountUniversalTransferApi {
	api.req.FromAccountType = GetPointer(FromAccountType)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) ToAccountType(ToAccountType string) *SpotSubAccountUniversalTransferApi {
	api.req.ToAccountType = GetPointer(ToAccountType)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) ClientTranId(ClientTranId string) *SpotSubAccountUniversalTransferApi {
	api.req.ClientTranId = GetPointer(ClientTranId)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) Symbol(Symbol string) *SpotSubAccountUniversalTransferApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) Asset(Asset string) *SpotSubAccountUniversalTransferApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) Amount(Amount decimal.Decimal) *SpotSubAccountUniversalTransferApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) RecvWindow(RecvWindow int64) *SpotSubAccountUniversalTransferApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) Timestamp(Timestamp int64) *SpotSubAccountUniversalTransferApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountFuturesEnableReq struct {
	Email      *string `json:"email"` //YES	子账户邮箱 备注
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SpotSubAccountFuturesEnableApi struct {
	client *SpotRestClient
	req    *SpotSubAccountFuturesEnableReq
}

func (api *SpotSubAccountFuturesEnableApi) Email(Email string) *SpotSubAccountFuturesEnableApi {
	api.req.Email = GetPointer(Email)
	return api
}
func (api *SpotSubAccountFuturesEnableApi) RecvWindow(RecvWindow int64) *SpotSubAccountFuturesEnableApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountFuturesEnableApi) Timestamp(Timestamp int64) *SpotSubAccountFuturesEnableApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountMarginEnableReq struct {
	Email      *string `json:"email"`      // YES 子账户邮箱 备注
	RecvWindow *int64  `json:"recvWindow"` // NO 接收窗口
	Timestamp  *int64  `json:"timestamp"`  // YES 时间戳
}

// YES 子账户邮箱 备注
func (api *SpotSubAccountMarginEnableApi) Email(Email string) *SpotSubAccountMarginEnableApi {
	api.req.Email = GetPointer(Email)
	return api
}

// NO 接收窗口
func (api *SpotSubAccountMarginEnableApi) RecvWindow(RecvWindow int64) *SpotSubAccountMarginEnableApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

// YES 时间戳
func (api *SpotSubAccountMarginEnableApi) Timestamp(Timestamp int64) *SpotSubAccountMarginEnableApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountMarginEnableApi struct {
	client *SpotRestClient
	req    *SpotSubAccountMarginEnableReq
}

// email	STRING	YES	Sub-account email
// subAccountApiKey	STRING	YES
// status	STRING	YES	IP限制状态。1或不填入(null) = IP未受限。2 = 仅限受信任IP访问。
// ipAddress	STRING	NO	可批量填入IP，以逗号区隔
// recvWindow	LONG	NO
// timestamp	LONG	YES
type SpotSubAccountSubAccountApiIpRestrictionPostReq struct {
	Email            *string `json:"email"`            // YES Sub-account email
	SubAccountApiKey *string `json:"subAccountApiKey"` // YES
	Status           *string `json:"status"`           // YES	IP限制状态。1或不填入(null) = IP未受限。2 = 仅限受信任IP访问。
	IpAddress        *string `json:"ipAddress"`        // NO	可批量填入IP，以逗号区隔
	RecvWindow       *int64  `json:"recvWindow"`       // NO
	Timestamp        *int64  `json:"timestamp"`        // YES
}

// YES	Sub-account email
func (api *SpotSubAccountSubAccountApiIpRestrictionPostApi) Email(Email string) *SpotSubAccountSubAccountApiIpRestrictionPostApi {
	api.req.Email = GetPointer(Email)
	return api
}

// YES
func (api *SpotSubAccountSubAccountApiIpRestrictionPostApi) SubAccountApiKey(SubAccountApiKey string) *SpotSubAccountSubAccountApiIpRestrictionPostApi {
	api.req.SubAccountApiKey = GetPointer(SubAccountApiKey)
	return api
}

// YES	IP限制状态。1或不填入(null) = IP未受限。2 = 仅限受信任IP访问。
func (api *SpotSubAccountSubAccountApiIpRestrictionPostApi) Status(Status string) *SpotSubAccountSubAccountApiIpRestrictionPostApi {
	api.req.Status = GetPointer(Status)
	return api
}

// NO	可批量填入IP，以逗号区隔
func (api *SpotSubAccountSubAccountApiIpRestrictionPostApi) IpAddress(IpAddress string) *SpotSubAccountSubAccountApiIpRestrictionPostApi {
	api.req.IpAddress = GetPointer(IpAddress)
	return api
}

// NO recvWindow
func (api *SpotSubAccountSubAccountApiIpRestrictionPostApi) RecvWindow(RecvWindow int64) *SpotSubAccountSubAccountApiIpRestrictionPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

// YES
func (api *SpotSubAccountSubAccountApiIpRestrictionPostApi) Timestamp(Timestamp int64) *SpotSubAccountSubAccountApiIpRestrictionPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountSubAccountApiIpRestrictionPostApi struct {
	client *SpotRestClient
	req    *SpotSubAccountSubAccountApiIpRestrictionPostReq
}

type SpotSubAccountSubAccountApiIpRestrictionGetReq struct {
	Email            *string `json:"email"`            // YES Sub-account email
	SubAccountApiKey *string `json:"subAccountApiKey"` // YES
	RecvWindow       *int64  `json:"recvWindow"`       // NO
	Timestamp        *int64  `json:"timestamp"`        // YES
}

// YES Sub-account email
func (api *SpotSubAccountSubAccountApiIpRestrictionGetApi) Email(Email string) *SpotSubAccountSubAccountApiIpRestrictionGetApi {
	api.req.Email = GetPointer(Email)
	return api
}

// YES
func (api *SpotSubAccountSubAccountApiIpRestrictionGetApi) SubAccountApiKey(SubAccountApiKey string) *SpotSubAccountSubAccountApiIpRestrictionGetApi {
	api.req.SubAccountApiKey = GetPointer(SubAccountApiKey)
	return api
}

// NO recvWindow
func (api *SpotSubAccountSubAccountApiIpRestrictionGetApi) RecvWindow(RecvWindow int64) *SpotSubAccountSubAccountApiIpRestrictionGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

// YES
func (api *SpotSubAccountSubAccountApiIpRestrictionGetApi) Timestamp(Timestamp int64) *SpotSubAccountSubAccountApiIpRestrictionGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountSubAccountApiIpRestrictionGetApi struct {
	client *SpotRestClient
	req    *SpotSubAccountSubAccountApiIpRestrictionGetReq
}

type SpotSubAccountSubAccountApiIpRestrictionDeleteReq struct {
	Email            *string `json:"email"`            // YES Sub-account email
	SubAccountApiKey *string `json:"subAccountApiKey"` // YES
	IpAddress        *string `json:"ipAddress"`        // YES 可批量填入IP，以逗号分隔
	RecvWindow       *int64  `json:"recvWindow"`       // NO 接收窗口
	Timestamp        *int64  `json:"timestamp"`        // YES 时间戳
}

// YES Sub-account email
func (api *SpotSubAccountSubAccountApiIpRestrictionDeleteApi) Email(Email string) *SpotSubAccountSubAccountApiIpRestrictionDeleteApi {
	api.req.Email = GetPointer(Email)
	return api
}

// YES
func (api *SpotSubAccountSubAccountApiIpRestrictionDeleteApi) SubAccountApiKey(SubAccountApiKey string) *SpotSubAccountSubAccountApiIpRestrictionDeleteApi {
	api.req.SubAccountApiKey = GetPointer(SubAccountApiKey)
	return api
}

// YES 可批量填入IP，以逗号分隔
func (api *SpotSubAccountSubAccountApiIpRestrictionDeleteApi) IpAddress(IpAddress string) *SpotSubAccountSubAccountApiIpRestrictionDeleteApi {
	api.req.IpAddress = GetPointer(IpAddress)
	return api
}

// NO 接收窗口
func (api *SpotSubAccountSubAccountApiIpRestrictionDeleteApi) RecvWindow(RecvWindow int64) *SpotSubAccountSubAccountApiIpRestrictionDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

// YES 时间戳
func (api *SpotSubAccountSubAccountApiIpRestrictionDeleteApi) Timestamp(Timestamp int64) *SpotSubAccountSubAccountApiIpRestrictionDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountSubAccountApiIpRestrictionDeleteApi struct {
	client *SpotRestClient
	req    *SpotSubAccountSubAccountApiIpRestrictionDeleteReq
}

type SpotSubAccountCapitalDepositSubAddressReq struct {
	Email      *string          `json:"email"`      // YES	子账户邮箱 备注
	Coin       *string          `json:"coin"`       // YES
	Network    *string          `json:"network"`    // NO	网络
	Amount     *decimal.Decimal `json:"amount"`     // NO	充值数量 使用LIGHTNING网络时，amount必须传入
	RecvWindow *int64           `json:"recvWindow"` // NO
	Timestamp  *int64           `json:"timestamp"`  // YES
}

// YES	子账户邮箱 备注
func (api *SpotSubAccountCapitalDepositSubAddressApi) Email(Email string) *SpotSubAccountCapitalDepositSubAddressApi {
	api.req.Email = GetPointer(Email)
	return api
}

// YES
func (api *SpotSubAccountCapitalDepositSubAddressApi) Coin(Coin string) *SpotSubAccountCapitalDepositSubAddressApi {
	api.req.Coin = GetPointer(Coin)
	return api
}

// NO	网络
func (api *SpotSubAccountCapitalDepositSubAddressApi) Network(Network string) *SpotSubAccountCapitalDepositSubAddressApi {
	api.req.Network = GetPointer(Network)
	return api
}

// NO	充值数量 使用LIGHTNING网络时，amount必须传入
func (api *SpotSubAccountCapitalDepositSubAddressApi) Amount(Amount decimal.Decimal) *SpotSubAccountCapitalDepositSubAddressApi {
	api.req.Amount = GetPointer(Amount)
	return api
}

// NO
func (api *SpotSubAccountCapitalDepositSubAddressApi) RecvWindow(RecvWindow int64) *SpotSubAccountCapitalDepositSubAddressApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

// YES
func (api *SpotSubAccountCapitalDepositSubAddressApi) Timestamp(Timestamp int64) *SpotSubAccountCapitalDepositSubAddressApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountCapitalDepositSubAddressApi struct {
	client *SpotRestClient
	req    *SpotSubAccountCapitalDepositSubAddressReq
}

type SpotSubAccountCapitalDepositSubHisrecReq struct {
	Email      *string `json:"email"`      // YES	子账户邮箱 备注
	Coin       *string `json:"coin"`       // NO
	Status     *int    `json:"status"`     // NO	0(0:pending,6: credited but cannot withdraw,7:Wrong Deposit,8:Waiting User confirm,1:success)
	StartTime  *int64  `json:"startTime"`  // NO
	EndTime    *int64  `json:"endTime"`    // NO
	Limit      *int    `json:"limit"`      // NO
	Offset     *int    `json:"offset"`     // NO	default:0
	RecvWindow *int64  `json:"recvWindow"` // NO
	Timestamp  *int64  `json:"timestamp"`  // YES
	TxId       *string `json:"txId"`       // NO
}

// YES	子账户邮箱 备注
func (api *SpotSubAccountCapitalDepositSubHisrecApi) Email(Email string) *SpotSubAccountCapitalDepositSubHisrecApi {
	api.req.Email = GetPointer(Email)
	return api
}

// NO
func (api *SpotSubAccountCapitalDepositSubHisrecApi) Coin(Coin string) *SpotSubAccountCapitalDepositSubHisrecApi {
	api.req.Coin = GetPointer(Coin)
	return api
}

// NO	0(0:pending,6: credited but cannot withdraw,7:Wrong Deposit,8:Waiting User confirm,1:success)
func (api *SpotSubAccountCapitalDepositSubHisrecApi) Status(Status int) *SpotSubAccountCapitalDepositSubHisrecApi {
	api.req.Status = GetPointer(Status)
	return api
}

// NO
func (api *SpotSubAccountCapitalDepositSubHisrecApi) StartTime(StartTime int64) *SpotSubAccountCapitalDepositSubHisrecApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}

// NO
func (api *SpotSubAccountCapitalDepositSubHisrecApi) EndTime(EndTime int64) *SpotSubAccountCapitalDepositSubHisrecApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}

// NO
func (api *SpotSubAccountCapitalDepositSubHisrecApi) Limit(Limit int) *SpotSubAccountCapitalDepositSubHisrecApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// NO	default:0
func (api *SpotSubAccountCapitalDepositSubHisrecApi) Offset(Offset int) *SpotSubAccountCapitalDepositSubHisrecApi {
	api.req.Offset = GetPointer(Offset)
	return api
}

// NO
func (api *SpotSubAccountCapitalDepositSubHisrecApi) RecvWindow(RecvWindow int64) *SpotSubAccountCapitalDepositSubHisrecApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

// YES
func (api *SpotSubAccountCapitalDepositSubHisrecApi) Timestamp(Timestamp int64) *SpotSubAccountCapitalDepositSubHisrecApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountCapitalDepositSubHisrecApi struct {
	client *SpotRestClient
	req    *SpotSubAccountCapitalDepositSubHisrecReq
}

type SpotSubAccountFuturesPositionRiskReq struct {
	Email       *string `json:"email"`       // YES	子账户邮箱 备注
	FuturesType *int    `json:"futuresType"` // YES	1:USDT Margined Futures, 2:COIN Margined Futures
	RecvWindow  *int64  `json:"recvWindow"`  // NO
	Timestamp   *int64  `json:"timestamp"`   // YES
}

// YES	子账户邮箱 备注
func (api *SpotSubAccountFuturesPositionRiskApi) Email(Email string) *SpotSubAccountFuturesPositionRiskApi {
	api.req.Email = GetPointer(Email)
	return api
}

// YES	1:USDT Margined Futures, 2:COIN Margined Futures
func (api *SpotSubAccountFuturesPositionRiskApi) FuturesType(FuturesType int) *SpotSubAccountFuturesPositionRiskApi {
	api.req.FuturesType = GetPointer(FuturesType)
	return api
}

// NO
func (api *SpotSubAccountFuturesPositionRiskApi) RecvWindow(RecvWindow int64) *SpotSubAccountFuturesPositionRiskApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

// YES
func (api *SpotSubAccountFuturesPositionRiskApi) Timestamp(Timestamp int64) *SpotSubAccountFuturesPositionRiskApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountFuturesPositionRiskApi struct {
	client *SpotRestClient
	req    *SpotSubAccountFuturesPositionRiskReq
}

type SpotSubAccountSpotSummaryReq struct {
	Email      *string `json:"email"`      // NO	子账户邮箱
	Page       *int    `json:"page"`       // NO	分页，默认 1
	Size       *int    `json:"size"`       // NO	单页条目数, 默认 10, 最大 20
	RecvWindow *int64  `json:"recvWindow"` // NO
	Timestamp  *int64  `json:"timestamp"`  // YES
}

// NO	子账户邮箱
func (api *SpotSubAccountSpotSummaryApi) Email(Email string) *SpotSubAccountSpotSummaryApi {
	api.req.Email = GetPointer(Email)
	return api
}

// NO	分页，默认 1
func (api *SpotSubAccountSpotSummaryApi) Page(Page int) *SpotSubAccountSpotSummaryApi {
	api.req.Page = GetPointer(Page)
	return api
}

// NO	单页条目数, 默认 10, 最大 20
func (api *SpotSubAccountSpotSummaryApi) Size(Size int) *SpotSubAccountSpotSummaryApi {
	api.req.Size = GetPointer(Size)
	return api
}

// NO
func (api *SpotSubAccountSpotSummaryApi) RecvWindow(RecvWindow int64) *SpotSubAccountSpotSummaryApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

// YES
func (api *SpotSubAccountSpotSummaryApi) Timestamp(Timestamp int64) *SpotSubAccountSpotSummaryApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountSpotSummaryApi struct {
	client *SpotRestClient
	req    *SpotSubAccountSpotSummaryReq
}
