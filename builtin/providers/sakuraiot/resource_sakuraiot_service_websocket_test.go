package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
	"testing"
)

func TestAccSakuraIoTServiceWebSocket(t *testing.T) {
	var service *models.WebSocketService
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSakuraIoTServiceWebSocketDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceWebSocketConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceWebSocketExists("sakuraiot_service_websocket.foobar", service),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_websocket.foobar", "name", "terraform_for_sakuraiot"),
				),
			},
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceWebSocketConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceWebSocketExists("sakuraiot_service_websocket.foobar", service),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_websocket.foobar", "name", "terraform_for_sakuraiot_upd"),
				),
			},
		},
	})
}

func TestAccSakuraIoTServiceWebSocket_Import(t *testing.T) {
	resourceName := "sakuraiot_service_websocket.foobar"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSakuraIoTServiceWebSocketDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceWebSocketConfig_basic,
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSakuraIoTServiceWebSocketExists(n string, service *models.WebSocketService) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SakuraIoT ServiceWebSocket ID is set")
		}
		client := testAccProvider.Meta().(*SakuraClient)
		p := client.NewReadWebSocketServiceParam(toInt64ID(rs.Primary.ID))
		service, err := client.ReadWebSocketService(p)

		if err != nil {
			return err
		}

		if toStrID(service.ID) != rs.Primary.ID {
			return fmt.Errorf("SakuraIoT ServiceWebSocket not found")
		}

		return nil
	}
}

func testAccCheckSakuraIoTServiceWebSocketDestroy(s *terraform.State) error {
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

const testAccCheckSakuraIoTServiceWebSocketConfig_basic = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot"
}
resource "sakuraiot_service_websocket" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "terraform_for_sakuraiot"
}`

const testAccCheckSakuraIoTServiceWebSocketConfig_update = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot"
}
resource "sakuraiot_service_websocket" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "terraform_for_sakuraiot_upd"
}`
