package main

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func field(t *testing.T, s apiStruct, name string) apiField {
	t.Helper()
	for _, f := range s.Fields {
		if f.Name == name {
			return f
		}
	}
	t.Fatalf("no field %s in %+v", name, s)
	return apiField{}
}

func TestExtractButtonProps(t *testing.T) {
	api, err := extract("../..")
	if err != nil {
		t.Fatal(err)
	}
	props, ok := api["ButtonProps"]
	if !ok {
		t.Fatalf("no ButtonProps in %v", api)
	}
	v := field(t, props, "Variant")
	if v.Type != "ButtonVariant" || v.Default != "Default" || v.Doc == "" {
		t.Errorf("Variant = %+v", v)
	}
	if want := []string{"Default", "Primary", "Dashed", "Filled", "Text", "Link"}; !slices.Equal(v.Values, want) {
		t.Errorf("Variant values = %v, want %v", v.Values, want)
	}
	if s := field(t, props, "Size"); s.Default != "Middle" || s.Doc == "" {
		t.Errorf("Size = %+v", s)
	}
}

func TestExtractEnumsAcrossBlocksAndPlainTypes(t *testing.T) {
	dir := t.TempDir()
	src := `package demo

type Tone int

const ToneSoft Tone = iota

const (
	ToneLoud Tone = iota + 1
)

type Count int

type CardProps struct {
	// Tone colors the card.
	Tone  Tone
	Count Count
	Title string
	Open  bool
}
`
	if err := os.WriteFile(filepath.Join(dir, "demo.go"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	api, err := extract(dir)
	if err != nil {
		t.Fatal(err)
	}
	card := api["CardProps"]
	if tone := field(t, card, "Tone"); !slices.Equal(tone.Values, []string{"ToneSoft", "ToneLoud"}) || tone.Default != "ToneSoft" || tone.Doc != "Tone colors the card." {
		t.Errorf("Tone = %+v", tone)
	}
	if count := field(t, card, "Count"); count.Values != nil || count.Default != "0" {
		t.Errorf("Count = %+v", count)
	}
	if title := field(t, card, "Title"); title.Default != `""` {
		t.Errorf("Title = %+v", title)
	}
	if open := field(t, card, "Open"); open.Default != "false" {
		t.Errorf("Open = %+v", open)
	}
}

func TestPageOf(t *testing.T) {
	for name, want := range map[string]string{
		"ButtonProps":     "button",
		"ActionIconProps": "action-icon",
		"TagProps":        "tag",
		"SVGIconProps":    "svg-icon",
	} {
		if got := pageOf(name); got != want {
			t.Errorf("pageOf(%s) = %q, want %q", name, got, want)
		}
	}
}

func TestExtractOnlyTreatsLocalIntegerTypesAsEnums(t *testing.T) {
	dir := t.TempDir()
	src := `package demo

import "time"

const MaxWidth float32 = 480

const FadeIn time.Duration = 200

type Tone int

const (
	ToneLoud Tone = iota + 1
	ToneSoft
)

type Mode int

const ModeNone = Mode(0)

const ModeAuto Mode = 1

type PanelProps struct {
	Width float32
	Fade  time.Duration
	Tone  Tone
	Mode  Mode
}
`
	if err := os.WriteFile(filepath.Join(dir, "demo.go"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	api, err := extract(dir)
	if err != nil {
		t.Fatal(err)
	}
	panel := api["PanelProps"]
	if w := field(t, panel, "Width"); w.Values != nil || w.Default != "0" {
		t.Errorf("Width = %+v, want no values and default 0", w)
	}
	if f := field(t, panel, "Fade"); f.Values != nil {
		t.Errorf("Fade = %+v, want no values", f)
	}
	if tone := field(t, panel, "Tone"); !slices.Equal(tone.Values, []string{"ToneLoud", "ToneSoft"}) || tone.Default != "0" {
		t.Errorf("Tone = %+v, want values [ToneLoud ToneSoft] and default 0", tone)
	}
	if mode := field(t, panel, "Mode"); !slices.Equal(mode.Values, []string{"ModeNone", "ModeAuto"}) || mode.Default != "ModeNone" {
		t.Errorf("Mode = %+v, want values [ModeNone ModeAuto] and default ModeNone", mode)
	}
}
