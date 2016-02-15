package areena

import (
	"net/http/httptest"
	"testing"
)

var (
	testID1 = "1-2313989"
	testID2 = "1-2442458"
)

func TestAPIClientProgramDetailsMissingMediaID(t *testing.T) {
	client, api := testClient(t, testID1)
	defer api.Close()
	_, err := client.ProgramDetails(testID1)
	if err != ErrMediaIDMissing {
		t.Fatal("wrong error, media id should not be available")
	}
}

func TestAPIClientProgramDetails(t *testing.T) {
	client, api := testClient(t, testID2)
	defer api.Close()
	pd, err := client.ProgramDetails(testID2)
	if err != nil {
		t.Fatal(err)
	}
	if pd.Slug != "tenavat-erikoisjaksot-siskon-kunnia" {
		t.Errorf("invalid slug %s", pd.Slug)
	}
	if pd.MediaID != "6-4fc65ce7b6a645eaa388d55ea62c1e8a" {
		t.Errorf("invalid media id %s", pd.MediaID)
	}
	if pd.HLSURL != "http://areenahdfi-vh.akamaihd.net/i/fi/d7/d71421c68cbda92d438326a5e76c2b50_,140288,350208,566272,914432,2268160,.mp4.csmil/master.m3u8?hdnea=st=1455554179~exp=1455597379~acl=/i/fi/d7/d71421c68cbda92d438326a5e76c2b50_*~hmac=758dadb78c16e32af70bd58b4e3761bbbadaba6366e0e69c4d68df81a7de95a2" {
		t.Errorf("invalid HLS URL %s", pd.HLSURL)
	}
	if pd.Title != "Tenavat erikoisjaksot - Siskon kunnia" {
		t.Errorf("invalid title %s", pd.Title)
	}
}

func testClient(t *testing.T, fixtureID string) (*Client, *httptest.Server) {
	api := createMockAPIServer(newFixture(fixtureID))
	cfg, err := NewConfig()
	cfg.APIBaseURL = api.URL
	client, err := NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	return client, api
}
