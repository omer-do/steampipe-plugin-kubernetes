package kubernetes

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type kubernetesConfig struct {
	ConfigPaths          []string `cty:"config_paths"`
	ConfigPath           *string  `cty:"config_path"`
	ConfigContext        *string  `cty:"config_context"`
	CustomResourceTables []string `cty:"custom_resource_tables"`
	ManifestFilePaths    []string `cty:"manifest_file_paths" steampipe:"watch"`
	SourceType           *string  `cty:"source_type"`
	ClusterCACertificate *string  `cty:"cluster_ca_certificate"`
	ClientCertificate    *string  `cty:"client_certificate"`
	Host                 *string  `cty:"host"`
	ClientKey            *string  `cty:"client_key"`
	Token                *string  `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"config_paths": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
	"config_path": {
		Type: schema.TypeString,
	},
	"config_context": {
		Type: schema.TypeString,
	},
	"custom_resource_tables": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
	"manifest_file_paths": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
	"source_type": {
		Type: schema.TypeString,
	},
	"cluster_ca_certificate": {
		Type: schema.TypeString,
	},
	"client_certificate": {
		Type: schema.TypeString,
	},
	"host": {
		Type: schema.TypeString,
	},
	"client_key": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &kubernetesConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) kubernetesConfig {
	if connection == nil || connection.Config == nil {
		return kubernetesConfig{}
	}
	config, _ := connection.Config.(kubernetesConfig)
	return config
}
