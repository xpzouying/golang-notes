package main

import (
	"fmt"
	"log"
)

type protoBuilder interface {
	buildHeader()
	buildBody()
}

type director struct {
	builder protoBuilder
}

func (d *director) build() {
	b := d.builder

	b.buildHeader()
	b.buildBody()
}

func main() {
	{

		b := newHTMLBuilder(
			withHTMLHeader("hi"),
			withHTMLBody("hello world"),
		)

		d := director{b}
		d.build()

		result := b.getResult()
		log.Println(result)
	}

	{
		b := newJSONBuilder(
			withJSONHeader("hi"),
			withJSONBody("hello world"),
		)

		d := director{b}
		d.build()

		result := b.getResult()
		log.Println(result)
	}

}

type htmlBuilder struct {
	header string
	body   string

	result string
}

type htmlBuildOption func(b *htmlBuilder)

func withHTMLHeader(header string) htmlBuildOption {
	return func(b *htmlBuilder) {
		b.header = header
	}
}

func withHTMLBody(body string) htmlBuildOption {
	return func(b *htmlBuilder) {
		b.body = body
	}
}

func newHTMLBuilder(opts ...htmlBuildOption) *htmlBuilder {
	builder := &htmlBuilder{
		header: "default header",
		body:   "default body",
	}

	for _, opt := range opts {
		opt(builder)
	}

	return builder
}

func (b *htmlBuilder) buildHeader() {
	b.result += fmt.Sprintf("<header>%s</header>", b.header)
}

func (b *htmlBuilder) buildBody() {
	b.result += fmt.Sprintf("<body>%s</body>", b.body)
}

func (b *htmlBuilder) getResult() string {
	return fmt.Sprintf("<html>%s</html>", b.result)
}

type jsonBuilder struct {
	header string
	body   string

	result string
}

type jsonBuildOption func(b *jsonBuilder)

func withJSONHeader(header string) jsonBuildOption {
	return func(b *jsonBuilder) {
		b.header = header
	}
}

func withJSONBody(body string) jsonBuildOption {
	return func(b *jsonBuilder) {
		b.body = body
	}
}

func newJSONBuilder(opts ...jsonBuildOption) *jsonBuilder {
	builder := &jsonBuilder{
		header: "default header",
		body:   "default body",
	}

	for _, opt := range opts {
		opt(builder)
	}

	return builder
}

func (b *jsonBuilder) buildHeader() {
	b.result += fmt.Sprintf(`"header": "%s", `, b.header)
}

func (b *jsonBuilder) buildBody() {
	b.result += fmt.Sprintf(`"body": "%s"`, b.body)
}

func (b *jsonBuilder) getResult() string {
	return fmt.Sprintf("{%s}", b.result)
}
