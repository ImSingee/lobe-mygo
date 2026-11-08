package image

import "testing"

func TestAssetsDecode(t *testing.T) {
	if wave() == nil || crest() == nil || fuji() == nil {
		t.Fatal("embedded jpegs do not decode")
	}
}
