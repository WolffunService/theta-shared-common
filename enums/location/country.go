package location

type Country int

const defaultRegion = RegionSouthAmerica

var CountryCodes []string
var CountryCodeRegion map[string]Region

func init() {
	CountryCodes = []string{
		"AR",
		"AU",
		"BR",
		"CA",
		"ES",
		"FR",
		"HK",
		"ID",
		"IN",
		"IR",
		"JP",
		"KR",
		"PH",
		"RU",
		"TH",
		"TW",
		"US",
		"VN",
	}

	CountryCodeRegion = make(map[string]Region)
	CountryCodeRegion["AR"] = RegionSouthAmerica
	CountryCodeRegion["AU"] = RegionAustralian
	CountryCodeRegion["BR"] = RegionSouthAmerica
	CountryCodeRegion["CA"] = RegionCanadaEast
	CountryCodeRegion["ES"] = RegionEuropean
	CountryCodeRegion["FR"] = RegionEuropean
	CountryCodeRegion["HK"] = RegionHongKong
	CountryCodeRegion["ID"] = RegionAsia
	CountryCodeRegion["IN"] = RegionIndia
	CountryCodeRegion["IR"] = RegionEuropean
	CountryCodeRegion["JP"] = RegionJapanese
	CountryCodeRegion["KR"] = RegionSouthKorea
	CountryCodeRegion["PH"] = RegionAsia
	CountryCodeRegion["RU"] = RegionRussia
	CountryCodeRegion["TH"] = RegionAsia
	CountryCodeRegion["TW"] = RegionHongKong
	CountryCodeRegion["US"] = RegionUS
	CountryCodeRegion["VN"] = RegionAsia

	MapRegion2Countries = make(map[Region][]string)
	for countryCode, region := range CountryCodeRegion {
		MapRegion2Countries[region] = append(MapRegion2Countries[region], countryCode)
	}
}
