package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

func resourceSakuraIoTServiceOutgoingWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceSakuraIoTServiceOutgoingWebhookCreate,
		Read:   resourceSakuraIoTServiceOutgoingWebhookRead,
		Update: resourceSakuraIoTServiceOutgoingWebhookUpdate,
		Delete: resourceSakuraIoTServiceOutgoingWebhookDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInt64IDType,
				ForceNew:     true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSakuraIoTServiceOutgoingWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	projectID := toInt64ID(d.Get("project_id").(string))
	name := d.Get("name").(string)
	secret := d.Get("secret").(string)
	url := d.Get("url").(string)

	p := client.NewCreateOutgoingWebhookServiceParam()
	p.SetOutgoingWebhookService(&models.OutgoingWebhookServiceUpdate{
		Name:    &name,
		Project: &projectID,
		Secret:  secret,
		URL:     &url,
	})

	service, err := client.CreateOutgoingWebhookService(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTServiceOutgoingWebhookCreate#CreateServiceOutgoingWebhook : %s", err)
	}

	d.SetId(toStrID(service.ID))
	return resourceSakuraIoTServiceOutgoingWebhookRead(d, meta)
}

func resourceSakuraIoTServiceOutgoingWebhookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewReadOutgoingWebhookServiceParam(toInt64ID(d.Id()))
	service, err := client.ReadOutgoingWebhookService(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTServiceOutgoingWebhookRead#ReadServiceOutgoingWebhook : %s", err)
	}

	d.Set("project_id", toStrID(service.Project))
	d.Set("name", service.Name)
	d.Set("secret", service.Secret)
	d.Set("url", service.URL)

	return nil
}

func resourceSakuraIoTServiceOutgoingWebhookUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*SakuraClient)

	if d.HasChange("name") || d.HasChange("secret") || d.HasChange("url") {

		projectID := toInt64ID(d.Get("project_id").(string))
		name := d.Get("name").(string)
		secret := d.Get("secret").(string)
		url := d.Get("url").(string)

		p := client.NewUpdateOutgoingWebhookServiceParam(toInt64ID(d.Id()))
		p.SetOutgoingWebhookService(&models.OutgoingWebhookServiceUpdate{
			Name:    &name,
			Project: &projectID,
			Secret:  secret,
			URL:     &url,
		})

		_, err := client.UpdateOutgoingWebhookService(p)
		if err != nil {
			return fmt.Errorf("Failed on resourceSakuraIoTServiceOutgoingWebhookUpdate#UpdateServiceOutgoingWebhook : %s", err)
		}
	}

	return resourceSakuraIoTServiceOutgoingWebhookRead(d, meta)

}

func resourceSakuraIoTServiceOutgoingWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewDeleteServiceParam(toInt64ID(d.Id()))
	err := client.DeleteService(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTServiceOutgoingWebhookDelete#DeleteServiceOutgoingWebhook : %s", err)
	}

	return nil
}
