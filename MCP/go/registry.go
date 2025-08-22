package main

import (
	"github.com/www-zoomconnect-com/mcp-server/config"
	"github.com/www-zoomconnect-com/mcp-server/models"
	tools_sms "github.com/www-zoomconnect-com/mcp-server/tools/sms"
	tools_messages "github.com/www-zoomconnect-com/mcp-server/tools/messages"
	tools_groups "github.com/www-zoomconnect-com/mcp-server/tools/groups"
	tools_voice "github.com/www-zoomconnect-com/mcp-server/tools/voice"
	tools_account "github.com/www-zoomconnect-com/mcp-server/tools/account"
	tools_contacts "github.com/www-zoomconnect-com/mcp-server/tools/contacts"
	tools_templates "github.com/www-zoomconnect-com/mcp-server/tools/templates"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_sms.CreateGet_api_rest_v1_sms_send_url_tokenTool(cfg),
		tools_sms.CreatePost_api_rest_v1_sms_send_url_tokenTool(cfg),
		tools_messages.CreateAnalyseTool(cfg),
		tools_messages.CreatePost_api_rest_v1_messages_messageid_markunreadTool(cfg),
		tools_messages.CreatePut_api_rest_v1_messages_messageid_markunreadTool(cfg),
		tools_sms.CreateGet_api_rest_v1_sms_sendTool(cfg),
		tools_sms.CreatePost_api_rest_v1_sms_sendTool(cfg),
		tools_groups.CreateGet_api_rest_v1_groups_allTool(cfg),
		tools_voice.CreateDelete_api_rest_v1_voice_messageidTool(cfg),
		tools_voice.CreateGet_api_rest_v1_voice_messageidTool(cfg),
		tools_messages.CreateAnalyse_message_lengthTool(cfg),
		tools_messages.CreateAnalyse_fullTool(cfg),
		tools_account.CreateGetbalanceTool(cfg),
		tools_voice.CreateSingle_textTool(cfg),
		tools_account.CreateTransferTool(cfg),
		tools_messages.CreateAnalyse_number_of_messagesTool(cfg),
		tools_messages.CreateAnalyse_message_credit_costTool(cfg),
		tools_messages.CreateGet_api_rest_v1_messages_allTool(cfg),
		tools_messages.CreatePost_api_rest_v1_messages_messageid_markreadTool(cfg),
		tools_messages.CreatePut_api_rest_v1_messages_messageid_markreadTool(cfg),
		tools_contacts.CreateDelete_api_rest_v1_contacts_contactidTool(cfg),
		tools_contacts.CreateGet_api_rest_v1_contacts_contactidTool(cfg),
		tools_contacts.CreatePost_api_rest_v1_contacts_contactidTool(cfg),
		tools_groups.CreatePost_api_rest_v1_groups_createTool(cfg),
		tools_templates.CreateDelete_api_rest_v1_templates_templateidTool(cfg),
		tools_templates.CreateGet_api_rest_v1_templates_templateidTool(cfg),
		tools_sms.CreatePost_api_rest_v1_sms_send_bulkTool(cfg),
		tools_sms.CreateGet_api_rest_v1_sms_send_bulkTool(cfg),
		tools_contacts.CreateGet_api_rest_v1_contacts_contactid_addfromgroup_groupidTool(cfg),
		tools_contacts.CreatePost_api_rest_v1_contacts_contactid_addfromgroup_groupidTool(cfg),
		tools_contacts.CreatePost_api_rest_v1_contacts_createTool(cfg),
		tools_messages.CreateDelete_api_rest_v1_messages_messageidTool(cfg),
		tools_messages.CreateGet_api_rest_v1_messages_messageidTool(cfg),
		tools_templates.CreateGet_api_rest_v1_templates_allTool(cfg),
		tools_groups.CreateGet_api_rest_v1_groups_groupid_addcontact_contactidTool(cfg),
		tools_groups.CreatePost_api_rest_v1_groups_groupid_addcontact_contactidTool(cfg),
		tools_account.CreateSearchTool(cfg),
		tools_account.CreatePut_api_rest_v1_account_userTool(cfg),
		tools_groups.CreatePost_api_rest_v1_groups_groupidTool(cfg),
		tools_groups.CreateDelete_api_rest_v1_groups_groupidTool(cfg),
		tools_groups.CreateGet_api_rest_v1_groups_groupidTool(cfg),
		tools_account.CreatePost_api_rest_v1_account_user_useridTool(cfg),
		tools_account.CreateGetuserTool(cfg),
		tools_account.CreateGetstatisticsTool(cfg),
		tools_messages.CreateAnalyse_message_encodingTool(cfg),
		tools_contacts.CreateGet_api_rest_v1_contacts_allTool(cfg),
		tools_voice.CreateGet_api_rest_v1_voice_allTool(cfg),
		tools_groups.CreateGet_api_rest_v1_groups_groupid_removecontact_contactidTool(cfg),
		tools_groups.CreatePost_api_rest_v1_groups_groupid_removecontact_contactidTool(cfg),
		tools_contacts.CreateGet_api_rest_v1_contacts_contactid_addtogroup_groupidTool(cfg),
		tools_contacts.CreatePost_api_rest_v1_contacts_contactid_addtogroup_groupidTool(cfg),
		tools_sms.CreateGet_api_rest_v1_sms_send_url_parametersTool(cfg),
		tools_sms.CreatePost_api_rest_v1_sms_send_url_parametersTool(cfg),
	}
}
