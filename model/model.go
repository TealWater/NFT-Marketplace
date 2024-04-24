package model

// StreamHeartBeat This struct is for creating a hear-beat event to keep the steam alive
type StreamHeartBeat struct {
	Topic   string `json:"topic"`
	Event   string `json:"event"`
	Payload struct {
	} `json:"payload"`
	Ref int `json:"ref"`
}
