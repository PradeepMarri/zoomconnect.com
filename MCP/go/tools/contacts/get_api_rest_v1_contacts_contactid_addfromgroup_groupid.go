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

func Get_api_rest_v1_contacts_contactid_addfromgroup_groupidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		contactIdVal, ok := args["contactId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: contactId"), nil
		}
		contactId, ok := contactIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: contactId"), nil
		}
		groupIdVal, ok := args["groupId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: groupId"), nil
		}
		groupId, ok := groupIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: groupId"), nil
		}
		url := fmt.Sprintf("%s/api/rest/v1/contacts/%s/addFromGroup/%s", cfg.BaseURL, contactId, groupId)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("token", cfg.APIKey)
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

func CreateGet_api_rest_v1_contacts_contactid_addfromgroup_groupidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_rest_v1_contacts_contactId_addFromGroup_groupId",
		mcp.WithDescription("removeFromGroup"),
		mcp.WithString("contactId", mcp.Required(), mcp.Description("contactId")),
		mcp.WithString("groupId", mcp.Required(), mcp.Description("groupId")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_api_rest_v1_contacts_contactid_addfromgroup_groupidHandler(cfg),
	}
}
