package g

var (
	KNOWN_IP_ADDRESSES []string
	AUTO_MIGRATE       bool
)

func init() {
	KNOWN_IP_ADDRESSES = []string{"127.0.0.1", "0.0.0.0", "localhost:8080"}
	// DEBUG = true
	AUTO_MIGRATE = true
}
