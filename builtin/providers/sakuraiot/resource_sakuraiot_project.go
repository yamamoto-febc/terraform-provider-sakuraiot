package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

func resourceSakuraIoTProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceSakuraIoTProjectCreate,
		Read:   resourceSakuraIoTProjectRead,
		Update: resourceSakuraIoTProjectUpdate,
		Delete: resourceSakuraIoTProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSakuraIoTProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	name := d.Get("name").(string)
	p := client.NewCreateProjectParam()
	p.SetProject(&models.ProjectUpdate{
		Name: &name,
	})

	project, err := client.CreateProject(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTProjectCreate#CreateProject : %s", err)
	}

	d.SetId(toStrID(project.ID))
	return resourceSakuraIoTProjectRead(d, meta)
}

func resourceSakuraIoTProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewReadProjectParam(toInt64ID(d.Id()))
	project, err := client.ReadProject(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTProjectRead#ReadProject : %s", err)
	}

	d.Set("name", project.Name)

	return nil
}

func resourceSakuraIoTProjectUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*SakuraClient)

	if d.HasChange("name") {
		name := d.Get("name").(string)

		p := client.NewUpdateProjectParam(toInt64ID(d.Id()))
		p.SetProject(&models.ProjectUpdate{
			Name: &name,
		})

		_, err := client.UpdateProject(p)
		if err != nil {
			return fmt.Errorf("Failed on resourceSakuraIoTProjectUpdate#UpdateProject : %s", err)
		}
	}

	return resourceSakuraIoTProjectRead(d, meta)

}

func resourceSakuraIoTProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewDeleteProjectParam(toInt64ID(d.Id()))
	err := client.DeleteProject(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTProjectDelete#DeleteProject : %s", err)
	}

	return nil
}
