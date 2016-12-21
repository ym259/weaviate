package events


// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaveEventsListParams creates a new WeaveEventsListParams object
// with the default values initialized.
func NewWeaveEventsListParams() WeaveEventsListParams {
	var (
		altDefault         = string("json")
		prettyPrintDefault = bool(true)
	)
	return WeaveEventsListParams{
		Alt: &altDefault,

		PrettyPrint: &prettyPrintDefault,
	}
}

// WeaveEventsListParams contains all the bound params for the weave events list operation
// typically these are obtained from a http.Request
//
// swagger:parameters weave.events.list
type WeaveEventsListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Data format for the response.
	  In: query
	  Default: "json"
	*/
	Alt *string
	/*Affected command id.
	  In: query
	  Collection Format: multi
	*/
	CommandID []string
	/*Sending or affected device id.
	  In: query
	  Collection Format: multi
	*/
	DeviceID []string
	/*End of time range in ms since epoch.
	  In: query
	*/
	EndTimeMs *string
	/*Selector specifying which fields to include in a partial response.
	  In: query
	*/
	Fields *string
	/*Specifies the language code that should be used for text values in the API response.
	  In: query
	*/
	Hl *string
	/*API key. Your API key identifies your project and provides you with API access, quota, and reports. Required unless you provide an OAuth 2.0 token.
	  In: query
	*/
	Key *string
	/*
	  In: query
	*/
	MaxResults *int64
	/*OAuth 2.0 token for the current user.
	  In: query
	*/
	OauthToken *string
	/*Returns response with indentations and line breaks.
	  In: query
	  Default: true
	*/
	PrettyPrint *bool
	/*Available to use for quota purposes for server-side applications. Can be any arbitrary string assigned to a user, but should not exceed 40 characters. Overrides userIp if both are provided.
	  In: query
	*/
	QuotaUser *string
	/*
	  In: query
	*/
	StartIndex *int64
	/*Start of time range in ms since epoch.
	  In: query
	*/
	StartTimeMs *string
	/*
	  In: query
	*/
	Token *string
	/*Event type.
	  In: query
	*/
	Type *string
	/*IP address of the site where the request originates. Use this if you want to enforce per-user limits.
	  In: query
	*/
	UserIP *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WeaveEventsListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAlt, qhkAlt, _ := qs.GetOK("alt")
	if err := o.bindAlt(qAlt, qhkAlt, route.Formats); err != nil {
		res = append(res, err)
	}

	qCommandID, qhkCommandID, _ := qs.GetOK("commandId")
	if err := o.bindCommandID(qCommandID, qhkCommandID, route.Formats); err != nil {
		res = append(res, err)
	}

	qDeviceID, qhkDeviceID, _ := qs.GetOK("deviceId")
	if err := o.bindDeviceID(qDeviceID, qhkDeviceID, route.Formats); err != nil {
		res = append(res, err)
	}

	qEndTimeMs, qhkEndTimeMs, _ := qs.GetOK("endTimeMs")
	if err := o.bindEndTimeMs(qEndTimeMs, qhkEndTimeMs, route.Formats); err != nil {
		res = append(res, err)
	}

	qFields, qhkFields, _ := qs.GetOK("fields")
	if err := o.bindFields(qFields, qhkFields, route.Formats); err != nil {
		res = append(res, err)
	}

	qHl, qhkHl, _ := qs.GetOK("hl")
	if err := o.bindHl(qHl, qhkHl, route.Formats); err != nil {
		res = append(res, err)
	}

	qKey, qhkKey, _ := qs.GetOK("key")
	if err := o.bindKey(qKey, qhkKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qMaxResults, qhkMaxResults, _ := qs.GetOK("maxResults")
	if err := o.bindMaxResults(qMaxResults, qhkMaxResults, route.Formats); err != nil {
		res = append(res, err)
	}

	qOauthToken, qhkOauthToken, _ := qs.GetOK("oauth_token")
	if err := o.bindOauthToken(qOauthToken, qhkOauthToken, route.Formats); err != nil {
		res = append(res, err)
	}

	qPrettyPrint, qhkPrettyPrint, _ := qs.GetOK("prettyPrint")
	if err := o.bindPrettyPrint(qPrettyPrint, qhkPrettyPrint, route.Formats); err != nil {
		res = append(res, err)
	}

	qQuotaUser, qhkQuotaUser, _ := qs.GetOK("quotaUser")
	if err := o.bindQuotaUser(qQuotaUser, qhkQuotaUser, route.Formats); err != nil {
		res = append(res, err)
	}

	qStartIndex, qhkStartIndex, _ := qs.GetOK("startIndex")
	if err := o.bindStartIndex(qStartIndex, qhkStartIndex, route.Formats); err != nil {
		res = append(res, err)
	}

	qStartTimeMs, qhkStartTimeMs, _ := qs.GetOK("startTimeMs")
	if err := o.bindStartTimeMs(qStartTimeMs, qhkStartTimeMs, route.Formats); err != nil {
		res = append(res, err)
	}

	qToken, qhkToken, _ := qs.GetOK("token")
	if err := o.bindToken(qToken, qhkToken, route.Formats); err != nil {
		res = append(res, err)
	}

	qType, qhkType, _ := qs.GetOK("type")
	if err := o.bindType(qType, qhkType, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserIP, qhkUserIP, _ := qs.GetOK("userIp")
	if err := o.bindUserIP(qUserIP, qhkUserIP, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaveEventsListParams) bindAlt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var altDefault string = string("json")
		o.Alt = &altDefault
		return nil
	}

	o.Alt = &raw

	if err := o.validateAlt(formats); err != nil {
		return err
	}

	return nil
}

func (o *WeaveEventsListParams) validateAlt(formats strfmt.Registry) error {

	if err := validate.Enum("alt", "query", *o.Alt, []interface{}{"json"}); err != nil {
		return err
	}

	return nil
}

func (o *WeaveEventsListParams) bindCommandID(rawData []string, hasKey bool, formats strfmt.Registry) error {

	commandIdIC := rawData

	if len(commandIdIC) == 0 {
		return nil
	}

	var commandIdIR []string
	for _, commandIdIV := range commandIdIC {
		commandIdI := commandIdIV

		commandIdIR = append(commandIdIR, commandIdI)
	}

	o.CommandID = commandIdIR

	return nil
}

func (o *WeaveEventsListParams) bindDeviceID(rawData []string, hasKey bool, formats strfmt.Registry) error {

	deviceIdIC := rawData

	if len(deviceIdIC) == 0 {
		return nil
	}

	var deviceIdIR []string
	for _, deviceIdIV := range deviceIdIC {
		deviceIdI := deviceIdIV

		deviceIdIR = append(deviceIdIR, deviceIdI)
	}

	o.DeviceID = deviceIdIR

	return nil
}

func (o *WeaveEventsListParams) bindEndTimeMs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.EndTimeMs = &raw

	return nil
}

func (o *WeaveEventsListParams) bindFields(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Fields = &raw

	return nil
}

func (o *WeaveEventsListParams) bindHl(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Hl = &raw

	return nil
}

func (o *WeaveEventsListParams) bindKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Key = &raw

	return nil
}

func (o *WeaveEventsListParams) bindMaxResults(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("maxResults", "query", "int64", raw)
	}
	o.MaxResults = &value

	return nil
}

func (o *WeaveEventsListParams) bindOauthToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.OauthToken = &raw

	return nil
}

func (o *WeaveEventsListParams) bindPrettyPrint(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var prettyPrintDefault bool = bool(true)
		o.PrettyPrint = &prettyPrintDefault
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("prettyPrint", "query", "bool", raw)
	}
	o.PrettyPrint = &value

	return nil
}

func (o *WeaveEventsListParams) bindQuotaUser(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.QuotaUser = &raw

	return nil
}

func (o *WeaveEventsListParams) bindStartIndex(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("startIndex", "query", "int64", raw)
	}
	o.StartIndex = &value

	return nil
}

func (o *WeaveEventsListParams) bindStartTimeMs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.StartTimeMs = &raw

	return nil
}

func (o *WeaveEventsListParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Token = &raw

	return nil
}

func (o *WeaveEventsListParams) bindType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Type = &raw

	if err := o.validateType(formats); err != nil {
		return err
	}

	return nil
}

func (o *WeaveEventsListParams) validateType(formats strfmt.Registry) error {

	if err := validate.Enum("type", "query", *o.Type, []interface{}{"adapterDeactivated", "commandCancelled", "commandCreated", "commandDeleted", "commandExpired", "commandUpdated", "deviceAclUpdated", "deviceConnectivityChange", "deviceCreated", "deviceDeleted", "devicePlaceUpdated", "deviceTransferred", "deviceUpdated", "deviceUseTimeUpdated", "deviceUserAclCreated", "deviceUserAclDeleted", "deviceUserAclUpdated", "eventsDeleted", "eventsRecordingDisabled", "eventsRecordingEnabled", "placeCreated", "placeDeleted", "placeMemberAdded", "placeMemberRemoved", "placeUpdated", "roomCreated", "roomDeleted", "roomUpdated"}); err != nil {
		return err
	}

	return nil
}

func (o *WeaveEventsListParams) bindUserIP(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.UserIP = &raw

	return nil
}
