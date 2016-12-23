package sakuraiot

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(APITokenParamName, nil),
				Description: "Your Sakura-IoT API token",
			},
			"secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(APISecretParamName, nil),
				Description: "Your Sakura-IoT API secret",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sakuraiot_project":                  resourceSakuraIoTProject(),
			"sakuraiot_module":                   resourceSakuraIoTModule(),
			"sakuraiot_service_incoming_webhook": resourceSakuraIoTServiceIncomingWebhook(),
			"sakuraiot_service_outgoing_webhook": resourceSakuraIoTServiceOutgoingWebhook(),
			"sakuraiot_service_websocket":        resourceSakuraIoTServiceWebSocket(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"sakuraiot_project":                  dataSourceSakuraIoTProject(),
			"sakuraiot_module":                   dataSourceSakuraIoTModule(),
			"sakuraiot_service_incoming_webhook": dataSourceSakuraIoTServiceIncomingWebhook(),
			"sakuraiot_service_outgoing_webhook": dataSourceSakuraIoTServiceOutgoingWebhook(),
			"sakuraiot_service_websocket":        dataSourceSakuraIoTServiceWebSocket(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		Token:  d.Get("token").(string),
		Secret: d.Get("secret").(string),
	}

	return config.NewClient()
}
