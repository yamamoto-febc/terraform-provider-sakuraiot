package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
	"testing"
)

func TestAccSakuraIoTModuleAlwaysOK(t *testing.T) {
	t.Logf("If you want to test the SakuraIoT module, prepare register_id and register_pass and uncomment/edit the following test code.")
}

//func TestAccSakuraIoTModule(t *testing.T) {
//	var module *models.Module
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckSakuraIoTModuleDestroy,
//		Steps: []resource.TestStep{
//			resource.TestStep{
//				Config: testAccCheckSakuraIoTModuleConfig_basic,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSakuraIoTModuleExists("sakuraiot_module.foobar", module),
//					resource.TestCheckResourceAttr(
//						"sakuraiot_module.foobar", "name", "module"),
//				),
//			},
//			resource.TestStep{
//				Config: testAccCheckSakuraIoTModuleConfig_update,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSakuraIoTModuleExists("sakuraiot_module.foobar", module),
//					resource.TestCheckResourceAttr(
//						"sakuraiot_module.foobar", "name", "module_upd"),
//				),
//			},
//		},
//	})
//}

func testAccCheckSakuraIoTModuleExists(n string, module *models.Module) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SakuraIoT Module ID is set")
		}
		client := testAccProvider.Meta().(*SakuraClient)
		p := client.NewReadModuleParam(rs.Primary.ID)
		module, err := client.ReadModule(p)

		if err != nil {
			return err
		}

		if module.ID != rs.Primary.ID {
			return fmt.Errorf("SakuraIoT Module not found")
		}

		return nil
	}
}

func testAccCheckSakuraIoTModuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*SakuraClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "sakuraiot_module" {
			continue
		}

		p := client.NewReadModuleParam(rs.Primary.ID)
		_, err := client.ReadModule(p)

		if err == nil {
			return fmt.Errorf("SakuraIoT Module still exists")
		}
	}

	return nil
}

const testAccCheckSakuraIoTModuleConfig_basic = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot"
}

resource "sakuraiot_module" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "module"
    register_id = "xxxxxxxx10"
    register_pass = "xxxxxxxx10"
}
`

const testAccCheckSakuraIoTModuleConfig_update = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot"
}

resource "sakuraiot_module" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "module_upd"
    register_id = "xxxxxxxx10"
    register_pass = "xxxxxxxx10"
}
`
