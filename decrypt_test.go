package areena

import (
	"testing"

	"github.com/segmentio/go-env"
)

func TestDecrypt(t *testing.T) {
	secret := env.MustGet("YLE_SECRET")
	enc := "1JXWjP87dHzBoDcxqKsqNCG7YRo/vbrQTmDpTp/+tZhxaGYMlYCd//eUrknu49COGYAXtvTDtazXVqn14YGtNXfTil0gMSNY/EorjJGnAHAjdITgOJCzMAnK5CUZMSAlFYe0PZOVwcoN7woZLTRdrDMcdHEWcWqnurypsGgFNW+aYTcYm3qryBQjoq2SYWnmMbBYMjNkwMe9wdPdclunwxx5mJoCMA94VaHFT6LMpZ1IVC3r/iTQjGtAhfAc5ZkfQ5j0ppVb0iixb+U2Bw1i7xBaRhyIeo0o2jZf6QvaIuroz/mweQ9CX0sV07Ox8Irrflo0v2s/Iosh+Vcz1NPoaW6+Q45gvuSp9ewJrtH+e/Hz+RhLNRWjRMLe2LGgXyN8IHOcLFu7GgideoA7tQDNQQ3xGdpeNg93ILPXzw+RqMo="
	dec, err := DecryptURL(secret, enc)
	if err != nil {
		t.Fatal(err)
	}
	if dec != "http://areenahdworld-vh.akamaihd.net/i/world/2d/2d72f64ba44eebf460210b8204801a3b_,148480,394240,664576,1021952,2764800,.mp4.csmil/master.m3u8?hdnea=st=1448653047~exp=1448653347~acl=/i/world/2d/2d72f64ba44eebf460210b8204801a3b_*~hmac=d958d712808a4931781845e4aea218e9ce2a51229b9df132982f7ce5aec2aa4b" {
		t.Error("decoded url was wrong")
	}
}
