package rules

type Special struct {
	Name                string `json:"name,omitempty"`
	RedealsAllowed      int    `json:"redealsAllowed,omitempty"`
	RotationsAllowed    int    `json:"rotationsAllowed,omitempty"`
	RotationTopToBottom bool   `json:"rotationTopToBottom,omitempty"`
	DrawsAllowed        int    `json:"drawsAllowed,omitempty"`
	DrawsAfterRedeals   bool   `json:"drawsAfterRedeals,omitempty"`
}
