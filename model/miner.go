package model

type Miner struct {
	Id                  int    `json:"id"`
	MinerFid            string `json:"miner_fid"`
	BidMode             int    `json:"bid_mode"`
	ExpectedSealingTime int    `json:"expected_sealing_time"`
	StartEpoch          int    `json:"start_epoch"`
	AutoBidDealPerDay   int    `toml:"auto_bid_deal_per_day"`
}
