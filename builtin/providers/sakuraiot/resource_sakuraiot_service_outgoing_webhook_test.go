package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
	"testing"
)

func TestAccSakuraIoTServiceOutgoingWebhook(t *testing.T) {
	var service *models.OutgoingWebhookService
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSakuraIoTServiceOutgoingWebhookDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceOutgoingWebhookConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceOutgoingWebhookExists("sakuraiot_service_outgoing_webhook.foobar", service),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_outgoing_webhook.foobar", "name", "terraform_for_sakuraiot"),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_outgoing_webhook.foobar", "url", "https://8.8.8.8/"),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_outgoing_webhook.foobar", "secret", "hogehoge"),
				),
			},
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceOutgoingWebhookConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceOutgoingWebhookExists("sakuraiot_service_outgoing_webhook.foobar", service),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_outgoing_webhook.foobar", "name", "terraform_for_sakuraiot_upd"),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_outgoing_webhook.foobar", "url", "https://8.8.4.4/"),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_outgoing_webhook.foobar", "secret", "hogehoge_upd"),
				),
			},
		},
	})
}

func TestAccSakuraIoTServiceOutgoingWebhook_Import(t *testing.T) {
	resourceName := "sakuraiot_service_outgoing_webhook.foobar"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSakuraIoTServiceOutgoingWebhookDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceOutgoingWebhookConfig_basic,
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSakuraIoTServiceOutgoingWebhookExists(n string, service *models.OutgoingWebhookService) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SakuraIoT ServiceOutgoingWebhook ID is set")
		}
		client := testAccProvider.Meta().(*SakuraClient)
		p := client.NewReadOutgoingWebhookServiceParam(toInt64ID(rs.Primary.ID))
		service, err := client.ReadOutgoingWebhookService(p)

		if err != nil {
			return err
		}

		if toStrID(service.ID) != rs.Primary.ID {
			return fmt.Errorf("SakuraIoT ServiceOutgoingWebhook not found")
		}

		return nil
	}
}

func testAccCheckSakuraIoTServiceOutgoingWebhookDestroy(s *terraform.State) error {
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

const testAccCheckSakuraIoTServiceOutgoingWebhookConfig_basic = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot"
}
resource "sakuraiot_service_outgoing_webhook" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "terraform_for_sakuraiot"
    url = "https://8.8.8.8/"
    secret = "hogehoge"
}`

const testAccCheckSakuraIoTServiceOutgoingWebhookConfig_update = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot"
}
resource "sakuraiot_service_outgoing_webhook" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "terraform_for_sakuraiot_upd"
    url = "https://8.8.4.4/"
    secret = "hogehoge_upd"
}`
