package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccSakuraIoTProjectDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		Providers:                 testAccProviders,
		PreventPostDestroyRefresh: true,
		CheckDestroy:              testAccCheckSakuraIoTProjectDataSourceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTProjectDataSourceConfig_base,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTProjectDataSourceID("sakuraiot_project.foobar"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTProjectDataSourceConfig_FindByName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTProjectDataSourceID("data.sakuraiot_project.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_project.foobar", "name", "01_project"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTProjectDataSourceConfig_WithSortAsc,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTProjectDataSourceID("data.sakuraiot_project.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_project.foobar", "name", "01_project"),
				),
			},
			resource.TestStep{
				Config: testAccCheckSakuraIoTProjectDataSourceConfig_WithSortDesc,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTProjectDataSourceID("data.sakuraiot_project.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_project.foobar", "name", "02_project"),
				),
			},
			resource.TestStep{
				Destroy: true,
				Config:  testAccCheckSakuraIoTProjectDataSourceConfig_NotFound,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTProjectDataSourceNotExists("data.sakuraiot_project.foobar"),
				),
			},
		},
	})
}

func testAccCheckSakuraIoTProjectDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find SakuraIoT Project data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SakuraIoT Project data source ID not set")
		}
		return nil
	}
}

func testAccCheckSakuraIoTProjectDataSourceNotExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if ok {
			return fmt.Errorf("Found SakuraIoT Project data source: %s", n)
		}
		return nil
	}
}

func testAccCheckSakuraIoTProjectDataSourceDestroy(s *terraform.State) error {
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

const testAccCheckSakuraIoTProjectDataSourceConfig_base = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_project" "foobar2" {
    name = "02_project"
}
`

const testAccCheckSakuraIoTProjectDataSourceConfig_FindByName = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_project" "foobar2" {
    name = "02_project"
}

data "sakuraiot_project" "foobar" {
    name = "01_project"
}
`

const testAccCheckSakuraIoTProjectDataSourceConfig_WithSortAsc = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_project" "foobar2" {
    name = "02_project"
}

data "sakuraiot_project" "foobar" {
    sort = "name"
}`

const testAccCheckSakuraIoTProjectDataSourceConfig_WithSortDesc = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_project" "foobar2" {
    name = "02_project"
}

data "sakuraiot_project" "foobar" {
    sort = "-name"
}`

const testAccCheckSakuraIoTProjectDataSourceConfig_NotFound = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_project" "foobar2" {
    name = "02_project"
}

data "sakuraiot_project" "foobar" {
    name = "xxxxxxxxxx"
}`
