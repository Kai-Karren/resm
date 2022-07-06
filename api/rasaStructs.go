package api

type Tracker struct {
	SenderId      string            `json:"sender_id"`
	Slots         map[string]string `json:"slots"`
	LatestMessage LatestMessage     `json:"latest_message"`
	Events        []Event           `json:"events"`
}

type LatestMessage struct {
	MessageId     string                 `json:"message_id"`
	Intent        Intent                 `json:"intent"`
	Entities      []interface{}          `json:"entities"`
	Text          string                 `json:"text"`
	Metadata      map[string]interface{} `json:"metadata"`
	IntentRanking []Intent               `json:"intent_ranking"`
}

type Intent struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Confidence float32 `json:"confidence"`
}

type Channel struct {
	Name string `json:"name"`
}

type Event struct {
	Event        string                 `json:"event"`
	Timestamp    float64                `json:"timestamp"`
	Name         string                 `json:"name"`
	Policy       string                 `json:"policy"`
	Confidence   float64                `json:"confidence"`
	ParseData    map[string]interface{} `json:"parse_data"`
	InputChannel string                 `json:"input_channel"`
	MessageId    string                 `json:"message_id"`
	Metadata     map[string]interface{} `json:"metadata"`
}
