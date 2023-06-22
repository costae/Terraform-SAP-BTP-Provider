package provider

import (
	"fmt"

	client "github.com/costae/Terraform-SAP-BTP-Provider/apiclient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// func validateName(v interface{}, k string) (ws []string, es []error) {
// 	var errs []error
// 	var warns []string
// 	value, ok := v.(string)
// 	if !ok {
// 		errs = append(errs, fmt.Errorf("Expected name to be string"))
// 		return warns, errs
// 	}
// 	whiteSpace := regexp.MustCompile(`\s+`)
// 	if whiteSpace.Match([]byte(value)) {
// 		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
// 		return warns, errs
// 	}
// 	return warns, errs
// }

func resourceSubaccount() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"SubaccountID": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"TechnicalName": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"DisplayName": {
				Type:     schema.TypeString,
				Required: true,
			},
			"GlobalAccountGUID": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ParentGUID": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"Region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"Subdomain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"BetaEnabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"UsedForProduction": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"State": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"StateMessage": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ParentType": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"CreatedDate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ModifiedDate": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: resourceCreateSubaccount,
		Read:   resourceReadSubaccount,
		Update: resourceUpdateSubaccount,
		Delete: resourceDeleteSubaccount,
	}
}

func resourceCreateSubaccount(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.APIClient)

	item := client.SubaccountInfo{
		Subaccount:        d.Get("Subaccount").(string),
		TechnicalName:     d.Get("TechnicalName").(string),
		DisplayName:       d.Get("DisplayName").(string),
		GlobalAccountGUID: d.Get("GlobalAccountGUID").(string),
		ParentGUID:        d.Get("ParentGUID").(string),
		ParentType:        d.Get("ParentType").(string),
		Region:            d.Get("Region").(string),
		Subdomain:         d.Get("Subdomain").(string),
		BetaEnabled:       d.Get("BetaEnabled").(bool),
		UsedForProduction: d.Get("UsedForProduction").(string),
		State:             d.Get("State").(string),
		StateMessage:      d.Get("StateMessage").(string),
		CreatedDate:       d.Get("CreatedDate").(string),
		ModifiedDate:      d.Get("ModifiedDate").(string),
	}
	f, err := apiClient.CreateSubaccountCommand(&item)

	if err != nil {
		return err
	}
	if f != nil {
		d = f
	}
	return nil
}

func resourceReadSubaccount(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.APIClient)

	// itemId := d.Id()
	item := client.SubaccountInfo{
		Subaccount:        d.Get("Subaccount").(string),
		TechnicalName:     d.Get("TechnicalName").(string),
		DisplayName:       d.Get("DisplayName").(string),
		GlobalAccountGUID: d.Get("GlobalAccountGUID").(string),
		ParentGUID:        d.Get("ParentGUID").(string),
		ParentType:        d.Get("ParentType").(string),
		Region:            d.Get("Region").(string),
		Subdomain:         d.Get("Subdomain").(string),
		BetaEnabled:       d.Get("BetaEnabled").(bool),
		UsedForProduction: d.Get("UsedForProduction").(string),
		State:             d.Get("State").(string),
		StateMessage:      d.Get("StateMessage").(string),
		CreatedDate:       d.Get("CreatedDate").(string),
		ModifiedDate:      d.Get("ModifiedDate").(string),
	}
	items, err := apiClient.GetSubaccountCommand(&item)
	if err != nil {
		return err
	}
	if items != nil {
		d = items
	}
	return nil
}

func resourceUpdateSubaccount(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.APIClient)

	// tfTags := d.Get("tags").(*schema.Set).List()
	// tags := make([]string, len(tfTags))
	// for i, tfTag := range tfTags {
	// 	tags[i] = tfTag.(string)
	// }

	// item := client.Item{
	// 	Name:        d.Get("name").(string),
	// 	Description: d.Get("description").(string),
	// 	Tags:        tags,
	// }
	item := client.SubaccountInfo{
		Subaccount:        d.Get("Subaccount").(string),
		TechnicalName:     d.Get("TechnicalName").(string),
		DisplayName:       d.Get("DisplayName").(string),
		GlobalAccountGUID: d.Get("GlobalAccountGUID").(string),
		ParentGUID:        d.Get("ParentGUID").(string),
		ParentType:        d.Get("ParentType").(string),
		Region:            d.Get("Region").(string),
		Subdomain:         d.Get("Subdomain").(string),
		BetaEnabled:       d.Get("BetaEnabled").(bool),
		UsedForProduction: d.Get("UsedForProduction").(string),
		State:             d.Get("State").(string),
		StateMessage:      d.Get("StateMessage").(string),
		CreatedDate:       d.Get("CreatedDate").(string),
		ModifiedDate:      d.Get("ModifiedDate").(string),
	}
	f, err := apiClient.UpdateSubaccountCommand(&item)
	if err != nil {
		return err
	}
	if f != nil {
		d = f
	}
	return nil
}

func resourceDeleteSubaccount(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.APIClient)

	// itemId := d.Id()
	item := client.SubaccountInfo{
		Subaccount:        d.Get("Subaccount").(string),
		TechnicalName:     d.Get("TechnicalName").(string),
		DisplayName:       d.Get("DisplayName").(string),
		GlobalAccountGUID: d.Get("GlobalAccountGUID").(string),
		ParentGUID:        d.Get("ParentGUID").(string),
		ParentType:        d.Get("ParentType").(string),
		Region:            d.Get("Region").(string),
		Subdomain:         d.Get("Subdomain").(string),
		BetaEnabled:       d.Get("BetaEnabled").(bool),
		UsedForProduction: d.Get("UsedForProduction").(string),
		State:             d.Get("State").(string),
		StateMessage:      d.Get("StateMessage").(string),
		CreatedDate:       d.Get("CreatedDate").(string),
		ModifiedDate:      d.Get("ModifiedDate").(string),
	}
	f, err := apiClient.DeleteItem(&item)
	if err != nil {
		return err
	}

	if f != nil {
		d = f
	}
	return nil
}
