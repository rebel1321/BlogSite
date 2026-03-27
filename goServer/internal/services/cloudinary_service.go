package services

import (
	"context"
	"mime/multipart"

	"goServer/config"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// Upload file to Cloudinary
func UploadFile(file multipart.File) (string, string, error) {
	ctx := context.Background()

	result, err := config.Cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "blog",
	})
	if err != nil {
		return "", "", err
	}

	return result.SecureURL, result.PublicID, nil
}

// Delete file from Cloudinary
func DeleteFile(publicID string) error {
	ctx := context.Background()

	_, err := config.Cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})

	return err
}