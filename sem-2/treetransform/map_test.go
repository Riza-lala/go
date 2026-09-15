package treetransform_test

import (
	"errors"
	"reflect"
	"slices"
	"strings"
	"testing"

	"seminars-02-language/treetransform"
)

func TestMapBuildsDeepCopyInPreOrder(t *testing.T) {
	source := sampleTree()
	before := sampleTree()
	var visited []int

	got, err := treetransform.Map(source, func(value int) (int, error) {
		visited = append(visited, value)
		return value * 10, nil
	})
	if err != nil {
		t.Fatalf("Map returned error: %v", err)
	}

	wantVisited := []int{1, 2, 3, 4}
	if !slices.Equal(visited, wantVisited) {
		t.Fatalf("transform order = %v, want %v", visited, wantVisited)
	}

	want := &treetransform.Node{
		Value: 10,
		Left:  &treetransform.Node{Value: 20},
		Right: &treetransform.Node{
			Value: 30,
			Left:  &treetransform.Node{Value: 40},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Map result = %#v, want %#v", got, want)
	}

	assertNoSharedNodes(t, source, got)
	if !reflect.DeepEqual(source, before) {
		t.Fatalf("Map changed the source tree: got %#v, want %#v", source, before)
	}
}

func TestMapNilTree(t *testing.T) {
	called := false

	got, err := treetransform.Map(nil, func(value int) (int, error) {
		called = true
		return value, nil
	})
	if err != nil {
		t.Fatalf("Map(nil, transform) returned error: %v", err)
	}
	if got != nil {
		t.Fatalf("Map(nil, transform) = %#v, want nil", got)
	}
	if called {
		t.Fatal("transform was called for an empty tree")
	}
}

func TestMapNilTransform(t *testing.T) {
	testCases := []struct {
		name   string
		root   *treetransform.Node
		before *treetransform.Node
	}{
		{name: "empty tree"},
		{name: "non-empty tree", root: sampleTree(), before: sampleTree()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := treetransform.Map(tc.root, nil)
			if !errors.Is(err, treetransform.ErrNilTransform) {
				t.Fatalf("Map(root, nil) error = %v, want ErrNilTransform", err)
			}
			if got != nil {
				t.Fatalf("Map(root, nil) = %#v, want nil", got)
			}
			if !reflect.DeepEqual(tc.root, tc.before) {
				t.Fatalf("Map(root, nil) changed source: got %#v, want %#v", tc.root, tc.before)
			}
		})
	}
}

func TestMapWrapsErrorAndStops(t *testing.T) {
	source := sampleTree()
	before := sampleTree()
	errStop := errors.New("stop")
	var visited []int

	got, err := treetransform.Map(source, func(value int) (int, error) {
		visited = append(visited, value)
		if value == 3 {
			return 0, errStop
		}
		return value, nil
	})

	if got != nil {
		t.Fatalf("Map returned partial tree %#v after an error", got)
	}
	if !errors.Is(err, errStop) {
		t.Fatalf("Map error = %v, want wrapped stop error", err)
	}
	if err == nil || !strings.Contains(err.Error(), "3") {
		t.Fatalf("Map error = %v, want context with failing value 3", err)
	}

	wantVisited := []int{1, 2, 3}
	if !slices.Equal(visited, wantVisited) {
		t.Fatalf("visited after error = %v, want %v", visited, wantVisited)
	}
	if !reflect.DeepEqual(source, before) {
		t.Fatalf("Map changed the source tree after an error: got %#v, want %#v", source, before)
	}
}

func sampleTree() *treetransform.Node {
	return &treetransform.Node{
		Value: 1,
		Left:  &treetransform.Node{Value: 2},
		Right: &treetransform.Node{
			Value: 3,
			Left:  &treetransform.Node{Value: 4},
		},
	}
}

func assertNoSharedNodes(t *testing.T, source, mapped *treetransform.Node) {
	t.Helper()

	if source == nil || mapped == nil {
		if source != nil || mapped != nil {
			t.Fatalf("tree shapes differ: source=%#v mapped=%#v", source, mapped)
		}
		return
	}
	if source == mapped {
		t.Fatalf("result reuses source node %p with value %d", source, source.Value)
	}

	assertNoSharedNodes(t, source.Left, mapped.Left)
	assertNoSharedNodes(t, source.Right, mapped.Right)
}
