build-proto:
	rm -f ./orch_proto/orch.pb.go ./orch_proto/orch_grpc.pb.go;
	protoc \
	--go_out=orch_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=orch_proto \
	--go-grpc_opt=paths=source_relative \
	orch.proto

run-terraform:
	terraform -chdir=infrastructure/terraform init;
	terraform -chdir=infrastructure/terraform plan;
	terraform -chdir=infrastructure/terraform apply;

run-bdd:
	docker build -f ./infrastructure/docker/Dockerfile.go_app_bdd -t hf-orch-bdd:latest .;
	docker run --rm --name hf-orch-bdd hf-orch-bdd:latest
	@docker rmi -f hf-orch-bdd >/dev/null 2>&1
	@docker rm $$(docker ps -a -f status=exited -q) -f >/dev/null 2>&1
