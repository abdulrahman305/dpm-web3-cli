package types

type mp map[string]string
type PackageJSON struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	Private         *bool  `json:"private,omitempty"`
	Scripts         mp     `json:"scripts,omitempty"`
	Dependencies    mp     `json:"dependencies,omitempty"`
	DevDependencies mp     `json:"devDependencies,omitempty"`
	Resolutions     mp     `json:"resolutions,omitempty"`
	Main            string `json:"main,omitempty"`
	License         string `json:"license,omitempty"`
}

func NewBasicPkgJson(name string) PackageJSON {
	private := true
	return PackageJSON{
		Name:    name,
		Version: "1.0.0",
		Private: &private,
		Scripts: mp{
			"test": "echo \"Error: no test specified\" && exit 1",
		},
		Main:    "index.js",
		License: "ISC",
	}
}
