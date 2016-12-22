package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/sakuraio-api/client"
)

func dataSourceSakuraIoTModule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSakuraIoTModuleRead,
		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateInt64IDType,
			},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sort": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateStringInWord(client.ModuleSortCols),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_online": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceSakuraIoTModuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewFindModulesParam()
	if projectID, ok := d.GetOk("project_id"); ok {
		v := toInt64ID(projectID.(string))
		p.SetProject(&v)
	}
	if model, ok := d.GetOk("model"); ok {
		v := model.(string)
		p.SetModel(&v)
	}
	if serial, ok := d.GetOk("serial_number"); ok {
		v := serial.(string)
		p.SetSerialNumber(&v)
	}
	if sort, ok := d.GetOk("sort"); ok {
		v := sort.(string)
		p.SetSort(&v)
	}

	modules, err := client.FindModules(p)
	if err != nil {
		return fmt.Errorf("Failed on dataSourceSakuraIoTModuleRead#FindModules : %s", err)
	}

	if len(modules) > 0 {
		d.SetId(modules[0].ID)
		d.Set("project_id", toStrID(modules[0].Project))
		d.Set("name", modules[0].Name)
		d.Set("is_online", modules[0].IsOnline)
	}
	return nil
}
