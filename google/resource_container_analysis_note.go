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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceContainerAnalysisNote() *schema.Resource {
	return &schema.Resource{
		Create: resourceContainerAnalysisNoteCreate,
		Read:   resourceContainerAnalysisNoteRead,
		Update: resourceContainerAnalysisNoteUpdate,
		Delete: resourceContainerAnalysisNoteDelete,

		Importer: &schema.ResourceImporter{
			State: resourceContainerAnalysisNoteImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"attestation_authority": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Note kind that represents a logical attestation "role" or "authority".
For example, an organization might have one AttestationAuthority for
"QA" and one for "build". This Note is intended to act strictly as a
grouping mechanism for the attached Occurrences (Attestations). This
grouping mechanism also provides a security boundary, since IAM ACLs
gate the ability for a principle to attach an Occurrence to a given
Note. It also provides a single point of lookup to find all attached
Attestation Occurrences, even if they don't all live in the same
project.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hint": {
							Type:     schema.TypeList,
							Required: true,
							Description: `This submessage provides human-readable hints about the purpose of
the AttestationAuthority. Because the name of a Note acts as its
resource reference, it is important to disambiguate the canonical
name of the Note (which might be a UUID for security purposes)
from "readable" names more suitable for debug output. Note that
these hints should NOT be used to look up AttestationAuthorities
in security sensitive contexts, such as when looking up
Attestations to verify.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"human_readable_name": {
										Type:     schema.TypeString,
										Required: true,
										Description: `The human readable name of this Attestation Authority, for
example "qa".`,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the note.`,
			},
			"expiration_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Time of expiration for this note. Leave empty if note does not expire.`,
			},
			"long_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A detailed description of the note`,
			},
			"related_note_names": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `Names of other notes related to this note.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"related_url": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `URLs associated with this note and related metadata.`,
				Elem:        containeranalysisNoteRelatedUrlSchema(),
				// Default schema.HashSchema is used.
			},
			"short_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A one sentence description of the note.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time this note was created.`,
			},
			"kind": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The type of analysis this note describes`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time this note was last updated.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func containeranalysisNoteRelatedUrlSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specific URL associated with the resource.`,
			},
			"label": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Label to describe usage of the URL`,
			},
		},
	}
}

func resourceContainerAnalysisNoteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandContainerAnalysisNoteName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	shortDescriptionProp, err := expandContainerAnalysisNoteShortDescription(d.Get("short_description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("short_description"); !isEmptyValue(reflect.ValueOf(shortDescriptionProp)) && (ok || !reflect.DeepEqual(v, shortDescriptionProp)) {
		obj["shortDescription"] = shortDescriptionProp
	}
	longDescriptionProp, err := expandContainerAnalysisNoteLongDescription(d.Get("long_description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("long_description"); !isEmptyValue(reflect.ValueOf(longDescriptionProp)) && (ok || !reflect.DeepEqual(v, longDescriptionProp)) {
		obj["longDescription"] = longDescriptionProp
	}
	relatedUrlProp, err := expandContainerAnalysisNoteRelatedUrl(d.Get("related_url"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("related_url"); !isEmptyValue(reflect.ValueOf(relatedUrlProp)) && (ok || !reflect.DeepEqual(v, relatedUrlProp)) {
		obj["relatedUrl"] = relatedUrlProp
	}
	expirationTimeProp, err := expandContainerAnalysisNoteExpirationTime(d.Get("expiration_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_time"); !isEmptyValue(reflect.ValueOf(expirationTimeProp)) && (ok || !reflect.DeepEqual(v, expirationTimeProp)) {
		obj["expirationTime"] = expirationTimeProp
	}
	relatedNoteNamesProp, err := expandContainerAnalysisNoteRelatedNoteNames(d.Get("related_note_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("related_note_names"); !isEmptyValue(reflect.ValueOf(relatedNoteNamesProp)) && (ok || !reflect.DeepEqual(v, relatedNoteNamesProp)) {
		obj["relatedNoteNames"] = relatedNoteNamesProp
	}
	attestationAuthorityProp, err := expandContainerAnalysisNoteAttestationAuthority(d.Get("attestation_authority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority"); !isEmptyValue(reflect.ValueOf(attestationAuthorityProp)) && (ok || !reflect.DeepEqual(v, attestationAuthorityProp)) {
		obj["attestationAuthority"] = attestationAuthorityProp
	}

	obj, err = resourceContainerAnalysisNoteEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "projects/{{project}}/notes/{{name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ContainerAnalysisBasePath}}projects/{{project}}/notes?noteId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Note: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Note: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/notes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Note %q: %#v", d.Id(), res)

	return resourceContainerAnalysisNoteRead(d, meta)
}

func resourceContainerAnalysisNoteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ContainerAnalysisBasePath}}projects/{{project}}/notes/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ContainerAnalysisNote %q", d.Id()))
	}

	res, err = resourceContainerAnalysisNoteDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ContainerAnalysisNote because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}

	if err := d.Set("name", flattenContainerAnalysisNoteName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("short_description", flattenContainerAnalysisNoteShortDescription(res["shortDescription"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("long_description", flattenContainerAnalysisNoteLongDescription(res["longDescription"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("kind", flattenContainerAnalysisNoteKind(res["kind"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("related_url", flattenContainerAnalysisNoteRelatedUrl(res["relatedUrl"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("expiration_time", flattenContainerAnalysisNoteExpirationTime(res["expirationTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("create_time", flattenContainerAnalysisNoteCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("update_time", flattenContainerAnalysisNoteUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("related_note_names", flattenContainerAnalysisNoteRelatedNoteNames(res["relatedNoteNames"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}
	if err := d.Set("attestation_authority", flattenContainerAnalysisNoteAttestationAuthority(res["attestationAuthority"], d, config)); err != nil {
		return fmt.Errorf("Error reading Note: %s", err)
	}

	return nil
}

func resourceContainerAnalysisNoteUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	shortDescriptionProp, err := expandContainerAnalysisNoteShortDescription(d.Get("short_description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("short_description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, shortDescriptionProp)) {
		obj["shortDescription"] = shortDescriptionProp
	}
	longDescriptionProp, err := expandContainerAnalysisNoteLongDescription(d.Get("long_description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("long_description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, longDescriptionProp)) {
		obj["longDescription"] = longDescriptionProp
	}
	relatedUrlProp, err := expandContainerAnalysisNoteRelatedUrl(d.Get("related_url"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("related_url"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, relatedUrlProp)) {
		obj["relatedUrl"] = relatedUrlProp
	}
	expirationTimeProp, err := expandContainerAnalysisNoteExpirationTime(d.Get("expiration_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expiration_time"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, expirationTimeProp)) {
		obj["expirationTime"] = expirationTimeProp
	}
	relatedNoteNamesProp, err := expandContainerAnalysisNoteRelatedNoteNames(d.Get("related_note_names"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("related_note_names"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, relatedNoteNamesProp)) {
		obj["relatedNoteNames"] = relatedNoteNamesProp
	}
	attestationAuthorityProp, err := expandContainerAnalysisNoteAttestationAuthority(d.Get("attestation_authority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, attestationAuthorityProp)) {
		obj["attestationAuthority"] = attestationAuthorityProp
	}

	obj, err = resourceContainerAnalysisNoteEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "projects/{{project}}/notes/{{name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ContainerAnalysisBasePath}}projects/{{project}}/notes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Note %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("short_description") {
		updateMask = append(updateMask, "shortDescription")
	}

	if d.HasChange("long_description") {
		updateMask = append(updateMask, "longDescription")
	}

	if d.HasChange("related_url") {
		updateMask = append(updateMask, "relatedUrl")
	}

	if d.HasChange("expiration_time") {
		updateMask = append(updateMask, "expirationTime")
	}

	if d.HasChange("related_note_names") {
		updateMask = append(updateMask, "relatedNoteNames")
	}

	if d.HasChange("attestation_authority") {
		updateMask = append(updateMask, "attestationAuthority")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Note %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Note %q: %#v", d.Id(), res)
	}

	return resourceContainerAnalysisNoteRead(d, meta)
}

func resourceContainerAnalysisNoteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "projects/{{project}}/notes/{{name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ContainerAnalysisBasePath}}projects/{{project}}/notes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Note %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Note")
	}

	log.Printf("[DEBUG] Finished deleting Note %q: %#v", d.Id(), res)
	return nil
}

func resourceContainerAnalysisNoteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/notes/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/notes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenContainerAnalysisNoteName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenContainerAnalysisNoteShortDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenContainerAnalysisNoteLongDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenContainerAnalysisNoteKind(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenContainerAnalysisNoteRelatedUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(containeranalysisNoteRelatedUrlSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"url":   flattenContainerAnalysisNoteRelatedUrlUrl(original["url"], d, config),
			"label": flattenContainerAnalysisNoteRelatedUrlLabel(original["label"], d, config),
		})
	}
	return transformed
}
func flattenContainerAnalysisNoteRelatedUrlUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenContainerAnalysisNoteRelatedUrlLabel(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenContainerAnalysisNoteExpirationTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenContainerAnalysisNoteCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenContainerAnalysisNoteUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenContainerAnalysisNoteRelatedNoteNames(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenContainerAnalysisNoteAttestationAuthority(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["hint"] =
		flattenContainerAnalysisNoteAttestationAuthorityHint(original["hint"], d, config)
	return []interface{}{transformed}
}
func flattenContainerAnalysisNoteAttestationAuthorityHint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["human_readable_name"] =
		flattenContainerAnalysisNoteAttestationAuthorityHintHumanReadableName(original["humanReadableName"], d, config)
	return []interface{}{transformed}
}
func flattenContainerAnalysisNoteAttestationAuthorityHintHumanReadableName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandContainerAnalysisNoteName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisNoteShortDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisNoteLongDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisNoteRelatedUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedUrl, err := expandContainerAnalysisNoteRelatedUrlUrl(original["url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUrl); val.IsValid() && !isEmptyValue(val) {
			transformed["url"] = transformedUrl
		}

		transformedLabel, err := expandContainerAnalysisNoteRelatedUrlLabel(original["label"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLabel); val.IsValid() && !isEmptyValue(val) {
			transformed["label"] = transformedLabel
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandContainerAnalysisNoteRelatedUrlUrl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisNoteRelatedUrlLabel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisNoteExpirationTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisNoteRelatedNoteNames(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandContainerAnalysisNoteAttestationAuthority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHint, err := expandContainerAnalysisNoteAttestationAuthorityHint(original["hint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHint); val.IsValid() && !isEmptyValue(val) {
		transformed["hint"] = transformedHint
	}

	return transformed, nil
}

func expandContainerAnalysisNoteAttestationAuthorityHint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHumanReadableName, err := expandContainerAnalysisNoteAttestationAuthorityHintHumanReadableName(original["human_readable_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHumanReadableName); val.IsValid() && !isEmptyValue(val) {
		transformed["humanReadableName"] = transformedHumanReadableName
	}

	return transformed, nil
}

func expandContainerAnalysisNoteAttestationAuthorityHintHumanReadableName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceContainerAnalysisNoteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Field was renamed in GA API
	obj["attestation"] = obj["attestationAuthority"]
	delete(obj, "attestationAuthority")

	return obj, nil
}

func resourceContainerAnalysisNoteDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Field was renamed in GA API
	res["attestationAuthority"] = res["attestation"]
	delete(res, "attestation")

	return res, nil
}
