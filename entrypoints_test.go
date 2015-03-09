package alchemyapi

import (
	"testing"
)

func TestEntryPointsHasArrange(t *testing.T) {
	eps := make(EntryPoints)
	if got := eps.hasArrange("foo"); got == true {
		t.Errorf("want %t, but %t", !got, got)
	}

	eps["foo"] = make(map[string]string)
	eps["foo"]["bar"] = "uri"
	if got := eps.hasArrange("foo"); got == false {
		t.Errorf("want %t, but %t", got, !got)
	}
}

func TestEntryPointsUpdate(t *testing.T) {
	eps := make(EntryPoints)
	eps.update("foo", "bar", "uri")

	if len(eps) != 1 {
		t.Error("The EntryPoints length shoule be 1.")
	}

	if _, foo := eps["foo"]; !foo {
		t.Errorf("want %t, but %t", foo, !foo)
	}

	if _, bar := eps["foo"]["bar"]; !bar {
		t.Errorf("want %t, but %t", bar, !bar)
	}

	if len(eps["foo"]) != 1 {
		t.Error("The EntryPoints 'foo' map length shoule be 1.")
	}

	if eps["foo"]["bar"] != "uri" {
		t.Errorf("want %s, but %s", "uri", eps["foo"]["bar"])
	}
}

func TestEntryPointsHasFlavor(t *testing.T) {
	eps := make(EntryPoints)
	if got := eps.hasFlavor("foo_bar", "foo"); got {
		t.Errorf("want %t, but %t", !got, got)
	}

	eps.update("foo_bar", "foo", "")
	if got := eps.hasFlavor("foo_bar", "foo"); got {
		t.Errorf("want %t, but %t", !got, got)
	}

	eps.update("foo_bar", "foo", "bar")
	if got := eps.hasFlavor("foo_bar", "foo"); !got {
		t.Errorf("want %t, but %t", !got, got)
	}
}

func TestEntryPointsInitialized(t *testing.T) {
	seeds := make(EntryPoints)
	seeds.update("sentiment", "url", "/url/URLGetTextSentiment")
	seeds.update("sentiment", "text", "/text/TextGetTextSentiment")
	seeds.update("sentiment", "html", "/html/HTMLGetTextSentiment")

	seeds.update("sentiment_targeted", "url", "/url/URLGetTargetedSentiment")
	seeds.update("sentiment_targeted", "text", "/text/TextGetTargetedSentiment")
	seeds.update("sentiment_targeted", "html", "/html/HTMLGetTargetedSentiment")

	seeds.update("taxonomy", "url", "/url/URLGetRankedTaxonomy")
	seeds.update("taxonomy", "text", "/text/TextGetRankedTaxonomy")
	seeds.update("taxonomy", "html", "/html/HTMLGetRankedTaxonomy")

	seeds.update("concepts", "url", "/url/URLGetRankedConcepts")
	seeds.update("concepts", "text", "/text/TextGetRankedConcepts")
	seeds.update("concepts", "html", "/html/HTMLGetRankedConcepts")

	seeds.update("entities", "url", "/url/URLGetRankedNamedEntities")
	seeds.update("entities", "text", "/text/TextGetRankedNamedEntities")
	seeds.update("entities", "html", "/html/HTMLGetRankedNamedEntities")

	seeds.update("keywords", "url", "/url/URLGetRankedKeywords")
	seeds.update("keywords", "text", "/text/TextGetRankedKeywords")
	seeds.update("keywords", "html", "/html/HTMLGetRankedKeywords")

	seeds.update("relations", "url", "/url/URLGetRelations")
	seeds.update("relations", "text", "/text/TextGetRelations")
	seeds.update("relations", "html", "/html/HTMLGetRelations")

	seeds.update("text", "url", "/url/URLGetText")
	seeds.update("text", "html", "/html/HTMLGetText")

	seeds.update("text_raw", "url", "/url/URLGetRawText")
	seeds.update("text_raw", "html", "/html/HTMLGetRawText")

	seeds.update("title", "url", "/url/URLGetTitle")
	seeds.update("title", "html", "/html/HTMLGetTitle")

	seeds.update("face", "url", "/url/URLGetRankedImageFaceTags")
	seeds.update("face", "image", "/image/ImageGetRankedImageFaceTags")

	seeds.update("image_extract", "url", "/url/URLGetImage")
	seeds.update("image_extract", "html", "/html/HTMLGetImage")

	seeds.update("image_tag", "url", "/url/URLGetRankedImageKeywords")
	seeds.update("image_tag", "image", "/image/ImageGetRankedImageKeywords")

	seeds.update("authors", "url", "/url/URLGetAuthors")
	seeds.update("authors", "html", "/html/HTMLGetAuthors")

	seeds.update("language", "url", "/url/URLGetLanguage")
	seeds.update("language", "text", "/text/TextGetLanguage")
	seeds.update("language", "html", "/html/HTMLGetLanguage")

	seeds.update("feeds", "url", "/url/URLGetFeedLinks")
	seeds.update("feeds", "html", "/html/HTMLGetFeedLinks")

	seeds.update("microformats", "url", "/url/URLGetMicroformatData")
	seeds.update("microformats", "html", "/html/HTMLGetMicroformatData")

	seeds.update("combined", "url", "/url/URLGetCombinedData")
	seeds.update("combined", "text", "/text/TextGetCombinedData")
	seeds.update("combined", "html", "/html/HTMLGetCombinedData")

	seeds.update("publication_date", "url", "/url/URLGetPubDate")
	seeds.update("publication_date", "html", "/html/HTMLGetPubDate")

	for k, v := range seeds {
		if entryPoints[k] == nil {
			t.Errorf("The %s should not be empty", k)
		}

		for flavor, uri := range v {
			if entryPoints[k][flavor] != uri {
				t.Errorf("EntryPoint %s, Flavor %s want %s, but %#v",
					k, flavor, uri,
					entryPoints[k][flavor],
				)
			}
		}
	}
}

func TestGetEntryPoints(t *testing.T) {
	if got := entryPoints.hasArrange("foo_bar"); got {
		t.Errorf("want %t, but %t", !got, got)
	}

	entryPoints.update("foo_bar", "foo", "bar")
	eps := GetEntryPoints()
	if got := eps.hasArrange("foo_bar"); !got {
		t.Errorf("want %t, but %t", !got, got)
	}

	if uri := eps["foo_bar"]["foo"]; uri != "bar" {
		t.Errorf("want %s, but %#v", "bar", uri)
	}
}
