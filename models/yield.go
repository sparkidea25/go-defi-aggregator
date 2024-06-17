package models

type Yield struct {
	ID           string  `json:"id"`
	Pool         string  `json:"pool"`
	Chain        string  `json:"chain"`
	Project      string  `json:"project"`
	APY          float64 `json:"apy"`
	TvlUsd       float64 `json:"tvlUsd"`
	Stablecoin   bool    `json:"stablecoin"`
	RewardTokens string  `json:"rewardTokens"`
	// Add other fields as needed
}
