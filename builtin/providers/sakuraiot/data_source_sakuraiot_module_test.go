package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccSakuraIoTModuleDataSourceAlwaysOK(t *testing.T) {
	t.Logf("If you want to test the SakuraIoT module, prepare register_id and register_pass and uncomment/edit the following test code.")
}

//func TestAccSakuraIoTModuleDataSource(t *testing.T) {
//	resource.Test(t, resource.TestCase{
//		PreCheck:                  func() { testAccPreCheck(t) },
//		Providers:                 testAccProviders,
//		PreventPostDestroyRefresh: true,
//		CheckDestroy:              testAccCheckSakuraIoTModuleDataSourceDestroy,
//		Steps: []resource.TestStep{
//			resource.TestStep{
//				Config: testAccCheckSakuraIoTModuleDataSourceConfig_base,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSakuraIoTModuleDataSourceID("sakuraiot_module.foobar"),
//				),
//			},
//
//			resource.TestStep{
//				Config: testAccCheckSakuraIoTModuleDataSourceConfig_FindByProjectID,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSakuraIoTModuleDataSourceID("data.sakuraiot_module.foobar"),
//					resource.TestCheckResourceAttr("data.sakuraiot_module.foobar", "name", "01_module"),
//				),
//			},
//
//			resource.TestStep{
//				Config: testAccCheckSakuraIoTModuleDataSourceConfig_FindBySerialNumber,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSakuraIoTModuleDataSourceID("data.sakuraiot_module.foobar"),
//					resource.TestCheckResourceAttr("data.sakuraiot_module.foobar", "name", "01_module"),
//				),
//			},
//
//			resource.TestStep{
//				Config: testAccCheckSakuraIoTModuleDataSourceConfig_FindByModel,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSakuraIoTModuleDataSourceID("data.sakuraiot_module.foobar"),
//					resource.TestCheckResourceAttr("data.sakuraiot_module.foobar", "name", "01_module"),
//				),
//			},
//
//			resource.TestStep{
//				Destroy: true,
//				Config:  testAccCheckSakuraIoTModuleDataSourceConfig_NotFound,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSakuraIoTModuleDataSourceNotExists("data.sakuraiot_module.foobar"),
//				),
//			},
//		},
//	})
//}

func testAccCheckSakuraIoTModuleDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find SakuraIoT Module data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SakuraIoT Module data source ID not set")
		}
		return nil
	}
}

func testAccCheckSakuraIoTModuleDataSourceNotExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if ok {
			return fmt.Errorf("Found SakuraIoT Module data source: %s", n)
		}
		return nil
	}
}

func testAccCheckSakuraIoTModuleDataSourceDestroy(s *terraform.State) error {
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

const testAccCheckSakuraIoTModuleDataSourceConfig_base = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_module" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "01_module"
    register_id = "xxxxxxxx10"
    register_pass = "xxxxxxxx10"
}
`

const testAccCheckSakuraIoTModuleDataSourceConfig_FindByProjectID = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_module" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "01_module"
    register_id = "xxxxxxxx10"
    register_pass = "xxxxxxxx10"
}

data "sakuraiot_module" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
}
`

const testAccCheckSakuraIoTModuleDataSourceConfig_FindBySerialNumber = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_module" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "01_module"
    register_id = "xxxxxxxx10"
    register_pass = "xxxxxxxx10"
}

data "sakuraiot_module" "foobar" {
    serial_number = "nnnnnnnn10"
}
`

const testAccCheckSakuraIoTModuleDataSourceConfig_FindByModel = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_module" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "01_module"
    register_id = "xxxxxxxx10"
    register_pass = "xxxxxxxx10"
}

data "sakuraiot_module" "foobar" {
    model = "SCM-LTE-BETA"
}
`

const testAccCheckSakuraIoTModuleDataSourceConfig_NotFound = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_module" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "01_module"
    register_id = "xxxxxxxx10"
    register_pass = "xxxxxxxx10"
}

data "sakuraiot_module" "foobar" {
    model = "xxxxxxxxxx"
}`
