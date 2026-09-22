package middleware

import (
	"os"
	"strings"
)

// FromEnviron builds an optional pipeline from WHALESHELL_MIDDLEWARE_* env vars.
//
//	WHALESHELL_MIDDLEWARE_JWT_AUD     — if set, require Bearer JWT with matching aud (unverified stub)
//	WHALESHELL_MIDDLEWARE_JWT_REQUIRED — "1"/"true" to require bearer even without aud
//	WHALESHELL_MIDDLEWARE_REMOTE_URL  — POST Decision JSON stage (fail-closed)
func FromEnviron(environ []string) *Pipeline {
	if environ == nil {
		environ = os.Environ()
	}
	env := map[string]string{}
	for _, e := range environ {
		k, v, ok := strings.Cut(e, "=")
		if ok {
			env[k] = v
		}
	}
	var stages []Stage
	aud := env["WHALESHELL_MIDDLEWARE_JWT_AUD"]
	reqJWT := strings.EqualFold(env["WHALESHELL_MIDDLEWARE_JWT_REQUIRED"], "1") ||
		strings.EqualFold(env["WHALESHELL_MIDDLEWARE_JWT_REQUIRED"], "true")
	if aud != "" || reqJWT {
		stages = append(stages, &JWTStub{Audience: aud, Required: reqJWT || aud != ""})
	}
	if u := strings.TrimSpace(env["WHALESHELL_MIDDLEWARE_REMOTE_URL"]); u != "" {
		stages = append(stages, &RemoteStage{URL: u, FailClosed: true, Label: "remote_env"})
	}
	if len(stages) == 0 {
		return nil
	}
	return &Pipeline{Stages: stages}
}
