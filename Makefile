DlVVERSION=$(v)

IMAGENAME_CERTMANAGER=dlv:$(DlVVERSION)

_buildEnv=CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64

build:
	@$(_buildEnv) go build \
			-o hello -gcflags "all=-N -l"
	@echo "hello build done"

docker:build
	@docker build -t dlv:$(DlVVERSION) ./
	@echo " docker build done"

save:docker
	@sed "s/%%VERSION%%/$(DlVVERSION)/g" install.sh > deploy/install.sh
	@chmod +x deploy/install.sh
	@docker save $(IMAGENAME_CERTMANAGER) > deploy/dlv_$(DlVVERSION).tar
	@scp -r deploy root@106.14.248.134:~/
	@echo "save done"
