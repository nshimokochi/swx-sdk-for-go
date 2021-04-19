package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	_nethttp "net/http"
	"os"
	"testing"
)

const NEG_SPACE = "test1negative2space3"
const NEG_DATA_ID = "00TEST00NEGATIVE00DATAID00"

var dataSource string

func check2XX(httpCode int, subject string) func(...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		// Parse args
		var resp interface{}
		var r *_nethttp.Response
		var err error
		switch len(args) {
		case 2:
			r = args[0].(*_nethttp.Response)
			if args[1] != nil {
				err = args[1].(error)
			}
		case 3:
			resp = args[0]
			r = args[1].(*_nethttp.Response)
			if args[2] != nil {
				err = args[2].(error)
			}
		default:
			return resp, fmt.Errorf("Unknown # of arguments for API call `%s`", subject)
		}

		// Check for known HTTP Status Code
		var httpStatus string
		switch httpCode {
		case 200:
			httpStatus = "200 - OK"
		case 201:
			httpStatus = "201 - Created"
		case 204:
			httpStatus = "204 - No Content"
		default:
			return resp, fmt.Errorf("No test defined for HTTP status (%d) in `%s`", httpCode, subject)
		}

		// Parse response body
		if os.Getenv("SHOW_TEST_LOG") == "true" && httpCode != 204 {
			body, readErr := ioutil.ReadAll(r.Body)
			if readErr != nil {
				return resp, fmt.Errorf("Error when reading `%s` body: %v", subject, readErr)
			}
			fmt.Printf("Response from `%s` body: %v\n", subject, string(body))
		}

		if r.StatusCode != httpCode {
			return resp, fmt.Errorf("`%s` HTTP status = %v; want %s\n", subject, r.StatusCode, httpStatus)
		}
		if err != nil {
			return resp, fmt.Errorf("`%s` Error response = %v; want %s\n", subject, err, httpStatus)
		}
		return resp, nil
	}
}

func check4XX(httpCode int, subject string) func(...interface{}) error {
	return func(args ...interface{}) error {
		var r *_nethttp.Response
		var err error

		switch len(args) {
		case 2:
			r = args[0].(*_nethttp.Response)
			err = args[1].(error)
		case 3:
			r = args[1].(*_nethttp.Response)
			err = args[2].(error)
		default:
			return fmt.Errorf("Unknown # of arguments for API call `%s`", subject)
		}

		// Check for known HTTP Status Code
		var httpStatus string
		switch httpCode {
		case 403:
			httpStatus = "403 - Forbidden"
		case 404:
			httpStatus = "404 - Not Found"
		default:
			return fmt.Errorf("No test defined for HTTP status (%d) in `%s`", httpCode, subject)
		}

		body, readErr := ioutil.ReadAll(r.Body)
		if readErr != nil {
			return fmt.Errorf("Error when reading `%s` body: %v", subject, readErr)
		}
		if os.Getenv("SHOW_TEST_LOG") == "true" {
			fmt.Printf("Response from `%s` body: %v\n", subject, string(body))
		}

		if r.StatusCode != httpCode {
			return fmt.Errorf("`%s` HTTP status = %v; want Error: %s\n", subject, r.StatusCode, httpStatus)
		}
		if err == nil {
			return fmt.Errorf("`%s` response = %v; want Error: %s\n", subject, string(body), httpStatus)
		}
		baseErr := BaseError{}
		err = json.Unmarshal(body, &baseErr)
		if err != nil {
			return fmt.Errorf("Error when decoding `%s` body to type `%T`: %v", subject, baseErr, err)
		}
		// Ideally it should not be nil. Thus, only check if code exists.
		if baseErr.Error.Code != nil && *baseErr.Error.Code != int32(httpCode) {
			return fmt.Errorf("`%s` Error code = %v; want Error: %s\n", subject, *baseErr.Error.Code, httpStatus)
		}
		return nil
	}
}

// Special check for negative data_id in ShowData. Use standard check404 once server is fixed.
func checkShow404(subject string) func(interface{}, *_nethttp.Response, error) error {
	return func(resp interface{}, r *_nethttp.Response, err error) error {
		body, readErr := ioutil.ReadAll(r.Body)
		if readErr != nil {
			return fmt.Errorf("Error when reading `%s` body: %v", subject, readErr)
		}
		if os.Getenv("SHOW_TEST_LOG") == "true" {
			fmt.Printf("Response from `%s` body: %v\n", subject, string(body))
		}

		if r.StatusCode != 404 {
			return fmt.Errorf("`%s` HTTP status = %v; want Error: 404 - Not Found\n", subject, r.StatusCode)
		}
		if err == nil {
			return fmt.Errorf("`%s` response = %v; want Error: 404 - Not Found\n", subject, string(body))
		}

		var objMap map[string]map[string]interface{}
		err = json.Unmarshal(body, &objMap)
		if err != nil {
			return fmt.Errorf("Error when decoding `%s` body to type `%T`: %v", subject, objMap, err)
		}

		if _, ok := objMap["error"]; !ok {
			return fmt.Errorf("`%s` response = %v; want error object\n", subject, string(body))
		}
		if val, ok := objMap["error"]["status"].(float64); !ok || val != 404 {
			return fmt.Errorf("`%s` response = %v; want error.status 404\n", subject, string(body))
		}
		return nil
	}
}

func TestNegativeCreateData(t *testing.T) {
	err := check4XX(403, "DataApi.CreateData")(apiClient.DataApi.CreateData(context.Background(), NEG_SPACE).Body("test").Execute())
	if err != nil {
		t.Error(err)
	}
}

func TestPositiveCreateData(t *testing.T) {
	contents := []interface{}{map[string]interface{}{
		"cpu":    map[string]interface{}{"percentage": "10.7"},
		"disk":   map[string]interface{}{"percentage": "52.1"},
		"memory": map[string]interface{}{"percentage": "65.9"},
	}, "{\n  \"hello\": \"world\"\n}", 12345, false, nil, map[string]interface{}{"hello": "world"}}

	for _, content := range contents {
		resp, err := check2XX(201, "DataApi.CreateData")(apiClient.DataApi.CreateData(context.Background(), os.Getenv("SPACE")).Body(content).Execute())
		if err != nil {
			t.Error(err)
		}
		response, ok := resp.(CreateDataResponse)
		if !ok {
			t.Error(fmt.Errorf("`DataApi.CreateData` response = %+v; want CreateDataResponse struct\n", resp))
		}
		if response.Result == nil {
			t.Error(fmt.Errorf("`DataApi.CreateData` Result = %v; want ID\n", response.Result))
		}
		dataSource = *response.Result
	}
}

func TestNegativeListData(t *testing.T) {
	err := check4XX(403, "DataApi.ListData")(apiClient.DataApi.ListData(context.Background(), NEG_SPACE).Source(os.Getenv("SOURCE")).Execute())
	if err != nil {
		t.Error(err)
	}
}

func TestPositiveListData(t *testing.T) {
	resp, err := check2XX(200, "DataApi.ListData")(apiClient.DataApi.ListData(context.Background(), os.Getenv("SPACE")).Source(os.Getenv("SOURCE")).Execute())
	if err != nil {
		t.Error(err)
	}
	response, ok := resp.(ListDataResponse)
	if !ok {
		t.Error(fmt.Errorf("`DataApi.ListData` response = %+v; want ListDataResponse struct\n", resp))
	}
	if response.Collection == nil {
		t.Error(fmt.Errorf("`DataApi.ListData` Collection = %+v; want non-empty Collection\n", response.Collection))
	}
}

func TestNegativeShowData(t *testing.T) {
	// Negative space
	err := check4XX(403, "DataApi.ShowData")(apiClient.DataApi.ShowData(context.Background(), NEG_SPACE, dataSource).Execute())
	if err != nil {
		t.Error(err)
	}
	// Negative data_id
	err = checkShow404("DataApi.ShowData")(apiClient.DataApi.ShowData(context.Background(), os.Getenv("SPACE"), NEG_DATA_ID).Execute())
	if err != nil {
		t.Error(err)
	}
}

func TestPositiveShowData(t *testing.T) {
	resp, err := check2XX(200, "DataApi.ShowData")(apiClient.DataApi.ShowData(context.Background(), os.Getenv("SPACE"), dataSource).Execute())
	if err != nil {
		t.Error(err)
	}

	respStr, err := json.Marshal(resp)
	if err != nil {
		t.Error(fmt.Errorf("Error when encoding `DataApi.ShowData` response of type `%T`: %v", resp, err))
	}
	dataEl := DataElement{}
	err = json.Unmarshal(respStr, &dataEl)
	if err != nil {
		t.Error(fmt.Errorf("Error when decoding `DataApi.ShowData` body to type `%T`: %v", dataEl, err))
	}
	if dataEl.At == nil {
		t.Error(fmt.Errorf("`DataApi.ShowData` at = %v; want Time\n", dataEl.At))
	}
	if dataEl.Content == nil {
		t.Error(fmt.Errorf("`DataApi.ShowData` Content = %v; want any-value Content\n", dataEl.Content))
	}
	if dataEl.Id == nil {
		t.Error(fmt.Errorf("`DataApi.ShowData` id = %v; want string\n", dataEl.Id))
	}
	if dataEl.SourceId == nil {
		t.Error(fmt.Errorf("`DataApi.ShowData` source_id = %v; want string\n", dataEl.SourceId))
	}
}

func TestNegativeDeleteData(t *testing.T) {
	err := check4XX(403, "DataApi.DeleteData")(apiClient.DataApi.DeleteData(context.Background(), NEG_SPACE, dataSource).Execute())
	if err != nil {
		t.Error(err)
	}
}

func TestPositiveDeleteData(t *testing.T) {
	_, err := check2XX(204, "DataApi.DeleteData")(apiClient.DataApi.DeleteData(context.Background(), os.Getenv("SPACE"), dataSource).Execute())
	if err != nil {
		t.Error(err)
	}
}

func TestNegativeDeleteDataFromSource(t *testing.T) {
	err := check4XX(403, "DataApi.DeleteDataFromSource")(apiClient.DataApi.DeleteDataFromSource(context.Background(), NEG_SPACE).Source(os.Getenv("SOURCE")).Execute())
	if err != nil {
		t.Error(err)
	}
}

func TestPositiveDeleteDataFromSource(t *testing.T) {
	_, err := check2XX(204, "DataApi.DeleteDataFromSource")(apiClient.DataApi.DeleteDataFromSource(context.Background(), os.Getenv("SPACE")).Source(os.Getenv("SOURCE")).Execute())
	if err != nil {
		t.Error(err)
	}
}
