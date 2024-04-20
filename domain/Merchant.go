package domain

type Merchant struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
