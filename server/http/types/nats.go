package types

type NatsAuthRequest struct {
	Token string `json:"token"`
}

type NatsAuthResponse struct {
	Publish   []string `json:"publish"`
	Subscribe []string `json:"subscribe"`
}
