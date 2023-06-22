package provider

import (
	client "github.com/costae/Terraform-SAP-BTP-Provider/apiclient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"BaseURL": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_URL", ""),
			},
			"Username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_USER", ""),
			},
			"Password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_PASS", ""),
			},
			"GlobalAccount": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_GA", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"example_subaccount": resourceSubaccount(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	BaseURL := d.Get("BaseURL").(string)
	Username := d.Get("Username").(string)
	Password := d.Get("Password").(string)
	GlobalAccount := d.Get("GlobalAccount").(string)
	return client.NewClient(BaseURL, Username, Password, GlobalAccount), nil

}
