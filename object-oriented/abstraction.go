package object_oriented

type IPictureStorage interface {
	SavePicture(picture Picture)
	GetPicture(pictureID string) Image
	DeletePicture(pictureID string)
	ModifyMetaInfo(pictureID string, metaInfo PictureMetaInfo)
}

type Picture struct{}

type PictureMetaInfo struct{}

type Image struct{}

// PictureStorage 实现 IPictureStorage 接口。
//
// Golang 隐式实现接口，结构体实现了该接口中定义的所有方法，那么这个结构体就实现了这个接口。
//
// 通过 interface 来实现抽象特性，调用者只需要知道接口中提供的功能，不需要知道具体实现逻辑。
type PictureStorage struct{}

// 抽象这个特性也可以不通过接口来实现，类的方法是通过编程语言中的函数这个语法机制实现的。
// 通过函数包裹具体的实现逻辑，本身就是一种抽象。在 Golang 中对应的就是结构体和它的方法。
//
// 调用者在使用函数的时候，并不需要研究内部具体的实现逻辑，通过函数的命名、注释和文档，了解功能后就可以直接使用。
func (p PictureStorage) SavePicture(picture Picture) {
	panic("implement me")
}

func (p PictureStorage) GetPicture(pictureID string) Image {
	panic("implement me")
}

func (p PictureStorage) DeletePicture(pictureID string) {
	panic("implement me")
}

func (p PictureStorage) ModifyMetaInfo(pictureID string, metaInfo PictureMetaInfo) {
	panic("implement me")
}
