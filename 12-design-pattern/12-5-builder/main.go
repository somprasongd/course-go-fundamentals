package main

import "fmt"

// Response โครงสร้างสำหรับ JSON response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseBuilder ใช้สำหรับสร้าง response ด้วยวิธี chain method
type ResponseBuilder struct {
	response Response
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{}
}

func (b *ResponseBuilder) SetStatus(status int) *ResponseBuilder {
	b.response.Status = status
	return b
}

func (b *ResponseBuilder) SetMessage(message string) *ResponseBuilder {
	b.response.Message = message
	return b
}

func (b *ResponseBuilder) SetData(data interface{}) *ResponseBuilder {
	b.response.Data = data
	return b
}

func (b *ResponseBuilder) Build() Response {
	return b.response
}

func main() {
	// สร้าง Response โดยใช้ Builder และ chain method ด้วย Fluent Interface
	response := NewResponseBuilder().
		SetStatus(200).
		SetMessage("Success").
		SetData(map[string]string{"key": "value"}).
		Build()

	fmt.Printf("Response: %+v\n", response)
}
