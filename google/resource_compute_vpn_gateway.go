// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeVpnGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeVpnGatewayCreate,
		Read:   resourceComputeVpnGatewayRead,
		Delete: resourceComputeVpnGatewayDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeVpnGatewayImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035.  Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
the first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The network this VPN gateway is accepting traffic for.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The region this gateway should sit in.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"gateway_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeVpnGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeVpnGatewayDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeVpnGatewayName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeVpnGatewayNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	regionProp, err := expandComputeVpnGatewayRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetVpnGateways")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new VpnGateway: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating VpnGateway: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating VpnGateway",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create VpnGateway: %s", err)
	}

	log.Printf("[DEBUG] Finished creating VpnGateway %q: %#v", d.Id(), res)

	return resourceComputeVpnGatewayRead(d, meta)
}

func resourceComputeVpnGatewayRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeVpnGateway %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading VpnGateway: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeVpnGatewayCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnGateway: %s", err)
	}
	if err := d.Set("description", flattenComputeVpnGatewayDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnGateway: %s", err)
	}
	if err := d.Set("name", flattenComputeVpnGatewayName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnGateway: %s", err)
	}
	if err := d.Set("gateway_id", flattenComputeVpnGatewayGatewayId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnGateway: %s", err)
	}
	if err := d.Set("network", flattenComputeVpnGatewayNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnGateway: %s", err)
	}
	if err := d.Set("region", flattenComputeVpnGatewayRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnGateway: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading VpnGateway: %s", err)
	}

	return nil
}

func resourceComputeVpnGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting VpnGateway %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "VpnGateway")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting VpnGateway",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting VpnGateway %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeVpnGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/targetVpnGateways/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeVpnGatewayCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnGatewayDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnGatewayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeVpnGatewayGatewayId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeVpnGatewayNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeVpnGatewayRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeVpnGatewayDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnGatewayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnGatewayNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnGatewayRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
