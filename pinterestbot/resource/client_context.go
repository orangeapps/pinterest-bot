package resource

type ClientContext struct {
	ActiveExperiments                Experiments        `json:"active_experiments"`
	AppType                          int64              `json:"app_type"`
	AppVersion                       string             `json:"app_version"`
	BrowserLocale                    string             `json:"browser_locale"`
	BrowserName                      string             `json:"browser_name"`
	BrowserType                      int64              `json:"browser_type"`
	BrowserVersion                   string             `json:"browser_version"`
	Country                          string             `json:"country"`
	CSPNonce                         string             `json:"csp_nonce"`
	CurrentURL                       string             `json:"current_url"`
	DeepLinkAction                   interface{}        `json:"deep_link_action"`
	DeepLinkDefault                  string             `json:"deep_link_default"`
	DeepLink                         string             `json:"deep_link"`
	EnabledAdvertiserCountries       []string           `json:"enabled_advertiser_countries"`
	FacebookToken                    interface{}        `json:"facebook_token"`
	FromOpenInAppClick               interface{}        `json:"from_open_in_app_click"`
	HasSterlingOnSteroidsWriteAccess bool               `json:"has_sterling_on_steroids_write_access"`
	HTTPReferrer                     string             `json:"http_referrer"`
	InMobileForkExp                  bool               `json:"in_mobile_fork_exp"`
	InviteCode                       interface{}        `json:"invite_code"`
	InviteSenderID                   string             `json:"invite_sender_id"`
	InviteSenderExperiments          Experiments        `json:"invite_sender_experiments"`
	IsAmp                            bool               `json:"is_amp"`
	IsAuthenticated                  bool               `json:"is_authenticated"`
	IsBot                            string             `json:"is_bot"`
	IsCarouselCampaignWhitelisted    bool               `json:"is_carousel_campaign_whitelisted"`
	IsDataSaver                      bool               `json:"is_data_saver"`
	IsInternalIP                     bool               `json:"is_internal_ip"`
	IsFullPage                       bool               `json:"is_full_page"`
	IsMobileAgent                    bool               `json:"is_mobile_agent"`
	IsMobileFork                     bool               `json:"is_mobile_fork"`
	IsRetina                         bool               `json:"is_retina"`
	IsSterlingOnSteroids             bool               `json:"is_sterling_on_steroids"`
	IsTabletAgent                    bool               `json:"is_tablet_agent"`
	Language                         string             `json:"language"`
	Locale                           string             `json:"locale"`
	MdlSchemeWhitelist               []string           `json:"mdl_scheme_whitelist"`
	MockDate                         interface{}        `json:"mock_date"`
	NagBrowserUnsupported            bool               `json:"nag_browser_unsupported"`
	Origin                           string             `json:"origin"`
	Path                             string             `json:"path"`
	PureReact                        bool               `json:"pure_react"`
	ReactOnly                        interface{}        `json:"react_only"`
	Referrer                         interface{}        `json:"referrer"`
	RequestHost                      string             `json:"request_host"`
	SEOExperiments                   map[string]*string `json:"seo_experiments"`
	SocialBot                        string             `json:"social_bot"`
	SiteType                         int64              `json:"site_type"`
	ShouldUseSterlingToken           bool               `json:"should_use_sterling_token"`
	SterlingOnSteroidsAccessToken    interface{}        `json:"sterling_on_steroids_access_token"`
	SterlingOnSteroidsLDAP           interface{}        `json:"sterling_on_steroids_ldap"`
	TriggerableExperiments           map[string]string  `json:"triggerable_experiments"`
	UnauthID                         string             `json:"unauth_id"`
	SEOExpID                         interface{}        `json:"seo_exp_id"`
	UserAgentCanUseNativeApp         bool               `json:"user_agent_can_use_native_app"`
	UserAgentIsIos9_OrAbove          bool               `json:"user_agent_is_ios_9_or_above"`
	UserAgentPlatform                string             `json:"user_agent_platform"`
	UserAgentPlatformVersion         interface{}        `json:"user_agent_platform_version"`
	UserAgent                        string             `json:"user_agent"`
	User                             UserInfo           `json:"user"`
	UtmCampaign                      interface{}        `json:"utm_campaign"`
	UtmMedium                        interface{}        `json:"utm_medium"`
	UtmSource                        interface{}        `json:"utm_source"`
	UtmTerm                          interface{}        `json:"utm_term"`
	UtmPai                           interface{}        `json:"utm_pai"`
	VisibleURL                       string             `json:"visible_url"`
	AnalysisUa                       AnalysisUa         `json:"analysis_ua"`
	RequestIdentifier                string             `json:"request_identifier"`
	RootRequestIdentifier            interface{}        `json:"root_request_identifier"`
	ParentRequestIdentifier          interface{}        `json:"parent_request_identifier"`
	FullPath                         string             `json:"full_path"`
	RealIP                           string             `json:"real_ip"`
	PlacedExperiences                interface{}        `json:"placed_experiences"`
	BatchExp                         bool               `json:"batch_exp"`
	ExperimentHash                   string             `json:"experiment_hash"`
}

type Experiments struct {
}

type AnalysisUa struct {
	OSName         string      `json:"os_name"`
	OSVersion      string      `json:"os_version"`
	BrowserName    string      `json:"browser_name"`
	BrowserVersion string      `json:"browser_version"`
	Device         string      `json:"device"`
	AppType        int64       `json:"app_type"`
	DeviceType     interface{} `json:"device_type"`
}

type UserInfo struct {
	UnauthID string `json:"unauth_id"`
}