build:
	go build -o ~/.terraform.d/plugins/terraform.local/ocitaskserv/ocitask/1.0.0/darwin_arm64/terraform-provider-ocitask main.go
	chmod 700 ~/.terraform.d/plugins/terraform.local/ocitaskserv/ocitask/1.0.0/darwin_arm64/terraform-provider-ocitask

test:
	go test -v ./ocitaskclient ./ocitaskprovider

clean:
	go clean -modcache
	rm ~/.terraform.d/plugins/terraform.local/ocitaskserv/ocitask/1.0.0/darwin_arm64/terraform-provider-ocitask
