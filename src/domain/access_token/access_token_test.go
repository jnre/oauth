package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	assert.EqualValues(t, "", at.AccessToken, "brand new access token should not be defined access token id")
	assert.True(t, at.UserId == 0, "new access token should not have an associated user Id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "acces token created three hours from now should not be expired")
}

func TestAccessTokenConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}
