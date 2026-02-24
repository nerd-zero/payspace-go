package payspace

import (
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Option configures the Client.
type Option func(*clientConfig)

type clientConfig struct {
	environment   Environment
	apiBaseURL    string
	identityURL   string
	scope         Scope
	httpClient    *http.Client
	logger        *zap.Logger
	logLevel      zapcore.Level
	timeout       time.Duration
	maxRetries    int
	retryBackoff  time.Duration
	rateLimitWait bool
	apiVersion    string
	userAgent     string
}

func defaultConfig() *clientConfig {
	return &clientConfig{
		environment:   Staging,
		scope:         ScopeFullAccess,
		logLevel:      zapcore.InfoLevel,
		timeout:       defaultTimeout,
		maxRetries:    3,
		retryBackoff:  500 * time.Millisecond,
		rateLimitWait: true,
		apiVersion:    defaultAPIVer,
		userAgent:     defaultUserAgent,
	}
}

// WithEnvironment sets the API environment (Staging or Production).
func WithEnvironment(env Environment) Option {
	return func(c *clientConfig) { c.environment = env }
}

// WithAPIBaseURL overrides the API base URL.
func WithAPIBaseURL(url string) Option {
	return func(c *clientConfig) { c.apiBaseURL = url }
}

// WithIdentityURL overrides the identity/auth base URL.
func WithIdentityURL(url string) Option {
	return func(c *clientConfig) { c.identityURL = url }
}

// WithScope sets the API access scope.
func WithScope(scope Scope) Option {
	return func(c *clientConfig) { c.scope = scope }
}

// WithHTTPClient provides a custom http.Client for requests.
func WithHTTPClient(client *http.Client) Option {
	return func(c *clientConfig) { c.httpClient = client }
}

// WithLogger provides a custom zap logger.
func WithLogger(logger *zap.Logger) Option {
	return func(c *clientConfig) { c.logger = logger }
}

// WithLogLevel sets the minimum log level.
func WithLogLevel(level zapcore.Level) Option {
	return func(c *clientConfig) { c.logLevel = level }
}

// WithTimeout sets the HTTP request timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *clientConfig) { c.timeout = d }
}

// WithRetry configures retry behavior for transient errors.
func WithRetry(maxRetries int, backoff time.Duration) Option {
	return func(c *clientConfig) {
		c.maxRetries = maxRetries
		c.retryBackoff = backoff
	}
}

// WithRateLimitWait enables automatic waiting when rate limited (429).
func WithRateLimitWait(wait bool) Option {
	return func(c *clientConfig) { c.rateLimitWait = wait }
}

// WithAPIVersion sets the API version (default "v2.0").
func WithAPIVersion(version string) Option {
	return func(c *clientConfig) { c.apiVersion = version }
}

// WithUserAgent sets the User-Agent header (default "payspace.com").
func WithUserAgent(ua string) Option {
	return func(c *clientConfig) { c.userAgent = ua }
}
