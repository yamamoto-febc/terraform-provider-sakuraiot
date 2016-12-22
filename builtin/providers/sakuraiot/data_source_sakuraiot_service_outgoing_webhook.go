package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/sakuraio-api/client"
)

func dataSourceSakuraIoTServiceOutgoingWebhook() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSakuraIoTServiceOutgoingWebhookRead,
		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInt64IDType,
			},
			"sort": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateStringInWord(client.ServiceSortCols),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSakuraIoTServiceOutgoingWebhookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewFindServicesParam()
	if projectID, ok := d.GetOk("project_id"); ok {
		v := toInt64ID(projectID.(string))
		p.SetProject(&v)
	}
	if sort, ok := d.GetOk("sort"); ok {
		v := sort.(string)
		p.SetSort(&v)
	}
	serviceType := "outgoing-webhook"
	p.SetType(&serviceType)

	services, err := client.FindServices(p)
	if err != nil {
		return fmt.Errorf("Failed on dataSourceSakuraIoTServiceOutgoingWebhookRead#FindServices : %s", err)
	}

	if len(services) > 0 {

		// Read OutgoingWebhook Detail
		service, err := client.ReadOutgoingWebhookService(
			client.NewReadOutgoingWebhookServiceParam(services[0].ID),
		)
		if err != nil {
			return fmt.Errorf("Failed on dataSourceSakuraIoTServiceOutgoingWebhookRead#ReadOutgoingWebhookService : %s", err)
		}

		d.SetId(toStrID(service.ID))
		d.Set("project_id", toStrID(service.Project))
		d.Set("name", service.Name)
		d.Set("secret", service.Secret)
		d.Set("url", service.URL)

	}

	return nil
}
