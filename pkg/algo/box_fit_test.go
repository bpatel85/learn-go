package algo_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bpatel85/learn-go/pkg/algo"
)

func TestTextBox(t *testing.T) {
	tb := algo.NewTextBox("hello", "@")
	fmt.Println(tb.Show())

	tb2 := tb.PadRight(3)
	fmt.Println(tb2.Show())

	tb3 := tb2.PadBelow(5)
	fmt.Println(tb3.Show())

	tb4 := tb.AddBelow("world, I am here", "+")
	fmt.Println(tb4.Show())

	tb5 := tb.AddRight("world", "*")
	fmt.Println(tb5.Show())

	tb6 := tb5.AddBelow("foobar", "+")
	fmt.Println(tb6.Show())

	tb7 := tb6.AddRight("100", "$")
	actual := tb7.Show()
	fmt.Println(actual)

	expected := `@@@@@@@@@*********$$$$$$$
@ hello @* world *$ 100 $
@@@@@@@@@*********$     $
++++++++++++++++++$     $
+ foobar         +$     $
++++++++++++++++++$$$$$$$`

	if !strings.EqualFold(expected, actual) {
		t.Error("not match")
	}
}
