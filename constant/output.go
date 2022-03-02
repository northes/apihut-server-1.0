package constant

type OutputCode string

const (
	JsonOutput   OutputCode = "json"
	ImageOutput  OutputCode = "img"
	Base64Output OutputCode = "base64"
	TextOutput   OutputCode = "text"
	SvgOutput    OutputCode = "svg"
)

func (o OutputCode) String() string {
	return string(o)
}
