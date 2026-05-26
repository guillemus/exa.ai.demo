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
		s := strings.Trim(string(bs), "\"")
		if s == "" {
			*x = 0
			return nil
		}
		v, err := strconv.Atoi(s)
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
	Query                  string    `json:"query"`
	SearchType             string    `json:"searchType"`
	NumResults             SignalInt `json:"numResults"`
	Category               string    `json:"category"`
	DeepModel              string    `json:"deepModel"`
	StructuredOutputs      bool      `json:"structuredOutputs"`
	StreamResponse         bool      `json:"streamResponse"`
	SystemPromptEnabled    bool      `json:"systemPromptEnabled"`
	SystemPrompt           string    `json:"systemPrompt"`
	Highlights             bool      `json:"highlights"`
	HighlightMaxCharacters SignalInt `json:"highlightMaxCharacters"`
	HighlightQuery         string    `json:"highlightQuery"`
	Text                   bool      `json:"text"`
	TextMaxCharacters      SignalInt `json:"textMaxCharacters"`
	TextMainContentOnly    bool      `json:"textMainContentOnly"`
	MaxAgeHours            SignalInt `json:"maxAgeHours"`
	LivecrawlTimeout       SignalInt `json:"livecrawlTimeout"`
	IncludeDomains         string    `json:"includeDomains"`
	ExcludeDomains         string    `json:"excludeDomains"`
	StartPublishedDate     string    `json:"startPublishedDate"`
	EndPublishedDate       string    `json:"endPublishedDate"`
	UserLocation           string    `json:"userLocation"`
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
	if f.HighlightMaxCharacters == 0 {
		f.HighlightMaxCharacters = 4000
	}
	if f.TextMaxCharacters == 0 {
		f.TextMaxCharacters = 20000
	}
	if f.LivecrawlTimeout == 0 {
		f.LivecrawlTimeout = 10000
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

func (f SearchForm) UsesOutputSchema() bool {
	return f.StructuredOutputs
}

func (f SearchForm) UsesStreaming() bool {
	return f.StructuredOutputs && f.StreamResponse
}

func (f SearchForm) UsesSystemPrompt() bool {
	return f.StructuredOutputs && f.SystemPromptEnabled && f.SystemPrompt != ""
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
	writePythonFilters(&b, f)
	if f.UsesOutputSchema() {
		b.WriteString("    output_schema = {\n")
		b.WriteString("        \"type\": \"text\"\n")
		b.WriteString("    },\n")
	}
	writePythonSynthesis(&b, f)
	writePythonContents(&b, f)
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
	writeJavaScriptFilters(&b, f)
	if f.UsesOutputSchema() {
		b.WriteString("  outputSchema: {\n")
		b.WriteString("    type: \"text\",\n")
		b.WriteString("  },\n")
	}
	writeJavaScriptSynthesis(&b, f)
	writeJavaScriptContents(&b, f)
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
	if f.UsesStreaming() {
		b.WriteString("  --header \"Accept: text/event-stream\" \\\n")
	}
	b.WriteString("  --header \"x-api-key: 47908a******************************\" \\\n")
	b.WriteString("  --data '{\n")
	b.WriteString("    \"query\": " + strconv.Quote(f.Query) + ",\n")
	b.WriteString("    \"category\": " + strconv.Quote(f.Category) + ",\n")
	fmt.Fprintf(&b, "    \"numResults\": %d,\n", f.NumResults)
	writeCurlFilters(&b, f)
	if f.UsesOutputSchema() {
		b.WriteString("    \"outputSchema\": {\n")
		b.WriteString("      \"type\": \"text\"\n")
		b.WriteString("    },\n")
	}
	writeCurlSynthesis(&b, f)
	writeCurlContents(&b, f)
	b.WriteString("    \"type\": " + strconv.Quote(f.EffectiveSearchType()) + "\n")
	b.WriteString("  }'")
	return b.String()
}

func writePythonSynthesis(b *strings.Builder, f SearchForm) {
	if f.UsesSystemPrompt() {
		b.WriteString("    system_prompt = " + strconv.Quote(f.SystemPrompt) + ",\n")
	}
	if f.UsesStreaming() {
		b.WriteString("    stream = True,\n")
	}
}

func writeJavaScriptSynthesis(b *strings.Builder, f SearchForm) {
	if f.UsesSystemPrompt() {
		b.WriteString("  systemPrompt: " + strconv.Quote(f.SystemPrompt) + ",\n")
	}
	if f.UsesStreaming() {
		b.WriteString("  stream: true,\n")
	}
}

func writeCurlSynthesis(b *strings.Builder, f SearchForm) {
	if f.UsesSystemPrompt() {
		b.WriteString("    \"systemPrompt\": " + strconv.Quote(f.SystemPrompt) + ",\n")
	}
	if f.UsesStreaming() {
		b.WriteString("    \"stream\": true,\n")
	}
}

func writePythonContents(b *strings.Builder, f SearchForm) {
	b.WriteString("    contents = {\n")
	if f.Highlights {
		fmt.Fprintf(b, "        \"highlights\": {\"max_characters\": %d", f.HighlightMaxCharacters)
		if f.HighlightQuery != "" {
			b.WriteString(", \"query\": " + strconv.Quote(f.HighlightQuery))
		}
		b.WriteString("},\n")
	}
	if f.Text {
		fmt.Fprintf(b, "        \"text\": {\"max_characters\": %d", f.TextMaxCharacters)
		if f.TextMainContentOnly {
			b.WriteString(", \"include_sections\": [\"body\"]")
		}
		b.WriteString("},\n")
	}
	if f.MaxAgeHours != 0 {
		fmt.Fprintf(b, "        \"max_age_hours\": %d,\n", f.MaxAgeHours)
	}
	fmt.Fprintf(b, "        \"livecrawl_timeout\": %d,\n", f.LivecrawlTimeout)
	b.WriteString("        \"extras\": {\"links\": 1}\n")
	b.WriteString("    },\n")
}

func writeJavaScriptContents(b *strings.Builder, f SearchForm) {
	b.WriteString("  contents: {\n")
	if f.Highlights {
		fmt.Fprintf(b, "    highlights: { maxCharacters: %d", f.HighlightMaxCharacters)
		if f.HighlightQuery != "" {
			b.WriteString(", query: " + strconv.Quote(f.HighlightQuery))
		}
		b.WriteString(" },\n")
	}
	if f.Text {
		fmt.Fprintf(b, "    text: { maxCharacters: %d", f.TextMaxCharacters)
		if f.TextMainContentOnly {
			b.WriteString(", includeSections: ['body']")
		}
		b.WriteString(" },\n")
	}
	if f.MaxAgeHours != 0 {
		fmt.Fprintf(b, "    maxAgeHours: %d,\n", f.MaxAgeHours)
	}
	fmt.Fprintf(b, "    livecrawlTimeout: %d,\n", f.LivecrawlTimeout)
	b.WriteString("    extras: { links: 1 },\n")
	b.WriteString("  },\n")
}

func writeCurlContents(b *strings.Builder, f SearchForm) {
	b.WriteString("    \"contents\": {\n")
	if f.Highlights {
		fmt.Fprintf(b, "      \"highlights\": {\"maxCharacters\": %d", f.HighlightMaxCharacters)
		if f.HighlightQuery != "" {
			b.WriteString(", \"query\": " + strconv.Quote(f.HighlightQuery))
		}
		b.WriteString("},\n")
	}
	if f.Text {
		fmt.Fprintf(b, "      \"text\": {\"maxCharacters\": %d", f.TextMaxCharacters)
		if f.TextMainContentOnly {
			b.WriteString(", \"includeSections\": [\"body\"]")
		}
		b.WriteString("},\n")
	}
	if f.MaxAgeHours != 0 {
		fmt.Fprintf(b, "      \"maxAgeHours\": %d,\n", f.MaxAgeHours)
	}
	fmt.Fprintf(b, "      \"livecrawlTimeout\": %d,\n", f.LivecrawlTimeout)
	b.WriteString("      \"extras\": {\"links\": 1}\n")
	b.WriteString("    },\n")
}

func writePythonFilters(b *strings.Builder, f SearchForm) {
	if f.IncludeDomains != "" {
		b.WriteString("    include_domains = " + pythonList(f.IncludeDomains) + ",\n")
	}
	if f.ExcludeDomains != "" {
		b.WriteString("    exclude_domains = " + pythonList(f.ExcludeDomains) + ",\n")
	}
	if f.StartPublishedDate != "" {
		b.WriteString("    start_published_date = " + strconv.Quote(f.StartPublishedDate) + ",\n")
	}
	if f.EndPublishedDate != "" {
		b.WriteString("    end_published_date = " + strconv.Quote(f.EndPublishedDate) + ",\n")
	}
	if f.UserLocation != "" {
		b.WriteString("    user_location = " + strconv.Quote(strings.ToUpper(f.UserLocation)) + ",\n")
	}
}

func writeJavaScriptFilters(b *strings.Builder, f SearchForm) {
	if f.IncludeDomains != "" {
		b.WriteString("  includeDomains: " + jsonList(f.IncludeDomains) + ",\n")
	}
	if f.ExcludeDomains != "" {
		b.WriteString("  excludeDomains: " + jsonList(f.ExcludeDomains) + ",\n")
	}
	if f.StartPublishedDate != "" {
		b.WriteString("  startPublishedDate: " + strconv.Quote(f.StartPublishedDate) + ",\n")
	}
	if f.EndPublishedDate != "" {
		b.WriteString("  endPublishedDate: " + strconv.Quote(f.EndPublishedDate) + ",\n")
	}
	if f.UserLocation != "" {
		b.WriteString("  userLocation: " + strconv.Quote(strings.ToUpper(f.UserLocation)) + ",\n")
	}
}

func writeCurlFilters(b *strings.Builder, f SearchForm) {
	if f.IncludeDomains != "" {
		b.WriteString("    \"includeDomains\": " + jsonList(f.IncludeDomains) + ",\n")
	}
	if f.ExcludeDomains != "" {
		b.WriteString("    \"excludeDomains\": " + jsonList(f.ExcludeDomains) + ",\n")
	}
	if f.StartPublishedDate != "" {
		b.WriteString("    \"startPublishedDate\": " + strconv.Quote(f.StartPublishedDate) + ",\n")
	}
	if f.EndPublishedDate != "" {
		b.WriteString("    \"endPublishedDate\": " + strconv.Quote(f.EndPublishedDate) + ",\n")
	}
	if f.UserLocation != "" {
		b.WriteString("    \"userLocation\": " + strconv.Quote(strings.ToUpper(f.UserLocation)) + ",\n")
	}
}

func pythonList(value string) string {
	parts := splitCSV(value)
	quoted := make([]string, 0, len(parts))
	for _, part := range parts {
		quoted = append(quoted, strconv.Quote(part))
	}
	return "[" + strings.Join(quoted, ", ") + "]"
}

func jsonList(value string) string {
	return pythonList(value)
}

func splitCSV(value string) []string {
	items := []string{}
	for part := range strings.SplitSeq(value, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			items = append(items, part)
		}
	}
	return items
}
