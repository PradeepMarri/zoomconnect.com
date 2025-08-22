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

func Get_api_rest_v1_sms_send_url_tokenHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		tokenVal, ok := args["token"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: token"), nil
		}
		token, ok := tokenVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: token"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["recipientNumber"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("recipientNumber=%v", val))
		}
		if val, ok := args["message"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("message=%v", val))
		}
		if val, ok := args["dateToSend"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dateToSend=%v", val))
		}
		if val, ok := args["campaign"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("campaign=%v", val))
		}
		if val, ok := args["dataField"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dataField=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/rest/v1/sms/send-url/%s%s", cfg.BaseURL, token, queryString)
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
		var result string
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

func CreateGet_api_rest_v1_sms_send_url_tokenTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_rest_v1_sms_send-url_token",
		mcp.WithDescription("send-url"),
		mcp.WithString("token", mcp.Required(), mcp.Description("token")),
		mcp.WithString("recipientNumber", mcp.Required(), mcp.Description("the phone number of the recipient to send to")),
		mcp.WithString("message", mcp.Required(), mcp.Description("the message to send")),
		mcp.WithString("dateToSend", mcp.Description("date format: yyyyMMddHHmm")),
		mcp.WithString("campaign", mcp.Description("optional campaign name")),
		mcp.WithString("dataField", mcp.Description("optional extra data")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_api_rest_v1_sms_send_url_tokenHandler(cfg),
	}
}
