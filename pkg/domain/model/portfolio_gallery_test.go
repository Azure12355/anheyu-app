package model

import (
	"encoding/json"
	"testing"
)

func TestGalleryImage_UnmarshalJSON_LegacyString(t *testing.T) {
	var item GalleryImage
	err := json.Unmarshal([]byte(`"https://example.com/legacy.jpg"`), &item)
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if item.Image != "https://example.com/legacy.jpg" {
		t.Fatalf("unexpected image: %s", item.Image)
	}
	if item.Link != "" {
		t.Fatalf("expected empty link, got: %s", item.Link)
	}
}

func TestGalleryImage_UnmarshalJSON_Object(t *testing.T) {
	var item GalleryImage
	err := json.Unmarshal(
		[]byte(`{"image":"https://example.com/new.jpg","link":"https://example.com/target"}`),
		&item,
	)
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if item.Image != "https://example.com/new.jpg" {
		t.Fatalf("unexpected image: %s", item.Image)
	}
	if item.Link != "https://example.com/target" {
		t.Fatalf("unexpected link: %s", item.Link)
	}
}

func TestGalleryImageSlice_UnmarshalJSON_MixedFormats(t *testing.T) {
	var list []GalleryImage
	err := json.Unmarshal(
		[]byte(`["https://example.com/old.jpg",{"image":"https://example.com/new.jpg","link":"https://example.com/jump"}]`),
		&list,
	)
	if err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if len(list) != 2 {
		t.Fatalf("unexpected length: %d", len(list))
	}

	if list[0].Image != "https://example.com/old.jpg" || list[0].Link != "" {
		t.Fatalf("unexpected legacy item: %+v", list[0])
	}

	if list[1].Image != "https://example.com/new.jpg" || list[1].Link != "https://example.com/jump" {
		t.Fatalf("unexpected object item: %+v", list[1])
	}
}
