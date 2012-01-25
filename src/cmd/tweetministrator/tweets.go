package main

type Tweet struct {
	Identifier string `json:"id_str"`
	Date string `json:"created_at"`
	Text string `json:"text"`
}

func (t *Tweet) UnixTimeStamp() int64 {
	return 0
}
