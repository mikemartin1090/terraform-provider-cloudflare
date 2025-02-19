package sdkv2provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareListItem_Basic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cloudflare.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &ListItem),
					resource.TestCheckResourceAttr(
						name, "ip", "192.0.2.0"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_MultipleItems(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cloudflare.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItemMultipleEntries(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name+"_1", rnd, &ListItem),
					resource.TestCheckResourceAttr(name+"_1", "ip", "192.0.2.0"),

					testAccCheckCloudflareListItemExists(name+"_2", rnd, &ListItem),
					resource.TestCheckResourceAttr(name+"_2", "ip", "192.0.2.1"),

					testAccCheckCloudflareListItemExists(name+"_3", rnd, &ListItem),
					resource.TestCheckResourceAttr(name+"_3", "ip", "192.0.2.2"),

					testAccCheckCloudflareListItemExists(name+"_4", rnd, &ListItem),
					resource.TestCheckResourceAttr(name+"_4", "ip", "192.0.2.3"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_Update(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var listItem cloudflare.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &listItem),
					resource.TestCheckResourceAttr(
						name, "ip", "192.0.2.0"),
				),
			},
			{
				Config: testAccCheckCloudflareIPListItem(rnd, rnd, rnd+"-updated", accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &listItem),
					resource.TestCheckResourceAttr(name, "comment", rnd+"-updated"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_BadListItemType(t *testing.T) {
	rnd := generateRandomResourceName()
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareBadListItemType(rnd, rnd, rnd, accountID),
				ExpectError: regexp.MustCompile(" can not be added to lists of type "),
			},
		},
	})
}

func TestAccCloudflareListItem_ASN(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cloudflare.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareASNListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &ListItem),
					resource.TestCheckResourceAttr(
						name, "asn", "1"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_Hostname(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cloudflare.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameListItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &ListItem),
					resource.TestCheckResourceAttr(
						name, "hostname.#", "1"),
					resource.TestCheckResourceAttr(
						name, "hostname.0.url_hostname", "example.com"),
				),
			},
		},
	})
}

func TestAccCloudflareListItem_Redirect(t *testing.T) {
	rnd := generateRandomResourceName()
	name := fmt.Sprintf("cloudflare_list_item.%s", rnd)
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	var ListItem cloudflare.ListItem

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAccount(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareHostnameRedirectItem(rnd, rnd, rnd, accountID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareListItemExists(name, rnd, &ListItem),
					resource.TestCheckResourceAttr(
						name, "redirect.#", "1"),
					resource.TestCheckResourceAttr(
						name, "redirect.0.source_url", "example.com/"),
					resource.TestCheckResourceAttr(
						name, "redirect.0.target_url", "https://example1.com"),
					resource.TestCheckResourceAttr(
						name, "redirect.0.status_code", "301"),
				),
			},
		},
	})
}

func testAccCheckCloudflareIPListItem(ID, name, comment, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "ip"
  }

  resource "cloudflare_list_item" "%[1]s" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	ip         = "192.0.2.0"
	comment    = "%[3]s"
  } `, ID, name, comment, accountID)
}

func testAccCheckCloudflareIPListItemMultipleEntries(ID, name, comment, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "ip"
  }

  resource "cloudflare_list_item" "%[1]s_1" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	ip         = "192.0.2.0"
	comment    = "%[3]s"
  }

  resource "cloudflare_list_item" "%[1]s_2" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	ip         = "192.0.2.1"
	comment    = "%[3]s"
  }

  resource "cloudflare_list_item" "%[1]s_3" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	ip         = "192.0.2.2"
	comment    = "%[3]s"
  }

  resource "cloudflare_list_item" "%[1]s_4" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	ip         = "192.0.2.3"
	comment    = "%[3]s"
  } `, ID, name, comment, accountID)
}

func testAccCheckCloudflareListItemExists(n string, name string, listItem *cloudflare.ListItem) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
		listRS := s.RootModule().Resources["cloudflare_list."+name]

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No List ID is set")
		}

		client := testAccProvider.Meta().(*cloudflare.API)
		foundList, err := client.GetListItem(context.Background(), cloudflare.AccountIdentifier(accountID), listRS.Primary.ID, rs.Primary.ID)
		if err != nil {
			return err
		}

		*listItem = foundList

		return nil
	}
}

func testAccCheckCloudflareBadListItemType(ID, name, comment, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "redirect"
  }

  resource "cloudflare_list_item" "%[2]s" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	ip         = "192.0.2.0"
	comment    = "%[3]s"
  } `, ID, name, comment, accountID)
}

func testAccCheckCloudflareASNListItem(ID, name, comment, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "asn"
  }

  resource "cloudflare_list_item" "%[1]s" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	asn = 1
	comment    = "%[3]s"
  } `, ID, name, comment, accountID)
}

func testAccCheckCloudflareHostnameListItem(ID, name, comment, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "hostname"
  }

  resource "cloudflare_list_item" "%[1]s" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	comment    = "%[3]s"
	hostname {
		url_hostname = "example.com"
	}
  } `, ID, name, comment, accountID)
}

func testAccCheckCloudflareHostnameRedirectItem(ID, name, comment, accountID string) string {
	return fmt.Sprintf(`
  resource "cloudflare_list" "%[2]s" {
    account_id          = "%[4]s"
    name                = "%[2]s"
    description         = "list named %[2]s"
    kind                = "redirect"
  }

  resource "cloudflare_list_item" "%[1]s" {
    account_id = "%[4]s"
	list_id    = cloudflare_list.%[2]s.id
	comment    = "%[3]s"
	redirect {
		source_url = "example.com/"
		target_url = "https://example1.com"
		status_code = 301
	}
  } `, ID, name, comment, accountID)
}
