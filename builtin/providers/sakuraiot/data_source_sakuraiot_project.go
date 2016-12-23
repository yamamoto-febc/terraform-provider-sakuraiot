package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/sakuraio-api/client"
)

func dataSourceSakuraIoTProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSakuraIoTProjectRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sort": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateStringInWord(client.ProjectSortCols),
			},
		},
	}
}

func dataSourceSakuraIoTProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*SakuraClient)

	p := client.NewFindProjectsParam()
	if name, ok := d.GetOk("name"); ok {
		v := name.(string)
		p.SetName(&v)
	}
	if sort, ok := d.GetOk("sort"); ok {
		v := sort.(string)
		p.SetSort(&v)
	}

	projects, err := client.FindProjects(p)
	if err != nil {
		return fmt.Errorf("Failed on dataSourceSakuraIoTProjectRead#FindProjects : %s", err)
	}

	if len(projects) > 0 {
		d.SetId(toStrID(projects[0].ID))
		d.Set("name", projects[0].Name)
	}

	return nil
}
