package jwt

import (
	"strings"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestNewTokenProducesStandardJWT(t *testing.T) {
	jwter := &jwt{
		conf: &Config{
			Secret:  "secret",
			Expire:  time.Hour,
			Renewal: time.Hour,
		},
	}

	token, err := jwter.NewToken(map[string]any{
		"userId":  uint32(2),
		"isAdmin": true,
	})
	if err != nil {
		t.Fatalf("NewToken error: %v", err)
	}

	if got := strings.Count(token, "."); got != 2 {
		t.Fatalf("expected standard JWT with 2 dots, got %d: %s", got, token)
	}

	claims := map[string]any{}
	if err := jwter.ParseByToken(token, &claims); err != nil {
		t.Fatalf("ParseByToken error: %v", err)
	}

	if got := claims["userId"]; got != float64(2) {
		t.Fatalf("expected userId=2, got %v", got)
	}
	if _, ok := claims["appKey"]; ok {
		t.Fatal("expected appKey to be absent")
	}
	if _, ok := claims["deptId"]; ok {
		t.Fatal("expected deptId to be absent")
	}
	if _, ok := claims["jobId"]; ok {
		t.Fatal("expected jobId to be absent")
	}
}

func TestUniqueTokenRoundTrip(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis run error: %v", err)
	}
	t.Cleanup(mr.Close)

	client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = client.Close() })

	jwter := &jwt{
		conf: &Config{
			Secret:  "secret",
			Expire:  time.Hour,
			Renewal: time.Hour,
		},
		redis: client,
	}

	if err := jwter.SetUnique("127.0.0.1", "token-value"); err != nil {
		t.Fatalf("SetUnique error: %v", err)
	}

	if !jwter.CompareUniqueToken("127.0.0.1", "token-value") {
		t.Fatal("expected unique token to match")
	}

	if jwter.CompareUniqueToken("127.0.0.1", "other-token") {
		t.Fatal("expected different token to mismatch")
	}
}
