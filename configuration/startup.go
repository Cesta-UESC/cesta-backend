package configuration

var (
	Auth = new(AuthConfiguration)
)

func SetDefault() {
	SetAuthDefault(Auth)
}
