package main

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

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	ai "github.com/elvuel/alchemyapi_go"
)

const (
	SampleText             = "Yesterday dumb Bob destroyed my fancy iPhone in beautiful Denver, Colorado. I guess I will have to head over to the Apple Store and buy a new one."
	SampleHtml             = "<html><head><title>Golang Demo | alchemyapi</title></head><body><h1>Did you know that alchemyapi works on HTML?</h1><p>Well, you do now.</p></body></html>"
	SampleUrl              = "http://www.npr.org/2013/11/26/247336038/dont-stuff-the-turkey-and-other-tips-from-americas-test-kitchen"
	SampleImageExtractHtml = "<html><head><title>Golang Demo | alchemyapi</title></head><body><img src='http://www.alchemyapi.com/sites/default/files/ET_Leader2.png'></body></html>"
	SampleHumanImageUrl    = "http://www.alchemyapi.com/sites/default/files/ET_Leader2.png"
	SampleObjectImageUrl   = "http://www.npr.org/2013/11/26/247336038/dont-stuff-the-turkey-and-other-tips-from-americas-test-kitchen"
)

var (
	analyzer *ai.Analyzer
)

func init() {
	var err error
	analyzer, err = ai.NewAnalyzer(os.Getenv("ALCHEMY_API_KEY"))
	if err != nil {
		panic(err)
	}
}

func main() {
	Sentiment()
	SentimentTargeted()
	Taxonomy()
	Concepts()
	Entities()
	Keywords()
	Relations()
	Text()
	TextRaw()
	Title()
	Face()
	ImageExtract()
	ImageTag()
	Authors()
	Language()
	Feeds()
	Microformats()
	Combined()
	PublicationDate()
}

func Sentiment() {
	fmt.Println("Sentiment#text")
	resp, err := analyzer.Sentiment("text", SampleText, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Sentiment#html")
	resp, err = analyzer.Sentiment("html", SampleHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}

	fmt.Println("Sentiment#url")
	resp, err = analyzer.Sentiment("url", SampleUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func SentimentTargeted() {
	fmt.Println("SentimentTargeted#text")
	resp, err := analyzer.SentimentTargeted("text", SampleText, "Denver", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("SentimentTargeted#html")
	resp, err = analyzer.SentimentTargeted("html", SampleHtml, "html", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}

	fmt.Println("SentimentTarget#url")
	resp, err = analyzer.SentimentTargeted("url", SampleUrl, "turkey", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Taxonomy() {
	fmt.Println("Taxonomy#text")
	resp, err := analyzer.Taxonomy("text", SampleText, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Taxonomy#html")
	resp, err = analyzer.Taxonomy("html", SampleHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}

	fmt.Println("Taxonomy#url")
	resp, err = analyzer.Taxonomy("url", SampleUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Concepts() {
	fmt.Println("Concepts#text")
	resp, err := analyzer.Concepts("text", SampleText, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Concepts#html")
	resp, err = analyzer.Concepts("html", SampleHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}

	fmt.Println("Concepts#url")
	resp, err = analyzer.Concepts("url", SampleUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Entities() {
	fmt.Println("Entities#text")
	optionsText := url.Values{}
	optionsText.Add("sentiment", "1")
	resp, err := analyzer.Entities("text", SampleText, optionsText)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Entities#html")
	optionsHtml := url.Values{}
	optionsHtml.Add("sentiment", "1")
	resp, err = analyzer.Entities("html", SampleHtml, optionsHtml)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}

	fmt.Println("Entities#url")
	optionsUrl := url.Values{}
	optionsUrl.Add("sentiment", "1")
	resp, err = analyzer.Entities("url", SampleUrl, optionsUrl)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Keywords() {
	fmt.Println("Keywords#text")
	optionsText := url.Values{}
	optionsText.Add("sentiment", "1")
	resp, err := analyzer.Keywords("text", SampleText, optionsText)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Keywords#html")
	optionsHtml := url.Values{}
	optionsHtml.Add("sentiment", "1")
	resp, err = analyzer.Keywords("html", SampleHtml, optionsHtml)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}

	fmt.Println("Keywords#url")
	optionsUrl := url.Values{}
	optionsUrl.Add("sentiment", "1")
	resp, err = analyzer.Keywords("url", SampleUrl, optionsUrl)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Relations() {
	fmt.Println("Relations#text")
	optionsText := url.Values{}
	optionsText.Add("sentiment", "1")
	resp, err := analyzer.Relations("text", SampleText, optionsText)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Relations#html")
	optionsHtml := url.Values{}
	optionsHtml.Add("sentiment", "1")
	resp, err = analyzer.Relations("html", SampleHtml, optionsHtml)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}

	fmt.Println("Relations#url")
	optionsUrl := url.Values{}
	optionsUrl.Add("sentiment", "1")
	resp, err = analyzer.Relations("url", SampleUrl, optionsUrl)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Text() {
	fmt.Println("Text#url")
	resp, err := analyzer.Text("url", SampleUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Text#html")
	resp, err = analyzer.Text("html", SampleHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func TextRaw() {
	fmt.Println("TextRaw#url")
	resp, err := analyzer.TextRaw("url", SampleUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("TextRaw#html")
	resp, err = analyzer.TextRaw("html", SampleHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Title() {
	fmt.Println("Title#url")
	resp, err := analyzer.Title("url", SampleUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Title#html")
	resp, err = analyzer.Title("html", SampleHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Face() {
	fmt.Println("Face#image")
	resp, err := analyzer.Face("image", "beckham.jpg", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Face#url")
	resp, err = analyzer.Face("url", SampleHumanImageUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func ImageExtract() {
	fmt.Println("ImageExtract#url")
	resp, err := analyzer.ImageExtract("url", "http://www.alchemyapi.com/company/leadership", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("ImageExtract#html")
	resp, err = analyzer.ImageExtract("html", SampleImageExtractHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func ImageTag() {
	fmt.Println("ImageTag#image")
	resp, err := analyzer.ImageTag("image", "beckham.jpg", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("ImageTag#url")
	resp, err = analyzer.ImageTag("url", "http://www.alchemyapi.com/sites/default/files/ET_Leader2.png", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Authors() {
	fmt.Println("Authors#url")
	resp, err := analyzer.Authors("url", SampleUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Authors#html")
	resp, err = analyzer.Authors("html", SampleHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Language() {
	fmt.Println("Language#text")
	resp, err := analyzer.Language("text", SampleText, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Language#html")
	resp, err = analyzer.Language("html", SampleHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}

	fmt.Println("Language#url")
	resp, err = analyzer.Language("url", SampleUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Feeds() {
	fmt.Println("Feeds#url")
	resp, err := analyzer.Feeds("url", SampleUrl, "", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Feeds#html")
	resp, err = analyzer.Feeds("html", SampleHtml, "test", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Microformats() {
	fmt.Println("Microformats#url")
	resp, err := analyzer.Microformats("url", SampleUrl, "", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Microformats#html")
	resp, err = analyzer.Microformats("html", SampleHtml, "test", url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func Combined() {
	fmt.Println("Combined#text")
	optionsText := url.Values{}
	optionsText.Add("extract", "page-image,image-kw,feed,entity,keyword,title,author,taxonomy,concept,relation,pub-date,doc-sentiment")
	resp, err := analyzer.Combined("text", SampleText, optionsText)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("Combined#html")
	optionsHtml := url.Values{}
	optionsHtml.Add("extract", "page-image,image-kw,feed,entity,keyword,title,author,taxonomy,concept,relation,pub-date,doc-sentiment")
	resp, err = analyzer.Combined("html", SampleHtml, optionsHtml)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}

	fmt.Println("Combined#url")
	optionsUrl := url.Values{}
	optionsUrl.Add("extract", "page-image,image-kw,feed,entity,keyword,title,author,taxonomy,concept,relation,pub-date,doc-sentiment")
	resp, err = analyzer.Combined("url", SampleUrl, optionsUrl)
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}

func PublicationDate() {
	fmt.Println("PublicationDate#url")
	resp, err := analyzer.PublicationDate("url", SampleUrl, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t" + string(data))
	}

	fmt.Println("PublicationDate#html")
	resp, err = analyzer.PublicationDate("html", SampleHtml, url.Values{})
	if err != nil {
		fmt.Println("\t Error: ", err.Error())
	} else {
		data, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println("\t", string(data))
	}
}
