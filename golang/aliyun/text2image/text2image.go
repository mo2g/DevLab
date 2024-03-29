package text2image

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ImageRequest struct {
	Model string `json:"model"`
	Input struct {
		Prompt         string `json:"prompt"`
		NegativePrompt string `json:"negative_prompt,omitempty"`
	} `json:"input"`
	Parameters struct {
		Size  string `json:"size,omitempty"`
		N     int    `json:"n,omitempty"`
		Steps string `json:"steps,omitempty"`
		Scale string `json:"scale,omitempty"`
	} `json:"parameters,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
}

type TaskResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	RequestId  string `json:"request_id,omitempty"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	Output     struct {
		TaskId     string `json:"task_id,omitempty"`
		TaskStatus string `json:"task_status,omitempty"`
		Code       string `json:"code,omitempty"`
		Message    string `json:"message,omitempty"`
		Results    []struct {
			B64Image string `json:"b64_image,omitempty"`
			Url      string `json:"url,omitempty"`
			Code     string `json:"code,omitempty"`
			Message  string `json:"message,omitempty"`
		} `json:"results,omitempty"`
		TaskMetrics struct {
			Total     int `json:"TOTAL,omitempty"`
			Succeeded int `json:"SUCCEEDED,omitempty"`
			Failed    int `json:"FAILED,omitempty"`
		} `json:"task_metrics,omitempty"`
	} `json:"output,omitempty"`
}

func (i *ImageRequest) ImageSynthesis(apikey string) (*TaskResponse, error) {
	aliTaskResponse := &TaskResponse{}
	url := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text2image/image-synthesis"

	jsonStr, err := json.Marshal(i)
	if err != nil {
		return aliTaskResponse, err
	}
	requestBody := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return aliTaskResponse, err
	}

	req.Header.Set("Authorization", "Bearer "+apikey)
	req.Header.Set("X-DashScope-Async", "enable")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return aliTaskResponse, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return aliTaskResponse, err
	}
	err = resp.Body.Close()
	if err != nil {
		return aliTaskResponse, err
	}
	err = json.Unmarshal(responseBody, &aliTaskResponse)
	if err != nil {
		return aliTaskResponse, err
	}

	aliResponse, err := i.asyncTaskWait(aliTaskResponse.Output.TaskId, apikey)
	if err != nil {
		return aliTaskResponse, err
	}

	return aliResponse, err
}

func (i *ImageRequest) asyncTask(taskID string, key string) (*TaskResponse, error) {
	url := fmt.Sprintf("https://dashscope.aliyuncs.com/api/v1/tasks/%s", taskID)

	var aliResponse TaskResponse

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &aliResponse, err
	}

	req.Header.Set("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &aliResponse, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)

	var response TaskResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return &aliResponse, err
	}

	return &response, nil
}

func (i *ImageRequest) asyncTaskWait(taskID string, key string) (*TaskResponse, error) {
	waitSeconds := 2
	step := 0

	var taskResponse TaskResponse

	for {
		step++
		rsp, err := i.asyncTask(taskID, key)
		if err != nil {
			return &taskResponse, err
		}

		if rsp.Output.TaskStatus == "" {
			return &taskResponse, nil
		}

		switch rsp.Output.TaskStatus {
		case "FAILED":
			fallthrough
		case "CANCELED":
			fallthrough
		case "SUCCEEDED":
			fallthrough
		case "UNKNOWN":
			return rsp, nil
		}

		time.Sleep(time.Duration(waitSeconds) * time.Second)
	}

	return &taskResponse, nil
}
