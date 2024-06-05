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

type TopOpenSeaNFTCollections struct {
	Collections []struct {
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
	} `json:"collections"`
	Next string `json:"next"`
}

type OpenSeaNFT struct {
	Nfts []struct {
		Identifier    string  `json:"identifier"`
		Collection    string  `json:"collection"`
		Contract      string  `json:"contract"`
		TokenStandard string  `json:"token_standard"`
		Name          string  `json:"name"`
		Description   string  `json:"description"`
		ImageURL      string  `json:"image_url"`
		MetadataURL   string  `json:"metadata_url"`
		OpenseaURL    string  `json:"opensea_url"`
		UpdatedAt     string  `json:"updated_at"`
		IsDisabled    bool    `json:"is_disabled"`
		IsNsfw        bool    `json:"is_nsfw"`
		Price         float64 `json:"price"`
		Currency      string  `json:"currency"`
	} `json:"nfts"`
	Next string `json:"next"`
}

type OpenSeaNFTListing struct {
	Listings []struct {
		OrderHash string `json:"order_hash"`
		Type      string `json:"type"`
		Price     struct {
			Current struct {
				Currency string `json:"currency"`
				Decimals int    `json:"decimals"`
				Value    string `json:"value"`
			} `json:"current"`
		} `json:"price"`
		ProtocolData struct {
			Parameters struct {
				Offerer string `json:"offerer"`
				Offer   []struct {
					ItemType             int    `json:"itemType"`
					Token                string `json:"token"`
					IdentifierOrCriteria string `json:"identifierOrCriteria"`
					StartAmount          string `json:"startAmount"`
					EndAmount            string `json:"endAmount"`
				} `json:"offer"`
				Consideration []struct {
					ItemType             int    `json:"itemType"`
					Token                string `json:"token"`
					IdentifierOrCriteria string `json:"identifierOrCriteria"`
					StartAmount          string `json:"startAmount"`
					EndAmount            string `json:"endAmount"`
					Recipient            string `json:"recipient"`
				} `json:"consideration"`
				StartTime                       string      `json:"startTime"`
				EndTime                         string      `json:"endTime"`
				OrderType                       int         `json:"orderType"`
				Zone                            string      `json:"zone"`
				ZoneHash                        string      `json:"zoneHash"`
				Salt                            string      `json:"salt"`
				ConduitKey                      string      `json:"conduitKey"`
				TotalOriginalConsiderationItems int         `json:"totalOriginalConsiderationItems"`
				Counter                         interface{} `json:"counter"`
			} `json:"parameters"`
			Signature string `json:"signature"`
		} `json:"protocol_data"`
		ProtocolAddress string `json:"protocol_address"`
	} `json:"listings"`
	Next string `json:"next"`
}

type SingleOpenSeaNFT struct {
	Nft struct {
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
		AnimationURL  any    `json:"animation_url"`
		IsSuspicious  bool   `json:"is_suspicious"`
		Creator       string `json:"creator"`
		Traits        []struct {
			TraitType   string `json:"trait_type"`
			DisplayType any    `json:"display_type"`
			MaxValue    any    `json:"max_value"`
			Value       any    `json:"value"`
		} `json:"traits"`
		Owners []struct {
			Address  string `json:"address"`
			Quantity int    `json:"quantity"`
		} `json:"owners"`
		Rarity struct {
			StrategyID      any    `json:"strategy_id"`
			StrategyVersion any    `json:"strategy_version"`
			Rank            int    `json:"rank"`
			Score           any    `json:"score"`
			CalculatedAt    string `json:"calculated_at"`
			MaxRank         any    `json:"max_rank"`
			TokensScored    int    `json:"tokens_scored"`
			RankingFeatures any    `json:"ranking_features"`
		} `json:"rarity"`
	} `json:"nft"`
}
