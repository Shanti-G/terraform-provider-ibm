// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMEnEventStreamsTemplateDataSourceBasic(t *testing.T) {
	templateInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	templateName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	templateType := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnEventStreamsTemplateDataSourceConfigBasic(templateInstanceID, templateName, templateType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "subscription_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "subscription_names.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "updated_at"),
				),
			},
		},
	})
}

func TestAccIBMEnEventStreamsTemplateDataSourceAllArgs(t *testing.T) {
	templateInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	templateName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	templateDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	templateType := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnPagerDutyTemplateDataSourceConfig(templateInstanceID, templateName, templateDescription, templateType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "subscription_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "subscription_names.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_event_streams_template.en_template_instance", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnEventStreamsTemplateDataSourceConfigBasic(templateInstanceID string, templateName string, templateType string) string {
	return fmt.Sprintf(`
		resource "ibm_en_event_streams_template" "en_template_instance" {
			instance_id = "%s"
			name = "%s"
			type = "%s"
		}

		data "ibm_en_event_streams_template" "en_template_instance" {
			instance_id = ibm_en_event_streams_template.en_template_instance.instance_id
			template_id = ibm_en_event_streams_template.en_template_instance.template_id
		}
	`, templateInstanceID, templateName, templateType)
}

func testAccCheckIBMEnEventStreamsTemplateDataSourceConfig(templateInstanceID string, templateName string, templateDescription string, templateType string) string {
	return fmt.Sprintf(`
		resource "ibm_en_event_streams_template" "en_template_instance" {
			instance_id = "%s"
			name = "%s"
			description = "%s"
			type = "%s"
			params {
			    body = "eyJuYW1lIjoie3tkYXRhLm5hbWV9fSIifQ=="
			}
		}

		data "ibm_en_event_streams_template" "en_template_instance" {
			instance_id = ibm_en_event_streams_template.en_template_instance.instance_id
			en_template_id = ibm_en_event_streams_template.en_template_instance.template_id
		}
	`, templateInstanceID, templateName, templateDescription, templateType)
}
