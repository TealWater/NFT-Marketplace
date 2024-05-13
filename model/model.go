package model

// StreamHeartBeat This struct is for creating a heat-beat event to keep the steam alive
type StreamHeartBeat struct {
	Topic   string `json:"topic"`
	Event   string `json:"event"`
	Payload struct {
	} `json:"payload"`
	Ref int `json:"ref"`
}

type OpenSeaCollectionStats struct {
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

type OpenSeaCollection struct {
	Collection              string `json:"collection"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	ImageURL                string `json:"image_url"`
	BannerImageURL          string `json:"banner_image_url"`
	Owner                   string `json:"owner"`
	SafelistStatus          string `json:"safelist_status"`
	Category                string `json:"category"`
	IsDisabled              bool   `json:"is_disabled"`
	IsNsfw                  bool   `json:"is_nsfw"`
	TraitOffersEnabled      bool   `json:"trait_offers_enabled"`
	CollectionOffersEnabled bool   `json:"collection_offers_enabled"`
	OpenseaURL              string `json:"opensea_url"`
	ProjectURL              string `json:"project_url"`
	WikiURL                 string `json:"wiki_url"`
	DiscordURL              string `json:"discord_url"`
	TelegramURL             string `json:"telegram_url"`
	TwitterUsername         string `json:"twitter_username"`
	InstagramUsername       string `json:"instagram_username"`
	Contracts               []struct {
		Address string `json:"address"`
	} `json:"contracts"`
	Editors []string `json:"editors"`
	Fees    []struct {
		Fee       float32 `json:"fee"`
		Recipient string  `json:"recipient"`
		Required  bool    `json:"required"`
	} `json:"fees"`
	RequiredZone string `json:"required_zone"`
	Rarity       struct {
		StrategyVersion string `json:"strategy_version"`
		CalculatedAt    string `json:"calculated_at"`
		MaxRank         int    `json:"max_rank"`
		TotalSupply     int    `json:"total_supply"`
	} `json:"rarity"`
	PaymentTokens []struct {
		Symbol   string `json:"symbol"`
		Address  string `json:"address"`
		Chain    string `json:"chain"`
		Image    string `json:"image"`
		Name     string `json:"name"`
		Decimals int    `json:"decimals"`
		EthPrice string `json:"eth_price"`
		UsdPrice string `json:"usd_price"`
	} `json:"payment_tokens"`
	TotalSupply int    `json:"total_supply"`
	CreatedDate string `json:"created_date"`
}

type OpenSeaCollectionEvent struct {
	AssetEvents []struct {
		EventType      string `json:"event_type"`
		OrderHash      string `json:"order_hash,omitempty"`
		Maker          string `json:"maker,omitempty"`
		EventTimestamp int    `json:"event_timestamp"`
		Nft            struct {
			Identifier    string `json:"identifier"`
			Collection    string `json:"collection"`
			Contract      string `json:"contract"`
			TokenStandard string `json:"token_standard"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			ImageURL      string `json:"image_url"`
			MetadataURL   string `json:"metadata_url"`
			OpenseaURL    string `json:"opensea_url"`
			UpdatedAt     string `json:"updated_at"`
			IsDisabled    bool   `json:"is_disabled"`
			IsNsfw        bool   `json:"is_nsfw"`
		} `json:"nft,omitempty"`
		OrderType       string `json:"order_type,omitempty"`
		ProtocolAddress string `json:"protocol_address,omitempty"`
		StartDate       int    `json:"start_date,omitempty"`
		ExpirationDate  int    `json:"expiration_date,omitempty"`
		Asset           struct {
			Identifier    string `json:"identifier"`
			Collection    string `json:"collection"`
			Contract      string `json:"contract"`
			TokenStandard string `json:"token_standard"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			ImageURL      string `json:"image_url"`
			MetadataURL   string `json:"metadata_url"`
			OpenseaURL    string `json:"opensea_url"`
			UpdatedAt     string `json:"updated_at"`
			IsDisabled    bool   `json:"is_disabled"`
			IsNsfw        bool   `json:"is_nsfw"`
		} `json:"asset,omitempty"`
		Quantity int    `json:"quantity,omitempty"`
		Taker    string `json:"taker,omitempty"`
		Payment  struct {
			Quantity     string `json:"quantity"`
			TokenAddress string `json:"token_address"`
			Decimals     int    `json:"decimals"`
			Symbol       string `json:"symbol"`
		} `json:"payment,omitempty"`
		Criteria struct {
			Collection struct {
				Slug string `json:"slug"`
			} `json:"collection"`
			Contract struct {
				Address string `json:"address"`
			} `json:"contract"`
			Trait struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"trait"`
			EncodedTokenIds string `json:"encoded_token_ids"`
		} `json:"criteria,omitempty"`
		IsPrivateListing bool   `json:"is_private_listing,omitempty"`
		ClosingDate      int    `json:"closing_date,omitempty"`
		Seller           string `json:"seller,omitempty"`
		Buyer            string `json:"buyer,omitempty"`
		Transaction      string `json:"transaction,omitempty"`
		FromAddress      string `json:"from_address,omitempty"`
		ToAddress        string `json:"to_address,omitempty"`
	} `json:"asset_events"`
	Next string `json:"next"`
}
