package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccSakuraIoTServiceOutgoingWebhookDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		Providers:                 testAccProviders,
		PreventPostDestroyRefresh: true,
		CheckDestroy:              testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_base,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceID("sakuraiot_service_outgoing_webhook.foobar1"),
					testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceID("sakuraiot_service_outgoing_webhook.foobar2"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_FindByProjectID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceID("data.sakuraiot_service_outgoing_webhook.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_service_outgoing_webhook.foobar", "name", "01_outgoing_webhook"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_SortByNameAsc,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceID("data.sakuraiot_service_outgoing_webhook.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_service_outgoing_webhook.foobar", "name", "01_outgoing_webhook"),
				),
			},

			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_SortByNameDesc,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceID("data.sakuraiot_service_outgoing_webhook.foobar"),
					resource.TestCheckResourceAttr("data.sakuraiot_service_outgoing_webhook.foobar", "name", "02_outgoing_webhook"),
				),
			},

			resource.TestStep{
				Destroy: true,
				Config:  testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_NotFound,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceNotExists("data.sakuraiot_service_outgoing_webhook.foobar"),
				),
			},
		},
	})
}

func testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find SakuraIoT ServiceOutgoingWebhook data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("SakuraIoT ServiceOutgoingWebhook data source ID not set")
		}
		return nil
	}
}

func testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceNotExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if ok {
			return fmt.Errorf("Found SakuraIoT ServiceOutgoingWebhook data source: %s", n)
		}
		return nil
	}
}

func testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*SakuraClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "sakuraiot_service_outgoing_webhook" {
			continue
		}

		p := client.NewReadOutgoingWebhookServiceParam(toInt64ID(rs.Primary.ID))
		_, err := client.ReadOutgoingWebhookService(p)

		if err == nil {
			return fmt.Errorf("SakuraIoT ServiceOutgoingWebhook still exists")
		}
	}

	return nil
}

const testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_base = `
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

var testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_FindByProjectID = fmt.Sprintf(`
%s

data "sakuraiot_service_outgoing_webhook" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    sort = "name"
}`, testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_base)

var testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_SortByNameAsc = fmt.Sprintf(`
%s

data "sakuraiot_service_outgoing_webhook" "foobar" {
    sort = "name"
}`, testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_base)

var testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_SortByNameDesc = fmt.Sprintf(`
%s

data "sakuraiot_service_outgoing_webhook" "foobar" {
    sort = "-name"
}`, testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_base)

var testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_NotFound = fmt.Sprintf(`
%s

data "sakuraiot_service_outgoing_webhook" "foobar" {
    project_id = 9999999999
}`, testAccCheckSakuraIoTServiceOutgoingWebhookDataSourceConfig_base)
