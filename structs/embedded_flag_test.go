package structs

import "testing"

func TestFieldIsEmbeddedUsesAnonymousFlag(t *testing.T) {
	type Embedded struct{ Promoted int }
	s := New(struct {
		Embedded
		Named Embedded
	}{})
	for _, tt := range []struct {
		name string
		want bool
	}{{"Embedded", true}, {"Promoted", true}, {"Named", false}} {
		field, ok := s.Field(tt.name)
		if !ok {
			t.Fatalf("field %s missing", tt.name)
		}
		if got := field.IsEmbedded(); got != tt.want {
			t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
		}
	}
}
