package api

import (
	// "bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	// "os/exec"

	// "strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (server *Server) updatenetworkdb(ctx *gin.Context) {
	os.Mkdir("/app/downloads", 0700)
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.Error(err)
		return
	}
	//Open received file
	fileToImport, err := fileHeader.Open()
	if err != nil {
		ctx.Error(err)
		return
	}
	defer fileToImport.Close()
	filenamenew := fileHeader.Filename
	tempFile, err := os.Create("/app/downloads/"+filenamenew)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()		

	//Delete temp file after importing
	defer os.Remove("/app/downloads/"+filenamenew)

	//Write data from received file to temp file
	fileBytes, err := io.ReadAll(fileToImport)
	if err != nil {
		ctx.Error(err)
		return
	}
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		ctx.Error(err)
		return
	}
	///////////////////////////

	//////////////////////////

	//////////////////////////
	ctx.JSON(http.StatusOK, "File "+filenamenew+" uploaded successfully")
	time.Sleep(1 * time.Second)
}
