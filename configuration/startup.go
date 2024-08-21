package configuration

var (
	Auth = new(AuthConfiguration)
)

func Configure() {
	SetAuthDefault(Auth)
}
