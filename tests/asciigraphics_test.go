package tests

import (
	"ascii-art/asciimodifier"
	"testing"
)

func TestShadow(t *testing.T) {
	input := "A"
	expected := "         \n"+
	"  _|_|   \n"+
	"_|    _| \n"+
	"_|_|_|_| \n"+
	"_|    _| \n"+
	"_|    _| \n"+
	"         \n"+
	"         \n"
	// Getting the map for the shadow file
	asciimap := asciimodifier.AsciiMapper("../bannerfiles/shadow.txt")
	got := asciimodifier.AsciiGraphics(input, asciimap)
	if got != expected{
		t.Errorf("expected %v and got %v ", expected,got)
	}
}

func TestStandard(t *testing.T){
	input := "A"
	expected := "           \n" +
	"    /\\     \n" +
	"   /  \\    \n" +
	"  / /\\ \\   \n" +
	" / ____ \\  \n" +
	"/_/    \\_\\ \n" +
	"           \n" +
	"           \n"
	asciimap := asciimodifier.AsciiMapper("../bannerfiles/standard.txt")
	got := asciimodifier.AsciiGraphics(input,asciimap)
	if got != expected{
		t.Errorf("expected %v and got %v ", expected,got)
	}
}

func TestThinkertoy(t *testing.T) {
	input := "A"
	expected := "      \n"+
	"  O   \n"+
	" / \\  \n"+
	"o---o \n"+
	"|   | \n"+
	"o   o \n"+
	"      \n"+
	"      \n"
	asciimap := asciimodifier.AsciiMapper("../bannerfiles/thinkertoy.txt")
	got := asciimodifier.AsciiGraphics(input, asciimap)
	if got != expected{
		t.Errorf("expected %v and got %v ", expected,got)
	}
}