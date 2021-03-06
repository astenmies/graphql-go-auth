package authUtils

// Config :
// - Configures http.Cookie creation.
type Config struct {
	// Name is the desired cookie name.
	Name string
	// Domain sets the cookie domain. Defaults to the host name of the responding
	// server when left zero valued.
	Domain string
	// Path sets the cookie path. Defaults to the path of the URL responding to
	// the request when left zero valued.
	Path string
	// MaxAge=0 means no 'Max-Age' attribute should be set.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	// Cookie 'Expires' will be set (or left unset) according to MaxAge
	MaxAge int
	// HTTPOnly indicates whether the browser should prohibit a cookie from
	// being accessible via Javascript. Recommended true.
	HTTPOnly bool
	// Secure flag indicating to the browser that the cookie should only be
	// transmitted over a TLS HTTPS connection. Recommended true in production.
	Secure bool
	// TriggerMutation is the mutation that triggers Oauth
	TriggerMutation string
}

// DefaultAuthConfig :
// - Configures short-lived temporary http.Cookie creation
var DefaultAuthConfig = &Config{
	Name:            "graphql-go-auth-temp-cookie",
	Path:            "/",
	MaxAge:          60, // 60 seconds
	HTTPOnly:        true,
	Secure:          true,           // HTTPS only
	TriggerMutation: "triggerOauth", // the mutation that triggers Oauth
}

// DebuggingAuthConfig :
// - Configures creation of short-lived temporary http.Cookie's
// - Does NOT require cookies be sent over HTTPS!
// - Use this config for development only
var DebuggingAuthConfig = &Config{
	Name:            "graphql-go-auth-temp-debug-cookie",
	Path:            "/",
	MaxAge:          60, // 60 seconds
	HTTPOnly:        true,
	Secure:          false,          // allows cookies to be send over HTTP
	TriggerMutation: "triggerOauth", // the mutation that triggers Oauth
}
