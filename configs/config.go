package configs

type Config struct {
	DataUrl            string `yaml:"DATA_URL"`
	TargetBuildVersion string `yaml:"BUILD_VERSION"`
	SavePath           string `yaml:"SAVE_PATH"`
}
