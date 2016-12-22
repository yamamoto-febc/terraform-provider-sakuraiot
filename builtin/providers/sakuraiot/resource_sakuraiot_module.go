package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
)

func resourceSakuraIoTModule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSakuraIoTModuleCreate,
		Read:   resourceSakuraIoTModuleRead,
		Update: resourceSakuraIoTModuleUpdate,
		Delete: resourceSakuraIoTModuleDelete,
		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateInt64IDType,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"register_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"register_pass": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_online": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceSakuraIoTModuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	projectID := toInt64ID(d.Get("project_id").(string))
	name := d.Get("name").(string)
	registerID := d.Get("register_id").(string)
	registerPass := d.Get("register_pass").(string)

	p := client.NewCreateModuleParam()
	p.SetRegisterInformation(&models.ModuleRegister{
		Project:      &projectID,
		Name:         &name,
		RegisterID:   &registerID,
		RegisterPass: &registerPass,
	})

	module, err := client.CreateModule(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTModuleCreate#CreateModule : %s", err)
	}

	d.SetId(module.ID)
	return resourceSakuraIoTModuleRead(d, meta)
}

func resourceSakuraIoTModuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewReadModuleParam(d.Id())
	module, err := client.ReadModule(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTModuleRead#ReadModule : %s", err)
	}

	d.Set("project_id", toStrID(module.Project))
	d.Set("name", module.Name)
	d.Set("is_online", module.IsOnline)

	return nil
}

func resourceSakuraIoTModuleUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*SakuraClient)

	if d.HasChange("name") || d.HasChange("project_id") {
		projectID := toInt64ID(d.Get("project_id").(string))
		name := d.Get("name").(string)

		p := client.NewUpdateModuleParam(d.Id())
		p.SetBody(&models.ModuleUpdate{
			Project: &projectID,
			Name:    &name,
		})

		_, err := client.UpdateModule(p)
		if err != nil {
			return fmt.Errorf("Failed on resourceSakuraIoTModuleUpdate#UpdateModule : %s", err)
		}
	}

	return resourceSakuraIoTModuleRead(d, meta)

}

func resourceSakuraIoTModuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewDeleteModuleParam(d.Id())
	err := client.DeleteModule(p)
	if err != nil {
		return fmt.Errorf("Failed on resourceSakuraIoTModuleDelete#DeleteModule : %s", err)
	}

	return nil
}
