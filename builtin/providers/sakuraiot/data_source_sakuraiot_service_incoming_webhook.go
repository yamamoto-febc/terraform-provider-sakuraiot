package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/sakuraio-api/client"
)

func dataSourceSakuraIoTServiceIncomingWebhook() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSakuraIoTServiceIncomingWebhookRead,
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

func dataSourceSakuraIoTServiceIncomingWebhookRead(d *schema.ResourceData, meta interface{}) error {
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
	serviceType := "incoming-webhook"
	p.SetType(&serviceType)

	services, err := client.FindServices(p)
	if err != nil {
		return fmt.Errorf("Failed on dataSourceSakuraIoTServiceIncomingWebhookRead#FindServices : %s", err)
	}

	if len(services) > 0 {

		// Read IncomingWebhook Detail
		service, err := client.ReadIncomingWebhookService(
			client.NewReadIncomingWebhookServiceParam(services[0].ID),
		)
		if err != nil {
			return fmt.Errorf("Failed on dataSourceSakuraIoTServiceIncomingWebhookRead#ReadIncomingWebhookService : %s", err)
		}

		d.SetId(toStrID(service.ID))
		d.Set("project_id", toStrID(service.Project))
		d.Set("name", service.Name)
		d.Set("secret", service.Secret)
		d.Set("url", service.URL)

	}

	return nil
}
