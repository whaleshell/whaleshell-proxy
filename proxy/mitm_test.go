package proxy

import (
	"testing"
	"time"
)

func TestLeafCachedWhileFresh(t *testing.T) {
	ca, err := GenerateMitmCA()
	if err != nil {
		t.Fatal(err)
	}
	first, err := ca.Leaf("github.com")
	if err != nil {
		t.Fatal(err)
	}
	second, err := ca.Leaf("GitHub.com")
	if err != nil {
		t.Fatal(err)
	}
	if first != second {
		t.Fatal("expected cached leaf for the same host")
	}
	if first.Leaf == nil {
		t.Fatal("expected parsed x509 leaf")
	}
}

func TestLeafReissuedNearExpiry(t *testing.T) {
	ca, err := GenerateMitmCA()
	if err != nil {
		t.Fatal(err)
	}
	stale, err := ca.Leaf("github.com")
	if err != nil {
		t.Fatal(err)
	}
	stale.Leaf.NotAfter = time.Now().Add(leafRenewBefore / 2)

	fresh, err := ca.Leaf("github.com")
	if err != nil {
		t.Fatal(err)
	}
	if fresh == stale {
		t.Fatal("expected a new leaf once the cached one is near expiry")
	}
	if time.Until(fresh.Leaf.NotAfter) <= leafRenewBefore {
		t.Fatalf("new leaf expires too soon: %s", fresh.Leaf.NotAfter)
	}
}
