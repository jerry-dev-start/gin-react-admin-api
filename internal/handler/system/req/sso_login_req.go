package req

type SsoLoginReq struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	Verifier string `json:"verifier"`
}
