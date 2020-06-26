package resource

import "fmt"

//const SCOPES  = [1]string{"pins"}

const UrlBase = "https://www.pinterest.com/";
const UrlSearch = "resource/BaseSearchResource/get/"
const UrlSearchWithPagination = "resource/SearchResource/get/"
const UrlUserInfo = "resource/UserResource/get/"
const UrlPinInfo = "resource/PinResource/get/"

func BuildSourceUrl(query string, scope string) string {
	return fmt.Sprintf("/search/%s/?q=%s&rs=typo_auto_original&auto_correction_disabled=true", scope, query)
}