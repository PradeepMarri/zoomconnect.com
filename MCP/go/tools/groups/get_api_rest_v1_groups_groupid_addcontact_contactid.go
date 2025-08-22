package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/www-zoomconnect-com/mcp-server/config"
	"github.com/www-zoomconnect-com/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_api_rest_v1_groups_groupid_addcontact_contactidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		groupIdVal, ok := args["groupId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: groupId"), nil
		}
		groupId, ok := groupIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: groupId"), nil
		}
		contactIdVal, ok := args["contactId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: contactId"), nil
		}
		contactId, ok := contactIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: contactId"), nil
		}
		url := fmt.Sprintf("%s/api/rest/v1/groups/%s/addContact/%s", cfg.BaseURL, groupId, contactId)
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
		var result map[string]interface{}
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

func CreateGet_api_rest_v1_groups_groupid_addcontact_contactidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_rest_v1_groups_groupId_addContact_contactId",
		mcp.WithDescription("addContact"),
		mcp.WithString("groupId", mcp.Required(), mcp.Description("groupId")),
		mcp.WithString("contactId", mcp.Required(), mcp.Description("contactId")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_api_rest_v1_groups_groupid_addcontact_contactidHandler(cfg),
	}
}
