package design_principles

import "net/http"

type HTML struct{}

func NewHTML(raw []byte) *HTML {
	return &HTML{}
}

type NetworkTransporter struct{}

func NewNetworkTransporter() *NetworkTransporter {
	return &NetworkTransporter{}
}

func (t *NetworkTransporter) SendV1(r *http.Request) []byte {
	return nil
}

// SendV2
// address 和 content 交给 NetworkTransporter，而非是直接把 http.Request 交给 NetworkTransporter。
func (t *NetworkTransporter) SendV2(address string, data []byte) []byte {
	return nil
}

type HtmlDownloader struct {
	transporter *NetworkTransporter
}

func NewHtmlDownloader(networkTransporter *NetworkTransporter) *HtmlDownloader {
	return &HtmlDownloader{transporter: networkTransporter}
}

func (d *HtmlDownloader) DownloadHTML(url string) (*HTML, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	//raw := d.transporter.SendV1(req)

	data := make([]byte, 0)
	_, err = req.Body.Read(data)
	if err != nil {
		return nil, err
	}

	raw := d.transporter.SendV2(req.Host, data)
	return NewHTML(raw), nil
}

type Document struct {
	html *HTML
	url  string
}

func NewDocument(url string) (*Document, error) {
	htmlDownloader := NewHtmlDownloader(NewNetworkTransporter())
	html, err := htmlDownloader.DownloadHTML(url)
	if err != nil {
		return nil, err
	}

	return &Document{url: url, html: html}, nil
}

func NewDocumentV2(url string, html *HTML) (*Document, error) {
	return &Document{}, nil
}

// DocumentFactory 通过工厂方法来创建Document
type DocumentFactory struct {
	downloader HtmlDownloader
}

func NewDocumentFactory(downloader HtmlDownloader) *DocumentFactory {
	return &DocumentFactory{downloader: downloader}
}

func (f *DocumentFactory) CreateDocument(url string) (*Document, error) {
	html, err := f.downloader.DownloadHTML(url)
	if err != nil {
		return nil, err
	}

	return NewDocumentV2(url, html)
}
