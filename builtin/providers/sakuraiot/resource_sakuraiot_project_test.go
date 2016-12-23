package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
	"testing"
)

func TestAccSakuraIoTProject(t *testing.T) {
	var project *models.Project
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSakuraIoTProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTProjectConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTProjectExists("sakuraiot_project.foobar", project),
					resource.TestCheckResourceAttr(
						"sakuraiot_project.foobar", "name", "terraform_for_sakuraiot"),
				),
			},
			resource.TestStep{
				Config: testAccCheckSakuraIoTProjectConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTProjectExists("sakuraiot_project.foobar", project),
					resource.TestCheckResourceAttr(
						"sakuraiot_project.foobar", "name", "terraform_for_sakuraiot_upd"),
				),
			},
		},
	})
}

func TestAccSakuraIoTProject_Import(t *testing.T) {
	resourceName := "sakuraiot_project.foobar"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSakuraIoTProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTProjectConfig_basic,
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSakuraIoTProjectExists(n string, project *models.Project) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SakuraIoT Project ID is set")
		}
		client := testAccProvider.Meta().(*SakuraClient)
		p := client.NewReadProjectParam(toInt64ID(rs.Primary.ID))
		project, err := client.ReadProject(p)

		if err != nil {
			return err
		}

		if toStrID(project.ID) != rs.Primary.ID {
			return fmt.Errorf("SakuraIoT Project not found")
		}

		return nil
	}
}

func testAccCheckSakuraIoTProjectDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*SakuraClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "sakuraiot_project" {
			continue
		}

		p := client.NewReadProjectParam(toInt64ID(rs.Primary.ID))
		_, err := client.ReadProject(p)

		if err == nil {
			return fmt.Errorf("SakuraIoT Project still exists")
		}
	}

	return nil
}

const testAccCheckSakuraIoTProjectConfig_basic = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot"
}`

const testAccCheckSakuraIoTProjectConfig_update = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot_upd"
}`
