package schemas

type LoginFormRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type ClientRegiterResponse struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
