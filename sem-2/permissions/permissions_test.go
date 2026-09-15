package permissions

import "testing"

func TestGrant(t *testing.T) {
	current := Read
	got := Grant(current, Write|Execute)
	want := Read | Write | Execute

	if got != want {
		t.Fatalf("Grant(%03b, %03b) = %03b, want %03b", current, Write|Execute, got, want)
	}

	if current != Read {
		t.Fatalf("Grant changed its argument: got %03b, want %03b", current, Read)
	}
}

func TestRevoke(t *testing.T) {
	all := Read | Write | Execute
	if got, want := Revoke(all, Write), Read|Execute; got != want {
		t.Fatalf("Revoke(%03b, %03b) = %03b, want %03b", all, Write, got, want)
	}

	if got := Revoke(Read, Write); got != Read {
		t.Fatalf("revoking an absent permission: got %03b, want %03b", got, Read)
	}
}

func TestHas(t *testing.T) {
	current := Read | Execute
	testCases := []struct {
		name     string
		required Permission
		want     bool
	}{
		{name: "single present", required: Read, want: true},
		{name: "single absent", required: Write, want: false},
		{name: "all present", required: Read | Execute, want: true},
		{name: "only some present", required: Read | Write, want: false},
		{name: "empty set", required: 0, want: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Has(current, tc.required); got != tc.want {
				t.Fatalf("Has(%03b, %03b) = %t, want %t", current, tc.required, got, tc.want)
			}
		})
	}
}
