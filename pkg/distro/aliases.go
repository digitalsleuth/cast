package distro

var aliases map[string]*Distro = map[string]*Distro{
	"sift": {
		Owner:   "teamdfir",
		Repo:    "sift-saltstack",
		Alias:   "sift",
		IsAlias: true,
	},
	"teamdfir/sift-saltstack": {
		Owner:   "teamdfir",
		Repo:    "sift-saltstack",
		Alias:   "sift",
		IsAlias: true,
	},
	"remnux": {
		Owner:   "remnux",
		Repo:    "salt-states",
		Alias:   "remnux",
		IsAlias: true,
	},
	"remnux/salt-states": {
		Owner:   "remnux",
		Repo:    "salt-states",
		Alias:   "remnux",
		IsAlias: true,
	},
        "mat": {
                Owner:   "digitalsleuth",
                Repo:    "mat-salt",
                Alias:   "mat",
                IsAlias: true,
        },
        "digitalsleuth/mat-salt": {
                Owner:   "digitalsleuth",
                Repo:    "mat-salt",
                Alias:   "mat",
                IsAlias: true,
        },
}

var manifests map[string]*Manifest = map[string]*Manifest{
	"sift": {
		Version: 1,
		Base:    "",
		Modes: []Mode{
			{
				Name:    "desktop",
				State:   "sift.desktop",
				Default: false,
			},
			{
				Name:    "server",
				State:   "sift.server",
				Default: true,
			},
			{
				Name:        "complete",
				State:       "sift.desktop",
				Deprecated:  true,
				Replacement: "desktop",
				Default:     false,
			},
			{
				Name:        "packages-only",
				State:       "sift.server",
				Deprecated:  true,
				Replacement: "server",
				Default:     false,
			},
		},
		SupportedOS: []SupportedOS{
			{
				ID:       "ubuntu",
				Release:  "20.04",
				Codename: "focal",
			},
		},
		Saltstack: Saltstack{
			Pillars: map[string]string{
				"sift_user_template": "{{ .User }}",
			},
		},
	},
	"remnux": {
		Version: 1,
		Base:    "",
		Modes: []Mode{
			{
				Name:    "dedicated",
				State:   "remnux.dedicated",
				Default: true,
			},
			{
				Name:    "addon",
				State:   "remnux.addon",
				Default: false,
			},
			{
				Name:    "cloud",
				State:   "remnux.cloud",
				Default: false,
			},
		},
		Saltstack: Saltstack{
			Pillars: map[string]string{
				"remnux_user_template": "{{ .User }}",
			},
		},
	},
        "mat": {
                Version: 1,
                Base:    "",
                Modes: []Mode{
                        {
                                Name:    "dedicated",
                                State:   "mat.dedicated",
                                Default: false,
                        },
                        {
                                Name:    "addon",
                                State:   "mat.addon",
                                Default: true,
                        },
                },
                SupportedOS: []SupportedOS{
                        {
                                ID:       "ubuntu",
                                Release:  "20.04",
                                Codename: "focal",
                        },
                },
                Saltstack: Saltstack{
                        Pillars: map[string]string{
                                "mat_user_template": "{{ .User }}",
                        },
                },
        },

}
