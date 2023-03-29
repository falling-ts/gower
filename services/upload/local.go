package upload

import (
	"mime/multipart"
	"path"
	"path/filepath"
	"time"

	"github.com/falling-ts/gower/services"

	"github.com/gin-gonic/gin"
)

const divTime = time.Hour * 24

// 本地仓库
type local struct {
	Host, Path          string
	imageDir, fileDir   string
	imagePath, filePath string
	updateAt            time.Time
}

func newLocal() services.Storage {
	l := &local{
		Host: config.Get("upload.local.host", "https://localhost").(string),
		Path: config.Get("upload.local.path", "storage/app").(string),
	}

	l.updateAt, _ = time.Parse(time.DateTime, "2023-01-01 00:00:00")
	l.updatePath()

	return l
}

// Image 上传图片
func (l *local) Image(c *gin.Context) (string, string, error) {
	image, err := c.FormFile("image")
	if err != nil {
		return "", "", err
	}

	l.updatePath()
	return l.save(c, image, l.imageDir, l.imagePath)
}

// File 上传文件
func (l *local) File(c *gin.Context) (string, string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", "", err
	}

	l.updatePath()
	return l.save(c, file, l.fileDir, l.filePath)
}

func (l *local) save(c *gin.Context, file *multipart.FileHeader, dir, _path string) (string, string, error) {
	filename := util.Nanoid() + filepath.Ext(file.Filename)
	dir = path.Join(dir, filename)
	_path = path.Join(_path, filename)
	if err := c.SaveUploadedFile(file, dir); err != nil {
		return "", "", err
	}

	return _path, path.Join(l.Host, _path), nil
}

func (l *local) updatePath() {
	if l.updateAt.After(time.Now()) {
		return
	}

	now := time.Now().Local()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	l.imagePath = path.Join("upload/images", year, month, day)
	l.filePath = path.Join("upload/files", year, month, day)

	l.imageDir = util.CreateDir(path.Join(l.Path, l.imagePath))
	l.fileDir = util.CreateDir(path.Join(l.Path, l.filePath))

	l.updateAt = time.Now().Add(divTime)
}
