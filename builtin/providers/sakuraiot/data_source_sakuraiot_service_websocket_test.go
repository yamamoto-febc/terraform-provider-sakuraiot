package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccSakuraIoTServiceWebSocketDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		Providers:                 testAccProviders,
		PreventPostDestroyRefresh: true,
		CheckDestroy:              testAccCheckSakuraIoTServiceWebSocketDataSourceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_base,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceWebSocketDataSourceID("sakuraiot_service_websocket.foobar1"),
					testAccCheckSakuraIoTServiceWebSocketDataSourceID("sakuraiot_service_websocket.foobar2"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_FindByProjectID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceWebSocketDataSourceID("data.sakuraiot_service_websocket.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_service_websocket.foobar", "name", "01_websocket"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_SortByNameAsc,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceWebSocketDataSourceID("data.sakuraiot_service_websocket.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_service_websocket.foobar", "name", "01_websocket"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_SortByNameDesc,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceWebSocketDataSourceID("data.sakuraiot_service_websocket.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_service_websocket.foobar", "name", "02_websocket"),
				),
			},

			resource.TestStep{
				Destroy: true,
				Config:  testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_NotFound,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceWebSocketDataSourceNotExists("data.sakuraiot_service_websocket.foobar"),
				),
			},
		},
	})
}

func testAccCheckSakuraIoTServiceWebSocketDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find SakuraIoT ServiceWebSocket data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SakuraIoT ServiceWebSocket data source ID not set")
		}
		return nil
	}
}

func testAccCheckSakuraIoTServiceWebSocketDataSourceNotExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if ok {
			return fmt.Errorf("Found SakuraIoT ServiceWebSocket data source: %s", n)
		}
		return nil
	}
}

func testAccCheckSakuraIoTServiceWebSocketDataSourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*SakuraClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "sakuraiot_service_websocket" {
			continue
		}

		p := client.NewReadWebSocketServiceParam(toInt64ID(rs.Primary.ID))
		_, err := client.ReadWebSocketService(p)

		if err == nil {
			return fmt.Errorf("SakuraIoT ServiceWebSocket still exists")
		}
	}

	return nil
}

const testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_base = `
resource "sakuraiot_project" "foobar" {
    name = "01_project"
}
resource "sakuraiot_service_incoming_webhook" "foobar1" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "01_incoming_webhook"
    secret = "secret"
}
resource "sakuraiot_service_incoming_webhook" "foobar2" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "02_incoming_webhook"
    secret = "secret"
}
resource "sakuraiot_service_outgoing_webhook" "foobar1" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "01_outgoing_webhook"
    url = "https://8.8.8.8/"
    secret = "secret"
}
resource "sakuraiot_service_outgoing_webhook" "foobar2" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "02_outgoing_webhook"
    url = "https://8.8.4.4/"
    secret = "secret"
}
resource "sakuraiot_service_websocket" "foobar1" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "01_websocket"
}
resource "sakuraiot_service_websocket" "foobar2" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "02_websocket"
}
`

var testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_FindByProjectID = fmt.Sprintf(`
%s

data "sakuraiot_service_websocket" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    sort = "name"
}`, testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_base)

var testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_SortByNameAsc = fmt.Sprintf(`
%s

data "sakuraiot_service_websocket" "foobar" {
    sort = "name"
}`, testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_base)

var testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_SortByNameDesc = fmt.Sprintf(`
%s

data "sakuraiot_service_websocket" "foobar" {
    sort = "-name"
}`, testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_base)

var testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_NotFound = fmt.Sprintf(`
%s

data "sakuraiot_service_websocket" "foobar" {
    project_id = 9999999999
}`, testAccCheckSakuraIoTServiceWebSocketDataSourceConfig_base)
