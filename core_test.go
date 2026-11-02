package core

import "testing"

func TestFusion(t *testing.T) {
	r, e := Fuse([][]string{{"a", "b", "a"}, {"b", "c"}}, 60)
	if e != nil || len(r) != 3 || r[0].ID != "b" {
		t.Fatal(r, e)
	}
	a, _ := Fuse([][]string{{"a"}}, 60)
	d, _ := Fuse([][]string{{"a", "a"}}, 60)
	if a[0].Score != d[0].Score {
		t.Fatal("duplicate boosted score")
	}
	if _, e = Fuse(nil, 0); e == nil {
		t.Fatal("zero k accepted")
	}
}
