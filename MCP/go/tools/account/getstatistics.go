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

func GetstatisticsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("from=%v", val))
		}
		if val, ok := args["to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to=%v", val))
		}
		if val, ok := args["userEmailAddress"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("userEmailAddress=%v", val))
		}
		if val, ok := args["campaign"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("campaign=%v", val))
		}
		if val, ok := args["includeRefundedAndOptout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeRefundedAndOptout=%v", val))
		}
		if val, ok := args["calculateCreditValue"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("calculateCreditValue=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/rest/v1/account/statistics%s", cfg.BaseURL, queryString)
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
		var result models.WebServiceAccountStatistics
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

func CreateGetstatisticsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_rest_v1_account_statistics",
		mcp.WithDescription("statistics"),
		mcp.WithString("from", mcp.Description("date format: dd-MM-yyyy")),
		mcp.WithString("to", mcp.Description("date format: dd-MM-yyyy")),
		mcp.WithString("userEmailAddress", mcp.Description("optional email address of user to return statistics for a single user, default is to return statistics for all users if administrator, or statistics for your own account if not an administrator")),
		mcp.WithString("campaign", mcp.Description("optional campaign name")),
		mcp.WithBoolean("includeRefundedAndOptout", mcp.Description("optionally include refunded and optout counts, default is false")),
		mcp.WithBoolean("calculateCreditValue", mcp.Description("optionally calculate using credit value rather than message count, default is false")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetstatisticsHandler(cfg),
	}
}
