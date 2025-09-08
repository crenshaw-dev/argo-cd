/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ArgoCDConfigSpec defines the desired state of ArgoCDConfig.
type ArgoCDConfigSpec struct {
	ResourceCustomizations ResourceCustomizations `json:"resourceCustomizations" protobuf:"bytes,1,opt,name=resourceCustomizations"`
	Manifests              Manifests              `json:"manifests" protobuf:"bytes,2,opt,name=manifests"`
	ResourceTracking       ResourceTracking       `json:"resourceTracking" protobuf:"bytes,3,opt,name=resourceTracking"`
	// InstallationID is an optional unique identifier for this Argo CD installation. It is not set by default.
	InstallationID string            `json:"installationID,omitempty" protobuf:"bytes,4,opt,name=installationID"`
	Server         Server            `json:"server" protobuf:"bytes,5,opt,name=server"`
	Application    ApplicationConfig `json:"application" protobuf:"bytes,6,opt,name=application"`
	GlobalProjects []GlobalProject   `json:"globalProjects,omitempty" protobuf:"bytes,7,rep,name=globalProjects"`
}

// ArgoCDConfigStatus defines the observed state of ArgoCDConfig.
type ArgoCDConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// ArgoCDConfig is the Schema for the argocdconfigs API.
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=argocdconfigs,shortName=argocdconfig;argocdconfigs
type ArgoCDConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ArgoCDConfigSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ArgoCDConfigStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ArgoCDConfigList contains a list of ArgoCDConfig.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ArgoCDConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ArgoCDConfig `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type GlobalProject struct {
	ProjectName   string               `json:"projectName,omitempty" protobuf:"bytes,1,opt,name=projectName"`
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty" protobuf:"bytes,2,opt,name=labelSelector"`
}

// ResourceCustomizations defines special treatment for specific resources.
type ResourceCustomizations struct {
	// Ignore defines resources and fields to ignore during comparison.
	Ignore Ignore `json:"ignore" protobuf:"bytes,1,opt,name=ignore"`
	// Health defines custom health checks for specific resources. Many health checks are built-in and are not defined
	// here. Checks defined here override built-in checks.
	Health []Health `json:"health,omitempty" protobuf:"bytes,2,rep,name=health"`
	// Actions defines custom actions that can be performed on specific resources. Many actions are built-in and are not
	// defined here. Actions defined here override built-in actions.
	Actions []ResourceActionsConfig `json:"actions,omitempty" protobuf:"bytes,3,rep,name=actions"`
	// Exclusions is a list of resources to exclude from Argo CD management. A number of resources are excluded by
	// default. Those defaults can be overridden by specifying them here. If you'd like to add to the defaults, you must
	// patch the field instead of replacing it.
	Exclusions []InclusionOrExclusion `json:"exclusions,omitempty" protobuf:"bytes,4,rep,name=exclusions"`
	Inclusions []InclusionOrExclusion `json:"inclusions,omitempty" protobuf:"bytes,5,rep,name=inclusions"`
}

type Ignore struct {
	Differences Differences `json:"differences,omitempty" protobuf:"bytes,1,opt,name=differences"`
	Updates     Updates     `json:"updates,omitempty" protobuf:"bytes,2,opt,name=updates"`
}

type Differences struct {
	// +listMapKey=group
	// +listMapKey=kind
	// +listType=map
	Rules []Rule `json:"rules" protobuf:"bytes,1,rep,name=rules"`
}

// Rule defines a rule for ignoring differences or allowing updates for specific resources and fields.
type Rule struct {
	Group                 string   `json:"group,omitempty" protobuf:"bytes,1,opt,name=group"`
	Kind                  string   `json:"kind" protobuf:"bytes,2,opt,name=kind"`
	JsonPointers          []string `json:"jsonPointers,omitempty" protobuf:"bytes,3,rep,name=jsonPointers"`
	JqPathExpressions     []string `json:"jqPathExpressions,omitempty" protobuf:"bytes,4,rep,name=jqPathExpressions"`
	ManagedFieldsManagers []string `json:"managedFieldsManagers,omitempty" protobuf:"bytes,5,rep,name=managedFieldsManagers"`
}

type Updates struct {
	// +kubebuilder:default=true
	Enabled bool `json:"enabled" protobuf:"bytes,1,opt,name=enabled"`
	// +listMapKey=group
	// +listMapKey=kind
	// +listType=map
	Rules []Rule `json:"rules" protobuf:"bytes,2,rep,name=rules"`
}

type Health struct {
	Group  string `json:"group" protobuf:"bytes,1,opt,name=group"`
	Kind   string `json:"kind" protobuf:"bytes,2,opt,name=kind"`
	Script string `json:"script" protobuf:"bytes,3,opt,name=script"`
}

type ResourceActionsConfig struct {
	Group     string    `json:"group,omitempty" protobuf:"bytes,1,opt,name=group"`
	Kind      string    `json:"kind" protobuf:"bytes,2,opt,name=kind"`
	Discovery Discovery `json:"discovery,omitempty" protobuf:"bytes,3,opt,name=discovery"`
	// +kubebuilder:validation:MinItems=1
	Actions []Action `json:"actions" protobuf:"bytes,4,rep,name=actions"`
}

type Discovery struct {
	// +kubebuilder:validation:MinLength=1
	Script string `json:"script" protobuf:"bytes,1,opt,name=script"`
}

type Action struct {
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// +kubebuilder:validation:MinLength=1
	Script string `json:"script" protobuf:"bytes,2,opt,name=script"`
}

type InclusionOrExclusion struct {
	Groups []string `json:"groups" protobuf:"bytes,1,rep,name=groups"`
	Kinds  []string `json:"kinds" protobuf:"bytes,2,rep,name=kinds"`
	// Clusters specifies which clusters this inclusion or exclusion applies to. If empty, applies to all clusters.
	Clusters []string `json:"clusters,omitempty" protobuf:"bytes,3,rep,name=clusters"`
}

type Manifests struct {
	Kustomize KustomizeOptions `json:"kustomize" protobuf:"bytes,1,opt,name=kustomize"`
	Jsonnet   Jsonnet          `json:"jsonnet" protobuf:"bytes,2,opt,name=jsonnet"`
	Helm      HelmOptions      `json:"helm" protobuf:"bytes,3,opt,name=helm"`
}

type Jsonnet struct {
	// +kubebuilder:default=true
	Enabled bool `json:"enabled" protobuf:"bytes,1,opt,name=enabled"`
}

type ResourceTracking struct {
	// Method specifies how Argo CD tracks resources. Supported values are "annotation", "label", and "annotation+label".
	// The default is "annotation".
	// +kubebuilder:default=annotation
	// +kubebuilder:validation:Enum=annotation;label;annotation+label
	Method string `json:"method" protobuf:"bytes,1,opt,name=method"`
	// InstanceLabelKey is the label key to use to uniquely identify the instance of an application. The default is "app.kubernetes.io/instance".
	// +kubebuilder:default=app.kubernetes.io/instance
	InstanceLabelKey string `json:"instanceLabelKey" protobuf:"bytes,2,opt,name=instanceLabelKey"`
}

type Server struct {
	// +kubebuilder:validation:Pattern=`^https?://`
	URL string `json:"url" protobuf:"bytes,1,opt,name=url"`
	// +kubebuilder:validation:item:Pattern=`^https?://`
	AdditionalUrls []string       `json:"additionalUrls" protobuf:"bytes,2,rep,name=additionalUrls"`
	Authentication Authentication `json:"authentication" protobuf:"bytes,3,opt,name=authentication"`
	UI             UI             `json:"ui" protobuf:"bytes,4,opt,name=ui"`
	// Max number of pod logs to render in the UI. Default is 10.
	// +kubebuilder:default=10
	MaxPodLogsToRender int64       `json:"maxPodLogsToRender" protobuf:"bytes,5,opt,name=maxPodLogsToRender"`
	Exec               Exec        `json:"exec" protobuf:"bytes,6,opt,name=exec"`
	Extensions         []Extension `json:"extensions" protobuf:"bytes,7,rep,name=extensions"`
	Webhook            Webhook     `json:"webhook" protobuf:"bytes,8,opt,name=webhook"`
	RBAC               RBAC        `json:"rbac" protobuf:"bytes,9,opt,name=rbac"`
}

type RBAC struct {
	// +kubebuilder:default=false
	FineGrainedInheritance FineGrainedInheritance `json:"fineGrainedInheritance" protobuf:"bytes,1,opt,name=fineGrainedInheritance"`
}

type FineGrainedInheritance struct {
	// +kubebuilder:default=false
	Enabled bool `json:"enabled" protobuf:"bytes,1,opt,name=enabled"`
}

type Authentication struct {
	AnonymousUsersEnabled bool `json:"anonymousUsersEnabled" protobuf:"bytes,1,opt,name=anonymousUsersEnabled"`
	// +kubebuilder:default=`24h`
	SessionDuration metav1.Duration `json:"sessionDuration" protobuf:"bytes,2,opt,name=sessionDuration"`
	// +kubebuilder:default=`^.{8,32}$`
	PasswordPattern string    `json:"passwordPattern" protobuf:"bytes,3,opt,name=passwordPattern"`
	AdminEnabled    bool      `json:"adminEnabled" protobuf:"bytes,4,opt,name=adminEnabled"`
	Accounts        []Account `json:"accounts" protobuf:"bytes,5,rep,name=accounts"`
	Dex             DexConfig `json:"dex" protobuf:"bytes,6,opt,name=dex"`
	OIDC            OIDC      `json:"oidc" protobuf:"bytes,7,opt,name=oidc"`
}

type Account struct {
	Name        string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Enabled     bool   `json:"enabled" protobuf:"bytes,2,opt,name=enabled"`
	AllowLogin  bool   `json:"allowLogin" protobuf:"bytes,3,opt,name=allowLogin"`
	AllowApiKey bool   `json:"allowApiKey" protobuf:"bytes,4,opt,name=allowApiKey"`
}

type DexConfig struct {
	Logger        LoggerConfig         `json:"logger,omitempty" protobuf:"bytes,4,opt,name=logger"`
	OAuth2        OAuth2Config         `json:"oauth2" protobuf:"bytes,7,opt,name=oauth2"`
	StaticClients []StaticClientConfig `json:"staticClients" protobuf:"bytes,8,rep,name=staticClients"`
	Connectors    []ConnectorConfig    `json:"connectors" protobuf:"bytes,9,rep,name=connectors"`
}

type StorageConfig struct {
	Type string `json:"type" protobuf:"bytes,1,opt,name=type"`
}

type WebConfig struct {
	HTTP    string `json:"http,omitempty" protobuf:"bytes,1,opt,name=http"`
	HTTPS   string `json:"https,omitempty" protobuf:"bytes,2,opt,name=https"`
	TLSCert string `json:"tlsCert,omitempty" protobuf:"bytes,3,opt,name=tlsCert"`
	TLSKey  string `json:"tlsKey,omitempty" protobuf:"bytes,4,opt,name=tlsKey"`
}

type LoggerConfig struct {
	Level  string `json:"level" protobuf:"bytes,1,opt,name=level"`
	Format string `json:"format" proto:"bytes,2,opt,name=format" protobuf:"bytes,2,opt,name=format"`
}

type OAuth2Config struct {
	SkipApprovalScreen bool `json:"skipApprovalScreen" protobuf:"bytes,1,opt,name=skipApprovalScreen"`
}

type StaticClientConfig struct {
	ID           string   `json:"id" protobuf:"bytes,1,opt,name=id"`
	Name         string   `json:"name" protobuf:"bytes,2,opt,name=name"`
	Secret       string   `json:"secret" protobuf:"bytes,3,opt,name=secret"`
	RedirectURIs []string `json:"redirectURIs" protobuf:"bytes,4,rep,name=redirectURIs"`
	Public       bool     `json:"public,omitempty" protobuf:"bytes,5,opt,name=public"`
}

type ConnectorConfig struct {
	Type   string                `json:"type" protobuf:"bytes,1,opt,name=type"`
	Config ConnectorConfigConfig `json:"config" protobuf:"bytes,2,opt,name=config"`
}

type ConnectorConfigConfig struct {
	RedirectURI string `json:"redirectURI" protobuf:"bytes,1,opt,name=redirectURI"`
}

type OIDC struct {
	Name                    string   `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Issuer                  string   `json:"issuer,omitempty" protobuf:"bytes,2,opt,name=issuer"`
	ClientID                string   `json:"clientID,omitempty" protobuf:"bytes,3,opt,name=clientID"`
	ClientSecret            string   `json:"clientSecret,omitempty" protobuf:"bytes,4,opt,name=clientSecret"`
	CLIClientID             string   `json:"cliClientID,omitempty" protobuf:"bytes,5,opt,name=cliClientID"`
	EnableUserInfoGroups    bool     `json:"enableUserInfoGroups,omitempty" protobuf:"bytes,6,opt,name=enableUserInfoGroups"`
	UserInfoPath            string   `json:"userInfoPath,omitempty" protobuf:"bytes,7,opt,name=userInfoPath"`
	UserInfoCacheExpiration string   `json:"userInfoCacheExpiration,omitempty" protobuf:"bytes,8,opt,name=userInfoCacheExpiration"`
	RequestedScopes         []string `json:"requestedScopes,omitempty" protobuf:"bytes,9,rep,name=requestedScopes"`
	// +listType=map
	// +listMapKey=name
	RequestedIDTokenClaims   []Claim         `json:"requestedIDTokenClaims,omitempty" protobuf:"bytes,10,rep,name=requestedIDTokenClaims"`
	LogoutURL                string          `json:"logoutURL,omitempty" protobuf:"bytes,11,opt,name=logoutURL"`
	RootCA                   string          `json:"rootCA,omitempty" protobuf:"bytes,12,opt,name=rootCA"`
	EnablePKCEAuthentication bool            `json:"enablePKCEAuthentication,omitempty" protobuf:"bytes,13,opt,name=enablePKCEAuthentication"`
	DomainHint               string          `json:"domainHint,omitempty" protobuf:"bytes,14,opt,name=domainHint"`
	Azure                    AzureOIDCConfig `json:"azure,omitempty" protobuf:"bytes,15,opt,name=azure"`
}

type Claim struct {
	Name      string   `json:"name" protobuf:"bytes,1,opt,name=name"`
	Essential bool     `json:"essential,omitempty" protobuf:"bytes,2,opt,name=essential"`
	Value     string   `json:"value,omitempty" protobuf:"bytes,3,opt,name=value"`
	Values    []string `json:"values,omitempty" protobuf:"bytes,4,rep,name=values"`
}

type AzureOIDCConfig struct {
	UseWorkloadIdentity bool `json:"useWorkloadIdentity,omitempty" protobuf:"bytes,1,opt,name=useWorkloadIdentity"`
}

type TLSConfig struct {
	InsecureSkipVerify string `json:"insecureSkipVerify" protobuf:"bytes,1,opt,name=insecureSkipVerify"`
}

type UI struct {
	Analytics                Analytics   `json:"analytics" protobuf:"bytes,1,opt,name=analytics"`
	Help                     Help        `json:"help" protobuf:"bytes,2,opt,name=help"`
	StatusBadge              StatusBadge `json:"statusBadge" protobuf:"bytes,3,opt,name=statusBadge"`
	CssUrl                   string      `json:"cssUrl" protobuf:"bytes,4,opt,name=cssUrl"`
	Banner                   Banner      `json:"banner" protobuf:"bytes,5,opt,name=banner"`
	DeepLinks                DeepLinks   `json:"deepLinks" protobuf:"bytes,6,opt,name=deepLinks"`
	MaskSensitiveAnnotations []string    `json:"maskSensitiveAnnotations" protobuf:"bytes,7,rep,name=maskSensitiveAnnotations"`
	CustomLabels             []string    `json:"customLabels" protobuf:"bytes,8,rep,name=customLabels"`
	// AllowedNodeLabels specifies which node labels may be shown in the UI. By default, none are shown.
	AllowedNodeLabels []string `json:"allowedNodeLabels,omitempty" protobuf:"bytes,9,rep,name=allowedNodeLabels"`
}

type Analytics struct {
	Google GoogleAnalytics `json:"google" protobuf:"bytes,1,opt,name=google"`
}

type GoogleAnalytics struct {
	// +kubebuilder:default=true
	AnonymizeUsers bool   `json:"anonymizeUsers" protobuf:"bytes,1,opt,name=anonymizeUsers"`
	TrackingId     string `json:"trackingId,omitempty" protobuf:"bytes,2,opt,name=trackingId"`
}

type Help struct {
	Chat ChatHelp `json:"chat,omitempty" protobuf:"bytes,1,opt,name=chat"`
}

type ChatHelp struct {
	Url  string `json:"url,omitempty" protobuf:"bytes,1,opt,name=url"`
	Text string `json:"text,omitempty" protobuf:"bytes,2,opt,name=text"`
	// +listType=map
	// +listMapKey=architecture
	Download []Download `json:"download,omitempty" protobuf:"bytes,3,rep,name=download"`
}

type Download struct {
	// +kubebuilder:validation:Enum=darwin-amd64;darwin-arm64;windows-amd64;linux-amd64;linux-arm64;linux-ppc64le;linux-s390x
	Architecture string `json:"architecture" protobuf:"bytes,1,opt,name=architecture"`
	Path         string `json:"path" protobuf:"bytes,2,opt,name=path"`
}

type StatusBadge struct {
	Enabled bool   `json:"enabled" protobuf:"bytes,1,opt,name=enabled"`
	Url     string `json:"url" protobuf:"bytes,2,opt,name=url"`
}

type Banner struct {
	Content string `json:"content" protobuf:"bytes,1,opt,name=content"`
	// +kubebuilder:validation:Pattern=`^https?://`
	Url       string `json:"url" protobuf:"bytes,2,opt,name=url"`
	Permanent bool   `json:"permanent" protobuf:"bytes,3,opt,name=permanent"`
	Position  string `json:"position" protobuf:"bytes,4,opt,name=position"`
}

// DeepLinks defines project, application, and resource-level deep links for the ArgoCD UI.
type DeepLinks struct {
	// Project contains deep links for projects.
	Project []DeepLink `json:"project" protobuf:"bytes,1,rep,name=project"`
	// Application contains deep links for applications.
	Application []DeepLink `json:"application" protobuf:"bytes,2,rep,name=application"`
	// Resource contains deep links for resources.
	Resource []DeepLink `json:"resource" protobuf:"bytes,3,rep,name=resource"`
}

// DeepLink defines a UI deep link, with optional conditional display and icon.
type DeepLink struct {
	// Url is the target URL for the link. Supports Go template expressions.
	Url string `json:"url" protobuf:"bytes,1,opt,name=url"`
	// Title is the display text for the link.
	Title string `json:"title" protobuf:"bytes,2,opt,name=title"`
	// Description is an optional description for the link.
	Description string `json:"description,omitempty" protobuf:"bytes,3,opt,name=description"`
	// IconClass is an optional icon CSS class for the link.
	IconClass string `json:"iconClass,omitempty" protobuf:"bytes,4,opt,name=iconClass"`
	// If is an optional condition for displaying the link, using expr-lang expressions.
	If string `json:"if,omitempty" protobuf:"bytes,5,opt,name=if"`
}

// Exec configures the UI exec feature and allowed shells.
type Exec struct {
	// Enabled indicates whether the UI exec feature is enabled.
	Enabled bool `json:"enabled" protobuf:"bytes,1,opt,name=enabled"`
	// Shells restricts which shells are allowed for exec, and in which order they are attempted.
	// +kubebuilder:default={bash,sh,powershell,cmd}
	Shells []string `json:"shells" protobuf:"bytes,2,rep,name=shells"`
}

// Extension defines an ArgoCD UI/API extension and its backend configuration.
type Extension struct {
	// Name defines the endpoint used to register the extension route.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Backend configures the extension backend connection and services.
	Backend Backend `json:"backend" protobuf:"bytes,2,opt,name=backend"`
}

// Backend configures connection parameters and services for an extension backend.
type Backend struct {
	// ConnectionTimeout is the maximum amount of time a dial to
	// the extension server will wait for a connect to complete.
	// Default: 2 seconds
	// +kubebuilder:default=`2s`
	ConnectionTimeout metav1.Duration `json:"connectionTimeout" protobuf:"bytes,1,opt,name=connectionTimeout"`

	// KeepAlive specifies the interval between keep-alive probes
	// for an active network connection between the API server and
	// the extension server.
	// Default: 15 seconds
	// +kubebuilder:default=`15s`
	KeepAlive metav1.Duration `json:"keepAlive" protobuf:"bytes,2,opt,name=keepAlive"`

	// IdleConnectionTimeout is the maximum amount of time an idle
	// (keep-alive) connection between the API server and the extension
	// server will remain idle before closing itself.
	// Default: 60 seconds
	// +kubebuilder:default=`60s`
	IdleConnectionTimeout metav1.Duration `json:"idleConnectionTimeout" protobuf:"bytes,3,opt,name=idleConnectionTimeout"`

	// MaxIdleConnections controls the maximum number of idle (keep-alive)
	// connections between the API server and the extension server.
	// Default: 30
	// +kubebuilder:default=30
	MaxIdleConnections int64 `json:"maxIdleConnections" protobuf:"bytes,4,opt,name=maxIdleConnections"`
	// Services lists the backend service URLs and cluster mappings.
	Services []ExtensionService `json:"services" protobuf:"bytes,5,rep,name=services"`
}

// ExtensionService defines a backend service URL and optional cluster mapping for extension routing.
type ExtensionService struct {
	// Url is the address where the extension backend service is available.
	Url string `json:"url" protobuf:"bytes,1,opt,name=url"`
	// Cluster maps the service to a specific cluster destination.
	Cluster *ClusterConfiguration `json:"cluster,omitempty" protobuf:"bytes,2,opt,name=cluster"`
	// Headers defines custom headers to include in requests to the backend.
	Headers []Header `json:"headers,omitempty" protobuf:"bytes,3,rep,name=headers"`
}

// Header defines a custom header key and value to include in requests to an extension backend.
type Header struct {
	// Name is the header key.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Value is the header value.
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

// ClusterConfiguration defines a cluster name and server address for service routing.
type ClusterConfiguration struct {
	// Name is the cluster name for service routing.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Server is the cluster server address for service routing.
	Server string `json:"server" protobuf:"bytes,2,opt,name=server"`
}

// Webhook configures the webhook server settings.
type Webhook struct {
	// MaxPayloadSize is the maximum size (in MB) of the payload that can be sent to the webhook server.
	MaxPayloadSize resource.Quantity `json:"maxPayloadSize" protobuf:"bytes,1,opt,name=maxPayloadSize"`
}

// ApplicationConfig configures application-level settings for reconciliation, sync, RBAC, events, and cluster options.
type ApplicationConfig struct {
	// Reconciliation configures the reconciliation timeout and jitter for applications.
	Reconciliation Reconciliation `json:"reconciliation" protobuf:"bytes,1,opt,name=reconciliation"`
	// Sync configures application sync options, including impersonation.
	Sync Sync `json:"sync" protobuf:"bytes,2,opt,name=sync"`
	// RespectRBAC instructs controller to only watch for resources it has permissions to list. Can be "disabled", "normal", or "strict".
	// Disabled by default.
	// +kubebuilder:default=disabled
	// +kubebuilder:validation:Enum=disabled;normal;strict
	RespectRBAC string `json:"respectRBAC" protobuf:"bytes,3,opt,name=respectRBAC"`
	// Events configures which metadata.label keys to include or exclude in Kubernetes events generated for Applications.
	Events Events `json:"events" protobuf:"bytes,4,opt,name=events"`
	// InClusterEnabled indicates whether to allow in-cluster server address. Enabled by default.
	// +kubebuilder:default=true
	InClusterEnabled bool `json:"inClusterEnabled" protobuf:"bytes,5,opt,name=inClusterEnabled"`
}

// Reconciliation configures reconciliation timeout and jitter for applications.
type Reconciliation struct {
	// Timeout is the amount of time before ArgoCD tries to discover new manifest versions. Disabled if set to 0.
	Timeout metav1.Duration `json:"timeout" protobuf:"bytes,1,opt,name=timeout"`
	// Jitter is the maximum duration added to the sync timeout to spread out refreshes. Disabled if set to 0.
	Jitter metav1.Duration `json:"jitter" protobuf:"bytes,2,opt,name=jitter"`
}

// Sync configures application sync options.
type Sync struct {
	// Impersonation enables application sync to use a custom service account via impersonation.
	Impersonation Impersonation `json:"impersonation" protobuf:"bytes,1,opt,name=impersonation"`
}

// Impersonation enables application sync to use a custom service account.
type Impersonation struct {
	// Enabled enables application sync impersonation.
	Enabled bool `json:"enabled" protobuf:"bytes,1,opt,name=enabled"`
}

// Events configures which metadata.label keys to include or exclude in Kubernetes events generated for Applications.
type Events struct {
	// Labels configures which metadata.label keys to include or exclude in events.
	Labels Labels `json:"labels" protobuf:"bytes,1,opt,name=labels"`
}

// Labels configures which metadata.label keys to include or exclude in events.
type Labels struct {
	// Include is a list of metadata.label keys to include in events. Supports wildcards.
	Include []string `json:"include" protobuf:"bytes,1,rep,name=include"`
	// Exclude is a list of metadata.label keys to exclude from events. Supports wildcards.
	Exclude []string `json:"exclude" protobuf:"bytes,2,rep,name=exclude"`
}
