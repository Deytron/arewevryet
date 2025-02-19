package utils

import (
	"crypto/tls"
	"fmt"
	"os/exec"
	"strings"

	"github.com/go-resty/resty/v2"
)

// APICall makes an HTTP request of the specified type (GET, POST, PATCH, PUT, DELETE)
// with the provided authorization type and token.
func APICall(method, url, token string, res interface{}, body any) (bool, string) {
	client := resty.New()

	// Skip SSL verification
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	var err error
	var resp *resty.Response

	// Create a new request
	req := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(res).
		SetBody(body).
		SetHeader("Authorization", `Mediabrowser Token="`+token+`"`)


	// Select the method to be used for the request
	switch method {
	case "GET":
		resp, err = req.Get(url)
	case "POST":
		resp, err = req.Post(url)
	case "PATCH":
		resp, err = req.Patch(url)
	case "PUT":
		resp, err = req.Put(url)
	case "DELETE":
		resp, err = req.Delete(url)
	default:
		return true, "Unsupported HTTP method"
	}

	// Debug mode logging
	fmt.Println("Raw Sent Body: ", fmt.Sprintf("%+v", body))
	fmt.Println("Header: ", fmt.Sprintf("%+v", req.Header))
	if resp != nil {
		fmt.Println("Method : " + method)
		fmt.Println("Endpoint : " + url)
		fmt.Println("Status Code: ", resp.StatusCode())
		fmt.Println("Response Body: ", resp.String())
		fmt.Println("Unmarshaled Result: ", fmt.Sprintf("%+v", res))
	} else {
		fmt.Println("No response received.")
	}

	if err != nil {
		if resp != nil {
			return resp.IsError(), resp.String()
		}
		return true, fmt.Sprintf("Request failed: %v", err)
	}

	if resp == nil {
		return true, "Response is nil; no data returned."
	}

	return resp.IsError(), resp.String()
}

// APICall makes an HTTP request of the specified type (GET, POST, PATCH, PUT, DELETE)
// with the provided authorization type and token.
func APICallBasicAuth(method, url, username, password string, res interface{}, body any) (bool, string) {
	client := resty.New()

	// Skip SSL verification
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	var err error
	var resp *resty.Response

	// Create a new request
	req := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(res).
		SetBody(body).
		SetBasicAuth(username, password)

	// Select the method to be used for the request
	switch method {
	case "GET":
		resp, err = req.Get(url)
	case "POST":
		resp, err = req.Post(url)
	case "PATCH":
		resp, err = req.Patch(url)
	case "PUT":
		resp, err = req.Put(url)
	case "DELETE":
		resp, err = req.Delete(url)
	default:
		return true, "Unsupported HTTP method"
	}

	// Debug mode logging
	fmt.Println("Raw Sent Body: ", fmt.Sprintf("%+v", body))
	if resp != nil {
		fmt.Println("Method : " + method)
		fmt.Println("Endpoint : " + url)
		fmt.Println("Status Code: ", resp.StatusCode())
		fmt.Println("Response Body: ", resp.String())
		fmt.Println("Unmarshaled Result: ", fmt.Sprintf("%+v", res))
	} else {
		fmt.Println("No response received.")
	}

	if err != nil {
		if resp != nil {
			return resp.IsError(), resp.String()
		}
		return true, fmt.Sprintf("Request failed: %v", err)
	}

	if resp == nil {
		return true, "Response is nil; no data returned."
	}

	return resp.IsError(), resp.String()
}

func APIBasicAuth(url, username, password string) (bool, string) {
	client := resty.New()

	// Skip SSL verification
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	var err error
	var resp *resty.Response

	// Create a new request
	req := client.R().
		SetBasicAuth(username, password)

	resp, err = req.Post(url)

	// On debug
	req.
		SetDebug(true).
		EnableGenerateCurlOnDebug()
	fmt.Println(resp.Request.GenerateCurlCommand())
	fmt.Println("Returned answer : ", resp)

	if err != nil {
		return resp.IsError(), resp.String()
	}

	return resp.IsError(), resp.String()
}

func PingIP(ip string) bool {
	// Execute the ping command
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	output, err := cmd.CombinedOutput()

	// Check if the ping was successful (no error and output contains "1 received")
	if err == nil && strings.Contains(string(output), "1 received") {
		return true
	}

	return false
}
