package vpn

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// ServerEnv represents the environment ('development' or 'production') in which the application is running.
var serverEnv string

// rootURL is a variable that represents the root URL of the Netbox API based on the environment.
var rootURL string

var id int

var data string

// anyStruct is an empty interface that can be used as a generic type placeholder for API response objects.
type anyStruct interface{}

// apiConnectionID is a function that establishes a connection with the Netbox API, retrieves a specific object
// identified by the given suffix and ID, and executes the specified HTTP method on it.
// It takes in three parameters: an empty interface object (r) that represents the API response object,
// the HTTP method to be executed (httpMethod), and the suffix of the API endpoint (suffix).
// It loads the configuration file, constructs the full API path, retrieves the API token,
// and then calls the executeAPIRequest function to perform the API request.
// The function does not return anything.
func apiConnectionID[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	s := strconv.Itoa(id)

	apiSuffix := config.GetString(suffix)
	fullAPIPath := rootURL + apiSuffix + s + "/"

	color.Yellow("\n  Getting Netbox API object from %s\n", fullAPIPath)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	executeAPIRequest(httpMethod, fullAPIPath, token, r)
	if err != nil {
		log.Fatalf("Error getting Netbox API objects: %s\n", err)
	}
}

func apiConnectionPatch(suffix string) {
	config := loadConfig()

	apiSuffix := config.GetString(suffix)
	fullAPIPath := rootURL + apiSuffix

	color.Yellow("\n  Patching Netbox API objects in %s\n", fullAPIPath)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	_, err = executeAPIPatch("PATCH", fullAPIPath, token, data)
	if err != nil {
		log.Fatalf("Error patching Netbox API objects: %s\n", err)
	}
}

func apiConnectionPatchID(suffix string) {
	config := loadConfig()

	s := strconv.Itoa(id)

	apiSuffix := config.GetString(suffix)
	fullAPIPath := rootURL + apiSuffix + s + "/"

	color.Yellow("\n  Patching Netbox API object from %s\n", fullAPIPath)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	_, err = executeAPIPatchID("PATCH", fullAPIPath, token, data)
	if err != nil {
		log.Fatalf("Error patching Netbox API object: %s\n", err)
	}
}

func apiConnectionPost(suffix string) {
	config := loadConfig()

	apiSuffix := config.GetString(suffix)
	fullAPIPath := rootURL + apiSuffix

	color.Yellow("\n  Posting Netbox API objects in %s\n", fullAPIPath)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	_, err = executeAPIPost("POST", fullAPIPath, token, data)
	if err != nil {
		log.Fatalf("Error posting Netbox API objects: %s\n", err)
	}
}

func apiConnectionDelete(suffix string) {
	config := loadConfig()

	apiSuffix := config.GetString(suffix)
	fullAPIPath := rootURL + apiSuffix

	color.Yellow("\n  Deleting Netbox API object from %s\n", fullAPIPath)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	_, err = executeAPIDelete("DELETE", fullAPIPath, token, data)
	if err != nil {
		log.Fatalf("Error deleting Netbox API object! %s", err)
	}
}

func apiConnectionDeleteID(suffix string) {
	config := loadConfig()

	s := strconv.Itoa(id)

	apiSuffix := config.GetString(suffix)
	fullAPIPath := rootURL + apiSuffix + s + "/"

	color.Yellow("\n  Deleting Netbox API object from %s\n", fullAPIPath)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	_, err = executeAPIDeleteID("DELETE", fullAPIPath, token)
	if err != nil {
		log.Fatalf("Error deleting Netbox API object! %s", err)
	}
}

func ApiConnectionNonID[T anyStruct](r T, httpMethod string, suffix string) {
	config := loadConfig()

	page := 1
	s := strconv.Itoa(page)

	apiSuffix := config.GetString(suffix)
	fullAPIPath := rootURL + apiSuffix + s

	color.Yellow("\n  Getting Netbox API objects from %s\n", fullAPIPath)
	token := config.GetString("cmd.token_key")

	err := CheckSSL(rootURL)
	if err != nil {
		fmt.Println("  SSL certificate is not valid: ", err)
	} else {
		color.Cyan("  SSL certificate is valid for: " + color.YellowString("%s", rootURL))
	}

	executeAPIRequest(httpMethod, fullAPIPath, token, r)
	if err != nil {
		log.Fatalf("Error patching Netbox API objects: %s\n", err)
	}
}

func executeAPIRequest[T anyStruct](method, url, token string, object T) {
	client := resty.New()

	var request = client.R()
	request.SetHeaders(map[string]string{
		"Authorization": token,
		"Accept":        "application/json",
	})

	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = request.Get(url)
	case "POST":
		resp, err = request.Post(url)
	case "PATCH":
		resp, err = request.Patch(url)
	case "DELETE":
		resp, err = request.Delete(url)
	default:
		log.Fatalf("Invalid HTTP method: %s\n", method)
		return
	}

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return
	}

	err = json.Unmarshal([]byte(resp.String()), &object)
	if err != nil {
		log.Println("Error while parsing the response bytes:", err)
	}
}

func executeAPIDelete(method, url, token string, data string) (*resty.Response, error) {
	client := resty.New()

	var request = client.R()
	request.SetHeaders(map[string]string{
		"Authorization": token,
		"Content-Type":  "application/json",
	}).
		SetBody(data)

	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = request.Get(url)
	case "POST":
		resp, err = request.Post(url)
	case "PATCH":
		resp, err = request.Patch(url)
	case "DELETE":
		resp, err = request.Delete(url)
	default:
		log.Fatalf("Invalid HTTP method: %s\n", method)
	}

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	if resp.StatusCode() == 200 {
		log.Fatalf(color.RedString("Error executing API request: %s\n", color.YellowString(resp.Status())))
	} else if resp.StatusCode() == 404 {
		fmt.Println(color.BlueString("  No such object on Netbox server."))
	} else if resp.StatusCode() == 204 {
		fmt.Println(color.GreenString("  Successfully deleted."))
	} else if resp.StatusCode() == 409 {
		fmt.Printf(color.RedString("  Dependency Error: there is a conflict with ID: "+color.YellowString("%d - HTTP Status Code: %v\n"), id, resp.Status()))
	} else {
		log.Fatalf("Unhandled response status code: %d - %s\n", resp.StatusCode(), resp.Status())
	}
	fmt.Println(resp)
	return resp, nil
}

func executeAPIDeleteID(method, url, token string) (*resty.Response, error) {
	client := resty.New()

	var request = client.R()
	request.SetHeaders(map[string]string{
		"Authorization": token,
		"Content-Type":  "application/json",
	})

	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = request.Get(url)
	case "POST":
		resp, err = request.Post(url)
	case "PATCH":
		resp, err = request.Patch(url)
	case "DELETE":
		resp, err = request.Delete(url)
	default:
		log.Fatalf("Invalid HTTP method: %s\n", method)
	}

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	if resp.StatusCode() == 200 {
		log.Fatalf(color.RedString("Error executing API request: %s\n", color.YellowString(resp.Status())))
	} else if resp.StatusCode() == 404 {
		fmt.Println(color.BlueString("  No such object on Netbox server."))
	} else if resp.StatusCode() == 204 {
		fmt.Println(color.GreenString("  Successfully deleted."))
	} else if resp.StatusCode() == 409 {
		fmt.Printf(color.RedString("  Dependency Error: there is a conflict with ID: "+color.YellowString("%d - HTTP Status Code: %v\n"), id, resp.Status()))
	} else {
		log.Fatalf("Unhandled response status code: %d - %s\n", resp.StatusCode(), resp.Status())
	}
	fmt.Println(resp)
	return resp, nil
}

func executeAPIPatchID(method, url, token string, data string) (*resty.Response, error) {
	client := resty.New()
	request := client.R().
		SetHeaders(map[string]string{
			"Authorization": token,
			"Content-Type":  "application/json",
		}).
		SetBody(data)

	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = request.Get(url)
	case "POST":
		resp, err = request.Post(url)
	case "PATCH":
		resp, err = request.Patch(url)
	case "DELETE":
		resp, err = request.Delete(url)
	default:
		log.Fatalf("Invalid HTTP method: %s\n", method)
	}

	if err != nil {
		return nil, fmt.Errorf("Error on response.\n[ERROR] - %s", err)
	}

	if resp.StatusCode() == 200 {
		fmt.Println(color.GreenString("  Successfully patched ID: " + color.YellowString("%d\n", id)))
	} else {
		log.Fatalf("Unhandled response status code: %d - %s\n", resp.StatusCode(), resp.Status())
	}
	return resp, nil
}

func executeAPIPatch(method, url, token string, data string) (*resty.Response, error) {
	client := resty.New()
	request := client.R().
		SetHeaders(map[string]string{
			"Authorization": token,
			"Content-Type":  "application/json",
		}).
		SetBody(data)

	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = request.Get(url)
	case "POST":
		resp, err = request.Post(url)
	case "PATCH":
		resp, err = request.Patch(url)
	case "DELETE":
		resp, err = request.Delete(url)
	default:
		log.Fatalf("Invalid HTTP method: %s\n", method)
	}

	if err != nil {
		return nil, fmt.Errorf("Error on response.\n[ERROR] - %s", err)
	}

	if resp.StatusCode() == 200 {
		fmt.Println(color.GreenString("  Successfully Patched data for: " + color.YellowString("%s\n", url)))
	} else {
		log.Fatalf("Unhandled response status code: %d - %s\n", resp.StatusCode(), resp.Status())
	}
	return resp, nil
}

func executeAPIPost(method, url, token string, data string) (*resty.Response, error) {
	client := resty.New()
	request := client.R().
		SetHeaders(map[string]string{
			"Authorization": token,
			"Content-Type":  "application/json",
		}).
		SetBody(data)

	var resp *resty.Response
	var _ error

	switch method {
	case "GET":
		resp, _ = request.Get(url)
	case "POST":
		resp, _ = request.Post(url)
	case "PATCH":
		resp, _ = request.Patch(url)
	case "DELETE":
		resp, _ = request.Delete(url)
	default:
		log.Fatalf("Invalid HTTP method: %s\n", method)
	}

	if resp.StatusCode() == 201 {
		fmt.Println(color.GreenString("  Successfully Posted data for: " + color.YellowString("%s\n", url)))
	} else {
		log.Fatalf("Unhandled response status code: %d - %s\n", resp.StatusCode(), resp.Status())
	}
	return resp, nil
}

// loadConfig is a function that loads the configuration file for the Netbox API.
// It returns a pointer to a *viper.Viper object that contains the configuration values.
// The configuration file is expected to be in YAML format and named "netbox_config.yaml".
// The function sets the config name, type, and path to find the configuration file.
// It also handles environment variable substitution.
// If there is an error reading the config file, it will panic and print the error.
// The function sets the rootURL variable based on the serverEnv variable value,
// which represents the environment ('development' or 'production') in which the application is running.
// The function finally returns the config object.
func loadConfig() *viper.Viper {
	vi := viper.New()
	vi.SetConfigName("netbox_config")
	vi.SetConfigType("yaml")
	vi.AddConfigPath(".")
	vi.AutomaticEnv()

	err := vi.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file! %s", err)
		panic(err)
	}

	switch env := serverEnv; env {
	case "development":
		rootURL = vi.GetString("cmd.netbox_dev_root_url")
	case "production":
		rootURL = vi.GetString("cmd.netbox_prod_root_url")
	default:
		log.Fatalf("Unrecognized environment: %s\n", env)
	}

	return vi
}

func CheckSSL(url string) error {
	// Create a new http client with the custom transport.
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
		Timeout: time.Second * 10,
	}

	// Request the URL
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body")
		}
	}(resp.Body)

	// Check if the status code is okay
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// No problem
	return nil
}
