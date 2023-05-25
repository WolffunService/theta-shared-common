package esmodel

type WalletStatMapping struct {
	User UserModel `json:"user"`
	THC  int64     `json:"thc"`
	THG  int64     `json:"thg"`
}
