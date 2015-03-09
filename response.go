package alchemyapi

/**
  Copyright 2015 AlchemyAPI
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

// jsonutils as tool, see https://github.com/bashtian/jsonutils.

type Concept struct {
	Census         string `json:"census"`
	CiaFactbook    string `json:"ciaFactbook"`
	Crunchbase     string `json:"crunchbase"`
	Dbpedia        string `json:"dbpedia"`
	Freebase       string `json:"freebase"`
	Geo            string `json:"geo"`
	Geonames       string `json:"geonames"`
	KnowledgeGraph struct {
		TypeHierarchy string `json:"typeHierarchy"`
	} `json:"knowledgeGraph"`
	MusicBrainz string `json:"musicBrainz"`
	Opencyc     string `json:"opencyc"`
	Relevance   string `json:"relevance"`
	Text        string `json:"text"`
	Website     string `json:"website"`
	Yago        string `json:"yago"`
}

type Entity struct {
	Count         string `json:"count"`
	Disambiguated struct {
		Census      string   `json:"census"`
		CiaFactbook string   `json:"ciaFactbook"`
		Crunchbase  string   `json:"crunchbase"`
		Dbpedia     string   `json:"dbpedia"`
		Freebase    string   `json:"freebase"`
		Geo         string   `json:"geo"`
		Geonames    string   `json:"geonames"`
		MusicBrainz string   `json:"musicBrainz"`
		Name        string   `json:"name"`
		Opencyc     string   `json:"opencyc"`
		SubType     []string `json:"subType"`
		Umbel       string   `json:"umbel"`
		Website     string   `json:"website"`
		Yago        string   `json:"yago"`
	} `json:"disambiguated"`
	KnowledgeGraph struct {
		TypeHierarchy string `json:"typeHierarchy"`
	} `json:"knowledgeGraph"`
	Quotations []struct {
		Quotation string `json:"quotation"`
	} `json:"quotations"`
	Relevance string    `json:"relevance"`
	Sentiment Sentiment `json:"sentiment"`
	Text      string    `json:"text"`
	Type      string    `json:"type"`
}

type Feed struct {
	Feed string `json:"feed"`
}

type ImageFace struct {
	Age struct {
		AgeRange string  `json:"ageRange"`
		Score    float64 `json:"score,string"`
	} `json:"age"`
	Gender struct {
		Gender string  `json:"gender"`
		Score  float64 `json:"score,string"`
	} `json:"gender"`
	Height   int64 `json:"height,string"`
	Identity struct {
		Disambiguated struct {
			Crunchbase  string   `json:"crunchbase"`
			Dbpedia     string   `json:"dbpedia"`
			Freebase    string   `json:"freebase"`
			MusicBrainz string   `json:"musicBrainz"`
			Name        string   `json:"name"`
			Opencyc     string   `json:"opencyc"`
			SubType     []string `json:"subType"`
			Website     string   `json:"website"`
			Yago        string   `json:"yago"`
		} `json:"disambiguated"`
		KnowledgeGraph struct {
			TypeHierarchy string `json:"typeHierarchy"`
		} `json:"knowledgeGraph"`
		Name  string  `json:"name"`
		Score float64 `json:"score,string"`
	} `json:"identity"`
	PositionX int64 `json:"positionX,string"`
	PositionY int64 `json:"positionY,string"`
	Width     int64 `json:"width,string"`
}

type ImageKeyword struct {
	KnowledgeGraph struct {
		TypeHierarchy string `json:"typeHierarchy"`
	} `json:"knowledgeGraph"`
	Score string `json:"score"`
	Text  string `json:"text"`
}

type Keyword struct {
	KnowledgeGraph struct {
		TypeHierarchy string `json:"typeHierarchy"`
	} `json:"knowledgeGraph"`
	Relevance string    `json:"relevance"`
	Sentiment Sentiment `json:"sentiment"`
	Text      string    `json:"text"`
}

type MicroFormat struct {
	FieldData string `json:"fieldData"`
	FieldName string `json:"fieldName"`
}

type Relation struct {
	Action struct {
		Lemmatized string `json:"lemmatized"`
		Text       string `json:"text"`
		Verb       struct {
			Negated string `json:"negated"`
			Tense   string `json:"tense"`
			Text    string `json:"text"`
		} `json:"verb"`
	} `json:"action"`
	Object struct {
		Entity               Entity    `json:"entity"`
		Sentiment            Sentiment `json:"sentiment"`
		SentimentFromSubject Sentiment `json:"sentimentFromSubject"`
		Text                 string    `json:"text"`
	} `json:"object"`
	Subject struct {
		Entity    Entity    `json:"entity"`
		Sentiment Sentiment `json:"sentiment"`
		Text      string    `json:"text"`
	} `json:"subject"`
}

type Sentiment struct {
	Mixed int64   `json:"mixed,string"`
	Score float64 `json:"score,string"`
	Type  string  `json:"type"`
}

type Taxonomy struct {
	Confident string  `json:"confident"`
	Label     string  `json:"label"`
	Score     float64 `json:"score,string"`
}

type PublicationDate struct {
	Confident string `json:"confident"`
	Date      string `json:"date"`
}

// Sentiment & SentimentTargeted Response
type SentimentResponse struct {
	DocSentiment      Sentiment `json:"docSentiment"`
	Language          string    `json:"language"`
	Status            string    `json:"status"`
	StatusInfo        string    `json:"statusInfo,omitempty"`
	Text              string    `json:"text,omitempty"`
	TotalTransactions int64     `json:"totalTransactions,string"`
	Url               string    `json:"url,omitempty"`
	Usage             string    `json:"usage,omitempty"`
}

// Taxonomy Response
type TaxonomyResponse struct {
	Language          string     `json:"language"`
	Status            string     `json:"status"`
	StatusInfo        string     `json:"statusInfo,omitempty"`
	Taxonomies        []Taxonomy `json:"taxonomy"`
	Text              string     `json:"text,omitempty"`
	TotalTransactions int64      `json:"totalTransactions,string"`
	Url               string     `json:"url,omitempty"`
	Usage             string     `json:"usage,omitempty"`
}

// Concepts Response
type ConceptsResponse struct {
	Concepts   []Concept `json:"concepts"`
	Language   string    `json:"language"`
	Status     string    `json:"status"`
	StatusInfo string    `json:"statusInfo,omitempty"`
	Text       string    `json:"text,omitempty"`
	Url        string    `json:"url,omitempty"`
	Usage      string    `json:"usage,omitempty"`
}

// Entities Response
type EntitiesResponse struct {
	Entities          []Entity `json:"entities"`
	Language          string   `json:"language"`
	Status            string   `json:"status"`
	StatusInfo        string   `json:"statusInfo,omitempty"`
	Text              string   `json:"text,omitempty"`
	TotalTransactions int64    `json:"totalTransactions,string"`
	Url               string   `json:"url,omitempty"`
	Usage             string   `json:"usage,omitempty"`
}

// KeywordsResponse
type KeywordsResponse struct {
	Keywords   []Keyword `json:"keywords"`
	Language   string    `json:"language"`
	Status     string    `json:"status"`
	StatusInfo string    `json:"statusInfo,omitempty"`
	Text       string    `json:"text,omitempty"`
	Url        string    `json:"url,omitempty"`
	Usage      string    `json:"usage,omitempty"`
}

type RelationsResponse struct {
	Language   string     `json:"language"`
	Relations  []Relation `json:"relations"`
	Status     string     `json:"status"`
	StatusInfo string     `json:"statusInfo,omitempty"`
	Text       string     `json:"text,omitempty"`
	Url        string     `json:"url,omitempty"`
	Usage      string     `json:"usage,omitempty"`
}

// text & title response
type TextTitleResponse struct {
	Language   string `json:"language"`
	Status     string `json:"status"`
	StatusInfo string `json:"statusInfo,omitempty"`
	Text       string `json:"text,omitempty"`
	Title      string `json:"title,omitempty"`
	Url        string `json:"url,omitempty"`
	Usage      string `json:"usage,omitempty"`
}

// FaceResponse
type FaceResponse struct {
	ImageFaces        []ImageFace `json:"imageFaces"`
	Status            string      `json:"status"`
	StatusInfo        string      `json:"statusInfo,omitempty"`
	TotalTransactions int64       `json:"totalTransactions,string"`
	Url               string      `json:"url,omitempty"`
	Usage             string      `json:"usage"`
}

// ImageExtractResponse
type ImageExtractResponse struct {
	Language   string `json:"language"`
	Status     string `json:"status"`
	StatusInfo string `json:"statusInfo,omitempty"`
	Image      string `json:"image,omitempty"`
	Url        string `json:"url,omitempty"`
	Usage      string `json:"usage,omitempty"`
}

// ImageTagResponse
type ImageTagResponse struct {
	ImageKeywords     []ImageKeyword `json:"imageKeywords"`
	Language          string         `json:"language"`
	Status            string         `json:"status"`
	StatusInfo        string         `json:"statusInfo,omitempty"`
	TotalTransactions int64          `json:"totalTransactions,string"`
	Url               string         `json:"url,omitempty"`
	Usage             string         `json:"usage,omitempty"`
}

// AuthorsResponse
type AuthorsResponse struct {
	Authors struct {
		Confident string   `json:"confident"`
		Names     []string `json:"names"`
	} `json:"authors"`
	Status     string `json:"status"`
	StatusInfo string `json:"statusInfo,omitempty"`
	Url        string `json:"url"`
	Usage      string `json:"usage"`
}

// LanguageResponse
type LanguageResponse struct {
	Ethnologue     string `json:"ethnologue"`
	Iso6391        string `json:"iso-639-1"`
	Iso6392        string `json:"iso-639-2"`
	Iso6393        string `json:"iso-639-3"`
	Language       string `json:"language"`
	NativeSpeakers string `json:"native-speakers"`
	Status         string `json:"status"`
	StatusInfo     string `json:"statusInfo,omitempty"`
	Url            string `json:"url"`
	Usage          string `json:"usage"`
	Wikipedia      string `json:"wikipedia"`
}

// FeedsResponse
type FeedsResponse struct {
	Feeds      []Feed `json:"feeds"`
	Status     string `json:"status"`
	StatusInfo string `json:"statusInfo,omitempty"`
	Url        string `json:"url"`
	Usage      string `json:"usage"`
}

// MicroFormatsResponse
type MicroFormatsResponse struct {
	Microformats []MicroFormat `json:"microformats"`
	Status       string        `json:"status"`
	StatusInfo   string        `json:"statusInfo,omitempty"`
	Url          string        `json:"url"`
	Usage        string        `json:"usage"`
}

// CombinedResponse
type CombinedResponse struct {
	Author            string          `json:"author"`
	Concepts          []Concept       `json:"concepts"`
	DocSentiment      Sentiment       `json:"docSentiment"`
	Entities          []Entity        `json:"entities"`
	Feeds             []Feed          `json:"feeds"`
	Image             string          `json:"image"`
	ImageKeywords     []ImageKeyword  `json:"imageKeywords"`
	Keywords          []Keyword       `json:"keywords`
	Language          string          `json:"language"`
	PublicationDate   PublicationDate `json:"publicationDate"`
	Relations         []Relation      `json:"relations"`
	Status            string          `json:"status"`
	StatusInfo        string          `json:"statusInfo,omitempty"`
	Taxonomies        []Taxonomy      `json:"taxonomy"`
	Title             string          `json:"title"`
	TotalTransactions int64           `json:"totalTransactions,string"`
	Url               string          `json:"url"`
	Usage             string          `json:"usage"`
}

// PublicationDateResponse
type PublicationDateResponse struct {
	PublicationDate PublicationDate `json:"publicationDate"`
	Status          string          `json:"status"`
	StatusInfo      string          `json:"statusInfo,omitempty"`
	Url             string          `json:"url"`
	Usage           string          `json:"usage"`
}
