package app

type Balance struct {
	Value int `json:"Value"`
}

type Config struct {
	TGToken       string  `mapstructure:"TG_TOKEN"`
	CancelText    string  `mapstructure:"CANCEL_TEXT"`
	QueryText     string  `mapstructure:"QUERY_TEXT"`
	QueryResponse string  `mapstructure:"QUERY_RESPONSE"`
	ErrorResponse string  `mapstructure:"ERROR_RESPONSE"`
	UserPick      string  `mapstructure:"USER_PICK"`
	BotAdmins     []int64 `mapstructure:"BOT_ADMINS"`
}
