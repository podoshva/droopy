// Package httpapi
package httpapi

import (
	"main/internal/handler/httpapi/handlers"
	"net/http"
)

type RouterDeps struct {
	Handlers handlers.Handlers
}

func NewRouter(deps RouterDeps) http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /")                                                          // Страница загрузки файла
	mux.HandleFunc("POST /api/uploads/init/", deps.Handlers.InitUpload)              // Нажали Upload
	mux.HandleFunc("POST /api/uploads/{id}/complete/", deps.Handlers.CompleteUpload) // Загрузили файл в S3

	// mux.HandleFunc("GET /s/{share_token}/")                                                // страница файла
	mux.HandleFunc("GET /api/share/{share_token}/", deps.Handlers.GetShareMeta)            // получаем метаданные файла
	mux.HandleFunc("GET /api/share/{share_token}/download/", deps.Handlers.GetDownloadURL) // получаем Presigned URL (нажали кнопу Download)

	return mux
}
