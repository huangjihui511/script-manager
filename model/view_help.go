package model

var (
	DemoYaml = `# This is a demo
config-name: jihui's config
command-groups:
  - group-name: shell
    commands:
	# required
    - script: echo "hello world"
	# optional
      info: change the world
    - script: echo "hello world again"
  - group-name: go
    commands:
    - script: go fmt ./...
    - script: go mod tidy
`
)
