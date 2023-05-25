package location

import (
	"reflect"

	"github.com/samber/lo"
)

type Region int

const (
	RegionEuropean     Region = 0
	RegionUS           Region = 1
	RegionAsia         Region = 2
	RegionJapanese     Region = 3
	RegionAustralian   Region = 5
	RegionUSAWest      Region = 6
	RegionSouthAmerica Region = 7
	RegionCanadaEast   Region = 8
	RegionSouthKorea   Region = 9
	RegionIndia        Region = 10
	RegionRussia       Region = 11
	RegionRussiaEast   Region = 12
	RegionHongKong     Region = 13

	RegionNone Region = 4
)

var MapRegion2NearRegion = map[Region][]Region{
	RegionEuropean:     {RegionRussia},
	RegionUS:           {RegionSouthAmerica, RegionUSAWest, RegionCanadaEast},
	RegionAsia:         {RegionHongKong, RegionIndia},
	RegionJapanese:     {RegionSouthKorea, RegionAsia},
	RegionAustralian:   {RegionHongKong, RegionAsia},
	RegionUSAWest:      {RegionSouthAmerica, RegionUS, RegionCanadaEast},
	RegionSouthAmerica: {RegionUS, RegionUSAWest, RegionCanadaEast},
	RegionCanadaEast:   {RegionSouthAmerica, RegionUS, RegionUSAWest},
	RegionSouthKorea:   {RegionJapanese, RegionAsia},
	RegionIndia:        {RegionAsia, RegionHongKong},
	RegionRussia:       {RegionRussiaEast, RegionEuropean},
	RegionRussiaEast:   {RegionRussia, RegionSouthKorea, RegionJapanese},
	RegionHongKong:     {RegionAsia, RegionAustralian},
}

var MapRegion2Countries map[Region][]string

func GetRegionByCountryCode(countryCode string) Region {
	if region, found := CountryCodeRegion[countryCode]; found {
		return region
	}
	return defaultRegion
}

func RandomRegion() Region {
	regions := reflect.ValueOf(MapRegion2Countries).MapKeys()
	region := Region(lo.Sample(regions).Int())
	if region.IsValid() {
		return region
	}

	return defaultRegion
}

func (region Region) IsValid() bool {
	if len(MapRegion2Countries[region]) == 0 {
		return false
	}

	return true
}
