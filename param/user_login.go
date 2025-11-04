package param

type LoginRequest struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

type LoginResponse struct {
	//User   UserInfo `json:"user"`
	Tokens Tokens `json:"tokens"`
}
