// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/bigmikesolutions/wingman/graphql/model/cursor"
)

type AddDatabaseError struct {
	Code    AddDatabaseClientErrorCode `json:"code"`
	Message *string                    `json:"message,omitempty"`
}

type AddDatabaseInput struct {
	MutationID *string    `json:"mutationId,omitempty"`
	Env        string     `json:"env"`
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	User       string     `json:"user"`
	Password   string     `json:"password"`
	Host       string     `json:"host"`
	Port       int        `json:"port"`
	Driver     DriverType `json:"driver"`
}

type AddDatabasePayload struct {
	MutationID *string           `json:"mutationId,omitempty"`
	ID         string            `json:"id"`
	Error      *AddDatabaseError `json:"error,omitempty"`
}

type AddDatabaseUserRole struct {
	ID             string                 `json:"id"`
	Description    *string                `json:"description,omitempty"`
	DatabaseAccess []*DatabaseAccessInput `json:"databaseAccess,omitempty"`
}

type AddDatabaseUserRoleError struct {
	Code    AddDatabaseUserRoleClientErrorCode `json:"code"`
	Message *string                            `json:"message,omitempty"`
}

type AddDatabaseUserRoleInput struct {
	MutationID  *string                `json:"mutationId,omitempty"`
	Environment string                 `json:"environment"`
	UserRoles   []*AddDatabaseUserRole `json:"userRoles"`
}

type AddDatabaseUserRolePayload struct {
	MutationID *string                   `json:"mutationId,omitempty"`
	UserRoles  []*UserRole               `json:"userRoles,omitempty"`
	Error      *AddDatabaseUserRoleError `json:"error,omitempty"`
}

type AddK8sUserRole struct {
	ID          *string    `json:"id,omitempty"`
	AccessType  AccessType `json:"accessType"`
	Description *string    `json:"description,omitempty"`
	Namespaces  []*string  `json:"namespaces,omitempty"`
	Pods        []*string  `json:"pods,omitempty"`
}

type AddK8sUserRoleError struct {
	Code    AddK8sUserRoleClientErrorCode `json:"code"`
	Message *string                       `json:"message,omitempty"`
}

type AddK8sUserRoleInput struct {
	MutationID *string           `json:"mutationId,omitempty"`
	UserRoles  []*AddK8sUserRole `json:"userRoles"`
}

type AddK8sUserRolePayload struct {
	MutationID *string              `json:"mutationId,omitempty"`
	UserRoles  []*UserRole          `json:"userRoles,omitempty"`
	Error      *AddK8sUserRoleError `json:"error,omitempty"`
}

type AddUserRoleBindingInput struct {
	MutationID *string                   `json:"mutationID,omitempty"`
	Bindings   []*NewUserRoleBindingData `json:"bindings"`
}

type AddUserRoleBindingOutputError struct {
	Code    AddUserRoleBindingClientErrorCode `json:"code"`
	Message *string                           `json:"message,omitempty"`
}

type AddUserRoleBindingPayload struct {
	MutationID *string                        `json:"mutationID,omitempty"`
	Bindings   []*UserRoleBinding             `json:"bindings"`
	Error      *AddUserRoleBindingOutputError `json:"error,omitempty"`
}

type Cluster struct {
	ID        string     `json:"id"`
	Namespace *Namespace `json:"namespace,omitempty"`
}

func (Cluster) IsEntity() {}

// Meta-info about connection, returned by query, holding information about pagination etc.
type ConnectionInfo struct {
	EndCursor   cursor.Cursor `json:"endCursor"`
	HasNextPage bool          `json:"hasNextPage"`
}

type CreateEnvironmentError struct {
	Code    CreateEnvironmentErrorCode `json:"code"`
	Message *string                    `json:"message,omitempty"`
}

type CreateEnvironmentInput struct {
	MutationID  *string `json:"mutationId,omitempty"`
	Env         string  `json:"env"`
	Description *string `json:"description,omitempty"`
}

type CreateEnvironmentPayload struct {
	MutationID *string                 `json:"mutationId,omitempty"`
	Env        string                  `json:"env"`
	Error      *CreateEnvironmentError `json:"error,omitempty"`
}

type Database struct {
	ID    string               `json:"id"`
	Info  *DatabaseInfo        `json:"info,omitempty"`
	Table *TableDataConnection `json:"table"`
}

func (Database) IsEntity() {}

type DatabaseAccess struct {
	ID     string                 `json:"id"`
	Info   *AccessType            `json:"info,omitempty"`
	Tables []*DatabaseTableAccess `json:"tables,omitempty"`
}

type DatabaseAccessInput struct {
	ID     string                      `json:"id"`
	Info   *AccessType                 `json:"info,omitempty"`
	Tables []*DatabaseTableAccessInput `json:"tables,omitempty"`
}

type DatabaseInfo struct {
	ID     string     `json:"id"`
	Host   string     `json:"host"`
	Port   int        `json:"port"`
	Driver DriverType `json:"driver"`
}

type DatabaseResource struct {
	ID    string           `json:"id"`
	Info  *AccessType      `json:"info,omitempty"`
	Table []*TableResource `json:"table,omitempty"`
}

type DatabaseTableAccess struct {
	Name       string     `json:"name"`
	Columns    []*string  `json:"columns,omitempty"`
	AccessType AccessType `json:"accessType"`
}

type DatabaseTableAccessInput struct {
	Name       string     `json:"name"`
	Columns    []*string  `json:"columns,omitempty"`
	AccessType AccessType `json:"accessType"`
}

type EnvGrantError struct {
	Code    EnvGrantErrorCode `json:"code"`
	Message *string           `json:"message,omitempty"`
}

type EnvGrantInput struct {
	MutationID *string               `json:"mutationId,omitempty"`
	Reason     *string               `json:"reason,omitempty"`
	IncidentID *string               `json:"incidentId,omitempty"`
	Resource   []*ResourceGrantInput `json:"resource,omitempty"`
}

type EnvGrantPayload struct {
	MutationID  *string        `json:"mutationId,omitempty"`
	Token       *string        `json:"token,omitempty"`
	Permissions []string       `json:"permissions,omitempty"`
	Error       *EnvGrantError `json:"error,omitempty"`
}

type Environment struct {
	ID          string     `json:"id"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	ModifiedAt  *time.Time `json:"modifiedAt,omitempty"`
	K8s         *Cluster   `json:"k8s,omitempty"`
	Database    *Database  `json:"database,omitempty"`
}

func (Environment) IsEntity() {}

type K8sResource struct {
	ID        *string              `json:"id,omitempty"`
	Namespace []*NamespaceResource `json:"namespace,omitempty"`
}

type Mutation struct {
}

type Namespace struct {
	Name string `json:"name"`
	Pod  *Pod   `json:"pod,omitempty"`
	Pods []*Pod `json:"pods,omitempty"`
}

func (Namespace) IsEntity() {}

type NamespaceResource struct {
	Name string    `json:"name"`
	Pods []*string `json:"pods,omitempty"`
}

type NewUserRoleBindingData struct {
	UserID      string   `json:"userID"`
	RoleIDs     []string `json:"roleIDs"`
	Description *string  `json:"description,omitempty"`
}

type Pod struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Image     string `json:"image"`
}

func (Pod) IsEntity() {}

type Query struct {
}

type ResourceGrantInput struct {
	Env      string               `json:"env"`
	K8s      []*NamespaceResource `json:"k8s,omitempty"`
	Database []*DatabaseResource  `json:"database,omitempty"`
}

type SignInError struct {
	Code    SignInErrorCode `json:"code"`
	Message *string         `json:"message,omitempty"`
}

type SignInInput struct {
	MutationID *string `json:"mutationId,omitempty"`
	Login      string  `json:"login"`
	Password   string  `json:"password"`
}

type SignInOutput struct {
	MutationID *string      `json:"mutationId,omitempty"`
	Token      *string      `json:"token,omitempty"`
	User       *User        `json:"user,omitempty"`
	Error      *SignInError `json:"error,omitempty"`
}

type TableDataConnection struct {
	ConnectionInfo *ConnectionInfo  `json:"connectionInfo"`
	Edges          []*TableDataEdge `json:"edges,omitempty"`
}

type TableDataEdge struct {
	Cursor cursor.Cursor `json:"cursor"`
	Node   *TableRow     `json:"node,omitempty"`
}

type TableFilter struct {
	Columns []*string `json:"columns,omitempty"`
}

type TableResource struct {
	Name       string      `json:"name"`
	Columns    []*string   `json:"columns,omitempty"`
	AccessType *AccessType `json:"accessType,omitempty"`
}

type TableRow struct {
	Index  *int      `json:"index,omitempty"`
	Values []*string `json:"values,omitempty"`
}

type User struct {
	ID          string      `json:"id"`
	Email       string      `json:"email"`
	FirstName   *string     `json:"firstName,omitempty"`
	LastName    *string     `json:"lastName,omitempty"`
	Description *string     `json:"description,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
	ModifiedAt  *time.Time  `json:"modifiedAt,omitempty"`
	Active      *bool       `json:"active,omitempty"`
	UserRoles   []*UserRole `json:"userRoles,omitempty"`
}

func (User) IsEntity() {}

type UserRole struct {
	ID          string            `json:"id"`
	AccessType  AccessType        `json:"accessType"`
	Description *string           `json:"description,omitempty"`
	CreatedAt   time.Time         `json:"createdAt"`
	ModifiedAt  *time.Time        `json:"modifiedAt,omitempty"`
	Namespaces  []*string         `json:"namespaces,omitempty"`
	Pods        []*Pod            `json:"pods,omitempty"`
	Databases   []*DatabaseAccess `json:"databases,omitempty"`
}

func (UserRole) IsEntity() {}

type UserRoleBinding struct {
	ID          *string     `json:"id,omitempty"`
	UserID      string      `json:"userID"`
	UserRoles   []*UserRole `json:"userRoles,omitempty"`
	Description *string     `json:"description,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
	ModifiedAt  *time.Time  `json:"modifiedAt,omitempty"`
}

func (UserRoleBinding) IsEntity() {}

type AccessType string

const (
	AccessTypeReadOnly  AccessType = "ReadOnly"
	AccessTypeWriteOnly AccessType = "WriteOnly"
	AccessTypeReadWrite AccessType = "ReadWrite"
)

var AllAccessType = []AccessType{
	AccessTypeReadOnly,
	AccessTypeWriteOnly,
	AccessTypeReadWrite,
}

func (e AccessType) IsValid() bool {
	switch e {
	case AccessTypeReadOnly, AccessTypeWriteOnly, AccessTypeReadWrite:
		return true
	}
	return false
}

func (e AccessType) String() string {
	return string(e)
}

func (e *AccessType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccessType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccessType", str)
	}
	return nil
}

func (e AccessType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AddDatabaseClientErrorCode string

const (
	AddDatabaseClientErrorCodeInvalidInput  AddDatabaseClientErrorCode = "INVALID_INPUT"
	AddDatabaseClientErrorCodeAlreadyExists AddDatabaseClientErrorCode = "ALREADY_EXISTS"
)

var AllAddDatabaseClientErrorCode = []AddDatabaseClientErrorCode{
	AddDatabaseClientErrorCodeInvalidInput,
	AddDatabaseClientErrorCodeAlreadyExists,
}

func (e AddDatabaseClientErrorCode) IsValid() bool {
	switch e {
	case AddDatabaseClientErrorCodeInvalidInput, AddDatabaseClientErrorCodeAlreadyExists:
		return true
	}
	return false
}

func (e AddDatabaseClientErrorCode) String() string {
	return string(e)
}

func (e *AddDatabaseClientErrorCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AddDatabaseClientErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddDatabaseClientErrorCode", str)
	}
	return nil
}

func (e AddDatabaseClientErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AddDatabaseUserRoleClientErrorCode string

const (
	AddDatabaseUserRoleClientErrorCodeInvalidInput     AddDatabaseUserRoleClientErrorCode = "INVALID_INPUT"
	AddDatabaseUserRoleClientErrorCodeUserNotFound     AddDatabaseUserRoleClientErrorCode = "USER_NOT_FOUND"
	AddDatabaseUserRoleClientErrorCodeUserRoleNotFound AddDatabaseUserRoleClientErrorCode = "USER_ROLE_NOT_FOUND"
	AddDatabaseUserRoleClientErrorCodeProviderError    AddDatabaseUserRoleClientErrorCode = "PROVIDER_ERROR"
	AddDatabaseUserRoleClientErrorCodeGenericError     AddDatabaseUserRoleClientErrorCode = "GENERIC_ERROR"
)

var AllAddDatabaseUserRoleClientErrorCode = []AddDatabaseUserRoleClientErrorCode{
	AddDatabaseUserRoleClientErrorCodeInvalidInput,
	AddDatabaseUserRoleClientErrorCodeUserNotFound,
	AddDatabaseUserRoleClientErrorCodeUserRoleNotFound,
	AddDatabaseUserRoleClientErrorCodeProviderError,
	AddDatabaseUserRoleClientErrorCodeGenericError,
}

func (e AddDatabaseUserRoleClientErrorCode) IsValid() bool {
	switch e {
	case AddDatabaseUserRoleClientErrorCodeInvalidInput, AddDatabaseUserRoleClientErrorCodeUserNotFound, AddDatabaseUserRoleClientErrorCodeUserRoleNotFound, AddDatabaseUserRoleClientErrorCodeProviderError, AddDatabaseUserRoleClientErrorCodeGenericError:
		return true
	}
	return false
}

func (e AddDatabaseUserRoleClientErrorCode) String() string {
	return string(e)
}

func (e *AddDatabaseUserRoleClientErrorCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AddDatabaseUserRoleClientErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddDatabaseUserRoleClientErrorCode", str)
	}
	return nil
}

func (e AddDatabaseUserRoleClientErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AddK8sUserRoleClientErrorCode string

const (
	AddK8sUserRoleClientErrorCodeInvalidInput     AddK8sUserRoleClientErrorCode = "INVALID_INPUT"
	AddK8sUserRoleClientErrorCodeUserNotFound     AddK8sUserRoleClientErrorCode = "USER_NOT_FOUND"
	AddK8sUserRoleClientErrorCodeUserRoleNotFound AddK8sUserRoleClientErrorCode = "USER_ROLE_NOT_FOUND"
	AddK8sUserRoleClientErrorCodeProviderError    AddK8sUserRoleClientErrorCode = "PROVIDER_ERROR"
	AddK8sUserRoleClientErrorCodeGenericError     AddK8sUserRoleClientErrorCode = "GENERIC_ERROR"
)

var AllAddK8sUserRoleClientErrorCode = []AddK8sUserRoleClientErrorCode{
	AddK8sUserRoleClientErrorCodeInvalidInput,
	AddK8sUserRoleClientErrorCodeUserNotFound,
	AddK8sUserRoleClientErrorCodeUserRoleNotFound,
	AddK8sUserRoleClientErrorCodeProviderError,
	AddK8sUserRoleClientErrorCodeGenericError,
}

func (e AddK8sUserRoleClientErrorCode) IsValid() bool {
	switch e {
	case AddK8sUserRoleClientErrorCodeInvalidInput, AddK8sUserRoleClientErrorCodeUserNotFound, AddK8sUserRoleClientErrorCodeUserRoleNotFound, AddK8sUserRoleClientErrorCodeProviderError, AddK8sUserRoleClientErrorCodeGenericError:
		return true
	}
	return false
}

func (e AddK8sUserRoleClientErrorCode) String() string {
	return string(e)
}

func (e *AddK8sUserRoleClientErrorCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AddK8sUserRoleClientErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddK8sUserRoleClientErrorCode", str)
	}
	return nil
}

func (e AddK8sUserRoleClientErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AddUserRoleBindingClientErrorCode string

const (
	AddUserRoleBindingClientErrorCodeInvalidInput     AddUserRoleBindingClientErrorCode = "INVALID_INPUT"
	AddUserRoleBindingClientErrorCodeUserNotFound     AddUserRoleBindingClientErrorCode = "USER_NOT_FOUND"
	AddUserRoleBindingClientErrorCodeUserRoleNotFound AddUserRoleBindingClientErrorCode = "USER_ROLE_NOT_FOUND"
	AddUserRoleBindingClientErrorCodeProviderError    AddUserRoleBindingClientErrorCode = "PROVIDER_ERROR"
	AddUserRoleBindingClientErrorCodeGenericError     AddUserRoleBindingClientErrorCode = "GENERIC_ERROR"
)

var AllAddUserRoleBindingClientErrorCode = []AddUserRoleBindingClientErrorCode{
	AddUserRoleBindingClientErrorCodeInvalidInput,
	AddUserRoleBindingClientErrorCodeUserNotFound,
	AddUserRoleBindingClientErrorCodeUserRoleNotFound,
	AddUserRoleBindingClientErrorCodeProviderError,
	AddUserRoleBindingClientErrorCodeGenericError,
}

func (e AddUserRoleBindingClientErrorCode) IsValid() bool {
	switch e {
	case AddUserRoleBindingClientErrorCodeInvalidInput, AddUserRoleBindingClientErrorCodeUserNotFound, AddUserRoleBindingClientErrorCodeUserRoleNotFound, AddUserRoleBindingClientErrorCodeProviderError, AddUserRoleBindingClientErrorCodeGenericError:
		return true
	}
	return false
}

func (e AddUserRoleBindingClientErrorCode) String() string {
	return string(e)
}

func (e *AddUserRoleBindingClientErrorCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AddUserRoleBindingClientErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddUserRoleBindingClientErrorCode", str)
	}
	return nil
}

func (e AddUserRoleBindingClientErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CreateEnvironmentErrorCode string

const (
	CreateEnvironmentErrorCodeInvalidInput  CreateEnvironmentErrorCode = "INVALID_INPUT"
	CreateEnvironmentErrorCodeUnauthorized  CreateEnvironmentErrorCode = "UNAUTHORIZED"
	CreateEnvironmentErrorCodeAlreadyExists CreateEnvironmentErrorCode = "ALREADY_EXISTS"
)

var AllCreateEnvironmentErrorCode = []CreateEnvironmentErrorCode{
	CreateEnvironmentErrorCodeInvalidInput,
	CreateEnvironmentErrorCodeUnauthorized,
	CreateEnvironmentErrorCodeAlreadyExists,
}

func (e CreateEnvironmentErrorCode) IsValid() bool {
	switch e {
	case CreateEnvironmentErrorCodeInvalidInput, CreateEnvironmentErrorCodeUnauthorized, CreateEnvironmentErrorCodeAlreadyExists:
		return true
	}
	return false
}

func (e CreateEnvironmentErrorCode) String() string {
	return string(e)
}

func (e *CreateEnvironmentErrorCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CreateEnvironmentErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CreateEnvironmentErrorCode", str)
	}
	return nil
}

func (e CreateEnvironmentErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DriverType string

const (
	DriverTypePostgres DriverType = "POSTGRES"
	DriverTypeMysql    DriverType = "MYSQL"
)

var AllDriverType = []DriverType{
	DriverTypePostgres,
	DriverTypeMysql,
}

func (e DriverType) IsValid() bool {
	switch e {
	case DriverTypePostgres, DriverTypeMysql:
		return true
	}
	return false
}

func (e DriverType) String() string {
	return string(e)
}

func (e *DriverType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DriverType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DriverType", str)
	}
	return nil
}

func (e DriverType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EnvGrantErrorCode string

const (
	EnvGrantErrorCodeInvalidInput  EnvGrantErrorCode = "INVALID_INPUT"
	EnvGrantErrorCodeUnauthorized  EnvGrantErrorCode = "UNAUTHORIZED"
	EnvGrantErrorCodeGrantRejected EnvGrantErrorCode = "GRANT_REJECTED"
)

var AllEnvGrantErrorCode = []EnvGrantErrorCode{
	EnvGrantErrorCodeInvalidInput,
	EnvGrantErrorCodeUnauthorized,
	EnvGrantErrorCodeGrantRejected,
}

func (e EnvGrantErrorCode) IsValid() bool {
	switch e {
	case EnvGrantErrorCodeInvalidInput, EnvGrantErrorCodeUnauthorized, EnvGrantErrorCodeGrantRejected:
		return true
	}
	return false
}

func (e EnvGrantErrorCode) String() string {
	return string(e)
}

func (e *EnvGrantErrorCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnvGrantErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnvGrantErrorCode", str)
	}
	return nil
}

func (e EnvGrantErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SignInErrorCode string

const (
	SignInErrorCodeInvalidInput      SignInErrorCode = "INVALID_INPUT"
	SignInErrorCodeWrongUserPassword SignInErrorCode = "WRONG_USER_PASSWORD"
)

var AllSignInErrorCode = []SignInErrorCode{
	SignInErrorCodeInvalidInput,
	SignInErrorCodeWrongUserPassword,
}

func (e SignInErrorCode) IsValid() bool {
	switch e {
	case SignInErrorCodeInvalidInput, SignInErrorCodeWrongUserPassword:
		return true
	}
	return false
}

func (e SignInErrorCode) String() string {
	return string(e)
}

func (e *SignInErrorCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SignInErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SignInErrorCode", str)
	}
	return nil
}

func (e SignInErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
