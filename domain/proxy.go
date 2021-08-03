package domain

type Proxie struct{
	Type string `json:"type"`
	IP string `json:"ip"`
	Port string `json:"port"'`
	Login string `json:"login"`
	Password string `json:"password"`
	Status string `json:"status"`
	LastErrorMessage string `json:"last_error_message"`
	LastLimitMessage string `json:"last_limit_message`
}

