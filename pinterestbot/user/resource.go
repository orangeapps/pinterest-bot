package user

import "github.com/jezerdave/go-pinterest-bot/pinterestbot/resource"

type Response struct {
	ResourceResponse  ResourceResponse       `json:"resource_response"`
	ClientContext     resource.ClientContext `json:"client_context"`
	Resource          resource.Resource      `json:"resource"`
	RequestIdentifier string                 `json:"request_identifier"`
}

type ResourceResponse struct {
	Status       string `json:"status"`
	Code         int64  `json:"code"`
	Data         Data   `json:"data"`
	Message      string `json:"message"`
	EndpointName string `json:"endpoint_name"`
	HTTPStatus   int64  `json:"http_status"`
}

//Data
type Data struct {
	Location                    string           `json:"location"`
	HasCustomBoardSortingOrder  bool             `json:"has_custom_board_sorting_order"`
	IsVerifiedMerchant          bool             `json:"is_verified_merchant"`
	Partner                     Partner          `json:"partner"`
	PinsDoneCount               int64            `json:"pins_done_count"`
	ImageXlargeURL              string           `json:"image_xlarge_url"`
	ProfileReach                int64            `json:"profile_reach"`
	StoryPinCount               int64            `json:"story_pin_count"`
	DomainURL                   string           `json:"domain_url"`
	ExplicitlyFollowedByMe      bool             `json:"explicitly_followed_by_me"`
	IsPartner                   bool             `json:"is_partner"`
	FollowingCount              int64            `json:"following_count"`
	LastPinSaveTime             string           `json:"last_pin_save_time"`
	ID                          string           `json:"id"`
	Indexed                     bool             `json:"indexed"`
	ShowCreatorProfile          bool             `json:"show_creator_profile"`
	NoindexReason               string           `json:"noindex_reason"`
	StorefrontManagementEnabled bool             `json:"storefront_management_enabled"`
	StorefrontSearchEnabled     bool             `json:"storefront_search_enabled"`
	BlockedByMe                 bool             `json:"blocked_by_me"`
	Type                        string           `json:"type"`
	VideoPinCount               int64            `json:"video_pin_count"`
	HasShoppingShowcase         interface{}      `json:"has_shopping_showcase"`
	LastName                    string           `json:"last_name"`
	NativePinCount              int64            `json:"native_pin_count"`
	FullName                    string           `json:"full_name"`
	GroupBoardCount             int64            `json:"group_board_count"`
	ShowDiscoveredFeed          bool             `json:"show_discovered_feed"`
	BoardCount                  int64            `json:"board_count"`
	HasCatalog                  bool             `json:"has_catalog"`
	VerifiedIdentity            VerifiedIdentity `json:"verified_identity"`
	FollowerCount               int64            `json:"follower_count"`
	HasShowcase                 bool             `json:"has_showcase"`
	About                       string           `json:"about"`
	PinCount                    int64            `json:"pin_count"`
	ImageLargeURL               string           `json:"image_large_url"`
	ImageSmallURL               string           `json:"image_small_url"`
	JoinedCommunitiesCount      int64            `json:"joined_communities_count"`
	ProfileDiscoveredPublic     bool             `json:"profile_discovered_public"`
	ShowImpressum               bool             `json:"show_impressum"`
	DomainVerified              bool             `json:"domain_verified"`
	ImpressumURL                interface{}      `json:"impressum_url"`
	WebsiteURL                  string           `json:"website_url"`
	Username                    string           `json:"username"`
	ProfileCover                ProfileCover     `json:"profile_cover"`
	FirstName                   string           `json:"first_name"`
	ImageMediumURL              string           `json:"image_medium_url"`
	IsTastemaker                bool             `json:"is_tastemaker"`
}

type Partner struct {
	ProfilePlace         interface{}         `json:"profile_place"`
	EnableProfileMessage bool                `json:"enable_profile_message"`
	EnableProfileAddress bool                `json:"enable_profile_address"`
	ContactPhone         string              `json:"contact_phone"`
	ContactPhoneCountry  ContactPhoneCountry `json:"contact_phone_country"`
	ContactEmail         interface{}         `json:"contact_email"`
}

type ContactPhoneCountry struct {
	Code      string `json:"code"`
	PhoneCode string `json:"phone_code"`
}

type ProfileCover struct {
	Video    interface{} `json:"video"`
	SourceID interface{} `json:"source_id"`
	Images   interface{} `json:"images"`
	Source   string      `json:"source"`
	Type     string      `json:"type"`
}

type VerifiedIdentity struct {
}
