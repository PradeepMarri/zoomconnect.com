package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/www-zoomconnect-com/mcp-server/config"
	"github.com/www-zoomconnect-com/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_api_rest_v1_contacts_contactidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.WebServiceContact
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/api/rest/v1/contacts/%s", cfg.BaseURL, contactId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.WebServiceContact
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

func CreatePost_api_rest_v1_contacts_contactidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_api_rest_v1_contacts_contactId",
		mcp.WithDescription("update"),
		mcp.WithString("contactId", mcp.Required(), mcp.Description("contactId")),
		mcp.WithString("lastName", mcp.Description("")),
		mcp.WithArray("links", mcp.Description("")),
		mcp.WithString("title", mcp.Description("")),
		mcp.WithString("contactId", mcp.Description("")),
		mcp.WithString("contactNumber", mcp.Description("")),
		mcp.WithString("firstName", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_api_rest_v1_contacts_contactidHandler(cfg),
	}
}
