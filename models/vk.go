package models

type ProfileInfo struct{
	FullName string `json:"fullname"`
	LastName string `json:"last_name"`
	FirstName string `json:"first_name"`
	Avatar string `json:"avatar"`
	Friends []string `json:"friends"`
	Subscriptions []string `json:"subscriptions"`
	Photos []string `json:"photos"`
	Posts []string `json:"posts"`
	Status string `json:"status"`
}

type PostInfo struct {
	Type string `json:"type"`
	ProfileID string `json:"profile_id"`
	PhotoURLs []string  `json:"photo_urls"`
	Text string `json:"text"`
	Likes []string `json:"likes"`
	Comments []string `json:"comments"`
}

type Account struct{
	Token string `json:"token"`
	Status string `json:"status"`
	LastErrorMessage string `json:"last_error_message"`
}




