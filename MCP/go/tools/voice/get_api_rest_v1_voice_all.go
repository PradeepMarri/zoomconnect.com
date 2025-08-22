package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/www-zoomconnect-com/mcp-server/config"
	"github.com/www-zoomconnect-com/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_api_rest_v1_voice_allHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		if val, ok := args["fromDateTimeSent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fromDateTimeSent=%v", val))
		}
		if val, ok := args["toDateTimeSent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("toDateTimeSent=%v", val))
		}
		if val, ok := args["toNumber"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("toNumber=%v", val))
		}
		if val, ok := args["message"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("message=%v", val))
		}
		if val, ok := args["campaign"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("campaign=%v", val))
		}
		if val, ok := args["dataField"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dataField=%v", val))
		}
		if val, ok := args["deleted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("deleted=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/rest/v1/voice/all%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("email", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.WebServiceVoiceMessages
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_api_rest_v1_voice_allTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_rest_v1_voice_all",
		mcp.WithDescription("all"),
		mcp.WithNumber("pageSize", mcp.Description("number of elements to return at a time")),
		mcp.WithNumber("page", mcp.Description("page number")),
		mcp.WithString("status", mcp.Description("filter by message status")),
		mcp.WithString("fromDateTimeSent", mcp.Description("date format: yyyyMMdd")),
		mcp.WithString("toDateTimeSent", mcp.Description("date format: yyyyMMdd")),
		mcp.WithString("toNumber", mcp.Description("phone number the message was sent to")),
		mcp.WithString("message", mcp.Description("search matching message text")),
		mcp.WithString("campaign", mcp.Description("search by campaign")),
		mcp.WithString("dataField", mcp.Description("search by data field")),
		mcp.WithBoolean("deleted", mcp.Description("return only deleted / not deleted messages")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_api_rest_v1_voice_allHandler(cfg),
	}
}
