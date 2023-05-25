package tradingeventenum

type ReportScoreType int

const (
	_             ReportScoreType = iota
	RST_SELL_HERO                 // người bán hero
	RST_BUY_HERO                  // người mua hero
	RST_RENT_OUT                  // người cho thuê
	RST_RENT                      // người thuê
	RST_BUY_BOX                   // người mua box
)
