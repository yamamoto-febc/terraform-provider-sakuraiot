package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/yamamoto-febc/sakuraio-api/gen/models"
	"testing"
)

func TestAccSakuraIoTServiceIncomingWebhook(t *testing.T) {
	var service *models.IncomingWebhookService
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSakuraIoTServiceIncomingWebhookDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceIncomingWebhookConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceIncomingWebhookExists("sakuraiot_service_incoming_webhook.foobar", service),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_incoming_webhook.foobar", "name", "terraform_for_sakuraiot"),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_incoming_webhook.foobar", "secret", "hogehoge"),
				),
			},
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceIncomingWebhookConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraIoTServiceIncomingWebhookExists("sakuraiot_service_incoming_webhook.foobar", service),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_incoming_webhook.foobar", "name", "terraform_for_sakuraiot_upd"),
					resource.TestCheckResourceAttr(
						"sakuraiot_service_incoming_webhook.foobar", "secret", "hogehoge_upd"),
				),
			},
		},
	})
}

func TestAccSakuraIoTServiceIncomingWebhook_Import(t *testing.T) {
	resourceName := "sakuraiot_service_incoming_webhook.foobar"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSakuraIoTServiceIncomingWebhookDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSakuraIoTServiceIncomingWebhookConfig_basic,
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSakuraIoTServiceIncomingWebhookExists(n string, service *models.IncomingWebhookService) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SakuraIoT ServiceIncomingWebhook ID is set")
		}
		client := testAccProvider.Meta().(*SakuraClient)
		p := client.NewReadIncomingWebhookServiceParam(toInt64ID(rs.Primary.ID))
		service, err := client.ReadIncomingWebhookService(p)

		if err != nil {
			return err
		}

		if toStrID(service.ID) != rs.Primary.ID {
			return fmt.Errorf("SakuraIoT ServiceIncomingWebhook not found")
		}

		return nil
	}
}

func testAccCheckSakuraIoTServiceIncomingWebhookDestroy(s *terraform.State) error {
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

const testAccCheckSakuraIoTServiceIncomingWebhookConfig_basic = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot"
}
resource "sakuraiot_service_incoming_webhook" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "terraform_for_sakuraiot"
    secret = "hogehoge"
}`

const testAccCheckSakuraIoTServiceIncomingWebhookConfig_update = `
resource "sakuraiot_project" "foobar" {
    name = "terraform_for_sakuraiot"
}
resource "sakuraiot_service_incoming_webhook" "foobar" {
    project_id = "${sakuraiot_project.foobar.id}"
    name = "terraform_for_sakuraiot_upd"
    secret = "hogehoge_upd"
}`
