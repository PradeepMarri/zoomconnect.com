package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// WebServiceAccount represents the WebServiceAccount schema from the OpenAPI specification
type WebServiceAccount struct {
	Creditbalance float64 `json:"creditBalance,omitempty"`
	Links []Link `json:"links,omitempty"`
}

// WebServiceVoiceMessage represents the WebServiceVoiceMessage schema from the OpenAPI specification
type WebServiceVoiceMessage struct {
	Message string `json:"message,omitempty"`
	Userdatafield string `json:"userDataField,omitempty"`
	Audiofileurl string `json:"audioFileUrl,omitempty"`
	Language string `json:"language,omitempty"`
	Links []Link `json:"links,omitempty"`
	Tonumber string `json:"toNumber,omitempty"`
	Campaign string `json:"campaign,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Messagestatus string `json:"messageStatus,omitempty"`
	Datetimesent string `json:"dateTimeSent,omitempty"`
	Voicemessageid string `json:"voiceMessageId,omitempty"`
}

// WebServiceSendSmsRequest represents the WebServiceSendSmsRequest schema from the OpenAPI specification
type WebServiceSendSmsRequest struct {
	Message string `json:"message,omitempty"`
	Recipientnumber string `json:"recipientNumber,omitempty"`
	Campaign string `json:"campaign,omitempty"`
	Datafield string `json:"dataField,omitempty"`
	Datetosend string `json:"dateToSend,omitempty"`
}

// WebServiceAnalyseMessageRequestMessageAndRecipientNumber represents the WebServiceAnalyseMessageRequestMessageAndRecipientNumber schema from the OpenAPI specification
type WebServiceAnalyseMessageRequestMessageAndRecipientNumber struct {
	Recipientnumber string `json:"recipientNumber,omitempty"`
	Message string `json:"message,omitempty"`
}

// WebServiceTemplates represents the WebServiceTemplates schema from the OpenAPI specification
type WebServiceTemplates struct {
	Webservicetemplates []WebServiceTemplate `json:"webServiceTemplates,omitempty"`
	Links []Link `json:"links,omitempty"`
}

// WebServiceUser represents the WebServiceUser schema from the OpenAPI specification
type WebServiceUser struct {
	Emailaddress string `json:"emailAddress,omitempty"`
	Firstname string `json:"firstName,omitempty"`
	Lastname string `json:"lastName,omitempty"`
	Password string `json:"password,omitempty"`
	Userid int64 `json:"userId,omitempty"`
	Company string `json:"company,omitempty"`
	Contactnumber string `json:"contactNumber,omitempty"`
	Creditbalance float64 `json:"creditBalance,omitempty"`
}

// WebServiceSendSmsRequests represents the WebServiceSendSmsRequests schema from the OpenAPI specification
type WebServiceSendSmsRequests struct {
	Defaultdatetosend string `json:"defaultDateToSend,omitempty"`
	Messagesperminute int `json:"messagesPerMinute,omitempty"`
	Sendsmsrequests []WebServiceSendSmsRequest `json:"sendSmsRequests,omitempty"`
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Templated bool `json:"templated,omitempty"`
	Href string `json:"href,omitempty"`
	Rel string `json:"rel,omitempty"`
}

// RestErrorDTO represents the RestErrorDTO schema from the OpenAPI specification
type RestErrorDTO struct {
	Code int `json:"code,omitempty"`
	Developermessage string `json:"developerMessage,omitempty"`
	Message string `json:"message,omitempty"`
	Moreinfourl string `json:"moreInfoUrl,omitempty"`
	Status string `json:"status,omitempty"`
}

// WebServiceCampaignStatistics represents the WebServiceCampaignStatistics schema from the OpenAPI specification
type WebServiceCampaignStatistics struct {
	Campaign string `json:"campaign,omitempty"`
	Statistics WebServiceStatistics `json:"statistics,omitempty"` // WebServiceStatistics
}

// WebServiceGroups represents the WebServiceGroups schema from the OpenAPI specification
type WebServiceGroups struct {
	Links []Link `json:"links,omitempty"`
	Webservicegroups []WebServiceGroup `json:"webServiceGroups,omitempty"`
}

// WebServiceSendSmsResponse represents the WebServiceSendSmsResponse schema from the OpenAPI specification
type WebServiceSendSmsResponse struct {
	ErrorField string `json:"error,omitempty"`
	Messageid string `json:"messageId,omitempty"`
}

// WebServiceAnalyseMessageRequestMessageOnly represents the WebServiceAnalyseMessageRequestMessageOnly schema from the OpenAPI specification
type WebServiceAnalyseMessageRequestMessageOnly struct {
	Message string `json:"message,omitempty"`
	Recipientnumber string `json:"recipientNumber,omitempty"`
}

// WebServiceContact represents the WebServiceContact schema from the OpenAPI specification
type WebServiceContact struct {
	Contactnumber string `json:"contactNumber,omitempty"`
	Firstname string `json:"firstName,omitempty"`
	Lastname string `json:"lastName,omitempty"`
	Links []Link `json:"links,omitempty"`
	Title string `json:"title,omitempty"`
	Contactid string `json:"contactId,omitempty"`
}

// WebServiceMessages represents the WebServiceMessages schema from the OpenAPI specification
type WebServiceMessages struct {
	Webservicemessages []WebServiceMessage `json:"webServiceMessages,omitempty"`
	Elements int `json:"elements,omitempty"`
	Links []Link `json:"links,omitempty"`
	Page int `json:"page,omitempty"`
	Pagesize int `json:"pageSize,omitempty"`
	Totalelements int64 `json:"totalElements,omitempty"`
	Totalpages int `json:"totalPages,omitempty"`
}

// WebServiceSendVoiceMessageResponse represents the WebServiceSendVoiceMessageResponse schema from the OpenAPI specification
type WebServiceSendVoiceMessageResponse struct {
	ErrorField string `json:"error,omitempty"`
	Voicemessageid string `json:"voiceMessageId,omitempty"`
}

// WebServiceSendSmsResponses represents the WebServiceSendSmsResponses schema from the OpenAPI specification
type WebServiceSendSmsResponses struct {
	Sendsmsresponses []WebServiceSendSmsResponse `json:"sendSmsResponses,omitempty"`
}

// WebServiceAccountStatistics represents the WebServiceAccountStatistics schema from the OpenAPI specification
type WebServiceAccountStatistics struct {
	Showingcreditvalue bool `json:"showingCreditValue,omitempty"`
	To string `json:"to,omitempty"`
	Users []WebServiceUserStatistics `json:"users,omitempty"`
	From string `json:"from,omitempty"`
	Grandtotal WebServiceStatistics `json:"grandTotal,omitempty"` // WebServiceStatistics
}

// WebServiceMessageLink represents the WebServiceMessageLink schema from the OpenAPI specification
type WebServiceMessageLink struct {
	Links []Link `json:"links,omitempty"`
	Messageid string `json:"messageId,omitempty"`
}

// WebServiceMessage represents the WebServiceMessage schema from the OpenAPI specification
type WebServiceMessage struct {
	Datetimereceived string `json:"dateTimeReceived,omitempty"`
	Messageid string `json:"messageId,omitempty"`
	Messagestatus string `json:"messageStatus,omitempty"`
	Tonumber string `json:"toNumber,omitempty"`
	Messagetype string `json:"messageType,omitempty"`
	Userdatafield string `json:"userDataField,omitempty"`
	Contact WebServiceContactLink `json:"contact,omitempty"` // WebServiceContactLink
	Datetimescheduled string `json:"dateTimeScheduled,omitempty"`
	Creditcost float64 `json:"creditCost,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Datetimesent string `json:"dateTimeSent,omitempty"`
	Repliedtomessage WebServiceMessageLink `json:"repliedToMessage,omitempty"` // WebServiceMessageLink
	Numberofmessages int `json:"numberOfMessages,omitempty"`
	Campaign string `json:"campaign,omitempty"`
	Fromnumber string `json:"fromNumber,omitempty"`
	Message string `json:"message,omitempty"`
	Links []Link `json:"links,omitempty"`
	Read bool `json:"read,omitempty"`
}

// WebServiceVoiceMessages represents the WebServiceVoiceMessages schema from the OpenAPI specification
type WebServiceVoiceMessages struct {
	Totalpages int `json:"totalPages,omitempty"`
	Elements int `json:"elements,omitempty"`
	Links []Link `json:"links,omitempty"`
	Messages []WebServiceVoiceMessage `json:"messages,omitempty"`
	Page int `json:"page,omitempty"`
	Pagesize int `json:"pageSize,omitempty"`
	Totalelements int64 `json:"totalElements,omitempty"`
}

// WebServiceVoiceMessageSendSingleTextRequest represents the WebServiceVoiceMessageSendSingleTextRequest schema from the OpenAPI specification
type WebServiceVoiceMessageSendSingleTextRequest struct {
	Datafield string `json:"dataField,omitempty"`
	Language string `json:"language,omitempty"`
	Message string `json:"message,omitempty"`
	Recipientnumber string `json:"recipientNumber,omitempty"`
	Retrycount int `json:"retryCount,omitempty"`
	Retrymaximuminterval int `json:"retryMaximumInterval,omitempty"`
	Retryminimuminterval int `json:"retryMinimumInterval,omitempty"`
	Campaign string `json:"campaign,omitempty"`
}

// WebServiceTransferCreditsRequest represents the WebServiceTransferCreditsRequest schema from the OpenAPI specification
type WebServiceTransferCreditsRequest struct {
	Transfertoemailaddress string `json:"transferToEmailAddress,omitempty"`
	Numberofcreditstotransfer int `json:"numberOfCreditsToTransfer,omitempty"`
	Transferfromemailaddress string `json:"transferFromEmailAddress,omitempty"`
}

// WebServiceTemplate represents the WebServiceTemplate schema from the OpenAPI specification
type WebServiceTemplate struct {
	Data string `json:"data,omitempty"`
	Links []Link `json:"links,omitempty"`
	Name string `json:"name,omitempty"`
	Templateid int64 `json:"templateId,omitempty"`
}

// WebServiceAnalyseMessageResponse represents the WebServiceAnalyseMessageResponse schema from the OpenAPI specification
type WebServiceAnalyseMessageResponse struct {
	Messagelengthwithinmaximumallowed bool `json:"messageLengthWithinMaximumAllowed,omitempty"`
	Numberofmessages int `json:"numberOfMessages,omitempty"`
	Characteranalysis [][]interface{} `json:"characterAnalysis,omitempty"`
	Messagecreditcost float64 `json:"messageCreditCost,omitempty"`
	Messageencoding string `json:"messageEncoding,omitempty"`
	Messagelength int `json:"messageLength,omitempty"`
}

// WebServiceUserStatistics represents the WebServiceUserStatistics schema from the OpenAPI specification
type WebServiceUserStatistics struct {
	Campaigns []WebServiceCampaignStatistics `json:"campaigns,omitempty"`
	Total WebServiceStatistics `json:"total,omitempty"` // WebServiceStatistics
	User WebServiceUser `json:"user,omitempty"` // WebServiceUser
}

// WebServiceStatistics represents the WebServiceStatistics schema from the OpenAPI specification
type WebServiceStatistics struct {
	Failed float64 `json:"failed,omitempty"`
	Failedoptout float64 `json:"failedOptout,omitempty"`
	Failedrefunded float64 `json:"failedRefunded,omitempty"`
	Sent float64 `json:"sent,omitempty"`
	Total float64 `json:"total,omitempty"`
	Delivered float64 `json:"delivered,omitempty"`
}

// WebServiceUsers represents the WebServiceUsers schema from the OpenAPI specification
type WebServiceUsers struct {
	Webserviceuserlist []WebServiceUser `json:"webServiceUserList,omitempty"`
}

// WebServiceGroup represents the WebServiceGroup schema from the OpenAPI specification
type WebServiceGroup struct {
	Groupid string `json:"groupId,omitempty"`
	Links []Link `json:"links,omitempty"`
	Name string `json:"name,omitempty"`
}

// WebServiceContacts represents the WebServiceContacts schema from the OpenAPI specification
type WebServiceContacts struct {
	Links []Link `json:"links,omitempty"`
	Webservicecontacts []WebServiceContact `json:"webServiceContacts,omitempty"`
}

// WebServiceContactLink represents the WebServiceContactLink schema from the OpenAPI specification
type WebServiceContactLink struct {
	Links []Link `json:"links,omitempty"`
	Contactid string `json:"contactId,omitempty"`
}
