package main

import (
	"fmt"
)
var (
	galleryArray [10] string
// Initialize a slice to store image filenames
 gallery = galleryArray[:0]
)
// Upload image function
func uploadImage(filename string) {
  if len(gallery) >= 10{
	fmt.Println("upload size exceeded")
	return
}
gallery = append(gallery,filename )
fmt.Println("uploaded:", filename)
}

// Delete image function
func deleteImage(filename string) {
index :=-1
for i, img := range gallery{
	if img ==filename{
		index = i
	break
	}
}
if index == -1{
	fmt.Println("failed to delete:",filename)
	return
}
copy(gallery[index:], gallery[index+1:])
gallery = gallery[:len(gallery)-1]
fmt.Println("deleted: ", filename)
}

// List images function
func listImages() {
if len(gallery) == 0{
	fmt.Println("gallery is empty")
	return
}

fmt.Println("gallery images")
for _, img := range gallery{
fmt.Printf("images in the gallery are %s\n", img)
}
}

func mian() {
  fmt.Println("Testing Image Gallery Management System")
	uploadImage("image1.jpg")
	uploadImage("image2.jpg")
	uploadImage("image3.jpg")
	listImages()
	deleteImage("image2.jpg")
	listImages()
	uploadImage("image4.jpg")
	listImages()
	deleteImage("image1.jpg")
	deleteImage("image2.jpg")
	deleteImage("image3.jpg")
	listImages()
}