build:
	go build -o bin/main ./cmd/main.go

run:
	go run main.go

setup-git-hooks:
ifeq ($(OS), Windows_NT)
	cmd.exe /c ./.git-hooks/setup/setup.bat
else
	chmod u+x ./.git-hooks/setup/setup.sh
	bash ./.git-hooks/setup/setup.sh
endif