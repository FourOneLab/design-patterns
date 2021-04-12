package interface_abstract

import (
	"encoding/base64"
	"image"
	"os"
)

const BucketName = "ai_images_bucket"

// AliyunImageStore 为了代码复用，封装图片存储相关逻辑，统一提供AliyunImageStore类，供整个系统使用
type AliyunImageStore struct{}

func CreateAliyunImageStore() *AliyunImageStore {
	return &AliyunImageStore{}
}

// CreateBucketIfNotExisting 创建 bucket，失败则返回错误
func (s *AliyunImageStore) CreateBucketIfNotExisting(bucketName string) error {
	return nil
}

// GenerateAccessToken 根据 accessKey / secretKey 等生成 access token
func (s *AliyunImageStore) GenerateAccessToken() string {
	return ""
}

// UploadToAliyun 上传图片到阿里云，返回图片存储的url
func (s *AliyunImageStore) UploadToAliyun(image image.Image, bucketName string, accessToken string) string {
	return ""
}

// DownloadFromAliyun 从阿里云下载图片
func (s *AliyunImageStore) DownloadFromAliyun(url string, accessToken string) image.Image {
	return nil
}

// ImageProcessingJob 是上传图片的一个示例
type ImageProcessingJob struct{}

func (j ImageProcessingJob) Process(filePath string) error {
	// filePath example = "testdata/video-001.q50.420.jpeg"
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := base64.NewDecoder(base64.StdEncoding, file)
	m, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	imageStore := CreateAliyunImageStore()
	// 创建 bucket 存储目录
	if err = imageStore.CreateBucketIfNotExisting(BucketName); err != nil {
		return err
	}

	// 生成访问凭证
	accessToken := imageStore.GenerateAccessToken()
	url := imageStore.UploadToAliyun(m, BucketName, accessToken)
	_ = url
	return nil
}

// 软件开发中，唯一不变的就是变化。

// 当图片存储从阿里云迁移到私有云上，为了满足整个需求，如何修改代码？
// 开发一个新的私有云存储类 PrivateImageStore，将原有的 AliyunImageStore 都替换掉。

// 细节是魔鬼。

// PrivateImageStore 需要将 AliyunImageStore 中所有导出的方法都逐个实现，
// 才能在尽量最小化代码修改的情况下进行替换。这样存在两个问题：
//
// - 有些方法暴露类实现细节：UploadToAliyun()/DownloadFromAliyun()中暴露aliyun，
//   没有接口意识和抽象思维时，很容易出现这样的问题，比较刚开始的需求只是存储到阿里云，
//   新的方法中肯定不能包含aliyun，那就要把项目中使用这两个方法的代码都修改。
//
// - 将图片存储到阿里云和私有云的流程可能不同：私有云不需要访问凭证，所有generateAccessToken()方法不同，
//   调用该方法的地方也就不同，那就要把项目中使用该方法的代码都修改。

// 解决上面两个问题的关键就是，在编写代码的时候遵从，基于接口而非实现编程：
// - 函数命名不要暴露任何实现细节：UploadToAliyun() --> Upload()
// - 封装具体实现：与具体存储相关的上传下载不要暴露给调用者，对外提供包裹所有细节的方法，给调用者使用
// - 为实现类定义抽象的接口：具体的实现类都依赖统一的接口定义，遵从一致的上传协议，使用者依赖接口而不是具体的某个实现类

type ImageStore interface {
	Upload(image image.Image, bucketName string) (string, error)
	Download(url string) image.Image
}

// NewAliyunImageStore 重构之后的代码，为了避免和原来的命名冲突，前面增加New
type NewAliyunImageStore struct{}

func CreateNewAliyunImageStore() *NewAliyunImageStore {
	return &NewAliyunImageStore{}
}

// Upload 上传图片到阿里云，并返回图片存储的 URL
func (n *NewAliyunImageStore) Upload(image image.Image, bucketName string) (string, error) {
	if err := n.createBucketIfNotExisting(bucketName); err != nil {
		return "", err
	}
	token := n.generateAccessToken()
	_ = token
	return "", nil
}

// Download 从阿里云下载图片
func (n *NewAliyunImageStore) Download(url string) image.Image {
	token := n.generateAccessToken()
	_ = token
	return nil
}

// 创建 bucket，失败则返回错误
func (n *NewAliyunImageStore) createBucketIfNotExisting(bucketName string) error {
	return nil
}

// 根据 accessKey / secretKey 等生成 access token
func (n *NewAliyunImageStore) generateAccessToken() string {
	return ""
}

type PrivateImageStore struct{}

func CreatePrivateImageStore() *PrivateImageStore {
	return &PrivateImageStore{}
}

// Upload 上传图片到私有云，并返回图片存储 URL
func (p *PrivateImageStore) Upload(image image.Image, bucketName string) (string, error) {
	if err := p.createBucketIfNotExisting(bucketName); err != nil {
		return "", err
	}
	return "", nil
}

// Download 从私有云下载图片
func (p *PrivateImageStore) Download(url string) image.Image {
	return nil
}

// 创建 bucket，失败则返回错误
func (p *PrivateImageStore) createBucketIfNotExisting(bucketName string) error {
	return nil
}

// NewImageProcessingJob 与上面的示例进行区分，命名前缀增加 New
type NewImageProcessingJob struct{}

func (j NewImageProcessingJob) Process(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := base64.NewDecoder(base64.StdEncoding, file)
	m, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	// 定义接口
	// TODO: 待优化的点，如果要替换图片的存储方式，这里的 ImageStore 还是需要修改
	var imageStore ImageStore
	// 实例化具体实现
	// imageStore = CreateNewAliyunImageStore()
	imageStore = CreatePrivateImageStore()
	_, err = imageStore.Upload(m, BucketName)
	if err != nil {
		return err
	}

	return nil
}
