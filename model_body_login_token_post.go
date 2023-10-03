/*
 * Data Repository Service
 *
 *  GET request:  - Fetch a DrsObject from the database by sending a unique ID through the request - Fetch an access url to the data which the object refers to - Fetch DrsObjects by doing a search on the aliases  POST request:  - Create a non-existing DrsObject in the database by giving an identifier  DELETE request:  - Delete a DrsObject from the database by unique identifier  PUT request:  - Update an existing DrsObject by unique identifier and the changes in the body 
 *
 * API version: 1.2.0
 * Contact: ict@cmgg.be
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package drs_api

type BodyLoginTokenPost struct {
	GrantType string `json:"grant_type,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Scope string `json:"scope,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}