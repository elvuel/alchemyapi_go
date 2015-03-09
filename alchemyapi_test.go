package alchemyapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	key := "foooooooooooooooooooooooooooooooooooobar"
	keyInvalid := key + "!"
	_, err := NewAnalyzer(key)
	if err != nil {
		t.Error("should not be error")
	}

	_, err = NewAnalyzer(keyInvalid)
	if err == nil {
		t.Error("should be error")
	}
}

func TestAnalyzerSetBaseUrl(t *testing.T) {
	analyzer, _ := NewAnalyzer("foooooooooooooooooooooooooooooooooooobar")
	analyzer.SetBaseUrl("baseuri")

	if analyzer.baseUrl != "baseuri" {
		t.Error("Should be setted")
	}
}

func TestAnalyzerSentiment(t *testing.T) {
	apiKey := "foooooooooooooooooooooooooooooooooooobar"
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			switch r.URL.String() {
			case entryPoints["sentiment"]["html"], entryPoints["sentiment"]["text"], entryPoints["sentiment"]["url"]:
				switch r.FormValue("test_eager") {
				case "":
					resp := &SentimentResponse{}
					resp.DocSentiment.Mixed = 1
					resp.DocSentiment.Score = -0.5
					resp.DocSentiment.Type = "negative"
					resp.Language = "english"
					resp.Status = "OK"
					resp.Text = r.FormValue("html") + r.FormValue("text") + r.FormValue("url")
					resp.TotalTransactions = 1
					resp.Usage = "Usage information"
					data, _ := json.Marshal(&resp)
					w.Write(data)
				default:
					w.Write([]byte("{\"status\":\"ERROR\",\"statusInfo\":\"malfunction\"}"))
				}
			default:
				w.Write([]byte("{\"status\":\"ERROR\",\"statusInfo\":\"unsupport\"}"))
			}
		} else {
			w.Write([]byte("{\"status\":\"ERROR\",\"statusInfo\":\"unsupport\"}"))
		}
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	analyzer, _ := NewAnalyzer(apiKey)
	analyzer.SetBaseUrl(server.URL)
	// normal
	for _, v := range []string{"html", "text", "url"} {
		resp, err := analyzer.Sentiment(
			v,
			"Yesterday dumb Bob destroyed my fancy iPhone in beautiful Denver, Colorado. ",
			url.Values{},
		)
		if err != nil {
			t.Errorf("With sentiment flavor %s, should not raise exception", v)
		}

		if resp == nil {
			t.Errorf("With sentiment flavor %s, should not be nil", v)
		}

		if resp.Status != "OK" {
			t.Errorf("With sentiment flavor %s, should be ok", v)
		}
	}
	// malfunction
	for _, v := range []string{"html", "text", "url"} {
		options := url.Values{}
		options.Add("test_eager", "malfunction")
		resp, err := analyzer.Sentiment(
			v,
			"foobar",
			options,
		)
		if err == nil {
			t.Errorf("With sentiment flavor %s, should raise exception", v)
		}

		if err.Error() != "malfunction" {
			t.Errorf("With sentiment flavor %s, should raise malfunction message", v)
		}

		if resp != nil {
			t.Errorf("With sentiment flavor %s, should be nil", v)
		}
	}
}
