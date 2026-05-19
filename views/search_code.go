package views

import (
	"fmt"
	"strconv"
	"strings"
)

const searchTypeDeep = "deep"

type SignalInt int

func (x *SignalInt) UnmarshalJSON(bs []byte) error {
	if len(bs) > 0 && bs[0] == '"' {
		v, err := strconv.Atoi(strings.Trim(string(bs), "\""))
		if err != nil {
			return err
		}
		*x = SignalInt(v)
		return nil
	}

	v, err := strconv.Atoi(string(bs))
	if err != nil {
		return err
	}
	*x = SignalInt(v)
	return nil
}

type SearchForm struct {
	Query      string    `json:"query"`
	SearchType string    `json:"searchType"`
	NumResults SignalInt `json:"numResults"`
	Category   string    `json:"category"`
	DeepModel  string    `json:"deepModel"`
}

func (f SearchForm) WithDefaults() SearchForm {
	if f.Query == "" {
		f.Query = "Latest news on Nvidia"
	}
	if f.SearchType == "" {
		f.SearchType = "auto"
	}
	if f.NumResults == 0 {
		f.NumResults = 10
	}
	if f.Category == "" {
		f.Category = "company"
	}
	if f.DeepModel == "" {
		f.DeepModel = "deep"
	}
	return f
}

func (f SearchForm) EffectiveSearchType() string {
	f = f.WithDefaults()
	if f.SearchType == searchTypeDeep {
		return f.DeepModel
	}
	return f.SearchType
}

func (f SearchForm) UsesDeepOutput() bool {
	return strings.HasPrefix(f.EffectiveSearchType(), searchTypeDeep)
}

func PythonSearchCode(f SearchForm) string {
	f = f.WithDefaults()

	var b strings.Builder
	b.WriteString("from exa_py import Exa\n\n")
	b.WriteString(`exa = Exa("47908a******************************")`)
	b.WriteString("\n\n")
	b.WriteString("result = exa.search(\n")
	b.WriteString("    " + strconv.Quote(f.Query) + ",\n")
	b.WriteString("    category = " + strconv.Quote(f.Category) + ",\n")
	fmt.Fprintf(&b, "    num_results = %d,\n", f.NumResults)
	if f.UsesDeepOutput() {
		b.WriteString("    output_schema = {\n")
		b.WriteString("        \"type\": \"text\"\n")
		b.WriteString("    },\n")
	}
	b.WriteString("    contents = {\n")
	b.WriteString("        \"highlights\": {\"max_characters\": 4000},\n")
	b.WriteString("        \"extras\": {\"links\": 1}\n")
	b.WriteString("    },\n")
	b.WriteString("    type = " + strconv.Quote(f.EffectiveSearchType()) + "\n")
	b.WriteString(")")
	return b.String()
}

func JavaScriptSearchCode(f SearchForm) string {
	f = f.WithDefaults()

	var b strings.Builder
	b.WriteString("import Exa from \"exa-js\";\n\n")
	b.WriteString(`const exa = new Exa("47908a******************************");`)
	b.WriteString("\n\n")
	b.WriteString("const result = await exa.search(" + strconv.Quote(f.Query) + ", {\n")
	b.WriteString("  category: " + strconv.Quote(f.Category) + ",\n")
	fmt.Fprintf(&b, "  numResults: %d,\n", f.NumResults)
	if f.UsesDeepOutput() {
		b.WriteString("  outputSchema: {\n")
		b.WriteString("    type: \"text\",\n")
		b.WriteString("  },\n")
	}
	b.WriteString("  contents: {\n")
	b.WriteString("    highlights: { maxCharacters: 4000 },\n")
	b.WriteString("    extras: { links: 1 },\n")
	b.WriteString("  },\n")
	b.WriteString("  type: " + strconv.Quote(f.EffectiveSearchType()) + ",\n")
	b.WriteString("});")
	return b.String()
}

func CurlSearchCode(f SearchForm) string {
	f = f.WithDefaults()

	var b strings.Builder
	b.WriteString("curl https://api.exa.ai/search \\\n")
	b.WriteString("  --request POST \\\n")
	b.WriteString("  --header \"Content-Type: application/json\" \\\n")
	b.WriteString("  --header \"x-api-key: 47908a******************************\" \\\n")
	b.WriteString("  --data '{\n")
	b.WriteString("    \"query\": " + strconv.Quote(f.Query) + ",\n")
	b.WriteString("    \"category\": " + strconv.Quote(f.Category) + ",\n")
	fmt.Fprintf(&b, "    \"numResults\": %d,\n", f.NumResults)
	if f.UsesDeepOutput() {
		b.WriteString("    \"outputSchema\": {\n")
		b.WriteString("      \"type\": \"text\"\n")
		b.WriteString("    },\n")
	}
	b.WriteString("    \"contents\": {\n")
	b.WriteString("      \"highlights\": {\"maxCharacters\": 4000},\n")
	b.WriteString("      \"extras\": {\"links\": 1}\n")
	b.WriteString("    },\n")
	b.WriteString("    \"type\": " + strconv.Quote(f.EffectiveSearchType()) + "\n")
	b.WriteString("  }'")
	return b.String()
}
