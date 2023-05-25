package wallettypeenum

type WalletType string

const (
	META_MASK      WalletType = "Metamask"
	COIN_BASE      WalletType = "Coinbase"
	WALLET_CONNECT WalletType = "WalletConnect"
	SEQUENCE       WalletType = "Sequence"
)

func (w WalletType) IsValid() bool {
	switch w {
	case META_MASK, COIN_BASE, WALLET_CONNECT, SEQUENCE:
		return true
	}
	return false
}
