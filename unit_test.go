package GetFolders

import (
	"testing"
)

func TestList(t *testing.T) {
	ret := List("/Users/zen/Downloads")
	t.Log(ret)
}
