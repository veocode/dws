package models

type DockerComposeFile struct {
	Version  string `yaml:"version"`
	Services struct {
		App struct {
			Build struct {
				Context string `yaml:"context"`
			} `yaml:"build"`
			Volumes []string `yaml:"volumes"`
			Links   []string `yaml:"links"`
		} `yaml:"app"`
		Webserver struct {
			Image   string   `yaml:"image"`
			Ports   []string `yaml:"ports"`
			Volumes []string `yaml:"volumes"`
			Links   []string `yaml:"links"`
		} `yaml:"webserver"`
		Pgsql struct {
			Image       string `yaml:"image"`
			Environment struct {
				POSTGRESDB       string `yaml:"POSTGRES_DB"`
				POSTGRESUSER     string `yaml:"POSTGRES_USER"`
				POSTGRESPASSWORD string `yaml:"POSTGRES_PASSWORD"`
			} `yaml:"environment"`
			Ports   []string `yaml:"ports"`
			Volumes []string `yaml:"volumes"`
		} `yaml:"pgsql"`
		Redis struct {
			Image   string   `yaml:"image"`
			Ports   []string `yaml:"ports"`
			Volumes []string `yaml:"volumes"`
		} `yaml:"redis"`
	} `yaml:"services"`
	Networks struct {
		Default struct {
			Ipam struct {
				Driver string `yaml:"driver"`
				Config []struct {
					Subnet string `yaml:"subnet"`
				} `yaml:"config"`
			} `yaml:"ipam"`
		} `yaml:"default"`
	} `yaml:"networks"`
}
