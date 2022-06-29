package api

// Follows Rasa NLG Server API https://rasa.com/docs/rasa/nlg/

type rasaNlgRequest struct {
	Response  string                 `json:"response"`
	Arguments map[string]interface{} `json:"arguments"`
	Tracker   map[string]interface{} `json:"tracker"`
	Channel   string                 `json:"channel"`
}

type rasaNlgResponse struct {
	Text string `json:"text"`
}
