package caseutil

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFromCase(t *testing.T) {
	example := func(name string, input string, funcName string, converter func(string) Words, expected ...string) {
		Convey(fmt.Sprintf("Given a %s string", name), t, func() {
			Convey(fmt.Sprintf("When %s is called", funcName), func() {
				words := converter(input)

				Convey("Then it returns the expected results", func() {
					So(len(words), ShouldEqual, len(expected))
					for i, word := range words {
						So(string(word), ShouldEqual, expected[i])
					}
				})
			})
		})
	}
	example("initial-cased", "TheQuickRedFox", "FromInitial", FromInitial, "the", "quick", "red", "fox")
	example("camel-cased", "theQuickRedFox", "FromInitial", FromInitial, "the", "quick", "red", "fox")
	example("kebab-cased", "the-quick-red-fox", "FromKebab", FromKebab, "the", "quick", "red", "fox")
	example("fat-kebab-cased", "THE-QUICK-RED-FOX", "FromKebab", FromKebab, "the", "quick", "red", "fox")
	example("snake-cased", "the_quick_red_fox", "FromSnake", FromSnake, "the", "quick", "red", "fox")
	example("screaming-snake-cased", "THE_QUICK_RED_FOX", "FromSnake", FromSnake, "the", "quick", "red", "fox")
}

func TestToCase(t *testing.T) {
	example := func(funcName string, actual string, expected string) {
		Convey(fmt.Sprintf("When %s is called", funcName), func() {
			Convey("Then it returns the expected string", func() {
				So(actual, ShouldEqual, expected)
			})
		})
	}
	Convey("Given a set of words", t, func() {
		words := Words([][]rune{
			[]rune("the"),
			[]rune("quick"),
			[]rune("red"),
			[]rune("fox"),
		})

		example("ToInitial(first=false)", words.ToInitial(false), "theQuickRedFox")
		example("ToInitial(first=true)", words.ToInitial(true), "TheQuickRedFox")
		example("ToKebab(fat=false)", words.ToKebab(false), "the-quick-red-fox")
		example("ToKebab(fat=true)", words.ToKebab(true), "THE-QUICK-RED-FOX")
		example("ToSnake(screaming=false)", words.ToSnake(false), "the_quick_red_fox")
		example("ToSnake(screaming=true)", words.ToSnake(true), "THE_QUICK_RED_FOX")
	})
}
