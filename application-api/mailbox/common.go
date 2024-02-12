package mailbox

type AuthConfig struct {
	Identity string `yaml:"identity"`
	Account  string `yaml:"account"`
	Pass     string `yaml:"pass"`
}

type SmtpConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Name string     `yaml:"name"`
	Auth AuthConfig `yaml:"auth"`
	Smtp SmtpConfig `yaml:"smtp"`
}

type MailBox struct {
	conf Config
}
