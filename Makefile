release_version=_1_0_1
target_path=./bin/script-manager${release_version}
source_path=./script-manager


build:
	go build ./
	mv ${source_path} ${target_path}