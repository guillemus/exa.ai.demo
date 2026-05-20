package exa

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	DefaultBaseURL = "https://api.exa.ai"
)

type Client struct {
	rc     *resty.Client
	apiKey string
}

func NewClient(apiKey string) *Client {
	rc := resty.New()
	rc.SetBaseURL(DefaultBaseURL)
	rc.SetHeader("x-api-key", apiKey)
	rc.SetHeader("Content-Type", "application/json")

	return &Client{
		rc:     rc,
		apiKey: apiKey,
	}
}

// Search performs a search with Exa's prompt-engineered query.
// Supports neural, fast, auto, deep, deep-reasoning, and instant search types.
// Optionally retrieves contents with text, highlights, and summaries.
func (c *Client) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	var resp SearchResponse

	r, err := c.rc.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/search")

	if err != nil {
		return nil, fmt.Errorf("search request failed: %w", err)
	}

	if r.IsError() {
		return nil, fmt.Errorf("search request failed with status %d: %s", r.StatusCode(), r.String())
	}

	return &resp, nil
}

func (c *Client) StreamSearch(ctx context.Context, req *SearchRequest, onChunk func(SearchStreamChunk) error) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return fmt.Errorf("encode stream search request: %w", err)
	}

	hreq, err := http.NewRequestWithContext(ctx, http.MethodPost, DefaultBaseURL+"/search", &body)
	if err != nil {
		return fmt.Errorf("create stream search request: %w", err)
	}
	hreq.Header.Set("x-api-key", c.apiKey)
	hreq.Header.Set("Content-Type", "application/json")
	hreq.Header.Set("Accept", "text/event-stream")

	resp, err := http.DefaultClient.Do(hreq)
	if err != nil {
		return fmt.Errorf("stream search request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bs, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("stream search request failed with status %d: %s", resp.StatusCode, string(bs))
	}

	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	dataLines := []string{}
	flush := func() error {
		if len(dataLines) == 0 {
			return nil
		}
		data := strings.TrimSpace(strings.Join(dataLines, "\n"))
		dataLines = dataLines[:0]
		if data == "[DONE]" {
			return io.EOF
		}
		var chunk SearchStreamChunk
		if err := json.Unmarshal([]byte(data), &chunk); err != nil {
			return fmt.Errorf("decode stream search chunk: %w", err)
		}
		if err := onChunk(chunk); err != nil {
			return err
		}
		return nil
	}

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\r")
		if line == "" {
			if err := flush(); err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}
			continue
		}
		if after, ok := strings.CutPrefix(line, "data:"); ok {
			dataLines = append(dataLines, after)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read stream search response: %w", err)
	}
	if err := flush(); err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}

// FindSimilar finds similar links to the provided URL.
// Optionally retrieves contents of the similar results.
func (c *Client) FindSimilar(ctx context.Context, req *FindSimilarRequest) (*FindSimilarResponse, error) {
	var resp FindSimilarResponse

	r, err := c.rc.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/findSimilar")

	if err != nil {
		return nil, fmt.Errorf("findSimilar request failed: %w", err)
	}

	if r.IsError() {
		return nil, fmt.Errorf("findSimilar request failed with status %d: %s", r.StatusCode(), r.String())
	}

	return &resp, nil
}

// GetContents retrieves contents from provided URLs.
// Supports text extraction, highlights, summaries, and subpage crawling.
func (c *Client) GetContents(ctx context.Context, req *ContentsRequest) (*ContentsResponse, error) {
	var resp ContentsResponse

	r, err := c.rc.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/contents")

	if err != nil {
		return nil, fmt.Errorf("contents request failed: %w", err)
	}

	if r.IsError() {
		return nil, fmt.Errorf("contents request failed with status %d: %s", r.StatusCode(), r.String())
	}

	return &resp, nil
}

// Answer performs a search and generates either a direct answer or detailed summary with citations.
// Supports streaming responses and structured output schemas.
func (c *Client) Answer(ctx context.Context, req *AnswerRequest) (*AnswerResponse, error) {
	var resp AnswerResponse

	r, err := c.rc.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/answer")

	if err != nil {
		return nil, fmt.Errorf("answer request failed: %w", err)
	}

	if r.IsError() {
		return nil, fmt.Errorf("answer request failed with status %d: %s", r.StatusCode(), r.String())
	}

	return &resp, nil
}

// CreateResearchTask creates a new research task with instructions and optional output schema.
// Research tasks run asynchronously and can be retrieved later by ID.
func (c *Client) CreateResearchTask(ctx context.Context, req *CreateResearchTaskRequest) (*ResearchTask, error) {
	var resp ResearchTask

	r, err := c.rc.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/research/v0/tasks")

	if err != nil {
		return nil, fmt.Errorf("create research task failed: %w", err)
	}

	if r.IsError() {
		return nil, fmt.Errorf("create research task failed with status %d: %s", r.StatusCode(), r.String())
	}

	return &resp, nil
}

// ListResearchTasks retrieves a paginated list of research tasks.
// Use cursor and limit for pagination.
func (c *Client) ListResearchTasks(ctx context.Context, cursor string, limit int) (*ListResearchTasksResponse, error) {
	var resp ListResearchTasksResponse

	req := c.rc.R().
		SetContext(ctx).
		SetResult(&resp)

	if cursor != "" {
		req.SetQueryParam("cursor", cursor)
	}
	if limit > 0 {
		req.SetQueryParam("limit", fmt.Sprintf("%d", limit))
	}

	r, err := req.Get("/research/v0/tasks")

	if err != nil {
		return nil, fmt.Errorf("list research tasks failed: %w", err)
	}

	if r.IsError() {
		return nil, fmt.Errorf("list research tasks failed with status %d: %s", r.StatusCode(), r.String())
	}

	return &resp, nil
}

// GetResearchTask retrieves a specific research task by its ID.
func (c *Client) GetResearchTask(ctx context.Context, taskID string) (*ResearchTask, error) {
	var resp ResearchTask

	r, err := c.rc.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/research/v0/tasks/%s", taskID))

	if err != nil {
		return nil, fmt.Errorf("get research task failed: %w", err)
	}

	if r.IsError() {
		return nil, fmt.Errorf("get research task failed with status %d: %s", r.StatusCode(), r.String())
	}

	return &resp, nil
}
