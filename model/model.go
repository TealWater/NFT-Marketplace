package model

// StreamHeartBeat This struct is for creating a heat-beat event to keep the steam alive
type StreamHeartBeat struct {
	Topic   string `json:"topic"`
	Event   string `json:"event"`
	Payload struct {
	} `json:"payload"`
	Ref int `json:"ref"`
}

type OpenSeaCollection struct {
	Total struct {
		Volume           float32 `json:"volume"`
		Sales            int     `json:"sales"`
		AveragePrice     float32 `json:"average_price"`
		NumOwners        int     `json:"num_owners"`
		MarketCap        float32 `json:"market_cap"`
		FloorPrice       float32 `json:"floor_price"`
		FloorPriceSymbol string  `json:"floor_price_symbol"`
	} `json:"total"`
	Intervals []struct {
		Interval     string  `json:"interval"`
		Volume       float32 `json:"volume"`
		VolumeDiff   float32 `json:"volume_diff"`
		VolumeChange float32 `json:"volume_change"`
		Sales        int     `json:"sales"`
		SalesDiff    float32 `json:"sales_diff"`
		AveragePrice float32 `json:"average_price"`
	} `json:"intervals"`
}
