package pin

import (
	"github.com/orangeapps/pinterest-bot/pinterestbot/resource"
)

type Response struct {
	ResourceResponse  ResourceResponse       `json:"resource_response"`
	ClientContext     resource.ClientContext `json:"client_context"`
	Resource          resource.Resource      `json:"resource"`
	RequestIdentifier string                 `json:"request_identifier"`
}

type SearchResponse struct {
	ResourceResponse  SearchResourceResponse `json:"resource_response"`
	ClientContext     resource.ClientContext `json:"client_context"`
	Resource          resource.Resource      `json:"resource"`
	RequestIdentifier string                 `json:"request_identifier"`
}

type ResourceResponse struct {
	Status         string `json:"status"`
	Code           int64  `json:"code"`
	Data           Data   `json:"data"`
	Message        string `json:"message"`
	DisplayShoptab bool   `json:"display_shoptab"`
	Terms          []Term `json:"terms"`
	EndpointName   string `json:"endpoint_name"`
	Bookmark       string `json:"bookmark"`
	HTTPStatus     int64  `json:"http_status"`
}

type SearchResourceResponse struct {
	Status       string    `json:"status"`
	Code         int64     `json:"code"`
	Data         []InfoPin `json:"data"`
	Message      string    `json:"message"`
	SearchNag    SearchNag `json:"search_nag"`
	EndpointName string    `json:"endpoint_name"`
	Bookmark     string    `json:"bookmark"`
	HTTPStatus   int64     `json:"http_status"`
}

type Data struct {
	Nag                      Nag       `json:"nag"`
	ShouldAppendGlobalSearch bool      `json:"should_append_global_search"`
	Results                  []InfoPin `json:"results"`
	Guides                   []Term    `json:"guides"`
	DisplayShopTab           bool      `json:"displayShopTab"`
	NoGiftWrap               bool      `json:"no_gift_wrap"`
}

type Term struct {
	Term          string `json:"term"`
	Display       string `json:"display"`
	Position      int64  `json:"position"`
	DominantColor string `json:"dominant_color"`
	Source        int64  `json:"source"`
	WebBlurredURL string `json:"web_blurred_url"`
}

type SearchNag struct {
	Nag Nag `json:"nag"`
}

type Nag struct {
}

type InfoPin struct {
	BuyableProduct                  interface{}       `json:"buyable_product"`
	VideoStatusMessage              interface{}       `json:"video_status_message"`
	DoneByMe                        bool              `json:"done_by_me"`
	ShoppingFlags                   []interface{}     `json:"shopping_flags"`
	IsRepin                         bool              `json:"is_repin"`
	PriceCurrency                   string            `json:"price_currency"`
	HasRequiredAttributionProvider  bool              `json:"has_required_attribution_provider"`
	LinkDomain                      LinkDomain        `json:"link_domain"`
	IsPromotable                    bool              `json:"is_promotable"`
	Method                          string            `json:"method"`
	IsPromoted                      bool              `json:"is_promoted"`
	RichRecipeTopIngredients        []interface{}     `json:"rich_recipe_top_ingredients"`
	CreatorAnalytics                interface{}       `json:"creator_analytics"`
	Embed                           interface{}       `json:"embed"`
	Category                        string            `json:"category"`
	BoardActivity                   interface{}       `json:"board_activity"`
	CloseupUserNote                 string            `json:"closeup_user_note"`
	OriginPinner                    interface{}       `json:"origin_pinner"`
	ContentSensitivity              Ity               `json:"content_sensitivity"`
	RichMetadata                    RichMetadata      `json:"rich_metadata"`
	Link                            string            `json:"link"`
	NativeCreator                   interface{}       `json:"native_creator"`
	CloseupUnifiedDescription       string            `json:"closeup_unified_description"`
	StoryPinData                    interface{}       `json:"story_pin_data"`
	CloseupDescription              string            `json:"closeup_description"`
	PerPinAnalytics                 PerPinAnalytics   `json:"per_pin_analytics"`
	TrackingParams                  string            `json:"tracking_params"`
	Description                     string            `json:"description"`
	NativePinStats                  interface{}       `json:"native_pin_stats"`
	Pinner                          Pinner            `json:"pinner"`
	StoryPinDataID                  interface{}       `json:"story_pin_data_id"`
	Promoter                        interface{}       `json:"promoter"`
	DominantColor                   string            `json:"dominant_color"`
	CommentCount                    int64             `json:"comment_count"`
	IsQuickPromotableByPinner       bool              `json:"is_quick_promotable_by_pinner"`
	Type                            string            `json:"type"`
	Images                          map[string]Image  `json:"images"`
	AggregatedPinData               AggregatedPinData `json:"aggregated_pin_data"`
	CommentsDisabled                bool              `json:"comments_disabled"`
	DidItDisabled                   bool              `json:"did_it_disabled"`
	CanDeleteDidItAndComments       bool              `json:"can_delete_did_it_and_comments"`
	TrackedLink                     string            `json:"tracked_link"`
	Board                           Board             `json:"board"`
	VideoStatus                     interface{}       `json:"video_status"`
	Access                          []interface{}     `json:"access"`
	ImageSignature                  string            `json:"image_signature"`
	GridTitle                       string            `json:"grid_title,omitempty"`
	CarouselData                    interface{}       `json:"carousel_data"`
	IsEligibleForAggregatedComments bool              `json:"is_eligible_for_aggregated_comments"`
	Domain                          string            `json:"domain"`
	Videos                          interface{}       `json:"videos"`
	ReactionCounts                  ReactionCounts    `json:"reaction_counts"`
	ProductPinData                  interface{}       `json:"product_pin_data"`
	ViaPinner                       interface{}       `json:"via_pinner"`
	Hashtags                        []interface{}     `json:"hashtags"`
	Title                           string            `json:"title,omitempty"`
	ID                              string            `json:"id"`
	PromotedIsRemovable             bool              `json:"promoted_is_removable"`
	Privacy                         string            `json:"privacy"`
	CreatedAt                       string            `json:"created_at"`
	ThirdPartyPinOwner              interface{}       `json:"third_party_pin_owner"`
	EditedFields                    []string          `json:"edited_fields"`
	BuyableProductAvailability      interface{}       `json:"buyable_product_availability"`
	DescriptionHTML                 string            `json:"description_html"`
	IsNative                        bool              `json:"is_native"`
	UnifiedUserNote                 string            `json:"unified_user_note"`
	MobileLink                      string            `json:"mobile_link"`
	IsEligibleForBrandCatalog       bool              `json:"is_eligible_for_brand_catalog"`
	IsVideo                         bool              `json:"is_video"`
	IsWhitelistedForTriedIt         bool              `json:"is_whitelisted_for_tried_it"`
	Attribution                     interface{}       `json:"attribution"`
	IsHidden                        bool              `json:"is_hidden"`
	IsQuickPromotable               bool              `json:"is_quick_promotable"`
	PriceValue                      float64           `json:"price_value"`
	Section                         interface{}       `json:"section"`
	RepinCount                      int64             `json:"repin_count"`
	IsPlayable                      bool              `json:"is_playable"`
}

type AggregatedPinData struct {
	DidItData        DidItData       `json:"did_it_data"`
	ID               string          `json:"id"`
	AggregatedStats  AggregatedStats `json:"aggregated_stats"`
	CreatorAnalytics interface{}     `json:"creator_analytics"`
	IsShopTheLook    bool            `json:"is_shop_the_look"`
	CommentCount     int64           `json:"comment_count"`
}

type AggregatedStats struct {
	Saves int64 `json:"saves"`
	Done  int64 `json:"done"`
}

type DidItData struct {
	Tags             []interface{}    `json:"tags"`
	VideosCount      int64            `json:"videos_count"`
	Type             string           `json:"type"`
	UserCount        int64            `json:"user_count"`
	RecommendedCount int64            `json:"recommended_count"`
	ResponsesCount   int64            `json:"responses_count"`
	Rating           int64            `json:"rating"`
	RecommendScores  []RecommendScore `json:"recommend_scores"`
	ImagesCount      int64            `json:"images_count"`
	DetailsCount     int64            `json:"details_count"`
}

type RecommendScore struct {
	Score float64 `json:"score"`
	Count int64   `json:"count"`
}

type Board struct {
	Name                    string          `json:"name"`
	URL                     string          `json:"url"`
	IsCollaborative         bool            `json:"is_collaborative"`
	Category                interface{}     `json:"category"`
	MapID                   string          `json:"map_id"`
	Layout                  string          `json:"layout"`
	Type                    string          `json:"type"`
	ID                      string          `json:"id"`
	Access                  []interface{}   `json:"access"`
	ShouldShowBoardActivity bool            `json:"should_show_board_activity"`
	CollaboratedByMe        bool            `json:"collaborated_by_me"`
	Owner                   PerPinAnalytics `json:"owner"`
	PinThumbnailUrls        []string        `json:"pin_thumbnail_urls"`
	Description             string          `json:"description"`
	ImageCoverURL           string          `json:"image_cover_url"`
	ImageThumbnailURL       string          `json:"image_thumbnail_url"`
	FollowedByMe            bool            `json:"followed_by_me"`
	Privacy                 string          `json:"privacy"`
}

type PerPinAnalytics struct {
	ID string `json:"id"`
}

type Ity struct {
}

type Image struct {
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
}

type LinkDomain struct {
	OfficialUser interface{} `json:"official_user"`
}

type Pinner struct {
	ImageMediumURL         string      `json:"image_medium_url"`
	ExplicitlyFollowedByMe bool        `json:"explicitly_followed_by_me"`
	Indexed                bool        `json:"indexed"`
	Type                   string      `json:"type"`
	DomainURL              interface{} `json:"domain_url"`
	ID                     string      `json:"id"`
	Location               string      `json:"location"`
	Username               string      `json:"username"`
	VerifiedIdentity       Ity         `json:"verified_identity"`
	DomainVerified         bool        `json:"domain_verified"`
	BlockedByMe            bool        `json:"blocked_by_me"`
	FullName               string      `json:"full_name"`
	ImageSmallURL          string      `json:"image_small_url"`
	IsDefaultImage         bool        `json:"is_default_image"`
	FirstName              string      `json:"first_name"`
	FollowedByMe           bool        `json:"followed_by_me"`
}

type ReactionCounts struct {
	The1 int64 `json:"1"`
}

type RichMetadata struct {
	HasPriceDrop         bool        `json:"has_price_drop"`
	URL                  string      `json:"url"`
	Title                string      `json:"title"`
	Tracker              interface{} `json:"tracker"`
	Article              Article     `json:"article"`
	Type                 string      `json:"type"`
	ID                   string      `json:"id"`
	AppleTouchIconImages IconImages  `json:"apple_touch_icon_images"`
	Description          string      `json:"description"`
	CanonicalURL         interface{} `json:"canonical_url"`
	Locale               string      `json:"locale"`
	AmpValid             bool        `json:"amp_valid"`
	AmpURL               string      `json:"amp_url"`
	FaviconImages        IconImages  `json:"favicon_images"`
	LinkStatus           int64       `json:"link_status"`
	SiteName             string      `json:"site_name"`
	FaviconLink          string      `json:"favicon_link"`
	AppleTouchIconLink   string      `json:"apple_touch_icon_link"`
}

type IconImages struct {
	Orig   string `json:"orig"`
	The50X string `json:"50x"`
}

type Article struct {
	Name          string        `json:"name"`
	Type          string        `json:"type"`
	ID            string        `json:"id"`
	Authors       []interface{} `json:"authors"`
	DatePublished string        `json:"date_published"`
	Description   string        `json:"description"`
}
