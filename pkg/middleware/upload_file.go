package middleware

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	dto "todo-list-thxrhmn/dto/result"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Check if the file extension is allowed
		allowedExts := map[string]bool{
			".txt": true,
			".pdf": true,
		}
		ext := filepath.Ext(file.Filename)
		if !allowedExts[ext] {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid file extension"})
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()

		tempFile, err := ioutil.TempFile("uploads", "file-*"+ext)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer tempFile.Close()

		if _, err = io.Copy(tempFile, src); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		data := tempFile.Name()
		filename := data[8:] // split uploads/

		c.Set("dataFile", filename)
		return next(c)
	}
}
