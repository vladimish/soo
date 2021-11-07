package torutils

import "testing"

func TestED25519ToAddress(t *testing.T) {
	got := ED25519ToAddress("DDD6C46CBDD5E043A4F69D668208F1CFD35D263C960D2E40894ED6ADF9B338DE")
	want := "3xlmi3f52xqehjhwtvtiechrz7jv2jr4sygs4qejj3lk36nthdpllhqd.onion"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
