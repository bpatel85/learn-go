package algo

import (
	"fmt"
	"strings"
)

/**

# TextBox class
# =============
# You will be implementing a library for displaying text
# boxes. To warm up, please define a class that can be
# initialized with some text and a border character, and
# will print out the boxed text:
#
#    box = TextBox("hello", "+")
#    box.show()
#    +++++++++
#    + hello +
#    +++++++++
#

# Private methods
# ===============
# Refactor your implementation of TextBox so that it
# includes the following private methods:
#
# - boxwidth()  Returns the width of the text box.
#
# - lines()     Returns a collection of the text lines
#               that make up a text box.
#
# Please implement the following TextBox padding API
#
# paddedRight()
# =============
# Add support for a paddedRight() method on the TextBox
# class that works as follows:
#
#    box = TextBox("hello", "+")
#    box.paddedRight(4).show()
#    +++++++++++++
#    + hello     +
#    +++++++++++++
#
#
# paddedBelow()
# =============
# Add support for a paddedBelow() method that works as shown
# below:
#
#    box = TextBox("hello", "+")
#    box.paddedBelow(3).show()
#    +++++++++
#    + hello +
#    +       +
#    +       +
#    +       +
#    +++++++++
#

# addRight()
# ==========
# Implement an addRight() method that works as shown
# below:
#
#    box = TextBox("hello", "+")
#    box.addRight("world", "@").show()
#    +++++++++@@@@@@@@@
#    + hello +@ world @
#    +++++++++@@@@@@@@@
#

# addBelow()
# ==========
# Implement an addBelow() method that works as follows:
#
#    box = TextBox("hello", "+")
#    box.addBelow("world !!!", "@").show()
#    +++++++++++++
#    + hello     +
#    +++++++++++++
#    @@@@@@@@@@@@@
#    @ world !!! @
#    @@@@@@@@@@@@@
#
# Notice how the initial textbox is padded by spaces
# on the right side so that both boxes have the same
# width.

# Method chaining
#================
# Add support so that you can chain addRight() and
# addBelow() calls indefinitely:
#
#    box = TextBox("hello", "+")
#    box.addBelow("world !!!", "@") \
#       .addRight(“10”, “$”).show()
#    +++++++++++++$$$$$$
#    + hello     +$ 10 $
#    +++++++++++++$    $
#    @@@@@@@@@@@@@$    $
#    @ world !!! @$    $
#    @@@@@@@@@@@@@$$$$$$
**/

const PaddingString = " "

type TextBox struct {
	texts []string
}

func (tb TextBox) Show() string {
	return fmt.Sprintf(strings.Join(tb.texts, "\n"))
}

func NewTextBox(text string, borderStr string) TextBox {
	// assume padding is a char
	// add validations for input
	len := len(text) + 4
	borderLine := strings.Repeat(borderStr, len)
	textStr := fmt.Sprintf("%s%s%s%s%s", borderStr, PaddingString, text, PaddingString, borderStr)
	return TextBox{
		texts: []string{borderLine, textStr, borderLine},
	}
}

func (tb TextBox) boxHeight() int {
	return len(tb.texts)
}

func (tb TextBox) boxWidth() int {
	return len(tb.texts[0])
}

func (tb TextBox) PadRight(step int) TextBox {
	newLines := make([]string, len(tb.texts))
	newWidth := len(tb.texts[0]) + step
	for i := 0; i < len(tb.texts); i += 3 {
		borderStr := tb.texts[i][:1]
		newBorder := strings.Repeat(borderStr, newWidth)

		newLines[i] = newBorder
		newLines[i+2] = newBorder
		newLines[i+1] = fmt.Sprintf("%s%s%s", tb.texts[i+1][:len(tb.texts[i+1])-1], strings.Repeat(PaddingString, step), borderStr)
	}

	return TextBox{texts: newLines}
}

func (tb TextBox) PadBelow(step int) TextBox {
	newLines := make([]string, 0)

	for i := 0; i < len(tb.texts)-1; i++ {
		newLines = append(newLines, tb.texts[i])
	}

	borderStr := tb.texts[len(tb.texts)-1][:1]
	newWidth := len(tb.texts[0]) - 2
	for i := 0; i < step; i++ {
		newLines = append(newLines, fmt.Sprintf("%s%s%s", borderStr, strings.Repeat(PaddingString, newWidth), borderStr))
	}

	newLines = append(newLines, fmt.Sprintf("%s%s%s", borderStr, strings.Repeat(borderStr, newWidth), borderStr))
	return TextBox{texts: newLines}
}

func (tb TextBox) AddBelow(text string, borderStr string) TextBox {
	tb2 := NewTextBox(text, borderStr)
	tb1 := tb

	if tb2.boxWidth() < tb1.boxWidth() {
		tb2 = tb2.PadRight(tb1.boxWidth() - tb2.boxWidth())
	} else {
		tb1 = tb1.PadRight(tb2.boxWidth() - tb1.boxWidth())
	}

	newLines := make([]string, 0)
	newLines = append(newLines, tb1.texts...)
	newLines = append(newLines, tb2.texts...)

	return TextBox{
		texts: newLines,
	}
}

func (tb TextBox) AddRight(text string, borderStr string) TextBox {
	tb2 := NewTextBox(text, borderStr)
	tb1 := tb

	if tb2.boxHeight() < tb1.boxWidth() {
		tb2 = tb2.PadBelow(tb1.boxHeight() - tb2.boxHeight())
	} else {
		tb1 = tb1.PadBelow(tb2.boxHeight() - tb1.boxHeight())
	}

	newLines := make([]string, 0)
	for idx := range tb2.texts {
		newLines = append(newLines, fmt.Sprintf("%s%s", tb1.texts[idx], tb2.texts[idx]))
	}

	return TextBox{
		texts: newLines,
	}
}
