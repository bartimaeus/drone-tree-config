/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucket

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type SshApiService service


/* SshApiService 
 Deletes a specific SSH public key from a user&#39;s account  Example: &#x60;&#x60;&#x60; $ curl -X DELETE https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/{b15b6026-9c02-4626-b4ad-b905f99f763a} &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username The account&#39;s username or UUID.
 @param keyId The SSH key&#39;s UUID value.
 @return */
func (a *SshApiService) UsersUsernameSshKeysDelete(ctx context.Context, username string, keyId string) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{username}/ssh-keys/"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", fmt.Sprintf("%v", keyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* SshApiService 
 Returns a specific SSH public key belonging to a user.  Example: &#x60;&#x60;&#x60; $ curl https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/{fbe4bbab-f6f7-4dde-956b-5c58323c54b3}  {     \&quot;comment\&quot;: \&quot;user@myhost\&quot;,     \&quot;created_on\&quot;: \&quot;2018-03-14T13:17:05.196003+00:00\&quot;,     \&quot;key\&quot;: \&quot;ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY\&quot;,     \&quot;label\&quot;: \&quot;\&quot;,     \&quot;last_used\&quot;: \&quot;2018-03-20T13:18:05.196003+00:00\&quot;,     \&quot;links\&quot;: {         \&quot;self\&quot;: {             \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a\&quot;         }     },     \&quot;owner\&quot;: {         \&quot;display_name\&quot;: \&quot;Mark Adams\&quot;,         \&quot;links\&quot;: {             \&quot;avatar\&quot;: {                 \&quot;href\&quot;: \&quot;https://bitbucket.org/account/markadams-atl/avatar/32/\&quot;             },             \&quot;html\&quot;: {                 \&quot;href\&quot;: \&quot;https://bitbucket.org/markadams-atl/\&quot;             },             \&quot;self\&quot;: {                 \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/users/markadams-atl\&quot;             }         },         \&quot;type\&quot;: \&quot;user\&quot;,         \&quot;username\&quot;: \&quot;markadams-atl\&quot;,         \&quot;nickname\&quot;: \&quot;markadams-atl\&quot;,         \&quot;uuid\&quot;: \&quot;{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}\&quot;     },     \&quot;type\&quot;: \&quot;ssh_key\&quot;,     \&quot;uuid\&quot;: \&quot;{b15b6026-9c02-4626-b4ad-b905f99f763a}\&quot; } &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username The account&#39;s username or UUID.
 @param keyId The SSH key&#39;s UUID value.
 @return SshAccountKey*/
func (a *SshApiService) UsersUsernameSshKeysGet(ctx context.Context, username string, keyId string) (SshAccountKey,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SshAccountKey
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{username}/ssh-keys/"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", fmt.Sprintf("%v", keyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* SshApiService 
 Returns a paginated list of the user&#39;s SSH public keys.  Example:  &#x60;&#x60;&#x60; $ curl https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys {     \&quot;page\&quot;: 1,     \&quot;pagelen\&quot;: 10,     \&quot;size\&quot;: 1,     \&quot;values\&quot;: [         {             \&quot;comment\&quot;: \&quot;user@myhost\&quot;,             \&quot;created_on\&quot;: \&quot;2018-03-14T13:17:05.196003+00:00\&quot;,             \&quot;key\&quot;: \&quot;ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY\&quot;,             \&quot;label\&quot;: \&quot;\&quot;,             \&quot;last_used\&quot;: \&quot;2018-03-20T13:18:05.196003+00:00\&quot;,             \&quot;links\&quot;: {                 \&quot;self\&quot;: {                     \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a\&quot;                 }             },             \&quot;owner\&quot;: {                 \&quot;display_name\&quot;: \&quot;Mark Adams\&quot;,                 \&quot;links\&quot;: {                     \&quot;avatar\&quot;: {                         \&quot;href\&quot;: \&quot;https://bitbucket.org/account/markadams-atl/avatar/32/\&quot;                     },                     \&quot;html\&quot;: {                         \&quot;href\&quot;: \&quot;https://bitbucket.org/markadams-atl/\&quot;                     },                     \&quot;self\&quot;: {                         \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/users/markadams-atl\&quot;                     }                 },                 \&quot;type\&quot;: \&quot;user\&quot;,                 \&quot;username\&quot;: \&quot;markadams-atl\&quot;,                 \&quot;nickname\&quot;: \&quot;markadams-atl\&quot;,                 \&quot;uuid\&quot;: \&quot;{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}\&quot;             },             \&quot;type\&quot;: \&quot;ssh_key\&quot;,             \&quot;uuid\&quot;: \&quot;{b15b6026-9c02-4626-b4ad-b905f99f763a}\&quot;         }     ] } &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username The account&#39;s UUID, account_id, or username. Note that username has been deprecated due to [privacy changes](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#removal-of-usernames-from-user-referencing-apis).
 @return PaginatedSshUserKeys*/
func (a *SshApiService) UsersUsernameSshKeysGet_1(ctx context.Context, username string) (PaginatedSshUserKeys,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PaginatedSshUserKeys
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{username}/ssh-keys"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* SshApiService 
 Adds a new SSH public key to the specified user account and returns the resulting key.  Example: &#x60;&#x60;&#x60; $ curl -X POST -H \&quot;Content-Type: application/json\&quot; -d &#39;{\&quot;key\&quot;: \&quot;ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY user@myhost\&quot;}&#39; https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys  {     \&quot;comment\&quot;: \&quot;user@myhost\&quot;,     \&quot;created_on\&quot;: \&quot;2018-03-14T13:17:05.196003+00:00\&quot;,     \&quot;key\&quot;: \&quot;ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY\&quot;,     \&quot;label\&quot;: \&quot;\&quot;,     \&quot;last_used\&quot;: \&quot;2018-03-20T13:18:05.196003+00:00\&quot;,     \&quot;links\&quot;: {         \&quot;self\&quot;: {             \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a\&quot;         }     },     \&quot;owner\&quot;: {         \&quot;display_name\&quot;: \&quot;Mark Adams\&quot;,         \&quot;links\&quot;: {             \&quot;avatar\&quot;: {                 \&quot;href\&quot;: \&quot;https://bitbucket.org/account/markadams-atl/avatar/32/\&quot;             },             \&quot;html\&quot;: {                 \&quot;href\&quot;: \&quot;https://bitbucket.org/markadams-atl/\&quot;             },             \&quot;self\&quot;: {                 \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/users/markadams-atl\&quot;             }         },         \&quot;type\&quot;: \&quot;user\&quot;,         \&quot;username\&quot;: \&quot;markadams-atl\&quot;,         \&quot;nickname\&quot;: \&quot;markadams-atl\&quot;,         \&quot;uuid\&quot;: \&quot;{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}\&quot;     },     \&quot;type\&quot;: \&quot;ssh_key\&quot;,     \&quot;uuid\&quot;: \&quot;{b15b6026-9c02-4626-b4ad-b905f99f763a}\&quot; } &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username The account&#39;s UUID, account_id, or username. Note that username has been deprecated due to [privacy changes](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#removal-of-usernames-from-user-referencing-apis).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (SshAccountKey) The new SSH key object. Note that the username property has been deprecated due to [privacy changes](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#removal-of-usernames-from-user-referencing-apis).
 @return SshAccountKey*/
func (a *SshApiService) UsersUsernameSshKeysPost(ctx context.Context, username string, localVarOptionals map[string]interface{}) (SshAccountKey,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SshAccountKey
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{username}/ssh-keys"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(SshAccountKey); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* SshApiService 
 Updates a specific SSH public key on a user&#39;s account  Note: Only the &#39;comment&#39; field can be updated using this API. To modify the key or comment values, you must delete and add the key again.  Example: &#x60;&#x60;&#x60; $ curl -X PUT -H \&quot;Content-Type: application/json\&quot; -d &#39;{\&quot;label\&quot;: \&quot;Work key\&quot;}&#39; https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/{b15b6026-9c02-4626-b4ad-b905f99f763a}  {     \&quot;comment\&quot;: \&quot;\&quot;,     \&quot;created_on\&quot;: \&quot;2018-03-14T13:17:05.196003+00:00\&quot;,     \&quot;key\&quot;: \&quot;ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY\&quot;,     \&quot;label\&quot;: \&quot;Work key\&quot;,     \&quot;last_used\&quot;: \&quot;2018-03-20T13:18:05.196003+00:00\&quot;,     \&quot;links\&quot;: {         \&quot;self\&quot;: {             \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a\&quot;         }     },     \&quot;owner\&quot;: {         \&quot;display_name\&quot;: \&quot;Mark Adams\&quot;,         \&quot;links\&quot;: {             \&quot;avatar\&quot;: {                 \&quot;href\&quot;: \&quot;https://bitbucket.org/account/markadams-atl/avatar/32/\&quot;             },             \&quot;html\&quot;: {                 \&quot;href\&quot;: \&quot;https://bitbucket.org/markadams-atl/\&quot;             },             \&quot;self\&quot;: {                 \&quot;href\&quot;: \&quot;https://api.bitbucket.org/2.0/users/markadams-atl\&quot;             }         },         \&quot;type\&quot;: \&quot;user\&quot;,         \&quot;username\&quot;: \&quot;markadams-atl\&quot;,         \&quot;nickname\&quot;: \&quot;markadams-atl\&quot;,         \&quot;uuid\&quot;: \&quot;{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}\&quot;     },     \&quot;type\&quot;: \&quot;ssh_key\&quot;,     \&quot;uuid\&quot;: \&quot;{b15b6026-9c02-4626-b4ad-b905f99f763a}\&quot; } &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username The account&#39;s username or UUID.
 @param keyId The SSH key&#39;s UUID value.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (SshAccountKey) The updated SSH key object
 @return SshAccountKey*/
func (a *SshApiService) UsersUsernameSshKeysPut(ctx context.Context, username string, keyId string, localVarOptionals map[string]interface{}) (SshAccountKey,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SshAccountKey
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{username}/ssh-keys/"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", fmt.Sprintf("%v", keyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(SshAccountKey); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

