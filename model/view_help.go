package model

var (
	QAs = []QA{
		{
			Question: "How to edit my script?",
			Answer: `
	1. Create environment variable named SM_CONF_PATH
	2. Let it be the script definition yaml file path on your machine
	3. Click the "Demo" botton which will copy the demo yaml file content to your clipboard
	4. Edit your script definition yaml file refering to the demo`,
		},
		{
			Question: "How to update the app after my edits",
			Answer:   `Click the botton named "Update"`,
		},
	}
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

type QA struct {
	Question string
	Answer   string
}
