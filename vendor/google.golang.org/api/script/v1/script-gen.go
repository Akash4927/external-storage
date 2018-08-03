// Package script provides access to the Apps Script API.
//
// See https://developers.google.com/apps-script/api/
//
// Usage example:
//
//   import "google.golang.org/api/script/v1"
//   ...
//   scriptService, err := script.New(oauthHttpClient)
package script // import "google.golang.org/api/script/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "script:v1"
const apiName = "script"
const apiVersion = "v1"
const basePath = "https://script.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Read, send, delete, and manage your email
	MailGoogleComScope = "https://mail.google.com/"

	// Manage your calendars
	WwwGoogleComCalendarFeedsScope = "https://www.google.com/calendar/feeds"

	// Manage your contacts
	WwwGoogleComM8FeedsScope = "https://www.google.com/m8/feeds"

	// View and manage the provisioning of groups on your domain
	AdminDirectoryGroupScope = "https://www.googleapis.com/auth/admin.directory.group"

	// View and manage the provisioning of users on your domain
	AdminDirectoryUserScope = "https://www.googleapis.com/auth/admin.directory.user"

	// View and manage your Google Docs documents
	DocumentsScope = "https://www.googleapis.com/auth/documents"

	// View and manage the files in your Google Drive
	DriveScope = "https://www.googleapis.com/auth/drive"

	// View and manage your forms in Google Drive
	FormsScope = "https://www.googleapis.com/auth/forms"

	// View and manage forms that this application has been installed in
	FormsCurrentonlyScope = "https://www.googleapis.com/auth/forms.currentonly"

	// View and manage your Google Groups
	GroupsScope = "https://www.googleapis.com/auth/groups"

	// View and manage your spreadsheets in Google Drive
	SpreadsheetsScope = "https://www.googleapis.com/auth/spreadsheets"

	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Processes = NewProcessesService(s)
	s.Projects = NewProjectsService(s)
	s.Scripts = NewScriptsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Processes *ProcessesService

	Projects *ProjectsService

	Scripts *ScriptsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProcessesService(s *Service) *ProcessesService {
	rs := &ProcessesService{s: s}
	return rs
}

type ProcessesService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Deployments = NewProjectsDeploymentsService(s)
	rs.Versions = NewProjectsVersionsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Deployments *ProjectsDeploymentsService

	Versions *ProjectsVersionsService
}

func NewProjectsDeploymentsService(s *Service) *ProjectsDeploymentsService {
	rs := &ProjectsDeploymentsService{s: s}
	return rs
}

type ProjectsDeploymentsService struct {
	s *Service
}

func NewProjectsVersionsService(s *Service) *ProjectsVersionsService {
	rs := &ProjectsVersionsService{s: s}
	return rs
}

type ProjectsVersionsService struct {
	s *Service
}

func NewScriptsService(s *Service) *ScriptsService {
	rs := &ScriptsService{s: s}
	return rs
}

type ScriptsService struct {
	s *Service
}

// Content: The Content resource.
type Content struct {
	// Files: The list of script project files.
	// One of the files is a script manifest; it must be named
	// "appsscript",
	// must have type of JSON, and include the manifest configurations for
	// the
	// project.
	Files []*File `json:"files,omitempty"`

	// ScriptId: The script project's Drive ID.
	ScriptId string `json:"scriptId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Files") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Files") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Content) MarshalJSON() ([]byte, error) {
	type NoMethod Content
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateProjectRequest: Request to create a script project.
type CreateProjectRequest struct {
	// ParentId: The Drive ID of a parent file that the created script
	// project is bound to.
	// This is usually the ID of a Google Doc, Google Sheet, Google Form,
	// or
	// Google Slides file. If not set, a standalone script project is
	// created.
	ParentId string `json:"parentId,omitempty"`

	// Title: The title for the project.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ParentId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ParentId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateProjectRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateProjectRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Deployment: Representation of a single script deployment.
type Deployment struct {
	// DeploymentConfig: The deployment configuration.
	DeploymentConfig *DeploymentConfig `json:"deploymentConfig,omitempty"`

	// DeploymentId: The deployment ID for this deployment.
	DeploymentId string `json:"deploymentId,omitempty"`

	// EntryPoints: The deployment's entry points.
	EntryPoints []*EntryPoint `json:"entryPoints,omitempty"`

	// UpdateTime: Last modified date time stamp.
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DeploymentConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeploymentConfig") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Deployment) MarshalJSON() ([]byte, error) {
	type NoMethod Deployment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeploymentConfig: Metadata the defines how a deployment is
// configured.
type DeploymentConfig struct {
	// Description: The description for this deployment.
	Description string `json:"description,omitempty"`

	// ManifestFileName: The manifest file name for this deployment.
	ManifestFileName string `json:"manifestFileName,omitempty"`

	// ScriptId: The script project's Drive ID.
	ScriptId string `json:"scriptId,omitempty"`

	// VersionNumber: The version number on which this deployment is based.
	VersionNumber int64 `json:"versionNumber,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeploymentConfig) MarshalJSON() ([]byte, error) {
	type NoMethod DeploymentConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// EntryPoint: A configuration that defines how a deployment is accessed
// externally.
type EntryPoint struct {
	// AddOn: Add-on properties.
	AddOn *GoogleAppsScriptTypeAddOnEntryPoint `json:"addOn,omitempty"`

	// EntryPointType: The type of the entry point.
	//
	// Possible values:
	//   "ENTRY_POINT_TYPE_UNSPECIFIED" - An unspecified entry point.
	//   "WEB_APP" - A web application entry point.
	//   "EXECUTION_API" - An API executable entry point.
	//   "ADD_ON" - An Add-On entry point.
	EntryPointType string `json:"entryPointType,omitempty"`

	// ExecutionApi: An entry point specification for Apps Script API
	// execution calls.
	ExecutionApi *GoogleAppsScriptTypeExecutionApiEntryPoint `json:"executionApi,omitempty"`

	// WebApp: An entry point specification for web apps.
	WebApp *GoogleAppsScriptTypeWebAppEntryPoint `json:"webApp,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AddOn") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AddOn") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EntryPoint) MarshalJSON() ([]byte, error) {
	type NoMethod EntryPoint
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExecutionError: An object that provides information about the nature
// of an error resulting
// from an attempted execution of a script function using the Apps
// Script API.
// If a run call
// succeeds but the script function (or Apps Script itself) throws an
// exception,
// the response body's error field
// contains a
// Status object. The `Status` object's `details` field
// contains an array with a single one of these `ExecutionError`
// objects.
type ExecutionError struct {
	// ErrorMessage: The error message thrown by Apps Script, usually
	// localized into the user's
	// language.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// ErrorType: The error type, for example `TypeError` or
	// `ReferenceError`. If the error
	// type is unavailable, this field is not included.
	ErrorType string `json:"errorType,omitempty"`

	// ScriptStackTraceElements: An array of objects that provide a stack
	// trace through the script to show
	// where the execution failed, with the deepest call first.
	ScriptStackTraceElements []*ScriptStackTraceElement `json:"scriptStackTraceElements,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ErrorMessage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ErrorMessage") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExecutionError) MarshalJSON() ([]byte, error) {
	type NoMethod ExecutionError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExecutionRequest: A request to run the function in a script. The
// script is identified by the
// specified `script_id`. Executing a function on a script returns
// results
// based on the implementation of the script.
type ExecutionRequest struct {
	// DevMode: If `true` and the user is an owner of the script, the script
	// runs at the
	// most recently saved version rather than the version deployed for use
	// with
	// the Apps Script API. Optional; default is `false`.
	DevMode bool `json:"devMode,omitempty"`

	// Function: The name of the function to execute in the given script.
	// The name does not
	// include parentheses or parameters.
	Function string `json:"function,omitempty"`

	// Parameters: The parameters to be passed to the function being
	// executed. The object type
	// for each parameter should match the expected type in Apps
	// Script.
	// Parameters cannot be Apps Script-specific object types (such as
	// a
	// `Document` or a `Calendar`); they can only be primitive types such
	// as
	// `string`, `number`, `array`, `object`, or `boolean`. Optional.
	Parameters []interface{} `json:"parameters,omitempty"`

	// SessionState: For Android add-ons only. An ID that represents the
	// user's current session
	// in the Android app for Google Docs or Sheets, included as extra data
	// in
	// the
	// [Intent](https://developer.android.com/guide/components/intents-fi
	// lters.html)
	// that launches the add-on. When an Android add-on is run with a
	// session
	// state, it gains the privileges of
	// a
	// [bound](https://developers.google.com/apps-script/guides/bound)
	// scri
	// pt&mdash;that is, it can access information like the user's
	// current
	// cursor position (in Docs) or selected cell (in Sheets). To retrieve
	// the
	// state,
	// call
	// `Intent.getStringExtra("com.google.android.apps.docs.addons.Sessi
	// onState")`.
	// Optional.
	SessionState string `json:"sessionState,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DevMode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DevMode") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExecutionRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ExecutionRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExecutionResponse: An object that provides the return value of a
// function executed using the
// Apps Script API. If the script function returns successfully, the
// response
// body's response field contains this
// `ExecutionResponse` object.
type ExecutionResponse struct {
	// Result: The return value of the script function. The type matches the
	// object type
	// returned in Apps Script. Functions called using the Apps Script API
	// cannot
	// return Apps Script-specific objects (such as a `Document` or a
	// `Calendar`);
	// they can only return primitive types such as a `string`, `number`,
	// `array`,
	// `object`, or `boolean`.
	Result interface{} `json:"result,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Result") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Result") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExecutionResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ExecutionResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// File: An individual file within a script project.
// A file is a third-party source code created by one or
// more
// developers. It can be a server-side JS code, HTML, or a
// configuration file. Each script project can contain multiple files.
type File struct {
	// CreateTime: Creation date timestamp.
	// This read-only field is only visible to users who have
	// WRITER
	// permission for the script project.
	CreateTime string `json:"createTime,omitempty"`

	// FunctionSet: The defined set of functions in the script file, if any.
	FunctionSet *GoogleAppsScriptTypeFunctionSet `json:"functionSet,omitempty"`

	// LastModifyUser: The user who modified the file most recently.
	// This read-only field is only visible to users who have
	// WRITER
	// permission for the script project.
	LastModifyUser *GoogleAppsScriptTypeUser `json:"lastModifyUser,omitempty"`

	// Name: The name of the file. The file extension is not part of the
	// file
	// name, which can be identified from the type field.
	Name string `json:"name,omitempty"`

	// Source: The file content.
	Source string `json:"source,omitempty"`

	// Type: The type of the file.
	//
	// Possible values:
	//   "ENUM_TYPE_UNSPECIFIED" - Undetermined file type; never actually
	// used.
	//   "SERVER_JS" - An Apps Script server-side code file.
	//   "HTML" - A file containing client-side HTML.
	//   "JSON" - A file in JSON format. This type is only used for the
	// script
	// project's manifest. The manifest file content must match
	// the
	// structure of a
	// valid
	// [ScriptManifest](/apps-script/concepts/manifests)
	Type string `json:"type,omitempty"`

	// UpdateTime: Last modified date timestamp.
	// This read-only field is only visible to users who have
	// WRITER
	// permission for the script project.
	UpdateTime string `json:"updateTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *File) MarshalJSON() ([]byte, error) {
	type NoMethod File
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAppsScriptTypeAddOnEntryPoint: An add-on entry point.
type GoogleAppsScriptTypeAddOnEntryPoint struct {
	// AddOnType: The add-on's required list of supported container types.
	//
	// Possible values:
	//   "UNKNOWN_ADDON_TYPE" - Default value, unknown add-on type.
	//   "GMAIL" - Add-on type for Gmail.
	//   "DATA_STUDIO" - Add-on type for Data Studio.
	AddOnType string `json:"addOnType,omitempty"`

	// Description: The add-on's optional description.
	Description string `json:"description,omitempty"`

	// HelpUrl: The add-on's optional help URL.
	HelpUrl string `json:"helpUrl,omitempty"`

	// PostInstallTipUrl: The add-on's required post install tip URL.
	PostInstallTipUrl string `json:"postInstallTipUrl,omitempty"`

	// ReportIssueUrl: The add-on's optional report issue URL.
	ReportIssueUrl string `json:"reportIssueUrl,omitempty"`

	// Title: The add-on's required title.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AddOnType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AddOnType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAppsScriptTypeAddOnEntryPoint) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAppsScriptTypeAddOnEntryPoint
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAppsScriptTypeExecutionApiConfig: API executable entry point
// configuration.
type GoogleAppsScriptTypeExecutionApiConfig struct {
	// Access: Who has permission to run the API executable.
	//
	// Possible values:
	//   "UNKNOWN_ACCESS" - Default value, should not be used.
	//   "MYSELF" - Only the user who deployed the web app or executable can
	// access it.
	// Note that this is not necessarily the owner of the script project.
	//   "DOMAIN" - Only users in the same domain as the user who deployed
	// the web app or
	// executable can access it.
	//   "ANYONE" - Any logged in user can access the web app or executable.
	//   "ANYONE_ANONYMOUS" - Any user, logged in or not, can access the web
	// app or executable.
	Access string `json:"access,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Access") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Access") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAppsScriptTypeExecutionApiConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAppsScriptTypeExecutionApiConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAppsScriptTypeExecutionApiEntryPoint: An API executable entry
// point.
type GoogleAppsScriptTypeExecutionApiEntryPoint struct {
	// EntryPointConfig: The entry point's configuration.
	EntryPointConfig *GoogleAppsScriptTypeExecutionApiConfig `json:"entryPointConfig,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntryPointConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntryPointConfig") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAppsScriptTypeExecutionApiEntryPoint) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAppsScriptTypeExecutionApiEntryPoint
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAppsScriptTypeFunction: Represents a function in a script
// project.
type GoogleAppsScriptTypeFunction struct {
	// Name: The function name in the script project.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAppsScriptTypeFunction) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAppsScriptTypeFunction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAppsScriptTypeFunctionSet: A set of functions. No duplicates
// are permitted.
type GoogleAppsScriptTypeFunctionSet struct {
	// Values: A list of functions composing the set.
	Values []*GoogleAppsScriptTypeFunction `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAppsScriptTypeFunctionSet) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAppsScriptTypeFunctionSet
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAppsScriptTypeProcess: Representation of a single script
// process execution that was started from
// the script editor, a trigger, an application, or using the Apps
// Script API.
// This is distinct from the `Operation`
// resource, which only represents executions started via the Apps
// Script API.
type GoogleAppsScriptTypeProcess struct {
	// Duration: Duration the execution spent executing.
	Duration string `json:"duration,omitempty"`

	// ExecutingUser: User-facing name for the user executing the script.
	ExecutingUser string `json:"executingUser,omitempty"`

	// FunctionName: Name of the function the started the execution.
	FunctionName string `json:"functionName,omitempty"`

	// ProcessStatus: The executions status.
	//
	// Possible values:
	//   "PROCESS_STATUS_UNSPECIFIED" - Unspecified status.
	//   "RUNNING" - The process is currently running.
	//   "PAUSED" - The process has paused.
	//   "COMPLETED" - The process has completed.
	//   "CANCELED" - The process was cancelled.
	//   "FAILED" - The process failed.
	//   "TIMED_OUT" - The process timed out.
	//   "UNKNOWN" - Process status unknown.
	//   "DELAYED" - The process is delayed, waiting for quota.
	ProcessStatus string `json:"processStatus,omitempty"`

	// ProcessType: The executions type.
	//
	// Possible values:
	//   "PROCESS_TYPE_UNSPECIFIED" - Unspecified type.
	//   "ADD_ON" - The process was started from an add-on entry point.
	//   "EXECUTION_API" - The process was started using the Apps Script
	// API.
	//   "TIME_DRIVEN" - The process was started from a time-based trigger.
	//   "TRIGGER" - The process was started from an event-based trigger.
	//   "WEBAPP" - The process was started from a web app entry point.
	//   "EDITOR" - The process was started using the Apps Script IDE.
	//   "SIMPLE_TRIGGER" - The process was started from a GSuite simple
	// trigger.
	//   "MENU" - The process was started from a GSuite menu item.
	ProcessType string `json:"processType,omitempty"`

	// ProjectName: Name of the script being executed.
	ProjectName string `json:"projectName,omitempty"`

	// StartTime: Time the execution started.
	StartTime string `json:"startTime,omitempty"`

	// UserAccessLevel: The executing users access level to the script.
	//
	// Possible values:
	//   "USER_ACCESS_LEVEL_UNSPECIFIED" - User access level unspecified
	//   "NONE" - The user has no access.
	//   "READ" - The user has read-only access.
	//   "WRITE" - The user has write access.
	//   "OWNER" - The user is an owner.
	UserAccessLevel string `json:"userAccessLevel,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Duration") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Duration") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAppsScriptTypeProcess) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAppsScriptTypeProcess
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAppsScriptTypeUser: A simple user profile resource.
type GoogleAppsScriptTypeUser struct {
	// Domain: The user's domain.
	Domain string `json:"domain,omitempty"`

	// Email: The user's identifying email address.
	Email string `json:"email,omitempty"`

	// Name: The user's display name.
	Name string `json:"name,omitempty"`

	// PhotoUrl: The user's photo.
	PhotoUrl string `json:"photoUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Domain") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Domain") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAppsScriptTypeUser) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAppsScriptTypeUser
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAppsScriptTypeWebAppConfig: Web app entry point configuration.
type GoogleAppsScriptTypeWebAppConfig struct {
	// Access: Who has permission to run the web app.
	//
	// Possible values:
	//   "UNKNOWN_ACCESS" - Default value, should not be used.
	//   "MYSELF" - Only the user who deployed the web app or executable can
	// access it.
	// Note that this is not necessarily the owner of the script project.
	//   "DOMAIN" - Only users in the same domain as the user who deployed
	// the web app or
	// executable can access it.
	//   "ANYONE" - Any logged in user can access the web app or executable.
	//   "ANYONE_ANONYMOUS" - Any user, logged in or not, can access the web
	// app or executable.
	Access string `json:"access,omitempty"`

	// ExecuteAs: Who to execute the web app as.
	//
	// Possible values:
	//   "UNKNOWN_EXECUTE_AS" - Default value, should not be used.
	//   "USER_ACCESSING" - The script runs as the user accessing the web
	// app.
	//   "USER_DEPLOYING" - The script runs as the user who deployed the web
	// app. Note that this is
	// not necessarily the owner of the script project.
	ExecuteAs string `json:"executeAs,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Access") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Access") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAppsScriptTypeWebAppConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAppsScriptTypeWebAppConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleAppsScriptTypeWebAppEntryPoint: A web application entry point.
type GoogleAppsScriptTypeWebAppEntryPoint struct {
	// EntryPointConfig: The entry point's configuration.
	EntryPointConfig *GoogleAppsScriptTypeWebAppConfig `json:"entryPointConfig,omitempty"`

	// Url: The URL for the web application.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntryPointConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntryPointConfig") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleAppsScriptTypeWebAppEntryPoint) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleAppsScriptTypeWebAppEntryPoint
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListDeploymentsResponse: Response with the list of deployments for
// the specified Apps Script project.
type ListDeploymentsResponse struct {
	// Deployments: The list of deployments.
	Deployments []*Deployment `json:"deployments,omitempty"`

	// NextPageToken: The token that can be used in the next call to get the
	// next page of
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Deployments") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Deployments") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListDeploymentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListDeploymentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListScriptProcessesResponse: Response with the list of
// Process resources.
type ListScriptProcessesResponse struct {
	// NextPageToken: Token for the next page of results. If empty, there
	// are no more pages
	// remaining.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Processes: List of processes matching request parameters.
	Processes []*GoogleAppsScriptTypeProcess `json:"processes,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListScriptProcessesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListScriptProcessesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListUserProcessesResponse: Response with the list of
// Process resources.
type ListUserProcessesResponse struct {
	// NextPageToken: Token for the next page of results. If empty, there
	// are no more pages
	// remaining.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Processes: List of processes matching request parameters.
	Processes []*GoogleAppsScriptTypeProcess `json:"processes,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListUserProcessesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListUserProcessesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListVersionsResponse: Response with the list of the versions for the
// specified script project.
type ListVersionsResponse struct {
	// NextPageToken: The token use to fetch the next page of records. if
	// not exist in the
	// response, that means no more versions to list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Versions: The list of versions.
	Versions []*Version `json:"versions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListVersionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListVersionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Metrics: Resource containing usage stats for a given script, based on
// the supplied
// filter and mask present in the request.
type Metrics struct {
	// ActiveUsers: Number of active users.
	ActiveUsers []*MetricsValue `json:"activeUsers,omitempty"`

	// FailedExecutions: Number of failed executions.
	FailedExecutions []*MetricsValue `json:"failedExecutions,omitempty"`

	// TotalExecutions: Number of total executions.
	TotalExecutions []*MetricsValue `json:"totalExecutions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ActiveUsers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActiveUsers") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Metrics) MarshalJSON() ([]byte, error) {
	type NoMethod Metrics
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MetricsValue: Metrics value that holds number of executions counted.
type MetricsValue struct {
	// EndTime: Required field indicating the end time of the interval.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: Required field indicating the start time of the interval.
	StartTime string `json:"startTime,omitempty"`

	// Value: Indicates the number of executions counted.
	Value uint64 `json:"value,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MetricsValue) MarshalJSON() ([]byte, error) {
	type NoMethod MetricsValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: A representation of an execution of an Apps Script
// function started with run. The execution response does not arrive
// until the function finishes executing. The maximum execution runtime
// is listed in the [Apps Script quotas
// guide](/apps-script/guides/services/quotas#current_limitations).
// <p>After execution has started, it can have one of four outcomes:</p>
// <ul> <li> If the script function returns successfully, the
//   response field contains an
//   ExecutionResponse object
//   with the function's return value in the object's `result`
// field.</li>
// <li> If the script function (or Apps Script itself) throws an
// exception, the
//   error field contains a
//   Status object. The `Status` object's `details`
//   field contains an array with a single
//   ExecutionError object that
//   provides information about the nature of the error.</li>
// <li> If the execution has not yet completed,
//   the done field is `false` and
//   the neither the `response` nor `error` fields are
// present.</li>
// <li> If the `run` call itself fails (for example, because of a
//   malformed request or an authorization error), the method returns an
// HTTP
//   response code in the 4XX range with a different format for the
// response
//   body. Client libraries automatically convert a 4XX response into
// an
//   exception class.</li>
// </ul>
type Operation struct {
	// Done: This field indicates whether the script execution has
	// completed. A completed execution has a populated `response` field
	// containing the ExecutionResponse from function that was executed.
	Done bool `json:"done,omitempty"`

	// Error: If a `run` call succeeds but the script function (or Apps
	// Script itself) throws an exception, this field contains a Status
	// object. The `Status` object's `details` field contains an array with
	// a single ExecutionError object that provides information about the
	// nature of the error.
	Error *Status `json:"error,omitempty"`

	// Response: If the script function returns successfully, this field
	// contains an ExecutionResponse object with the function's return
	// value.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Project: The script project resource.
type Project struct {
	// CreateTime: When the script was created.
	CreateTime string `json:"createTime,omitempty"`

	// Creator: User who originally created the script.
	Creator *GoogleAppsScriptTypeUser `json:"creator,omitempty"`

	// LastModifyUser: User who last modified the script.
	LastModifyUser *GoogleAppsScriptTypeUser `json:"lastModifyUser,omitempty"`

	// ParentId: The parent's Drive ID that the script will be attached to.
	// This is usually
	// the ID of a Google Document or Google Sheet. This filed is optional,
	// and
	// if not set, a stand-alone script will be created.
	ParentId string `json:"parentId,omitempty"`

	// ScriptId: The script project's Drive ID.
	ScriptId string `json:"scriptId,omitempty"`

	// Title: The title for the project.
	Title string `json:"title,omitempty"`

	// UpdateTime: When the script was last updated.
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Project) MarshalJSON() ([]byte, error) {
	type NoMethod Project
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ScriptStackTraceElement: A stack trace through the script that shows
// where the execution failed.
type ScriptStackTraceElement struct {
	// Function: The name of the function that failed.
	Function string `json:"function,omitempty"`

	// LineNumber: The line number where the script failed.
	LineNumber int64 `json:"lineNumber,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Function") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Function") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ScriptStackTraceElement) MarshalJSON() ([]byte, error) {
	type NoMethod ScriptStackTraceElement
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: If a `run` call succeeds but the script function (or Apps
// Script itself) throws an exception, the response body's error field
// contains this `Status` object.
type Status struct {
	// Code: The status code. For this API, this value either: <ul> <li> 10,
	// indicating a `SCRIPT_TIMEOUT` error,</li> <li> 3, indicating an
	// `INVALID_ARGUMENT` error, or</li> <li> 1, indicating a `CANCELLED`
	// execution.</li> </ul>
	Code int64 `json:"code,omitempty"`

	// Details: An array that contains a single ExecutionError object that
	// provides information about the nature of the error.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which is in English. Any
	// user-facing error message is localized and sent in the details field,
	// or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UpdateDeploymentRequest: Request with deployment information to
// update an existing deployment.
type UpdateDeploymentRequest struct {
	// DeploymentConfig: The deployment configuration.
	DeploymentConfig *DeploymentConfig `json:"deploymentConfig,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeploymentConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeploymentConfig") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *UpdateDeploymentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UpdateDeploymentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Version: A resource representing a script project version. A version
// is a "snapshot"
// of a script project and is similar to a read-only branched release.
// When
// creating deployments, the version to use must be specified.
type Version struct {
	// CreateTime: When the version was created.
	CreateTime string `json:"createTime,omitempty"`

	// Description: The description for this version.
	Description string `json:"description,omitempty"`

	// ScriptId: The script project's Drive ID.
	ScriptId string `json:"scriptId,omitempty"`

	// VersionNumber: The incremental ID that is created by Apps Script when
	// a version is
	// created. This is system assigned number and is immutable once
	// created.
	VersionNumber int64 `json:"versionNumber,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Version) MarshalJSON() ([]byte, error) {
	type NoMethod Version
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "script.processes.list":

type ProcessesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List information about processes made by or on behalf of a
// user,
// such as process type and current status.
func (r *ProcessesService) List() *ProcessesListCall {
	c := &ProcessesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of returned processes per page of results. Defaults to
// 50.
func (c *ProcessesListCall) PageSize(pageSize int64) *ProcessesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The token for
// continuing a previous list request on the next page. This
// should be set to the value of `nextPageToken` from a previous
// response.
func (c *ProcessesListCall) PageToken(pageToken string) *ProcessesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// UserProcessFilterDeploymentId sets the optional parameter
// "userProcessFilter.deploymentId": Optional field used to limit
// returned processes to those originating from
// projects with a specific deployment ID.
func (c *ProcessesListCall) UserProcessFilterDeploymentId(userProcessFilterDeploymentId string) *ProcessesListCall {
	c.urlParams_.Set("userProcessFilter.deploymentId", userProcessFilterDeploymentId)
	return c
}

// UserProcessFilterEndTime sets the optional parameter
// "userProcessFilter.endTime": Optional field used to limit returned
// processes to those that completed
// on or before the given timestamp.
func (c *ProcessesListCall) UserProcessFilterEndTime(userProcessFilterEndTime string) *ProcessesListCall {
	c.urlParams_.Set("userProcessFilter.endTime", userProcessFilterEndTime)
	return c
}

// UserProcessFilterFunctionName sets the optional parameter
// "userProcessFilter.functionName": Optional field used to limit
// returned processes to those originating from
// a script function with the given function name.
func (c *ProcessesListCall) UserProcessFilterFunctionName(userProcessFilterFunctionName string) *ProcessesListCall {
	c.urlParams_.Set("userProcessFilter.functionName", userProcessFilterFunctionName)
	return c
}

// UserProcessFilterProjectName sets the optional parameter
// "userProcessFilter.projectName": Optional field used to limit
// returned processes to those originating from
// projects with project names containing a specific string.
func (c *ProcessesListCall) UserProcessFilterProjectName(userProcessFilterProjectName string) *ProcessesListCall {
	c.urlParams_.Set("userProcessFilter.projectName", userProcessFilterProjectName)
	return c
}

// UserProcessFilterScriptId sets the optional parameter
// "userProcessFilter.scriptId": Optional field used to limit returned
// processes to those originating from
// projects with a specific script ID.
func (c *ProcessesListCall) UserProcessFilterScriptId(userProcessFilterScriptId string) *ProcessesListCall {
	c.urlParams_.Set("userProcessFilter.scriptId", userProcessFilterScriptId)
	return c
}

// UserProcessFilterStartTime sets the optional parameter
// "userProcessFilter.startTime": Optional field used to limit returned
// processes to those that were
// started on or after the given timestamp.
func (c *ProcessesListCall) UserProcessFilterStartTime(userProcessFilterStartTime string) *ProcessesListCall {
	c.urlParams_.Set("userProcessFilter.startTime", userProcessFilterStartTime)
	return c
}

// UserProcessFilterStatuses sets the optional parameter
// "userProcessFilter.statuses": Optional field used to limit returned
// processes to those having one of
// the specified process statuses.
//
// Possible values:
//   "PROCESS_STATUS_UNSPECIFIED"
//   "RUNNING"
//   "PAUSED"
//   "COMPLETED"
//   "CANCELED"
//   "FAILED"
//   "TIMED_OUT"
//   "UNKNOWN"
//   "DELAYED"
func (c *ProcessesListCall) UserProcessFilterStatuses(userProcessFilterStatuses ...string) *ProcessesListCall {
	c.urlParams_.SetMulti("userProcessFilter.statuses", append([]string{}, userProcessFilterStatuses...))
	return c
}

// UserProcessFilterTypes sets the optional parameter
// "userProcessFilter.types": Optional field used to limit returned
// processes to those having one of
// the specified process types.
//
// Possible values:
//   "PROCESS_TYPE_UNSPECIFIED"
//   "ADD_ON"
//   "EXECUTION_API"
//   "TIME_DRIVEN"
//   "TRIGGER"
//   "WEBAPP"
//   "EDITOR"
//   "SIMPLE_TRIGGER"
//   "MENU"
func (c *ProcessesListCall) UserProcessFilterTypes(userProcessFilterTypes ...string) *ProcessesListCall {
	c.urlParams_.SetMulti("userProcessFilter.types", append([]string{}, userProcessFilterTypes...))
	return c
}

// UserProcessFilterUserAccessLevels sets the optional parameter
// "userProcessFilter.userAccessLevels": Optional field used to limit
// returned processes to those having one of
// the specified user access levels.
//
// Possible values:
//   "USER_ACCESS_LEVEL_UNSPECIFIED"
//   "NONE"
//   "READ"
//   "WRITE"
//   "OWNER"
func (c *ProcessesListCall) UserProcessFilterUserAccessLevels(userProcessFilterUserAccessLevels ...string) *ProcessesListCall {
	c.urlParams_.SetMulti("userProcessFilter.userAccessLevels", append([]string{}, userProcessFilterUserAccessLevels...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProcessesListCall) Fields(s ...googleapi.Field) *ProcessesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProcessesListCall) IfNoneMatch(entityTag string) *ProcessesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProcessesListCall) Context(ctx context.Context) *ProcessesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProcessesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProcessesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/processes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.processes.list" call.
// Exactly one of *ListUserProcessesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListUserProcessesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProcessesListCall) Do(opts ...googleapi.CallOption) (*ListUserProcessesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListUserProcessesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List information about processes made by or on behalf of a user,\nsuch as process type and current status.",
	//   "flatPath": "v1/processes",
	//   "httpMethod": "GET",
	//   "id": "script.processes.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of returned processes per page of results. Defaults to\n50.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token for continuing a previous list request on the next page. This\nshould be set to the value of `nextPageToken` from a previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProcessFilter.deploymentId": {
	//       "description": "Optional field used to limit returned processes to those originating from\nprojects with a specific deployment ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProcessFilter.endTime": {
	//       "description": "Optional field used to limit returned processes to those that completed\non or before the given timestamp.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProcessFilter.functionName": {
	//       "description": "Optional field used to limit returned processes to those originating from\na script function with the given function name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProcessFilter.projectName": {
	//       "description": "Optional field used to limit returned processes to those originating from\nprojects with project names containing a specific string.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProcessFilter.scriptId": {
	//       "description": "Optional field used to limit returned processes to those originating from\nprojects with a specific script ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProcessFilter.startTime": {
	//       "description": "Optional field used to limit returned processes to those that were\nstarted on or after the given timestamp.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProcessFilter.statuses": {
	//       "description": "Optional field used to limit returned processes to those having one of\nthe specified process statuses.",
	//       "enum": [
	//         "PROCESS_STATUS_UNSPECIFIED",
	//         "RUNNING",
	//         "PAUSED",
	//         "COMPLETED",
	//         "CANCELED",
	//         "FAILED",
	//         "TIMED_OUT",
	//         "UNKNOWN",
	//         "DELAYED"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "userProcessFilter.types": {
	//       "description": "Optional field used to limit returned processes to those having one of\nthe specified process types.",
	//       "enum": [
	//         "PROCESS_TYPE_UNSPECIFIED",
	//         "ADD_ON",
	//         "EXECUTION_API",
	//         "TIME_DRIVEN",
	//         "TRIGGER",
	//         "WEBAPP",
	//         "EDITOR",
	//         "SIMPLE_TRIGGER",
	//         "MENU"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "userProcessFilter.userAccessLevels": {
	//       "description": "Optional field used to limit returned processes to those having one of\nthe specified user access levels.",
	//       "enum": [
	//         "USER_ACCESS_LEVEL_UNSPECIFIED",
	//         "NONE",
	//         "READ",
	//         "WRITE",
	//         "OWNER"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/processes",
	//   "response": {
	//     "$ref": "ListUserProcessesResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProcessesListCall) Pages(ctx context.Context, f func(*ListUserProcessesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "script.processes.listScriptProcesses":

type ProcessesListScriptProcessesCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// ListScriptProcesses: List information about a script's executed
// processes, such as process type
// and current status.
func (r *ProcessesService) ListScriptProcesses() *ProcessesListScriptProcessesCall {
	c := &ProcessesListScriptProcessesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of returned processes per page of results. Defaults to
// 50.
func (c *ProcessesListScriptProcessesCall) PageSize(pageSize int64) *ProcessesListScriptProcessesCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The token for
// continuing a previous list request on the next page. This
// should be set to the value of `nextPageToken` from a previous
// response.
func (c *ProcessesListScriptProcessesCall) PageToken(pageToken string) *ProcessesListScriptProcessesCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// ScriptId sets the optional parameter "scriptId": The script ID of the
// project whose processes are listed.
func (c *ProcessesListScriptProcessesCall) ScriptId(scriptId string) *ProcessesListScriptProcessesCall {
	c.urlParams_.Set("scriptId", scriptId)
	return c
}

// ScriptProcessFilterDeploymentId sets the optional parameter
// "scriptProcessFilter.deploymentId": Optional field used to limit
// returned processes to those originating from
// projects with a specific deployment ID.
func (c *ProcessesListScriptProcessesCall) ScriptProcessFilterDeploymentId(scriptProcessFilterDeploymentId string) *ProcessesListScriptProcessesCall {
	c.urlParams_.Set("scriptProcessFilter.deploymentId", scriptProcessFilterDeploymentId)
	return c
}

// ScriptProcessFilterEndTime sets the optional parameter
// "scriptProcessFilter.endTime": Optional field used to limit returned
// processes to those that completed
// on or before the given timestamp.
func (c *ProcessesListScriptProcessesCall) ScriptProcessFilterEndTime(scriptProcessFilterEndTime string) *ProcessesListScriptProcessesCall {
	c.urlParams_.Set("scriptProcessFilter.endTime", scriptProcessFilterEndTime)
	return c
}

// ScriptProcessFilterFunctionName sets the optional parameter
// "scriptProcessFilter.functionName": Optional field used to limit
// returned processes to those originating from
// a script function with the given function name.
func (c *ProcessesListScriptProcessesCall) ScriptProcessFilterFunctionName(scriptProcessFilterFunctionName string) *ProcessesListScriptProcessesCall {
	c.urlParams_.Set("scriptProcessFilter.functionName", scriptProcessFilterFunctionName)
	return c
}

// ScriptProcessFilterStartTime sets the optional parameter
// "scriptProcessFilter.startTime": Optional field used to limit
// returned processes to those that were
// started on or after the given timestamp.
func (c *ProcessesListScriptProcessesCall) ScriptProcessFilterStartTime(scriptProcessFilterStartTime string) *ProcessesListScriptProcessesCall {
	c.urlParams_.Set("scriptProcessFilter.startTime", scriptProcessFilterStartTime)
	return c
}

// ScriptProcessFilterStatuses sets the optional parameter
// "scriptProcessFilter.statuses": Optional field used to limit returned
// processes to those having one of
// the specified process statuses.
//
// Possible values:
//   "PROCESS_STATUS_UNSPECIFIED"
//   "RUNNING"
//   "PAUSED"
//   "COMPLETED"
//   "CANCELED"
//   "FAILED"
//   "TIMED_OUT"
//   "UNKNOWN"
//   "DELAYED"
func (c *ProcessesListScriptProcessesCall) ScriptProcessFilterStatuses(scriptProcessFilterStatuses ...string) *ProcessesListScriptProcessesCall {
	c.urlParams_.SetMulti("scriptProcessFilter.statuses", append([]string{}, scriptProcessFilterStatuses...))
	return c
}

// ScriptProcessFilterTypes sets the optional parameter
// "scriptProcessFilter.types": Optional field used to limit returned
// processes to those having one of
// the specified process types.
//
// Possible values:
//   "PROCESS_TYPE_UNSPECIFIED"
//   "ADD_ON"
//   "EXECUTION_API"
//   "TIME_DRIVEN"
//   "TRIGGER"
//   "WEBAPP"
//   "EDITOR"
//   "SIMPLE_TRIGGER"
//   "MENU"
func (c *ProcessesListScriptProcessesCall) ScriptProcessFilterTypes(scriptProcessFilterTypes ...string) *ProcessesListScriptProcessesCall {
	c.urlParams_.SetMulti("scriptProcessFilter.types", append([]string{}, scriptProcessFilterTypes...))
	return c
}

// ScriptProcessFilterUserAccessLevels sets the optional parameter
// "scriptProcessFilter.userAccessLevels": Optional field used to limit
// returned processes to those having one of
// the specified user access levels.
//
// Possible values:
//   "USER_ACCESS_LEVEL_UNSPECIFIED"
//   "NONE"
//   "READ"
//   "WRITE"
//   "OWNER"
func (c *ProcessesListScriptProcessesCall) ScriptProcessFilterUserAccessLevels(scriptProcessFilterUserAccessLevels ...string) *ProcessesListScriptProcessesCall {
	c.urlParams_.SetMulti("scriptProcessFilter.userAccessLevels", append([]string{}, scriptProcessFilterUserAccessLevels...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProcessesListScriptProcessesCall) Fields(s ...googleapi.Field) *ProcessesListScriptProcessesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProcessesListScriptProcessesCall) IfNoneMatch(entityTag string) *ProcessesListScriptProcessesCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProcessesListScriptProcessesCall) Context(ctx context.Context) *ProcessesListScriptProcessesCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProcessesListScriptProcessesCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProcessesListScriptProcessesCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/processes:listScriptProcesses")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.processes.listScriptProcesses" call.
// Exactly one of *ListScriptProcessesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListScriptProcessesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProcessesListScriptProcessesCall) Do(opts ...googleapi.CallOption) (*ListScriptProcessesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListScriptProcessesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List information about a script's executed processes, such as process type\nand current status.",
	//   "flatPath": "v1/processes:listScriptProcesses",
	//   "httpMethod": "GET",
	//   "id": "script.processes.listScriptProcesses",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of returned processes per page of results. Defaults to\n50.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token for continuing a previous list request on the next page. This\nshould be set to the value of `nextPageToken` from a previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scriptId": {
	//       "description": "The script ID of the project whose processes are listed.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scriptProcessFilter.deploymentId": {
	//       "description": "Optional field used to limit returned processes to those originating from\nprojects with a specific deployment ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scriptProcessFilter.endTime": {
	//       "description": "Optional field used to limit returned processes to those that completed\non or before the given timestamp.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scriptProcessFilter.functionName": {
	//       "description": "Optional field used to limit returned processes to those originating from\na script function with the given function name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scriptProcessFilter.startTime": {
	//       "description": "Optional field used to limit returned processes to those that were\nstarted on or after the given timestamp.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scriptProcessFilter.statuses": {
	//       "description": "Optional field used to limit returned processes to those having one of\nthe specified process statuses.",
	//       "enum": [
	//         "PROCESS_STATUS_UNSPECIFIED",
	//         "RUNNING",
	//         "PAUSED",
	//         "COMPLETED",
	//         "CANCELED",
	//         "FAILED",
	//         "TIMED_OUT",
	//         "UNKNOWN",
	//         "DELAYED"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "scriptProcessFilter.types": {
	//       "description": "Optional field used to limit returned processes to those having one of\nthe specified process types.",
	//       "enum": [
	//         "PROCESS_TYPE_UNSPECIFIED",
	//         "ADD_ON",
	//         "EXECUTION_API",
	//         "TIME_DRIVEN",
	//         "TRIGGER",
	//         "WEBAPP",
	//         "EDITOR",
	//         "SIMPLE_TRIGGER",
	//         "MENU"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "scriptProcessFilter.userAccessLevels": {
	//       "description": "Optional field used to limit returned processes to those having one of\nthe specified user access levels.",
	//       "enum": [
	//         "USER_ACCESS_LEVEL_UNSPECIFIED",
	//         "NONE",
	//         "READ",
	//         "WRITE",
	//         "OWNER"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/processes:listScriptProcesses",
	//   "response": {
	//     "$ref": "ListScriptProcessesResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProcessesListScriptProcessesCall) Pages(ctx context.Context, f func(*ListScriptProcessesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "script.projects.create":

type ProjectsCreateCall struct {
	s                    *Service
	createprojectrequest *CreateProjectRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Create: Creates a new, empty script project with no script files and
// a base
// manifest file.
func (r *ProjectsService) Create(createprojectrequest *CreateProjectRequest) *ProjectsCreateCall {
	c := &ProjectsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.createprojectrequest = createprojectrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsCreateCall) Fields(s ...googleapi.Field) *ProjectsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsCreateCall) Context(ctx context.Context) *ProjectsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createprojectrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.create" call.
// Exactly one of *Project or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Project.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsCreateCall) Do(opts ...googleapi.CallOption) (*Project, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Project{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new, empty script project with no script files and a base\nmanifest file.",
	//   "flatPath": "v1/projects",
	//   "httpMethod": "POST",
	//   "id": "script.projects.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/projects",
	//   "request": {
	//     "$ref": "CreateProjectRequest"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   }
	// }

}

// method id "script.projects.get":

type ProjectsGetCall struct {
	s            *Service
	scriptId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a script project's metadata.
func (r *ProjectsService) Get(scriptId string) *ProjectsGetCall {
	c := &ProjectsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetCall) Fields(s ...googleapi.Field) *ProjectsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetCall) IfNoneMatch(entityTag string) *ProjectsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetCall) Context(ctx context.Context) *ProjectsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId": c.scriptId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.get" call.
// Exactly one of *Project or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Project.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGetCall) Do(opts ...googleapi.CallOption) (*Project, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Project{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a script project's metadata.",
	//   "flatPath": "v1/projects/{scriptId}",
	//   "httpMethod": "GET",
	//   "id": "script.projects.get",
	//   "parameterOrder": [
	//     "scriptId"
	//   ],
	//   "parameters": {
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}",
	//   "response": {
	//     "$ref": "Project"
	//   }
	// }

}

// method id "script.projects.getContent":

type ProjectsGetContentCall struct {
	s            *Service
	scriptId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetContent: Gets the content of the script project, including the
// code source and
// metadata for each script file.
func (r *ProjectsService) GetContent(scriptId string) *ProjectsGetContentCall {
	c := &ProjectsGetContentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	return c
}

// VersionNumber sets the optional parameter "versionNumber": The
// version number of the project to retrieve. If not provided,
// the
// project's HEAD version is returned.
func (c *ProjectsGetContentCall) VersionNumber(versionNumber int64) *ProjectsGetContentCall {
	c.urlParams_.Set("versionNumber", fmt.Sprint(versionNumber))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetContentCall) Fields(s ...googleapi.Field) *ProjectsGetContentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetContentCall) IfNoneMatch(entityTag string) *ProjectsGetContentCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetContentCall) Context(ctx context.Context) *ProjectsGetContentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetContentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetContentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/content")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId": c.scriptId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.getContent" call.
// Exactly one of *Content or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Content.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGetContentCall) Do(opts ...googleapi.CallOption) (*Content, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Content{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the content of the script project, including the code source and\nmetadata for each script file.",
	//   "flatPath": "v1/projects/{scriptId}/content",
	//   "httpMethod": "GET",
	//   "id": "script.projects.getContent",
	//   "parameterOrder": [
	//     "scriptId"
	//   ],
	//   "parameters": {
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "versionNumber": {
	//       "description": "The version number of the project to retrieve. If not provided, the\nproject's HEAD version is returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/content",
	//   "response": {
	//     "$ref": "Content"
	//   }
	// }

}

// method id "script.projects.getMetrics":

type ProjectsGetMetricsCall struct {
	s            *Service
	scriptId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetMetrics: Get metrics data for scripts, such as number of
// executions and
// active users.
func (r *ProjectsService) GetMetrics(scriptId string) *ProjectsGetMetricsCall {
	c := &ProjectsGetMetricsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	return c
}

// MetricsFilterDeploymentId sets the optional parameter
// "metricsFilter.deploymentId": Optional field indicating a specific
// deployment to retrieve metrics from.
func (c *ProjectsGetMetricsCall) MetricsFilterDeploymentId(metricsFilterDeploymentId string) *ProjectsGetMetricsCall {
	c.urlParams_.Set("metricsFilter.deploymentId", metricsFilterDeploymentId)
	return c
}

// MetricsGranularity sets the optional parameter "metricsGranularity":
// Required field indicating what granularity of metrics are returned.
//
// Possible values:
//   "UNSPECIFIED_GRANULARITY"
//   "WEEKLY"
//   "DAILY"
func (c *ProjectsGetMetricsCall) MetricsGranularity(metricsGranularity string) *ProjectsGetMetricsCall {
	c.urlParams_.Set("metricsGranularity", metricsGranularity)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetMetricsCall) Fields(s ...googleapi.Field) *ProjectsGetMetricsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetMetricsCall) IfNoneMatch(entityTag string) *ProjectsGetMetricsCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetMetricsCall) Context(ctx context.Context) *ProjectsGetMetricsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetMetricsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetMetricsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/metrics")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId": c.scriptId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.getMetrics" call.
// Exactly one of *Metrics or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Metrics.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGetMetricsCall) Do(opts ...googleapi.CallOption) (*Metrics, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Metrics{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get metrics data for scripts, such as number of executions and\nactive users.",
	//   "flatPath": "v1/projects/{scriptId}/metrics",
	//   "httpMethod": "GET",
	//   "id": "script.projects.getMetrics",
	//   "parameterOrder": [
	//     "scriptId"
	//   ],
	//   "parameters": {
	//     "metricsFilter.deploymentId": {
	//       "description": "Optional field indicating a specific deployment to retrieve metrics from.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "metricsGranularity": {
	//       "description": "Required field indicating what granularity of metrics are returned.",
	//       "enum": [
	//         "UNSPECIFIED_GRANULARITY",
	//         "WEEKLY",
	//         "DAILY"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scriptId": {
	//       "description": "Required field indicating the script to get metrics for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/metrics",
	//   "response": {
	//     "$ref": "Metrics"
	//   }
	// }

}

// method id "script.projects.updateContent":

type ProjectsUpdateContentCall struct {
	s          *Service
	scriptId   string
	content    *Content
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// UpdateContent: Updates the content of the specified script
// project.
// This content is stored as the HEAD version, and is used when the
// script is
// executed as a trigger, in the script editor, in add-on preview mode,
// or as
// a web app or Apps Script API in development mode. This clears all
// the
// existing files in the project.
func (r *ProjectsService) UpdateContent(scriptId string, content *Content) *ProjectsUpdateContentCall {
	c := &ProjectsUpdateContentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	c.content = content
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUpdateContentCall) Fields(s ...googleapi.Field) *ProjectsUpdateContentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsUpdateContentCall) Context(ctx context.Context) *ProjectsUpdateContentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsUpdateContentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsUpdateContentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.content)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/content")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId": c.scriptId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.updateContent" call.
// Exactly one of *Content or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Content.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsUpdateContentCall) Do(opts ...googleapi.CallOption) (*Content, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Content{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the content of the specified script project.\nThis content is stored as the HEAD version, and is used when the script is\nexecuted as a trigger, in the script editor, in add-on preview mode, or as\na web app or Apps Script API in development mode. This clears all the\nexisting files in the project.",
	//   "flatPath": "v1/projects/{scriptId}/content",
	//   "httpMethod": "PUT",
	//   "id": "script.projects.updateContent",
	//   "parameterOrder": [
	//     "scriptId"
	//   ],
	//   "parameters": {
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/content",
	//   "request": {
	//     "$ref": "Content"
	//   },
	//   "response": {
	//     "$ref": "Content"
	//   }
	// }

}

// method id "script.projects.deployments.create":

type ProjectsDeploymentsCreateCall struct {
	s                *Service
	scriptId         string
	deploymentconfig *DeploymentConfig
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Create: Creates a deployment of an Apps Script project.
func (r *ProjectsDeploymentsService) Create(scriptId string, deploymentconfig *DeploymentConfig) *ProjectsDeploymentsCreateCall {
	c := &ProjectsDeploymentsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	c.deploymentconfig = deploymentconfig
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeploymentsCreateCall) Fields(s ...googleapi.Field) *ProjectsDeploymentsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDeploymentsCreateCall) Context(ctx context.Context) *ProjectsDeploymentsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDeploymentsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDeploymentsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deploymentconfig)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/deployments")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId": c.scriptId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.deployments.create" call.
// Exactly one of *Deployment or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Deployment.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsDeploymentsCreateCall) Do(opts ...googleapi.CallOption) (*Deployment, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Deployment{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a deployment of an Apps Script project.",
	//   "flatPath": "v1/projects/{scriptId}/deployments",
	//   "httpMethod": "POST",
	//   "id": "script.projects.deployments.create",
	//   "parameterOrder": [
	//     "scriptId"
	//   ],
	//   "parameters": {
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/deployments",
	//   "request": {
	//     "$ref": "DeploymentConfig"
	//   },
	//   "response": {
	//     "$ref": "Deployment"
	//   }
	// }

}

// method id "script.projects.deployments.delete":

type ProjectsDeploymentsDeleteCall struct {
	s            *Service
	scriptId     string
	deploymentId string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Delete: Deletes a deployment of an Apps Script project.
func (r *ProjectsDeploymentsService) Delete(scriptId string, deploymentId string) *ProjectsDeploymentsDeleteCall {
	c := &ProjectsDeploymentsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	c.deploymentId = deploymentId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeploymentsDeleteCall) Fields(s ...googleapi.Field) *ProjectsDeploymentsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDeploymentsDeleteCall) Context(ctx context.Context) *ProjectsDeploymentsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDeploymentsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDeploymentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/deployments/{deploymentId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId":     c.scriptId,
		"deploymentId": c.deploymentId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.deployments.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDeploymentsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a deployment of an Apps Script project.",
	//   "flatPath": "v1/projects/{scriptId}/deployments/{deploymentId}",
	//   "httpMethod": "DELETE",
	//   "id": "script.projects.deployments.delete",
	//   "parameterOrder": [
	//     "scriptId",
	//     "deploymentId"
	//   ],
	//   "parameters": {
	//     "deploymentId": {
	//       "description": "The deployment ID to be undeployed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/deployments/{deploymentId}",
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "script.projects.deployments.get":

type ProjectsDeploymentsGetCall struct {
	s            *Service
	scriptId     string
	deploymentId string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a deployment of an Apps Script project.
func (r *ProjectsDeploymentsService) Get(scriptId string, deploymentId string) *ProjectsDeploymentsGetCall {
	c := &ProjectsDeploymentsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	c.deploymentId = deploymentId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeploymentsGetCall) Fields(s ...googleapi.Field) *ProjectsDeploymentsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDeploymentsGetCall) IfNoneMatch(entityTag string) *ProjectsDeploymentsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDeploymentsGetCall) Context(ctx context.Context) *ProjectsDeploymentsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDeploymentsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDeploymentsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/deployments/{deploymentId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId":     c.scriptId,
		"deploymentId": c.deploymentId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.deployments.get" call.
// Exactly one of *Deployment or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Deployment.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsDeploymentsGetCall) Do(opts ...googleapi.CallOption) (*Deployment, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Deployment{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a deployment of an Apps Script project.",
	//   "flatPath": "v1/projects/{scriptId}/deployments/{deploymentId}",
	//   "httpMethod": "GET",
	//   "id": "script.projects.deployments.get",
	//   "parameterOrder": [
	//     "scriptId",
	//     "deploymentId"
	//   ],
	//   "parameters": {
	//     "deploymentId": {
	//       "description": "The deployment ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/deployments/{deploymentId}",
	//   "response": {
	//     "$ref": "Deployment"
	//   }
	// }

}

// method id "script.projects.deployments.list":

type ProjectsDeploymentsListCall struct {
	s            *Service
	scriptId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the deployments of an Apps Script project.
func (r *ProjectsDeploymentsService) List(scriptId string) *ProjectsDeploymentsListCall {
	c := &ProjectsDeploymentsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of deployments on each returned page. Defaults to 50.
func (c *ProjectsDeploymentsListCall) PageSize(pageSize int64) *ProjectsDeploymentsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The token for
// continuing a previous list request on the next page. This
// should be set to the value of `nextPageToken` from a previous
// response.
func (c *ProjectsDeploymentsListCall) PageToken(pageToken string) *ProjectsDeploymentsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeploymentsListCall) Fields(s ...googleapi.Field) *ProjectsDeploymentsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDeploymentsListCall) IfNoneMatch(entityTag string) *ProjectsDeploymentsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDeploymentsListCall) Context(ctx context.Context) *ProjectsDeploymentsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDeploymentsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDeploymentsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/deployments")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId": c.scriptId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.deployments.list" call.
// Exactly one of *ListDeploymentsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListDeploymentsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDeploymentsListCall) Do(opts ...googleapi.CallOption) (*ListDeploymentsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListDeploymentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the deployments of an Apps Script project.",
	//   "flatPath": "v1/projects/{scriptId}/deployments",
	//   "httpMethod": "GET",
	//   "id": "script.projects.deployments.list",
	//   "parameterOrder": [
	//     "scriptId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of deployments on each returned page. Defaults to 50.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token for continuing a previous list request on the next page. This\nshould be set to the value of `nextPageToken` from a previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/deployments",
	//   "response": {
	//     "$ref": "ListDeploymentsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsDeploymentsListCall) Pages(ctx context.Context, f func(*ListDeploymentsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "script.projects.deployments.update":

type ProjectsDeploymentsUpdateCall struct {
	s                       *Service
	scriptId                string
	deploymentId            string
	updatedeploymentrequest *UpdateDeploymentRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// Update: Updates a deployment of an Apps Script project.
func (r *ProjectsDeploymentsService) Update(scriptId string, deploymentId string, updatedeploymentrequest *UpdateDeploymentRequest) *ProjectsDeploymentsUpdateCall {
	c := &ProjectsDeploymentsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	c.deploymentId = deploymentId
	c.updatedeploymentrequest = updatedeploymentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeploymentsUpdateCall) Fields(s ...googleapi.Field) *ProjectsDeploymentsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDeploymentsUpdateCall) Context(ctx context.Context) *ProjectsDeploymentsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDeploymentsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDeploymentsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.updatedeploymentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/deployments/{deploymentId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId":     c.scriptId,
		"deploymentId": c.deploymentId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.deployments.update" call.
// Exactly one of *Deployment or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Deployment.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsDeploymentsUpdateCall) Do(opts ...googleapi.CallOption) (*Deployment, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Deployment{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a deployment of an Apps Script project.",
	//   "flatPath": "v1/projects/{scriptId}/deployments/{deploymentId}",
	//   "httpMethod": "PUT",
	//   "id": "script.projects.deployments.update",
	//   "parameterOrder": [
	//     "scriptId",
	//     "deploymentId"
	//   ],
	//   "parameters": {
	//     "deploymentId": {
	//       "description": "The deployment ID for this deployment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/deployments/{deploymentId}",
	//   "request": {
	//     "$ref": "UpdateDeploymentRequest"
	//   },
	//   "response": {
	//     "$ref": "Deployment"
	//   }
	// }

}

// method id "script.projects.versions.create":

type ProjectsVersionsCreateCall struct {
	s          *Service
	scriptId   string
	version    *Version
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new immutable version using the current code, with
// a unique
// version number.
func (r *ProjectsVersionsService) Create(scriptId string, version *Version) *ProjectsVersionsCreateCall {
	c := &ProjectsVersionsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	c.version = version
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsVersionsCreateCall) Fields(s ...googleapi.Field) *ProjectsVersionsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsVersionsCreateCall) Context(ctx context.Context) *ProjectsVersionsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsVersionsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsVersionsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.version)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/versions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId": c.scriptId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.versions.create" call.
// Exactly one of *Version or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Version.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsVersionsCreateCall) Do(opts ...googleapi.CallOption) (*Version, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Version{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new immutable version using the current code, with a unique\nversion number.",
	//   "flatPath": "v1/projects/{scriptId}/versions",
	//   "httpMethod": "POST",
	//   "id": "script.projects.versions.create",
	//   "parameterOrder": [
	//     "scriptId"
	//   ],
	//   "parameters": {
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/versions",
	//   "request": {
	//     "$ref": "Version"
	//   },
	//   "response": {
	//     "$ref": "Version"
	//   }
	// }

}

// method id "script.projects.versions.get":

type ProjectsVersionsGetCall struct {
	s             *Service
	scriptId      string
	versionNumber int64
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
	header_       http.Header
}

// Get: Gets a version of a script project.
func (r *ProjectsVersionsService) Get(scriptId string, versionNumber int64) *ProjectsVersionsGetCall {
	c := &ProjectsVersionsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	c.versionNumber = versionNumber
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsVersionsGetCall) Fields(s ...googleapi.Field) *ProjectsVersionsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsVersionsGetCall) IfNoneMatch(entityTag string) *ProjectsVersionsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsVersionsGetCall) Context(ctx context.Context) *ProjectsVersionsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsVersionsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsVersionsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/versions/{versionNumber}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId":      c.scriptId,
		"versionNumber": strconv.FormatInt(c.versionNumber, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.versions.get" call.
// Exactly one of *Version or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Version.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsVersionsGetCall) Do(opts ...googleapi.CallOption) (*Version, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Version{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a version of a script project.",
	//   "flatPath": "v1/projects/{scriptId}/versions/{versionNumber}",
	//   "httpMethod": "GET",
	//   "id": "script.projects.versions.get",
	//   "parameterOrder": [
	//     "scriptId",
	//     "versionNumber"
	//   ],
	//   "parameters": {
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "versionNumber": {
	//       "description": "The version number.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/versions/{versionNumber}",
	//   "response": {
	//     "$ref": "Version"
	//   }
	// }

}

// method id "script.projects.versions.list":

type ProjectsVersionsListCall struct {
	s            *Service
	scriptId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List the versions of a script project.
func (r *ProjectsVersionsService) List(scriptId string) *ProjectsVersionsListCall {
	c := &ProjectsVersionsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of versions on each returned page. Defaults to 50.
func (c *ProjectsVersionsListCall) PageSize(pageSize int64) *ProjectsVersionsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The token for
// continuing a previous list request on the next page. This
// should be set to the value of `nextPageToken` from a previous
// response.
func (c *ProjectsVersionsListCall) PageToken(pageToken string) *ProjectsVersionsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsVersionsListCall) Fields(s ...googleapi.Field) *ProjectsVersionsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsVersionsListCall) IfNoneMatch(entityTag string) *ProjectsVersionsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsVersionsListCall) Context(ctx context.Context) *ProjectsVersionsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsVersionsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsVersionsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{scriptId}/versions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId": c.scriptId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.projects.versions.list" call.
// Exactly one of *ListVersionsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListVersionsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsVersionsListCall) Do(opts ...googleapi.CallOption) (*ListVersionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListVersionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List the versions of a script project.",
	//   "flatPath": "v1/projects/{scriptId}/versions",
	//   "httpMethod": "GET",
	//   "id": "script.projects.versions.list",
	//   "parameterOrder": [
	//     "scriptId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of versions on each returned page. Defaults to 50.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token for continuing a previous list request on the next page. This\nshould be set to the value of `nextPageToken` from a previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scriptId": {
	//       "description": "The script project's Drive ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{scriptId}/versions",
	//   "response": {
	//     "$ref": "ListVersionsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsVersionsListCall) Pages(ctx context.Context, f func(*ListVersionsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "script.scripts.run":

type ScriptsRunCall struct {
	s                *Service
	scriptId         string
	executionrequest *ExecutionRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Run: Runs a function in an Apps Script project. The project must be
// deployed
// for use with the Apps Script API.
//
// This method requires authorization with an OAuth 2.0 token that
// includes at
// least one of the scopes listed in the
// [Authorization](#authorization)
// section; script projects that do not require authorization cannot
// be
// executed through this API. To find the correct scopes to include in
// the
// authentication token, open the project in the script editor, then
// select
// **File > Project properties** and click the **Scopes** tab.
func (r *ScriptsService) Run(scriptId string, executionrequest *ExecutionRequest) *ScriptsRunCall {
	c := &ScriptsRunCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.scriptId = scriptId
	c.executionrequest = executionrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ScriptsRunCall) Fields(s ...googleapi.Field) *ScriptsRunCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ScriptsRunCall) Context(ctx context.Context) *ScriptsRunCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ScriptsRunCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ScriptsRunCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.executionrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/scripts/{scriptId}:run")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"scriptId": c.scriptId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "script.scripts.run" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ScriptsRunCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Runs a function in an Apps Script project. The project must be deployed\nfor use with the Apps Script API.\n\nThis method requires authorization with an OAuth 2.0 token that includes at\nleast one of the scopes listed in the [Authorization](#authorization)\nsection; script projects that do not require authorization cannot be\nexecuted through this API. To find the correct scopes to include in the\nauthentication token, open the project in the script editor, then select\n**File \u003e Project properties** and click the **Scopes** tab.",
	//   "flatPath": "v1/scripts/{scriptId}:run",
	//   "httpMethod": "POST",
	//   "id": "script.scripts.run",
	//   "parameterOrder": [
	//     "scriptId"
	//   ],
	//   "parameters": {
	//     "scriptId": {
	//       "description": "The script ID of the script to be executed. To find the script ID, open\nthe project in the script editor and select **File \u003e Project properties**.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/scripts/{scriptId}:run",
	//   "request": {
	//     "$ref": "ExecutionRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://mail.google.com/",
	//     "https://www.google.com/calendar/feeds",
	//     "https://www.google.com/m8/feeds",
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/documents",
	//     "https://www.googleapis.com/auth/drive",
	//     "https://www.googleapis.com/auth/forms",
	//     "https://www.googleapis.com/auth/forms.currentonly",
	//     "https://www.googleapis.com/auth/groups",
	//     "https://www.googleapis.com/auth/spreadsheets",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}
