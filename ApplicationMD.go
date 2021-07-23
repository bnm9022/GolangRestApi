package main

type ApplicationMD struct {
    Id      string `yaml:"-"`
    Title   string `yaml:"title"`
    Version    string `yaml:"version"`
    Maintainers []Maintainers `yaml:"maintainers"`
    Company string `yaml:"company"`
    Website string `yaml:"website"`
    Source string `yaml:"source"`
    License string `yaml:"license"`
    Description string `yaml:"description"`
}

type Maintainers struct {
    Name string `yaml:"name"`
    Email string `yaml:"email"`
}

