package rivalitem

// ItemAppeared of field, group and selector
//
// Example:
// DAILY_DEALS:
//	full-evolve_half-evolve:
//		full-evolve: 80
// 		half-evolve: 20
type ItemAppeared map[string]map[string]map[string]int
