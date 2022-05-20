package stdlib

import (
	"fmt"
	"testing"
)

func Test_hash(t *testing.T) {
	in := "galloping cattle"
	tests := []struct {
		fn  string
		out string
	}{
		{"base32", "M5QWY3DPOBUW4ZZAMNQXI5DMMU======"},
		{"base64", "Z2FsbG9waW5nIGNhdHRsZQ=="},
		{"md5", "aee95ca37cfc14f4e623bb7effabbe1d"},
		{"sha1", "45f159ccf56cf30d2dacfe4ce02ce0de88838344"},
		{"sha256", "b618391489bd9629f7338f01015f9145b929eb845fd72f2145e83561741963e9"},
		{"sha512", "ec875557fe4cd2c9f8370f31c97f9c3956286efa18b4146c46a99a11c13f3aac947e8cb6a9acce29901f08151340ba2bcbe52c32e6acd8f0ca7475f3472baac3"},
		{"sha3_256", "04ec6e35949de98dfa5f56b93a02fadf15e188f1421d45a1d082c9e803b7a05d"},
		{"sha3_512", "f22e4b213c152cb1081585d9408847b016b05b352a9442faa66a55cff02ac20cec181e2376a641f54ce77b922cd14a0a8d11a4b1e153b329a1d8b5eb80be2300"},
		{"blake2b_256", "d76d404e37d5db7f0ddfe62a99ac5d8744d77b58a0a6e7ed5d37de7a372f8f28"},
		{"blake2b_512", "ce3df894669848b1a513454b0e6254e61694d42f9996f20fdacde0daea61444b961f33cbcc594f9290752aa3724bdbdc7be2932f14765d53b9f20baec91274cc"},
	}

	for _, test := range tests {
		assertQuery(t, fmt.Sprintf("SELECT %s('%s')", test.fn, in), test.out)
	}
}

func Test_unhash(t *testing.T) {
	out := "galloping cattle"
	tests := []struct {
		fn string
		in string
	}{
		{"from_base32", "M5QWY3DPOBUW4ZZAMNQXI5DMMU======"},
		{"from_base64", "Z2FsbG9waW5nIGNhdHRsZQ=="},
	}

	for _, test := range tests {
		assertQuery(t, fmt.Sprintf("SELECT %s('%s')", test.fn, test.in), out)
	}
}
