package alchemyapi

/**
  Copyright 20153 AlchemyAPI
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	Version = "1.0.0"
)

var (
	entryPoints   EntryPoints
	ApiKeyInvalid = errors.New("It appears that the key is invalid.")
)

// Analyzer
type Analyzer struct {
	apiKey  string
	baseUrl string
}

// initialize the entrypoints
// same order with the http://www.alchemyapi.com/api intro
func init() {
	entryPoints = make(EntryPoints)
	entryPoints.update("sentiment", "url", "/url/URLGetTextSentiment")
	entryPoints.update("sentiment", "text", "/text/TextGetTextSentiment")
	entryPoints.update("sentiment", "html", "/html/HTMLGetTextSentiment")

	entryPoints.update("sentiment_targeted", "url", "/url/URLGetTargetedSentiment")
	entryPoints.update("sentiment_targeted", "text", "/text/TextGetTargetedSentiment")
	entryPoints.update("sentiment_targeted", "html", "/html/HTMLGetTargetedSentiment")

	entryPoints.update("taxonomy", "url", "/url/URLGetRankedTaxonomy")
	entryPoints.update("taxonomy", "text", "/text/TextGetRankedTaxonomy")
	entryPoints.update("taxonomy", "html", "/html/HTMLGetRankedTaxonomy")

	entryPoints.update("concepts", "url", "/url/URLGetRankedConcepts")
	entryPoints.update("concepts", "text", "/text/TextGetRankedConcepts")
	entryPoints.update("concepts", "html", "/html/HTMLGetRankedConcepts")

	entryPoints.update("entities", "url", "/url/URLGetRankedNamedEntities")
	entryPoints.update("entities", "text", "/text/TextGetRankedNamedEntities")
	entryPoints.update("entities", "html", "/html/HTMLGetRankedNamedEntities")

	entryPoints.update("keywords", "url", "/url/URLGetRankedKeywords")
	entryPoints.update("keywords", "text", "/text/TextGetRankedKeywords")
	entryPoints.update("keywords", "html", "/html/HTMLGetRankedKeywords")

	entryPoints.update("relations", "url", "/url/URLGetRelations")
	entryPoints.update("relations", "text", "/text/TextGetRelations")
	entryPoints.update("relations", "html", "/html/HTMLGetRelations")

	entryPoints.update("text", "url", "/url/URLGetText")
	entryPoints.update("text", "html", "/html/HTMLGetText")

	entryPoints.update("text_raw", "url", "/url/URLGetRawText")
	entryPoints.update("text_raw", "html", "/html/HTMLGetRawText")

	entryPoints.update("title", "url", "/url/URLGetTitle")
	entryPoints.update("title", "html", "/html/HTMLGetTitle")

	entryPoints.update("face", "url", "/url/URLGetRankedImageFaceTags")
	entryPoints.update("face", "image", "/image/ImageGetRankedImageFaceTags")

	entryPoints.update("image_extract", "url", "/url/URLGetImage")
	entryPoints.update("image_extract", "html", "/html/HTMLGetImage")

	entryPoints.update("image_tag", "url", "/url/URLGetRankedImageKeywords")
	entryPoints.update("image_tag", "image", "/image/ImageGetRankedImageKeywords")

	entryPoints.update("authors", "url", "/url/URLGetAuthors")
	entryPoints.update("authors", "html", "/html/HTMLGetAuthors")

	entryPoints.update("language", "url", "/url/URLGetLanguage")
	entryPoints.update("language", "text", "/text/TextGetLanguage")
	entryPoints.update("language", "html", "/html/HTMLGetLanguage")

	entryPoints.update("feeds", "url", "/url/URLGetFeedLinks")
	entryPoints.update("feeds", "html", "/html/HTMLGetFeedLinks")

	entryPoints.update("microformats", "url", "/url/URLGetMicroformatData")
	entryPoints.update("microformats", "html", "/html/HTMLGetMicroformatData")

	entryPoints.update("combined", "url", "/url/URLGetCombinedData")
	entryPoints.update("combined", "text", "/text/TextGetCombinedData")
	entryPoints.update("combined", "html", "/html/HTMLGetCombinedData")

	entryPoints.update("publication_date", "url", "/url/URLGetPubDate")
	entryPoints.update("publication_date", "html", "/html/HTMLGetPubDate")
}

// Creates new Analyzer
func NewAnalyzer(apiKey string) (*Analyzer, error) {
	analyzer := &Analyzer{apiKey: apiKey, baseUrl: "http://access.alchemyapi.com/calls"}
	if err := analyzer.validate(); err != nil {
		return nil, err
	} else {
		return analyzer, nil
	}
}

//Validates the Api Key length
func (analyzer *Analyzer) validate() error {
	if len(analyzer.apiKey) != 40 {
		return ApiKeyInvalid
	} else {
		return nil
	}
}

// Allow to reset the baseurl
func (analyzer *Analyzer) SetBaseUrl(url string) {
	analyzer.baseUrl = url
}

/*
	Calculates the sentiment for text, a URL or HTML.
	For an overview, please refer to: http://www.alchemyapi.com/products/features/sentiment-analysis/
	For the docs, please refer to: http://www.alchemyapi.com/api/sentiment-analysis/

	INPUT:
	flavor -> which version of the call, i.e. text, url or html.
	payload -> the data to analyze, either the text, the url or html code.
	options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

	Available Options:
	showSourceText -> 0: disabled (default), 1: enabled

	OUTPUT:
	The response, already converted from JSON to a SentimentResponse Object

*/
func (analyzer *Analyzer) Sentiment(flavor, payload string, options url.Values) (*SentimentResponse, error) {
	if !entryPoints.hasFlavor("sentiment", flavor) {
		return nil, errors.New(fmt.Sprintf("sentiment analysis for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "sentiment", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(SentimentResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
	Calculates the targeted sentiment for text, a URL or HTML.
	For an overview, please refer to: http://www.alchemyapi.com/products/features/sentiment-analysis/
	For the docs, please refer to: http://www.alchemyapi.com/api/sentiment-analysis/

	INPUT:
	flavor -> which version of the call, i.e. text, url or html.
	payload -> the data to analyze, either the text, the url or html code.
	target -> the word or phrase to run sentiment analysis on.
	options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

	Available Options:
	showSourceText	-> 0: disabled, 1: enabled

	OUTPUT:
	The response, already converted from JSON to a SentimentResponse Object.
*/
func (analyzer *Analyzer) SentimentTargeted(flavor, payload, target string, options url.Values) (*SentimentResponse, error) {
	if target == "" {
		return nil, errors.New("targeted sentiment requires a non-null target.")
	}

	if !entryPoints.hasFlavor("sentiment_targeted", flavor) {
		return nil, errors.New(fmt.Sprintf("sentiment targeted analysis for %s not available", flavor))
	}

	options.Add(flavor, payload)
	options.Add("target", target)
	url := entryPoints.urlFor(analyzer.baseUrl, "sentiment_targeted", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(SentimentResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
	Categorizes the text for a URL, text or HTML.
	For an overview, please refer to: http://www.alchemyapi.com/products/features/text-categorization/
	For the docs, please refer to: http://www.alchemyapi.com/api/taxonomy/

	INPUT:
	flavor -> which version of the call, i.e.  url, text or html.
	payload -> the data to analyze, either the the url, text or html code.
	options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

	Available Options:
	showSourceText -> 0: disabled (default), 1: enabled.

	OUTPUT:
	The response, already converted from JSON to a TaxonomyResponse object.
*/
func (analyzer *Analyzer) Taxonomy(flavor, payload string, options url.Values) (*TaxonomyResponse, error) {
	if !entryPoints.hasFlavor("taxonomy", flavor) {
		return nil, errors.New(fmt.Sprintf("Taxonomy info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "taxonomy", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(TaxonomyResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Tags the concepts for text, a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/concept-tagging/
   For the docs, please refer to: http://www.alchemyapi.com/api/concept-tagging/

   INPUT:
   flavor -> which version of the call, i.e. text, url or html.
   payload -> the data to analyze, either the text, the url or html code.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   maxRetrieve -> the maximum number of concepts to retrieve (default: 8)
   linkedData -> include linked data, 0: disabled, 1: enabled (default)
   showSourceText -> 0:disabled (default), 1: enabled

   OUTPUT:
   The response, already converted from JSON to a ConceptsResponse.
*/
func (analyzer *Analyzer) Concepts(flavor, payload string, options url.Values) (*ConceptsResponse, error) {
	if !entryPoints.hasFlavor("concepts", flavor) {
		return nil, errors.New(fmt.Sprintf("concepts info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "concepts", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(ConceptsResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Extracts the entities for text, a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/entity-extraction/
   For the docs, please refer to: http://www.alchemyapi.com/api/entity-extraction/

   INPUT:
   flavor -> which version of the call, i.e. text, url or html.
   payload -> the data to analyze, either the text, the url or html code.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   disambiguate -> disambiguate entities (i.e. Apple the company vs. apple the fruit). 0: disabled, 1: enabled (default)
   linkedData -> include linked data on disambiguated entities. 0: disabled, 1: enabled (default)
   coreference -> resolve coreferences (i.e. the pronouns that correspond to named entities). 0: disabled, 1: enabled (default)
   quotations -> extract quotations by entities. 0: disabled (default), 1: enabled.
   sentiment -> analyze sentiment for each entity. 0: disabled (default), 1: enabled. Requires 1 additional API transction if enabled.
   showSourceText -> 0: disabled (default), 1: enabled
   maxRetrieve -> the maximum number of entities to retrieve (default: 50)

   OUTPUT:
   The response, already converted from JSON to a EntitiesResponse object.
*/
func (analyzer *Analyzer) Entities(flavor, payload string, options url.Values) (*EntitiesResponse, error) {
	if !entryPoints.hasFlavor("entities", flavor) {
		return nil, errors.New(fmt.Sprintf("entities info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "entities", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(EntitiesResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Extracts the keywords from text, a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/keyword-extraction/
   For the docs, please refer to: http://www.alchemyapi.com/api/keyword-extraction/

   INPUT:
   flavor -> which version of the call, i.e. text, url or html.
   payload -> the data to analyze, either the text, the url or html code.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   keywordExtractMode -> normal (default), strict
   sentiment -> analyze sentiment for each keyword. 0: disabled (default), 1: enabled. Requires 1 additional API transaction if enabled.
   showSourceText -> 0: disabled (default), 1: enabled.
   maxRetrieve -> the max number of keywords returned (default: 50)

   OUTPUT:
   The response, already converted from JSON to a KeywordsResponse object.
*/
func (analyzer *Analyzer) Keywords(flavor, payload string, options url.Values) (*KeywordsResponse, error) {
	if !entryPoints.hasFlavor("keywords", flavor) {
		return nil, errors.New(fmt.Sprintf("keywords info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "keywords", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(KeywordsResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Extracts the relations for text, a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/relation-extraction/
   For the docs, please refer to: http://www.alchemyapi.com/api/relation-extraction/

   INPUT:
   flavor -> which version of the call, i.e. text, url or html.
   payload -> the data to analyze, either the text, the url or html code.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   sentiment -> 0: disabled (default), 1: enabled. Requires one additional API transaction if enabled.
   keywords -> extract keywords from the subject and object. 0: disabled (default), 1: enabled. Requires one additional API transaction if enabled.
   entities -> extract entities from the subject and object. 0: disabled (default), 1: enabled. Requires one additional API transaction if enabled.
   requireEntities -> only extract relations that have entities. 0: disabled (default), 1: enabled.
   sentimentExcludeEntities -> exclude full entity name in sentiment analysis. 0: disabled, 1: enabled (default)
   disambiguate -> disambiguate entities (i.e. Apple the company vs. apple the fruit). 0: disabled, 1: enabled (default)
   linkedData -> include linked data with disambiguated entities. 0: disabled, 1: enabled (default).
   coreference -> resolve entity coreferences. 0: disabled, 1: enabled (default)
   showSourceText -> 0: disabled (default), 1: enabled.
   maxRetrieve -> the maximum number of relations to extract (default: 50, max: 100)

   OUTPUT:
   The response, already converted from JSON to a Relation object.
*/
func (analyzer *Analyzer) Relations(flavor, payload string, options url.Values) (*RelationsResponse, error) {
	if !entryPoints.hasFlavor("relations", flavor) {
		return nil, errors.New(fmt.Sprintf("relations info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "relations", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(RelationsResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Extracts the cleaned text (removes ads, navigation, etc.) for text, a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/text-extraction/
   For the docs, please refer to: http://www.alchemyapi.com/api/text-extraction/

   INPUT:
   flavor -> which version of the call, i.e. text, url or html.
   data -> the data to analyze, either the text, the url or html code.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   useMetadata -> utilize meta description data, 0: disabled, 1: enabled (default)
   extractLinks -> include links, 0: disabled (default), 1: enabled.

   OUTPUT:
   The response, already converted from JSON to a Title & Text Response object.
*/
func (analyzer *Analyzer) Text(flavor, payload string, options url.Values) (*TextTitleResponse, error) {
	if !entryPoints.hasFlavor("text", flavor) {
		return nil, errors.New(fmt.Sprintf("text info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "text", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(TextTitleResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

// see Text
func (analyzer *Analyzer) TextRaw(flavor, payload string, options url.Values) (*TextTitleResponse, error) {
	if !entryPoints.hasFlavor("text_raw", flavor) {
		return nil, errors.New(fmt.Sprintf("text_raw info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "text_raw", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(TextTitleResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

// see Text
func (analyzer *Analyzer) Title(flavor, payload string, options url.Values) (*TextTitleResponse, error) {
	if !entryPoints.hasFlavor("title", flavor) {
		return nil, errors.New(fmt.Sprintf("title info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "title", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(TextTitleResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Tag image from a URL or raw image data.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/api/face-detection/
   For the docs, please refer to: http://www.alchemyapi.com/api/face-detection

   INPUT:
   flavor -> which version of the call, i.e.  url or image.
   payload -> the data to analyze, the url or image path with depends on flavor
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   extractMode -> trust-metadata: less CPU-intensive and less accurate, always-infer: more CPU-intensive and more accurate
   (image flavor only)
   imagePostMode -> how you will post the image
       raw: pass an unencoded image file using POST
			for flavor 'image' will be set to 'raw' automatically

   OUTPUT:
   The response, already converted from JSON to a FaceResponse object.
*/
func (analyzer *Analyzer) Face(flavor, payload string, options url.Values) (*FaceResponse, error) {
	if !entryPoints.hasFlavor("face", flavor) {
		return nil, errors.New(fmt.Sprintf("face info for %s not available", flavor))
	}

	var binData io.Reader

	switch flavor {
	case "url":
		options.Add(flavor, payload)
		binData = nil
	case "image":
		imageData, err := ioutil.ReadFile(payload)
		if err != nil {
			return nil, err
		}
		binData = bytes.NewReader(imageData)
		options.Set("imagePostMode", "raw")
	default:
		return nil, errors.New(fmt.Sprintf("face flavor for %s not support", flavor))
	}

	url := entryPoints.urlFor(analyzer.baseUrl, "face", flavor)
	data, err := analyzer.analyze(url, options, binData)

	if err != nil {
		return nil, err
	} else {
		response := new(FaceResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Extract image from a URL or HTML.
   For the docs, please refer to: http://www.alchemyapi.com/api/image-link-extraction

   INPUT:
   flavor -> which version of the call, i.e.  url or html.
   payload -> the data to analyze, i.e. the url or html.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   extractMode -> trust-metadata: less CPU-intensive and less accurate, always-infer: more CPU-intensive and more accurate

   OUTPUT:
   The response, already converted from JSON to a ImageExtractResponse object.
*/
func (analyzer *Analyzer) ImageExtract(flavor, payload string, options url.Values) (*ImageExtractResponse, error) {
	if !entryPoints.hasFlavor("image_extract", flavor) {
		return nil, errors.New(fmt.Sprintf("image_extract info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "image_extract", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		fmt.Println(string(data))
		response := new(ImageExtractResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Tag image from a URL or raw image data.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/image-tagging/
   For the docs, please refer to: http://www.alchemyapi.com/api/image-tagging/

   INPUT:
   flavor -> which version of the call, i.e.  url or image.
   payload -> the data to analyze, the url or image path which depends on flavor
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   extractMode -> trust-metadata: less CPU-intensive and less accurate, always-infer: more CPU-intensive and more accurate
   (image flavor only)
   imagePostMode -> how you will post the image
       raw: pass an unencoded image file using POST

   OUTPUT:
   The response, already converted from JSON to a ImageTagResponse object.
*/
func (analyzer *Analyzer) ImageTag(flavor, payload string, options url.Values) (*ImageTagResponse, error) {
	if !entryPoints.hasFlavor("image_tag", flavor) {
		return nil, errors.New(fmt.Sprintf("image_tag info for %s not available", flavor))
	}

	var binData io.Reader

	switch flavor {
	case "url":
		options.Add(flavor, payload)
		binData = nil
	case "image":
		imageData, err := ioutil.ReadFile(payload)
		if err != nil {
			return nil, err
		}
		binData = bytes.NewReader(imageData)
		options.Set("imagePostMode", "raw")
	default:
		return nil, errors.New(fmt.Sprintf("image_tag flavor for %s not support", flavor))
	}

	url := entryPoints.urlFor(analyzer.baseUrl, "image_tag", flavor)
	data, err := analyzer.analyze(url, options, binData)

	if err != nil {
		return nil, err
	} else {
		response := new(ImageTagResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Extracts the authors from a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/authors-extraction/
   For the docs, please refer to: http://www.alchemyapi.com/api/authors-extraction/

   INPUT:
   flavor -> which version of the call, i.e. text, url or html.
   payload -> the data to analyze, either the text, the url or html code.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   none

   OUTPUT:
   The response, already converted from JSON to a AuthorsResponse object.
*/
func (analyzer *Analyzer) Authors(flavor, payload string, options url.Values) (*AuthorsResponse, error) {
	if !entryPoints.hasFlavor("authors", flavor) {
		return nil, errors.New(fmt.Sprintf("authors info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "authors", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(AuthorsResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Detects the language for text, a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/api/language-detection/
   For the docs, please refer to: http://www.alchemyapi.com/products/features/language-detection/

   INPUT:
   flavor -> which version of the call, i.e. text, url or html.
   data -> the data to analyze, either the text, the url or html code.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   none

   OUTPUT:
   The response, already converted from JSON to a LanguageResponse object.
*/
func (analyzer *Analyzer) Language(flavor, payload string, options url.Values) (*LanguageResponse, error) {
	if !entryPoints.hasFlavor("language", flavor) {
		return nil, errors.New(fmt.Sprintf("language analysis for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "language", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(LanguageResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Detects the RSS/ATOM feeds for a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/feed-detection/
   For the docs, please refer to: http://www.alchemyapi.com/api/feed-detection/

   INPUT:
   flavor -> which version of the call, i.e.  url or html.
   payload -> the data to analyze, either the the url or html code.
   urlParam -> only used for flavor html to setup the url parameter(requested see the doc)
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   none

   OUTPUT:
   The response, already converted from JSON to a FeedsResponse object.
*/
func (analyzer *Analyzer) Feeds(flavor, payload, urlParam string, options url.Values) (*FeedsResponse, error) {
	if !entryPoints.hasFlavor("feeds", flavor) {
		return nil, errors.New(fmt.Sprintf("feeds info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	if flavor == "html" {
		options.Add("url", urlParam)
	}

	url := entryPoints.urlFor(analyzer.baseUrl, "feeds", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(FeedsResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Parses the microformats for a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/microformats-parsing/
   For the docs, please refer to: http://www.alchemyapi.com/api/microformats-parsing/

   INPUT:
   flavor -> which version of the call, i.e.  url or html.
   payload -> the data to analyze, either the the url or html code.
   urlParam -> only used for flavor html to setup the url parameter(requested see the doc)
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   none

   OUTPUT:
   The response, already converted from JSON to a MicroFormatsResponse object.
*/
func (analyzer *Analyzer) Microformats(flavor, payload, urlParam string, options url.Values) (*MicroFormatsResponse, error) {
	if !entryPoints.hasFlavor("microformats", flavor) {
		return nil, errors.New(fmt.Sprintf("microformats info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	if flavor == "html" {
		options.Add("url", urlParam)
	}

	url := entryPoints.urlFor(analyzer.baseUrl, "microformats", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(MicroFormatsResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Combined call (see options below for available extractions) for a URL or text or HTML.

   INPUT:
   flavor -> which version of the call, i.e.  url or text or html.
   data -> the data to analyze, either the the url or text or html.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   extract -> VALUE,VALUE,VALUE,... (possible VALUEs: page-image,entity,keyword,title,author,taxonomy,concept,relation,doc-sentiment)
   extractMode -> (only applies when 'page-image' VALUE passed to 'extract' option)
       trust-metadata: less CPU-intensive, less accurate
       always-infer: more CPU-intensive, more accurate
   disambiguate -> whether to disambiguate detected entities, 0: disabled, 1: enabled (default)
   linkedData -> whether to include Linked Data content links with disambiguated entities, 0: disabled, 1: enabled (default). disambiguate must be enabled to use this.
   coreference -> whether to he/she/etc coreferences into detected entities, 0: disabled, 1: enabled (default)
   quotations -> whether to enable quotations extraction, 0: disabled (default), 1: enabled
   sentiment -> whether to enable entity-level sentiment analysis, 0: disabled (default), 1: enabled. Requires one additional API transaction if enabled.
   showSourceText -> 0: disabled (default), 1: enabled.
   maxRetrieve -> maximum number of named entities to extract (default: 50)

   OUTPUT:
   The response, already converted from JSON to a CombinedResponse object.
*/
func (analyzer *Analyzer) Combined(flavor, payload string, options url.Values) (*CombinedResponse, error) {
	if !entryPoints.hasFlavor("combined", flavor) {
		return nil, errors.New(fmt.Sprintf("combined analysis for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "combined", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(CombinedResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

/*
   Extracts the publication_date from a URL or HTML.
   For an overview, please refer to: http://www.alchemyapi.com/products/features/publication_date-extraction/
   For the docs, please refer to: http://www.alchemyapi.com/api/publication_date-extraction/

   INPUT:
   flavor -> which version of the call, i.e. text, url or html.
   payload -> the data to analyze, either the text, the url or html code.
   options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

   Available Options:
   none

   OUTPUT:
   The response, already converted from JSON to a PublicationDateResponse object.
*/
func (analyzer *Analyzer) PublicationDate(flavor, payload string, options url.Values) (*PublicationDateResponse, error) {
	if !entryPoints.hasFlavor("publication_date", flavor) {
		return nil, errors.New(fmt.Sprintf("publication_date info for %s not available", flavor))
	}

	options.Add(flavor, payload)
	url := entryPoints.urlFor(analyzer.baseUrl, "publication_date", flavor)
	data, err := analyzer.analyze(url, options, nil)

	if err != nil {
		return nil, err
	} else {
		response := new(PublicationDateResponse)
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		} else {
			if response.Status != "OK" {
				return nil, errors.New(response.StatusInfo)
			} else {
				return response, nil
			}
		}
	}
}

// Send request
func (analyzer *Analyzer) analyze(url string, payload url.Values, binData io.Reader) ([]byte, error) {
	client := http.Client{}
	payload.Add("apikey", analyzer.apiKey)
	payload.Add("outputMode", "json")
	var req *http.Request

	if binData == nil {
		req, _ = http.NewRequest("POST", url, strings.NewReader(payload.Encode()))
	} else {
		url += fmt.Sprintf("?%s", payload.Encode())
		req, _ = http.NewRequest("POST", url, binData)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		var data []byte
		var err error

		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(resp.Body)
			defer reader.Close()
			data, err = ioutil.ReadAll(reader)
		default:
			data, err = ioutil.ReadAll(resp.Body)
		}

		if err != nil {
			return nil, err
		}
		return data, nil
	}
}
