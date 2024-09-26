package integrationtests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenURLHandler__GetValidShortUrl__ExpectedSuccessMessage(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	// FIXTURES
	longUrl := "https://en.wikipedia.org/"

	body := map[string]string{
		"long_url": longUrl,
	}

	url := baseTest.BaseTestURL + "/shorten"
	bodyJSON, _ := json.Marshal(body)
	bodyReader := bytes.NewBuffer(bodyJSON)

	// EXERCISE
	resp, err := baseTest.Request("POST", url, bodyReader)

	// ASSERT
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respData map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	parseErr := json.Unmarshal(respBody, &respData)

	assert.NoError(t, parseErr)
	assert.Contains(t, respData["short_url"].(string), baseTest.BaseTestURL)
}

func TestRedirectHandler__GetValidLongUrl__ExpectedSuccessMessage(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	// FIXTURES
	longUrl := "https://github.com/claudiobalbin"
	shortUrlKey := "ABC123"
	baseTest.CacheRepository.SetUrl(shortUrlKey, longUrl)

	url := baseTest.BaseTestURL + "/" + shortUrlKey

	// EXERCISE
	resp, err := baseTest.Request("GET", url, nil)

	// ASSERT
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, longUrl, resp.Request.URL.String())
}
