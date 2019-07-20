package ngerest

import (
	"strconv"
	"strings"
)

// StringFloat json string format float64
type StringFloat float64

// UnmarshalJSON unmarshal float from json
func (f *StringFloat) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	dataStr = strings.Trim(dataStr, "\" ")

	if dataStr == "" {
		*f = 0
		return nil
	}

	value, err := strconv.ParseFloat(dataStr, 64)
	if err == nil {
		return err
	}

	*f = value
	return nil
}

// Instrument Tradeable Contracts, Indices, and History
type Instrument struct {
	Symbol                         string      `json:"symbol"`
	RootSymbol                     string      `json:"rootSymbol,omitempty"`
	State                          string      `json:"state,omitempty"`
	Typ                            string      `json:"typ,omitempty"`
	Listing                        NGETime     `json:"listing,omitempty"`
	Front                          NGETime     `json:"front,omitempty"`
	Expiry                         NGETime     `json:"expiry,omitempty"`
	Settle                         NGETime     `json:"settle,omitempty"`
	RelistInterval                 NGETime     `json:"relistInterval,omitempty"`
	InverseLeg                     string      `json:"inverseLeg,omitempty"`
	SellLeg                        string      `json:"sellLeg,omitempty"`
	BuyLeg                         string      `json:"buyLeg,omitempty"`
	OptionStrikePcnt               float64     `json:"optionStrikePcnt,omitempty"`
	OptionStrikeRound              float64     `json:"optionStrikeRound,omitempty"`
	OptionStrikePrice              float64     `json:"optionStrikePrice,omitempty"`
	OptionMultiplier               float64     `json:"optionMultiplier,omitempty"`
	PositionCurrency               string      `json:"positionCurrency,omitempty"`
	Underlying                     string      `json:"underlying,omitempty"`
	QuoteCurrency                  string      `json:"quoteCurrency,omitempty"`
	UnderlyingSymbol               string      `json:"underlyingSymbol,omitempty"`
	Reference                      string      `json:"reference,omitempty"`
	ReferenceSymbol                string      `json:"referenceSymbol,omitempty"`
	CalcInterval                   NGETime     `json:"calcInterval,omitempty"`
	PublishInterval                NGETime     `json:"publishInterval,omitempty"`
	PublishTime                    NGETime     `json:"publishTime,omitempty"`
	MaxOrderQty                    float32     `json:"maxOrderQty,omitempty"`
	MaxPrice                       float64     `json:"maxPrice,omitempty"`
	LotSize                        float32     `json:"lotSize,omitempty"`
	TickSize                       float64     `json:"tickSize,omitempty"`
	Multiplier                     float32     `json:"multiplier,omitempty"`
	SettlCurrency                  string      `json:"settlCurrency,omitempty"`
	UnderlyingToPositionMultiplier float32     `json:"underlyingToPositionMultiplier,omitempty"`
	UnderlyingToSettleMultiplier   float32     `json:"underlyingToSettleMultiplier,omitempty"`
	QuoteToSettleMultiplier        float32     `json:"quoteToSettleMultiplier,omitempty"`
	IsQuanto                       bool        `json:"isQuanto,omitempty"`
	IsInverse                      bool        `json:"isInverse,omitempty"`
	InitMargin                     float64     `json:"initMargin,omitempty"`
	MaintMargin                    float64     `json:"maintMargin,omitempty"`
	RiskLimit                      float32     `json:"riskLimit,omitempty"`
	RiskStep                       float32     `json:"riskStep,omitempty"`
	Limit                          StringFloat `json:"limit,omitempty"`
	Capped                         bool        `json:"capped,omitempty"`
	Taxed                          bool        `json:"taxed,omitempty"`
	Deleverage                     bool        `json:"deleverage,omitempty"`
	MakerFee                       float64     `json:"makerFee,omitempty"`
	TakerFee                       float64     `json:"takerFee,omitempty"`
	SettlementFee                  float64     `json:"settlementFee,omitempty"`
	InsuranceFee                   float64     `json:"insuranceFee,omitempty"`
	FundingBaseSymbol              string      `json:"fundingBaseSymbol,omitempty"`
	FundingQuoteSymbol             string      `json:"fundingQuoteSymbol,omitempty"`
	FundingPremiumSymbol           string      `json:"fundingPremiumSymbol,omitempty"`
	FundingTimestamp               NGETime     `json:"fundingTimestamp,omitempty"`
	FundingInterval                NGETime     `json:"fundingInterval,omitempty"`
	FundingRate                    float64     `json:"fundingRate,omitempty"`
	IndicativeFundingRate          float64     `json:"indicativeFundingRate,omitempty"`
	RebalanceTimestamp             NGETime     `json:"rebalanceTimestamp,omitempty"`
	RebalanceInterval              NGETime     `json:"rebalanceInterval,omitempty"`
	OpeningTimestamp               NGETime     `json:"openingTimestamp,omitempty"`
	ClosingTimestamp               NGETime     `json:"closingTimestamp,omitempty"`
	SessionInterval                NGETime     `json:"sessionInterval,omitempty"`
	PrevClosePrice                 float64     `json:"prevClosePrice,omitempty"`
	LimitDownPrice                 float64     `json:"limitDownPrice,omitempty"`
	LimitUpPrice                   float64     `json:"limitUpPrice,omitempty"`
	BankruptLimitDownPrice         StringFloat `json:"bankruptLimitDownPrice,omitempty"`
	BankruptLimitUpPrice           StringFloat `json:"bankruptLimitUpPrice,omitempty"`
	PrevTotalVolume                float32     `json:"prevTotalVolume,omitempty"`
	TotalVolume                    float32     `json:"totalVolume,omitempty"`
	Volume                         float32     `json:"volume,omitempty"`
	Volume24h                      float32     `json:"volume24h,omitempty"`
	PrevTotalTurnover              float32     `json:"prevTotalTurnover,omitempty"`
	TotalTurnover                  float32     `json:"totalTurnover,omitempty"`
	Turnover                       float32     `json:"turnover,omitempty"`
	Turnover24h                    float32     `json:"turnover24h,omitempty"`
	PrevPrice24h                   float64     `json:"prevPrice24h,omitempty"`
	Vwap                           float64     `json:"vwap,omitempty"`
	HighPrice                      float64     `json:"highPrice,omitempty"`
	LowPrice                       float64     `json:"lowPrice,omitempty"`
	LastPrice                      float64     `json:"lastPrice,omitempty"`
	LastPriceProtected             float64     `json:"lastPriceProtected,omitempty"`
	LastTickDirection              string      `json:"lastTickDirection,omitempty"`
	LastChangePcnt                 float64     `json:"lastChangePcnt,omitempty"`
	BidPrice                       float64     `json:"bidPrice,omitempty"`
	MidPrice                       float64     `json:"midPrice,omitempty"`
	AskPrice                       float64     `json:"askPrice,omitempty"`
	ImpactBidPrice                 float64     `json:"impactBidPrice,omitempty"`
	ImpactMidPrice                 float64     `json:"impactMidPrice,omitempty"`
	ImpactAskPrice                 float64     `json:"impactAskPrice,omitempty"`
	HasLiquidity                   bool        `json:"hasLiquidity,omitempty"`
	OpenInterest                   float32     `json:"openInterest,omitempty"`
	OpenValue                      float32     `json:"openValue,omitempty"`
	FairMethod                     string      `json:"fairMethod,omitempty"`
	FairBasisRate                  float64     `json:"fairBasisRate,omitempty"`
	FairBasis                      float64     `json:"fairBasis,omitempty"`
	FairPrice                      float64     `json:"fairPrice,omitempty"`
	MarkMethod                     string      `json:"markMethod,omitempty"`
	MarkPrice                      float64     `json:"markPrice,omitempty"`
	IndicativeTaxRate              float64     `json:"indicativeTaxRate,omitempty"`
	IndicativeSettlePrice          float64     `json:"indicativeSettlePrice,omitempty"`
	OptionUnderlyingPrice          float64     `json:"optionUnderlyingPrice,omitempty"`
	SettledPrice                   float64     `json:"settledPrice,omitempty"`
	Timestamp                      NGETime     `json:"timestamp,omitempty"`
	BestPrice                      float64     `json:"bestPrice,omitempty"`
	ForeignNotional24H             float64     `json:"foreignNotional24h,omitempty"`
	HomeNotional24H                float64     `json:"homeNotional24h,omitempty"`
	LiqMarkPrice                   float64     `json:"liqMarkPrice,omitempty"`
	OrderMarkPrice                 float64     `json:"orderMarkPrice,omitempty"`
	PrevMarkPrice24H               float64     `json:"prevMarkPrice24h,omitempty"`
}
