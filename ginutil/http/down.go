package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

// FileDataStream 封装返回文件流
func FileDataStream(ctx *gin.Context, fileName string) {
	ctx.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Writer.Header().Set("Pragma", "no-cache")
	ctx.Writer.Header().Set("Expires", "0")
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.Writer.Header().Add("Content-Disposition", "attachment;filename*=utf-8''"+path.Base(fileName))
	ctx.Writer.Header().Add("Content-Transfer-Encoding", "binary")

	bFile, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	ctx.Data(http.StatusOK, "application/octet-stream", bFile)
}

// FileStream 封装返回文件流
func FileStream(ctx *gin.Context, fileName string) {
	ctx.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Writer.Header().Set("Pragma", "no-cache")
	ctx.Writer.Header().Set("Expires", "0")
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.Writer.Header().Add("Content-Disposition", "attachment;filename*=utf-8''"+path.Base(fileName))
	ctx.Writer.Header().Add("Content-Transfer-Encoding", "binary")
	ctx.File(fileName)
}
