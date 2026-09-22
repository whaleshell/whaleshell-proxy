package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJWTStubAudience(t *testing.T) {
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"aud":"whaleshell"}`))
	token := "x." + payload + ".y"
	st := &jwtStub{Audience: "whaleshell", Required: true}
	dec, err := st.Evaluate(context.Background(), Request{
		Headers: map[string]string{"Authorization": "Bearer " + token},
	})
	if err != nil || !dec.Allow {
		t.Fatalf("dec=%v err=%v", dec, err)
	}
	dec, _ = st.Evaluate(context.Background(), Request{})
	if dec.Allow {
		t.Fatal("expected deny without token")
	}
}

func TestRemoteStage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(Decision{Allow: false, Reason: "nope"})
	}))
	defer srv.Close()
	p := &Pipeline{Stages: []Stage{
		&remoteStage{URL: srv.URL, FailClosed: true},
	}}
	dec, err := p.Run(context.Background(), Request{Host: "x", Method: "GET", Path: "/"})
	if err != nil || dec.Allow {
		t.Fatalf("dec=%v err=%v", dec, err)
	}
}
