package utils

import (
	"net/url"

	"google.golang.org/api/firebasedynamiclinks/v1"
)

type GenShortLinkRequest struct {
	Data                  interface{}
	ThetanDomainLink      string
	ThetanDomainUriPrefix string
	EnableAndroidInfo     bool
	AndroidPackageName    string
	EnableIOSInfo         bool
	IOSBundleID           string
	IOSAppStoreID         string
	SocialMediaInfo       *firebasedynamiclinks.SocialMetaTagInfo
}

func genLink(domainLink string, data interface{}) string {
	base, err := url.Parse(domainLink)
	if err != nil {
		return ""
	}
	base.RawQuery = ConvertRawQuery(data)
	return base.String()
}

func GenDynamicLinkInfo(genShortLinkRequest GenShortLinkRequest) *firebasedynamiclinks.DynamicLinkInfo {
	data := &firebasedynamiclinks.DynamicLinkInfo{
		DomainUriPrefix: genShortLinkRequest.ThetanDomainUriPrefix,
		Link:            genLink(genShortLinkRequest.ThetanDomainLink, genShortLinkRequest.Data),
	}

	if genShortLinkRequest.EnableAndroidInfo {
		data.AndroidInfo = &firebasedynamiclinks.AndroidInfo{
			AndroidPackageName: genShortLinkRequest.AndroidPackageName,
		}
	}

	if genShortLinkRequest.EnableIOSInfo {
		data.IosInfo = &firebasedynamiclinks.IosInfo{
			IosBundleId:   genShortLinkRequest.IOSBundleID,
			IosAppStoreId: genShortLinkRequest.IOSAppStoreID,
		}
	}

	if genShortLinkRequest.SocialMediaInfo != nil {
		data.SocialMetaTagInfo = genShortLinkRequest.SocialMediaInfo
	}

	return data
}
