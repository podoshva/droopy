package main

import (
	"context"
	"log"
	"main/internal/adapter/postgres"
	"main/internal/adapter/postgres/repos"
	"main/internal/adapter/s3"
	"main/internal/app"
	"main/internal/handler/httpapi"
	"main/internal/handler/httpapi/handlers"
	"main/pkg"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cfg pkg.Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	client, err := s3.NewClient(ctx, cfg.S3Config)
	if err != nil {
		panic(err)
	}
	pool, err := postgres.NewPool(ctx, cfg.DBConfig)
	if err != nil {
		panic(err)
	}

	shareRepo := repos.NewShareRepo(pool)
	shareService := app.NewShareService(shareRepo)

	uploadRepo := repos.NewUploadRepo(pool)
	uploadService := app.NewUploadService(uploadRepo)

	router := httpapi.NewRouter(httpapi.RouterDeps{
		Handlers: handlers.Handlers{
			ShareService:  *shareService,
			UploadService: *uploadService,
		},
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// objects := client.ListObjects(ctx, "b6ae0db60774-podoshva", minio.ListObjectsOptions{
	// 	Recursive: true,
	// })

	// for object := range objects {
	// 	if object.Err != nil {
	// 		fmt.Println(object.Err)
	// 		return
	// 	}
	// 	fmt.Printf("Name: %v\nSize: %v\n\n", object.Key, object.Size)
	// }

	// opts := minio.RemoveObjectOptions{
	// 	GovernanceBypass: true,
	// }
	// err = client.RemoveObject(ctx, "b6ae0db60774-podoshva", "test/", opts)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// uploadFile(ctx, client)

	// getPresignedURL(ctx, client)

}

// func getPresignedURL(ctx context.Context, client *minio.Client) {
// 	reqParams := make(url.Values)
// 	reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")

// 	// Generates a presigned url which expires in a day.
// 	presignedURL, err := client.PresignedGetObject(ctx, "b6ae0db60774-podoshva", "readme", time.Second*24*60*60, reqParams)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Successfully generated presigned URL", presignedURL)
// }

// func printAvailableBuckets(ctx context.Context, client *minio.Client) {
// 	buckets, err := client.ListBuckets(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	for _, bucket := range buckets {
// 		fmt.Println(bucket)
// 	}
// }

// func uploadFile(ctx context.Context, client *minio.Client) {
// 	file, err := os.Open("./README.md")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fileStat, err := file.Stat()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	uploadInfo, err := client.PutObject(ctx, "b6ae0db60774-podoshva", file.Name(), file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("Successfuly uploaded: ", uploadInfo)
// }

// short_link | original_file_name | s3_key | created_at |
