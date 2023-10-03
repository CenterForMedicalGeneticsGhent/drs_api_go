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

// The checksum of the ```DrsObject```. At least one checksum must be provided.         For blobs, the checksum is computed over the bytes in the blob.         For bundles, the checksum is computed over a sorted concatenation of the checksums of its         top-level contained objects (not recursive, names not included).         The list of checksums is sorted alphabetically (hex-code) before concatenation         and a further checksum is performed on the concatenated checksum value.         For example, if a bundle contains blobs with the following checksums:         md5(blob1) = 72794b6d md5(blob2) = 5e089d29 Then the checksum of the bundle is:         md5( concat( sort( md5(blob1), md5(blob2) ) ) ) = md5( concat( sort( 72794b6d, 5e089d29 ) ) ) =         md5( concat( 5e089d29, 72794b6d ) ) = md5( 5e089d2972794b6d ) = f7a29a04
type AllOfDrsObjectChecksums struct {
	// The hex-string encoded checksum for the data
	Checksum string `json:"checksum"`
	// The digest method used to create the checksum. The         value (e.g. ```sha-256```) SHOULD be listed as ```Hash Name String``` in the         https://www.iana.org/assignments/named-information/named-information.xhtml#hash-alg[IANA Named Information Hash Algorithm Registry].         Other values MAY be used, as long as implementors are aware of the issues         discussed in https://tools.ietf.org/html/rfc6920#section-9.4[RFC6920].         GA4GH may provide more explicit guidance for use of non-IANA-registered algorithms in the future.         Until then, if implementors do choose such an algorithm (e.g. because it's implemented by their storage provider),         they SHOULD use an existing standard ```type``` value such as ```md5```, ```etag```, ```crc32c```, ```trunc512```, or ```sha1```.
	Type_ string `json:"type"`
}