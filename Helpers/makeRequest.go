package Helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

type ChartData struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						End       int    `json:"end"`
						Start     int    `json:"start"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						End       int    `json:"end"`
						Start     int    `json:"start"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						End       int    `json:"end"`
						Start     int    `json:"start"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
				DataGranularity string   `json:"dataGranularity"`
				Range           string   `json:"range"`
				ValidRanges     []string `json:"validRanges"`
			} `json:"meta"`
			Timestamp  []int `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Volume []int     `json:"volume"`
					Open   []float64 `json:"open"`
					Close  []float64 `json:"close"`
					Low    []float64 `json:"low"`
					High   []float64 `json:"high"`
				} `json:"quote"`
				Adjclose []struct {
					Adjclose []float64 `json:"adjclose"`
				} `json:"adjclose"`
			} `json:"indicators"`
		} `json:"result"`
		Error struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"error"`
	} `json:"chart"`
}

type ChannelJsonData struct {
	Data           Data   `json:"data"`
	ErrDescription string `json:"errDescription"`
}
type Data struct {
	Symbol    string  `json:"symbol"`
	Currency  string  `json:"currency"`
	Price     float64 `json:"price"`
	Change    float64 `json:"change"`
	PerChange float64 `json:"perChange"`
}

func getNearestFloat(number float64) float64 {
	return (math.Round(number*100) / 100)
}

func MakeRequest(ticker string, ch chan ChannelJsonData) {
	decodedChartData := ChartData{}
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%v?interval=1d&events=div,splits", ticker)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	// getting random user agent
	GetUserAgent()
	req.Header.Add("User-Agent", *UserAgent)

	// initialising request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// decoding the json response body

	json.NewDecoder(resp.Body).Decode(&decodedChartData)

	// checking if there is any errors in json response
	if decodedChartData.Chart.Error.Code != "" {
		errResult := ChannelJsonData{
			Data: Data{
				Symbol:    ticker,
				Currency:  "",
				Price:     0,
				Change:    0,
				PerChange: 0,
			},
			ErrDescription: decodedChartData.Chart.Error.Description,
		}

		ch <- errResult
		return
	}

	//when sending request to SPX ticker getting length of price is 0 therefore to handle that
	if len(decodedChartData.Chart.Result[0].Indicators.Quote[0].Close) == 0 {
		errResult := ChannelJsonData{
			Data: Data{
				Symbol:    ticker,
				Currency:  "",
				Price:     0,
				Change:    0,
				PerChange: 0,
			},
			ErrDescription: "No data found please enter valid ticker",
		}

		ch <- errResult
		return
	}

	symbol := decodedChartData.Chart.Result[0].Meta.Symbol
	price := getNearestFloat(decodedChartData.Chart.Result[0].Indicators.Quote[0].Close[0])
	prevClose := getNearestFloat(decodedChartData.Chart.Result[0].Meta.ChartPreviousClose)
	change := getNearestFloat(price - prevClose)
	perChange := getNearestFloat((change / prevClose) * 100)
	currency := decodedChartData.Chart.Result[0].Meta.Currency
	successResult := ChannelJsonData{
		Data: Data{
			Symbol:    symbol,
			Currency:  currency,
			Price:     price,
			Change:    change,
			PerChange: perChange,
		},
		ErrDescription: "",
	}
	ch <- successResult

}
