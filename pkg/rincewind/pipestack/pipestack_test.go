package pipestack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var expect_error = errors.New("Want :), have :(")

type testPipeAndMatcher struct {
	t     *testing.T
	items []testItem
	size  int
	idx   int
}

// Implements DataPipe.Next
func (tpm *testPipeAndMatcher) Next() interface{} {

	if tpm.idx >= tpm.size {
		return nil
	}

	item := tpm.items[tpm.idx]
	tpm.idx++
	return item
}

// Implements Matcher.Match
func (tpm *testPipeAndMatcher) Match(next, other interface{}) bool {

	if next == nil {
		return false
	}

	a, ok := next.(testItem)
	require.True(tpm.t, ok, "Wrong 'next' type passed to Match function")

	if b, ok := other.(testItem); ok {
		return a == b
	}

	if b, ok := other.(string); ok {
		return a.name == b
	}

	require.Fail(tpm.t, "Wrong 'other' type passed to Match function")
	return false
}

// Implements Matcher.Expect
func (tpm *testPipeAndMatcher) Expect(next, other interface{}) error {

	if tpm.Match(next, other) {
		return nil
	}

	return expect_error
}

type testItem struct {
	name string
	age  int
}

func setup(t *testing.T, items []testItem) *PipeStack {
	tpm := &testPipeAndMatcher{t, items, len(items), 0}
	return NewPipeStack(tpm, tpm)
}

func Test1_1(t *testing.T) {

	alice := testItem{"Alice", 28}
	bob := testItem{"Bob", 45}

	ps := setup(t, []testItem{alice, bob})

	require.NotNil(t, ps)
	require.False(t, ps.Empty())
	require.True(t, ps.EmptyStack())

	require.True(t, ps.MatchNext(alice))
	require.False(t, ps.MatchNext(bob))
	require.True(t, ps.MatchNext("Alice"))
	require.False(t, ps.MatchNext("Bob"))
	require.Equal(t, alice, ps.PeekNext())

	require.Equal(t, alice, ps.Next())
	require.False(t, ps.Empty())
	require.True(t, ps.EmptyStack())

	require.True(t, ps.MatchNext(bob))
	require.False(t, ps.MatchNext(alice))
	require.True(t, ps.MatchNext("Bob"))
	require.False(t, ps.MatchNext("Alice"))
	require.Equal(t, bob, ps.PeekNext())

	require.Equal(t, bob, ps.Next())
	require.True(t, ps.Empty())
}

func Test2_1(t *testing.T) {

	alice := testItem{"Alice", 28}
	bob := testItem{"Bob", 45}
	charlie := testItem{"Charlie", 21}

	ps := setup(t, []testItem{alice, bob, charlie})

	require.True(t, ps.AcceptPush(alice))
	require.False(t, ps.Empty())
	require.False(t, ps.EmptyStack())
	require.Equal(t, alice, ps.PeekTop())
	require.True(t, ps.MatchTop(alice))

	require.Equal(t, alice, ps.Pop())
	require.False(t, ps.Empty())
	require.True(t, ps.EmptyStack())

	require.Nil(t, ps.ExpectPush(bob))
	require.False(t, ps.Empty())
	require.False(t, ps.EmptyStack())
	require.Equal(t, bob, ps.PeekTop())
	require.True(t, ps.MatchTop(bob))

	require.Nil(t, ps.ExpectPush("Charlie"))
	require.False(t, ps.Empty())
	require.False(t, ps.EmptyStack())
	require.Equal(t, charlie, ps.PeekTop())
	require.True(t, ps.MatchTop(charlie))

	require.Nil(t, ps.AcceptPop(bob))
	_, e := ps.ExpectPop(bob)
	require.Equal(t, expect_error, e)

	require.Equal(t, charlie, ps.AcceptPop("Charlie"))
	require.False(t, ps.Empty())
	require.False(t, ps.EmptyStack())
	require.Equal(t, bob, ps.PeekTop())
	require.True(t, ps.MatchTop(bob))

	data, e := ps.ExpectPop(bob)
	require.Nil(t, e)
	require.Equal(t, bob, data)
	require.True(t, ps.Empty())
	require.True(t, ps.EmptyStack())
}
