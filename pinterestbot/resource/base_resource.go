package resource

type Response struct {
	ResourceResponse  interface{} `json:"resource_response"`
	ClientContext     `json:"client_context"`
	Resource          `json:"resource"`
	RequestIdentifier string `json:"request_identifier"`
}

type Resource struct {
	Name    string  `json:"name"`
	Options Options `json:"options"`
}

type Options struct {
	Bookmarks   []string `json:"bookmarks,omitempty" url:"bookmarks,omitempty"`
	PageSize    int64    `json:"page_size,omitempty" url:"page_size,omitempty"`
	IsPrefetch  bool     `json:"isPrefetch,omitempty" url:"isPrefetch,omitempty"`
	Query       string   `json:"query,omitempty" url:"query,omitempty"`
	Scope       string   `json:"scope,omitempty" url:"scope,omitempty"`
	Count       string   `json:"count,omitempty" url:"count,omitempty"`
	FieldSetKey string   `json:"field_set_key,omitempty" url:"field_set_key,omitempty"`
	UserID      string   `json:"user_id,omitempty" url:"user_id,omitempty"`
}
