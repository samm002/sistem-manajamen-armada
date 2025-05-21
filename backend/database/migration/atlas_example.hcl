// rename to atlas.hcl 
// fill placeholders with actual value
// remove this line and all before it

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "../model",
    "--dialect", "postgres",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "<dev database url>"
  url = "<main database url>"
  migration {
    dir = "file://."
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}