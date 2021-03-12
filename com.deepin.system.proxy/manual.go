package proxy

// scope proxy
type ScopeProxies struct {
	Proxies map[string][]Proxy `yaml:"proxies"` // map[http,sock4,sock5][]proxy
	// proxy setting
	ProxyProgram   []string `yaml:"proxy-program"`    // global proxy will ignore
	NoProxyProgram []string `yaml:"no-proxy-program"` // app proxy will ignore

	// white list
	WhiteList []string `yaml:"whitelist"` // white site dont use proxy, not use this time
	TPort     int      `yaml:"t-port"`
}

// proxy type
type Proxy struct {
	// proxy proto type
	ProtoType string `json:"type"` // http sock4 sock5

	// [proto]&[name] as ident
	Name string `yaml:"name"`

	// proxy server
	Server string `yaml:"server"`
	Port   int    `yaml:"port"`

	// auth message
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}
