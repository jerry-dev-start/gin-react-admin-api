package req

type SsoLoginReq struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	Verifier string `json:"verifier"`
}

type SsoRedirectUrlReq struct {
	RedirectUrl string `json:"redirect_url" form:"redirect_url"`
}
