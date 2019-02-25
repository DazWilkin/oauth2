package particle

import (
	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.particle.io/oauth/authorize",
	TokenURL: "https://api.particle.io/oauth/token",
}
