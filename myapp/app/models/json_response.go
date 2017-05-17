package models

type JsonError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type JsonResponse struct {
    Success string      `json:"success"`
    Data    interface{} `json:"data"`
    Error   JsonError   `json:"error"`
}
