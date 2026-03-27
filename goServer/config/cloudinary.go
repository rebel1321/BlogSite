package config

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

var Cloudinary *cloudinary.Cloudinary

func InitCloudinary() {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_SECRET_KEY"),
	)

	if err != nil {
		log.Fatal("❌ Cloudinary init failed:", err)
	}

	Cloudinary = cld
	fmt.Println("✅ Cloudinary initialized")
}