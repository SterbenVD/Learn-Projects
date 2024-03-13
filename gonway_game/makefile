bin_file = ./bin/exec
settings_file = ./bin/settings.json

run: $(bin_file) $(settings_file)
	$(bin_file) -settings $(settings_file)

build:
	go build -o $(bin_file) ./main.go

build_wayland:
	go build -tags wayland -o $(bin_file) ./main.go

clean:
	rm -rf ./bin