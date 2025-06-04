package context

const (
	DB_Return_Enough_NFT = 50213
)

const (
	RankType_PurchaseReferral = int64(iota + 1)
	RankType_PurchaseNFT
)

// 상품 구매 시작 proc에 사용
type PreSales struct {
	SalesID                    int64  `json:"sales_id"`
	OpenStartSDT               string `json:"open_start_sdt"`
	OpenEndSDT                 string `json:"open_end_sdt"`
	TotalExchangePointQuantity int64  `json:"total_exchange_point_quantity"`
	CurrExchangePointQuantity  int64  `json:"curr_exchange_point_quantity"`
	ReferralStartSDT           string `json:"referral_start_sdt"`
	ReferralEndSDT             string `json:"referral_end_sdt"`
	ReferralLimitCount         int64  `json:"referral_limit_count"`
	ReferralRewardAppID        int64  `json:"referral_reward_app_id"`
	ReferralRewardPointID      int64  `json:"referral_reward_point_id"`
	ReferrerPointQuantity      int64  `json:"referrer_point_quantity"`
	RefereePointQuantity       int64  `json:"referee_point_quantity"`
	ActiveGame                 string `json:"active_game"`
	MetaURL                    string `json:"meta_url"`
}

// 추천 등록 여부
type RefreralResponse struct {
	SalesID      int64 `json:"sales_id"`
	IsRegistered bool  `json:"is_registered"`
}

// type PreSalesClaim struct {
// 	SalesID          int64  `json:"sales_id"`
// 	ClaimID          int64  `json:"claim_id"`
// 	DisplayArea      int64  `json:"display_area"`
// 	DisplaySortOrder int64  `json:"display_sort_order"`
// 	IsEnabled        bool   `json:"is_enabled"`
// 	ClaimStartSDT    string `json:"claim_start_sdt"`
// 	ClaimEndSDT      string `json:"claim_end_sdt"`
// 	ImageURL         string `json:"image_url"`
// 	ImageRewardID    int64  `json:"image_reward_id"`
// 	RewardType       int64  `json:"reward_type"`
// 	NFTPackID        int64  `json:"nft_pack_id"`
// 	LWAPrice         string `json:"lwa_price"`
// 	LWADiscountRatio string `json:"lwq_discount_ratio"`
// 	LWADiscountPrice string `json:"lwq_discount_price"`
// 	SUIPrice         string `json:"sui_price"`
// 	SUIDiscountRatio string `json:"sui_discount_ratio"`
// 	SUIDiscountPrice string `json:"sui_discount_price"`
// }

// type ReqGetPreSales struct {
// }

// func NewReqGetPreSales() *ReqGetPreSales {
// 	return new(ReqGetPreSales)
// }

// func (o *ReqGetPreSales) CheckValidate() *base.BaseResponse {
// 	return nil
// }

// type ReqGetPreSalesClaims struct {
// 	SalesID int64 `json:"sales_id" query:"sales_id"`
// }

// func NewReqGetPreSalesClaims() *ReqGetPreSalesClaims {
// 	return new(ReqGetPreSalesClaims)
// }

// func (o *ReqGetPreSalesClaims) CheckValidate() *base.BaseResponse {
// 	return nil
// }

// type GetPreSalesRank struct {
// 	SalesID int64 `json:"sales_id" query:"sales_id"`
// }

// func NewGetPreSalesRank() *GetPreSalesRank {
// 	return new(GetPreSalesRank)
// }

// func (o *GetPreSalesRank) CheckValidate() *base.BaseResponse {
// 	if o.SalesID == 0 {
// 		return base.MakeBaseResponse(resultcode.Result_Require_SalesID)
// 	}

// 	return nil
// }

// type ResGetPreSalesRank struct {
// 	RankGenerateDT      string          `json:"rank_generate_dt"`
// 	PurchaseReferalRank []*PurchaseRank `json:"purchase_referal_rank"`
// 	PurchaseNFTRank     []*PurchaseRank `json:"purchase_nft_rank"`
// }

// type PurchaseRank struct {
// 	Rank          int64   `json:"rank"`
// 	InnoUID       string  `json:"inno_uid"`
// 	PointQuantity float64 `json:"point_quantity"`
// }

// type ReqGetPreSalesDashboard struct {
// 	// accesstoken 에서 추출
// 	AUID    int64  `json:"au_id"`
// 	InnoUID string `json:"inno_uid"`

// 	SalesID int64 `json:"sales_id" query:"sales_id"`
// }

// func NewReqGetPreSalesDashboard() *ReqGetPreSalesDashboard {
// 	return new(ReqGetPreSalesDashboard)
// }

// func (o *ReqGetPreSalesDashboard) CheckValidate(ctx *InnoMarketContext) *base.BaseResponse {
// 	if ctx.GetValue() != nil {
// 		o.AUID = ctx.GetValue().AUID
// 		o.InnoUID = ctx.GetValue().InnoUID
// 	}

// 	if o.SalesID == 0 {
// 		return base.MakeBaseResponse(resultcode.Result_Require_SalesID)
// 	}

// 	return nil
// }

// type ResGetPreSalesDashboard struct {
// 	PurchaseReferalPoint *PreSalesDashboard `json:"purchase_referal_point"`
// 	PurchaseNFTPoint     *PreSalesDashboard `json:"purchase_nft_point"`
// }

// type PreSalesDashboard struct {
// 	MyPoint      int64   `json:"my_point"`
// 	MyRank       int64   `json:"my_rank"`
// 	MyPrizeRatio float64 `json:"my_prize_ratio"`
// }

// type PreSalesAccountPoint struct {
// 	PurchasePointQuantity int64   `json:"purchase_point_quantity"`
// 	PaidPointQuantity     int64   `json:"paid_point_quantity"`
// 	ReferralPointQuantity int64   `json:"referral_point_quantity"`
// 	ReferrerInnoUID       string  `json:"refferer_inno_uid"`
// 	PurchasePointRankID   int64   `json:"purchase_point_rank_id"`
// 	MyPurchasePrizeRatio  float64 `json:"my_purchase_prize_ratio"`
// 	PaidPointRankID       int64   `json:"paid_point_rank_id"`
// 	MyPaidPrizeRatio      float64 `json:"my_paid_prize_ratio"`
// }

// type PreSalesLeaderBoard struct {
// 	RankGenerateDT string                     `json:"rank_generate_dt"`
// 	Ranks          []*PreSalesLeaderBoardRank `json:"ranks"`
// }

// type PreSalesLeaderBoardRank struct {
// 	RankType      int64  `json:"rank_type"`
// 	InnoUID       string `json:"inno_uid"`
// 	RankID        int64  `json:"rank_id"`
// 	PointQuantity int64  `json:"point_quantity"`
// }
