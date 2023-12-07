package unionfind

import "testing"

func TestUf(t *testing.T) {
	slice := []string{
		"first",
		"second",
		"second",
		"third",
		"giraffe",
		"zebra",
		"cat",
	}
	uf := New(slice)
	if len(uf.parents) != len(slice)-1 {
		t.Errorf("expected uf.parents to have size %d, it is %d instead", len(slice)-1, len(uf.parents))
	}

	if uf.Find("third") == uf.Find("first") {
		t.Error("expected third and first to be in different trees")
	}

	united := uf.Union("first", "second")
	if uf.Find("first") != united {
		t.Errorf("Expected first to be in %d, it is in %d insted", united, uf.Find("first"))
	}
	if uf.Find("second") != united {
		t.Errorf("Expected second to be in %d, it is in %d insted", united, uf.Find("second"))
	}

	united = uf.Union("second", "third")
	if uf.Find("first") != united {
		t.Errorf("Expected first to be in %d, it is in %d insted", united, uf.Find("first"))
	}
	if uf.Find("second") != united {
		t.Errorf("Expected second to be in %d, it is in %d insted", united, uf.Find("second"))
	}
	if uf.Find("third") != united {
		t.Errorf("Expected third to be in %d, it is in %d insted", united, uf.Find("third"))
	}

	uf.Union("giraffe", "zebra")
	uf.Union("cat", "zebra")

	if uf.Find("cat") != uf.Find("zebra") {
		t.Errorf("Expected root of cat to be same as root of zebra")
	}
}
