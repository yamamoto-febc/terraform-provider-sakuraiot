package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

func resourceSakuraIoTServiceWebSocket() *schema.Resource {
	return &schema.Resource{
		Create: resourceSakuraIoTServiceWebSocketCreate,
		Read:   resourceSakuraIoTServiceWebSocketRead,
		Update: resourceSakuraIoTServiceWebSocketUpdate,
		Delete: resourceSakuraIoTServiceWebSocketDelete,
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
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSakuraIoTServiceWebSocketCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	projectID := toInt64ID(d.Get("project_id").(string))
	name := d.Get("name").(string)
	p := client.NewCreateWebSocketServiceParam()
	p.SetWebSocketService(&models.WebSocketServiceUpdate{
		Name:    &name,
		Project: &projectID,
	})

	service, err := client.CreateWebSocketService(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTServiceWebSocketCreate#CreateServiceWebSocket : %s", err)
	}

	d.SetId(toStrID(service.ID))
	return resourceSakuraIoTServiceWebSocketRead(d, meta)
}

func resourceSakuraIoTServiceWebSocketRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewReadWebSocketServiceParam(toInt64ID(d.Id()))
	service, err := client.ReadWebSocketService(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTServiceWebSocketRead#ReadServiceWebSocket : %s", err)
	}

	d.Set("project_id", toStrID(service.Project))
	d.Set("name", service.Name)
	d.Set("url", service.URL)

	return nil
}

func resourceSakuraIoTServiceWebSocketUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*SakuraClient)

	if d.HasChange("name") {

		projectID := toInt64ID(d.Get("project_id").(string))
		name := d.Get("name").(string)

		p := client.NewUpdateWebSocketServiceParam(toInt64ID(d.Id()))
		p.SetWebSocketService(&models.WebSocketServiceUpdate{
			Name:    &name,
			Project: &projectID,
		})

		_, err := client.UpdateWebSocketService(p)
		if err != nil {
			return fmt.Errorf("Failed on resourceSakuraIoTServiceWebSocketUpdate#UpdateServiceWebSocket : %s", err)
		}
	}

	return resourceSakuraIoTServiceWebSocketRead(d, meta)

}

func resourceSakuraIoTServiceWebSocketDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewDeleteServiceParam(toInt64ID(d.Id()))
	err := client.DeleteService(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTServiceWebSocketDelete#DeleteServiceWebSocket : %s", err)
	}

	return nil
}
