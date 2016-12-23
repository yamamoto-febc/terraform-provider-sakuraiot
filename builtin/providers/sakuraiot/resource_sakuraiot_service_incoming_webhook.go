package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

func resourceSakuraIoTServiceIncomingWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceSakuraIoTServiceIncomingWebhookCreate,
		Read:   resourceSakuraIoTServiceIncomingWebhookRead,
		Update: resourceSakuraIoTServiceIncomingWebhookUpdate,
		Delete: resourceSakuraIoTServiceIncomingWebhookDelete,
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
				Computed: true,
			},
		},
	}
}

func resourceSakuraIoTServiceIncomingWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	projectID := toInt64ID(d.Get("project_id").(string))
	name := d.Get("name").(string)
	secret := d.Get("secret").(string)

	p := client.NewCreateIncomingWebhookServiceParam()
	p.SetIncomingWebhookService(&models.IncomingWebhookServiceUpdate{
		Name:    &name,
		Project: &projectID,
		Secret:  secret,
	})

	service, err := client.CreateIncomingWebhookService(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTServiceIncomingWebhookCreate#CreateServiceIncomingWebhook : %s", err)
	}
	d.SetId(toStrID(service.ID))
	return resourceSakuraIoTServiceIncomingWebhookRead(d, meta)
}

func resourceSakuraIoTServiceIncomingWebhookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewReadIncomingWebhookServiceParam(toInt64ID(d.Id()))
	service, err := client.ReadIncomingWebhookService(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTServiceIncomingWebhookRead#ReadServiceIncomingWebhook : %s", err)
	}

	d.Set("project_id", toStrID(service.Project))
	d.Set("name", service.Name)
	d.Set("secret", service.Secret)
	d.Set("url", service.URL)

	return nil
}

func resourceSakuraIoTServiceIncomingWebhookUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*SakuraClient)

	if d.HasChange("name") || d.HasChange("secret") {

		projectID := toInt64ID(d.Get("project_id").(string))
		name := d.Get("name").(string)
		secret := d.Get("secret").(string)

		p := client.NewUpdateIncomingWebhookServiceParam(toInt64ID(d.Id()))
		p.SetIncomingWebhookService(&models.IncomingWebhookServiceUpdate{
			Name:    &name,
			Project: &projectID,
			Secret:  secret,
		})

		_, err := client.UpdateIncomingWebhookService(p)
		if err != nil {
			return fmt.Errorf("Failed on resourceSakuraIoTServiceIncomingWebhookUpdate#UpdateServiceIncomingWebhook : %s", err)
		}
	}

	return resourceSakuraIoTServiceIncomingWebhookRead(d, meta)

}

func resourceSakuraIoTServiceIncomingWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewDeleteServiceParam(toInt64ID(d.Id()))
	err := client.DeleteService(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTServiceIncomingWebhookDelete#DeleteServiceIncomingWebhook : %s", err)
	}

	return nil
}
