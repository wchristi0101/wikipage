/*
 * Wikipedia API
 *
 * This is an api that searches a wikipage
 *
 * API version: 1.0.1
 * Contact: bchristian14@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type WebsiteDataInner struct {
	// optional header property
	Header string `json:"header,omitempty"`

	Label []interface{} `json:"label"`
}
