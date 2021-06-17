package apiclient

import "io"

//APIClient interface for all exchange api
type APIClient interface {
	// INIT

	/* Init struture. In parameters it takes private values ​​for
	 * access to api. If any exchange implementation does not use some parameter, when
	 * calling a function, it remains empty.
	 *
	 * returns:
	 *   Returns error.
	 */
	Init(accountID string, apiKey string, apiSecret string) error

	// PUBLIC API

	/* Retrieves the last price on symbol
	 *
	 * arguments:
	 *   symbol  pair text-id in apiclient format.
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *
	 * returns:
	 *   Returns lastPrice and error
	 */
	GetLastPrice(symbol string) (lastPrice float64, err error)

	/* Get excnhage Order book and sort it.
	 *
	 * arguments:
	 *   symbol   pair text-id in apiclient format
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *   depth   not required argument (default value 50)
	 *      (how many orders you whant to see on Buy/Sell side)
	 *
	 * bids - Buy orders, sorted by price descending.
	 * asks - Sell orders, sorted by price ascending.
	 *
	 * returns:
	 *   Returns a pointer to apiclient.OrderBook structure and error.
	 */
	GetOrderBook(symbol string, depth ...int) (*OrderBook, error)

	/* Retrieves the number of decimal after comma for price and amount
	 *
	 * arguments:
	 *   symbol  pair text-id in apiclient format.
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *
	 * returns:
	 *   Returns a pointer to apiclient.Decimals structure and error
	 */
	GetDecs(symbol string) (*Decimals, error)

	/* Retrieves last volume and price candles
	 *
	 * arguments:
	 *   symbol        pair text-id in apiclient format.
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *   candlePeriod  candle period in minutes (1m = 1, 1h = 60, ...)
	 *   number        how many candles you want to recive.
	 *
	 * returns:
	 *   Returns a pointer to filled and time-sorted structure of type apiclient.KLine and error
	 */
	GetKLine(symbol string, candlePeriod int, number int) (*KLine, error)

	/* Retrieves trade history
	 *
	 * arguments:
	 *   symbol  pair text-id in apiclient format.
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *   number  how many trades you want to recive.
	 *
	 * returns:
	 *   Returns a pointer to an array of filled and time-sorted structures of type apiclient.Trade and error
	 */
	GetTradeHistory(symbol string, number int) (*[]Trade, error)

	/* Retrieves market data information
	 *
	 * arguments:
	 *   symbol  pair text-id in apiclient format.
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *
	 * returns:
	 *   Returns a pointer to apiclient.MarketData structure and error.
	 */
	GetMarketData(symbol string) (*MarketData, error)

	/* Retrieves all trading pairs from exchange
	 *
	 * returns:
	 *   Return a pointer to an array of symbols and error.
	 *   symbol  pair text-id in apiclient format.
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 */
	GetTradingPairs(symbol string) (*[]string, error)

	// PRIVATE API

	/* Receives non-zero currency balances.
	 *
	 * returns:
	 *   Returns a pointer to the implementation balance map and error.
	 */
	GetBalances() (*map[string]Balance, error) 

	/* Get information about order.
	 *
	 * arguments:
	 *   symbol  pair text-id in apiclient format, not required argument (depends on exchange)
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *   id       MakedOrder.ID
	 *
	 * returns:
	 *   Returns a pointer to apiclient.MakedOrder structure and error.
	 */
	GetOrderStatus(id string, symbol string) (*MakedOrder, error)

	/* Put Limit Sell order.
	 *
	 * arguments:
	 *   symbol:   pair text-id in apiclient format
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *   amount   token amount
	 *   price    one token price
	 *	 check    not required argument (default value false)
	 *      (if true after Sell request call GetOrderStatus and fulfill apiclient.MakedOrder strucure)
	 *
	 * returns:
	 *   Returns a pointer to apiclient.MakedOrder structure and error.
	 */
	Sell(symbol string, amount float64, price float64, check ...bool) (*MakedOrder, error)

	/* Put Limit Buy order.
	 *
	 * arguments:
	 *   symbol:   pair text-id in apiclient format
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *   amount   token amount
	 *   price    one token price
	 *	 check    not required argument (default value false)
	 *      (if true after Buy request call GetOrderStatus and fulfill apiclient.MakedOrder strucure)
	 *
	 * returns:
	 *   Returns a pointer to apiclient.MakedOrder structure and error.
	 */
	Buy(symbol string, amount float64, price float6, check ...bool4) (*MakedOrder, error)

	/* Cancel order.
	 *
	 * arguments:
	 *   symbol   pair text-id in apiclient format, not required argument (depends on exchange)
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *   id       MakedOrder.ID
	 *
	 * returns:
	 *   Returns error.
	 */
	CancelOrder(symbol string, id string) error

	/* Cancel all open orders on symbol.
	 *
	 * arguments:
	 *   symbol   pair text-id in apiclient format, not required argument (depends on exchange)
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *
	 * returns:
	 *   Returns error.
	 */
	CancelAll(symbol string) error

	/* Get a list of all open orders on this account on symbol.
	 *
	 * arguments:
	 *   symbol  pair text-id in apiclient format, not required argument (if empty return all open orders)
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *
	 * returns:
	 *   Returns a pointer to an array of apiclient.MakedOrder structure and error.
	 */
	GetMyOpenOrders(symbol string) (*[]MakedOrder, error)

	/* Get a list of all trades from this account on symbol per period.
	 *
	 * arguments:
	 *   symbol  pair text-id in apiclient format, not required argument (if empty return all trades)
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *   period   enum (1d, 1w, 1m)
	 *
	 * returns:
	 *   Returns a pointer to an array of apiclient.MakedOrder structure and error.
	 */
	GetMyTradeHistory(symbol string, period string) (*[]MakedOrder, error)

	/* Get a list of all orders from this account on symbol per period.
	 *
	 * arguments:
	 *   symbol  pair text-id in apiclient format, not required argument (if empty return all trades)
	 * 		(example "USDT_BTC" main currency on the left, on the right is the currency that we buy when calling Buy)
	 *   period   enum (1d, 1w, 1m)
	 *
	 * returns:
	 *   Returns a pointer to an array of apiclient.MakedOrder structure and error.
	 */
	GetMyOrderHistory(symbol string, period string) (*[]MakedOrder, error)

	/* Creates a withdraw request.
	 *
	 * arguments:
	 *   asset     currency id in string format: <raw-line, uppercase> (BTC)
	 *   address   blockchain address
	 *   amount    withdrawal amount
	 *   chain     blockchain
	 *
	 * returns:
	 *   id    id of the withdrawal request in the exchange string format
	 *   err   error
	 */
	Withdraw(asset string, address string, amount float64, chain string) (id string, err error)
	
	/* Get withdraw list.
	 *
	 * returns:
	 *   Returns a pointer to an array of apiclient.Transfer structure and error.
	 */
	GetWithdrawList() (*[]Transfer, err error)
	
	/* Get deposit list.
	 *
	 * returns:
	 *   Returns a pointer to an array of apiclient.Transfer structure and error.
	 */
	GetDepositList() (*[]Transfer, err error)
}

//Status type for enum about order status
type Status string

//Side type for enum about order buy or sell types or something like that
type Side string

//Color type for volume candle color
type Color string

//constants about Status and Side
const (
	Buy  Side = "BUY"
	Sell Side = "SELL"

	Filled          Status = "FILLED"
	NotFilled       Status = "NOTFILLED"
	PartiallyFilled Status = "PATIALLYFILLED"
	Undefined       Status = "UNDEFINED"

	Red   Color = "rgba(255, 82, 82, 0.5)" // DOWN or SELL
	Green Color = "rgba(0, 150, 136, 0.5)" // UP or BUY
)

//Balance help struct for APIClient
type Balance struct {
	Free   float64 `json:"free"`   // Available balance for use in new orders
	Locked float64 `json:"locked"` // Locked balance in orders or withdrawals
}

//OrderBook help struct for APIClient
type OrderBook struct {
	Asks []Order `json:"asks"` // asks.Price > any bids.Price
	Bids []Order `json:"bids"` // bids.Price < any asks.Price
}

//Order help struct for APIClient
type Order struct {
	Quantity float64 `json:"quantity"` // token amount
	Price    float64 `json:"price"`    // one token price
}

//MakedOrder help struct for APIClient
type MakedOrder struct {
	Time   int64   `json:"time"` // UNIX time in seconds (10 digits)
	
	ID string `json:"id"`

	Status Status `json:"status"` // Status Should be one of apiclient.Status constants(Filled, NotFilled, PartiallyFilled, Undefined)

	LeftAmount  float64 `json:"leftAmount"`
	RightAmount float64 `json:"rightAmount"`

	LeftAmountExecuted  float64 `json:"leftAmountExecuted"`
	RightAmountExecuted float64 `json:"rightAmountExecuted"`

	// Commission is factically not used
	Commission   float64 `json:"commission"`
	Rate         float64 `json:"rate"`
	RateExecuted float64 `json:"rateExecuted"`

	// Side Should be one of apiclient.Side constants(Buy, Sell)
	Side Side `json:"side"`
}

//KLine help struct for APIClient
type KLine struct {
	PriceCandles  []PriceCandle  `json:"priceCandles"`
	VolumeCandles []VolumeCandle `json:"volumeCandles"`
}

//PriceCandle help struct for APIClient
type PriceCandle struct {
	Time  int64   `json:"time"`  // UNIX time in seconds (10 digits)
	Open  float64 `json:"open"`  // open price
	Close float64 `json:"close"` // close price
	High  float64 `json:"high"`  // high price
	Low   float64 `json:"low"`   // low price
}

//VolumeCandle help struct for APIClient
type VolumeCandle struct {
	Time  int64   `json:"time"`  // UNIX time in seconds (10 digits)
	Value float64 `json:"value"` // volume
	Color Color   `json:"color"` // apiclient.Green if Close > Open, apiclient.Red if Close < Open
}

type Trade struct {
	Time   int64   `json:"time"` // UNIX time in seconds (10 digits)
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	Side   Side    `json:"side"` // if side doesnt defined, UP = BUY / DOWN = SELL
}

type Decimals struct {
	PriceDecs  int `json:"priceDecs"`
	AmountDecs int `json:"amountDecs"`
}

type MarketData struct {
	VolumeLeft      float64 `json:"volumeLeft"`
	VolumeRight     float64 `json:"volumeRight"`
	Price           float64 `json:"price"`
	PriceChnagePerc float64 `json:"priceChnagePerc"`
	PriceChnageAbs  float64 `json:"priceChnageAbs"`
	SpreadPerc      float64 `json:"spreadPerc"`
	MinSell         float64 `json:"minSell"`
	MaxBuy          float64 `json:"maxBuy"`
	DayPriceHigh    float64 `json:"dayPriceHigh"`
	DayPriceLow     float64 `json:"dayPriceLow"`
}

type Transfer struct {
	Time     int64   `json:"time"` // UNIX time in seconds (10 digits)
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Txid     string  `json:"txid"`
}
