package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccSakuraIoTServiceIncomingWebhookDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		Providers:                 testAccProviders,
		PreventPostDestroyRefresh: true,
		CheckDestroy:              testAccCheckSakuraIoTServiceIncomingWebhookDataSourceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_base,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceIncomingWebhookDataSourceID("sakuraiot_service_incoming_webhook.foobar1"),
					testAccCheckSakuraIoTServiceIncomingWebhookDataSourceID("sakuraiot_service_incoming_webhook.foobar2"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_FindByProjectID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceIncomingWebhookDataSourceID("data.sakuraiot_service_incoming_webhook.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_service_incoming_webhook.foobar", "name", "01_incoming_webhook"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_SortByNameAsc,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceIncomingWebhookDataSourceID("data.sakuraiot_service_incoming_webhook.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_service_incoming_webhook.foobar", "name", "01_incoming_webhook"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_SortByNameDesc,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceIncomingWebhookDataSourceID("data.sakuraiot_service_incoming_webhook.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_service_incoming_webhook.foobar", "name", "02_incoming_webhook"),
				),
			},

			resource.TestStep{
				Destroy: true,
				Config:  testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_NotFound,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceIncomingWebhookDataSourceNotExists("data.sakuraiot_service_incoming_webhook.foobar"),
				),
			},
		},
	})
}

func testAccCheckSakuraIoTServiceIncomingWebhookDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find SakuraIoT ServiceIncomingWebhook data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SakuraIoT ServiceIncomingWebhook data source ID not set")
		}
		return nil
	}
}

func testAccCheckSakuraIoTServiceIncomingWebhookDataSourceNotExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if ok {
			return fmt.Errorf("Found SakuraIoT ServiceIncomingWebhook data source: %s", n)
		}
		return nil
	}
}

func testAccCheckSakuraIoTServiceIncomingWebhookDataSourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*SakuraClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "sakuraiot_service_incoming_webhook" {
			continue
		}

		p := client.NewReadIncomingWebhookServiceParam(toInt64ID(rs.Primary.ID))
		_, err := client.ReadIncomingWebhookService(p)

		if err == nil {
			return fmt.Errorf("SakuraIoT ServiceIncomingWebhook still exists")
		}
	}

	return nil
}

const testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_base = `
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

var testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_FindByProjectID = fmt.Sprintf(`
%s

data "sakuraiot_service_incoming_webhook" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    sort = "name"
}`, testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_base)

var testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_SortByNameAsc = fmt.Sprintf(`
%s

data "sakuraiot_service_incoming_webhook" "foobar" {
    sort = "name"
}`, testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_base)

var testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_SortByNameDesc = fmt.Sprintf(`
%s

data "sakuraiot_service_incoming_webhook" "foobar" {
    sort = "-name"
}`, testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_base)

var testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_NotFound = fmt.Sprintf(`
%s

data "sakuraiot_service_incoming_webhook" "foobar" {
    project_id = 9999999999
}`, testAccCheckSakuraIoTServiceIncomingWebhookDataSourceConfig_base)
