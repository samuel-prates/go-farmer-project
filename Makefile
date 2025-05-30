submodule-init:
	git submodule update --init --remote
	git submodule foreach --recursive git checkout master

submodule-update:
	git submodule update --remote
	git submodule foreach --recursive git checkout master
	git submodule foreach --recursive git pull

build-backend:
	cd backend && \
		make install && \
		make build

docker-backend-up:
	cd backend && \
		make docker-up

docker-backend-down:
	cd backend && \
		make docker-down

docker-backend-logs:
	cd backend && \
		make docker-log

docker-backend-erase:
	cd backend && \
		make docker-erase

docker-backend-restart:
	cd backend && \
		make docker-restart

build-frontend:
	cd frontend && \
		make install && \
		make build

docker-frontend-up:
	cd frontend && \
		make docker-up

docker-frontend-down:
	cd frontend && \
		make docker-down

docker-frontend-logs:
	cd frontend && \
		make docker-log

docker-frontend-erase:
	cd frontend && \
		make docker-erase

docker-frontend-restart:	
	cd frontend && \
		make docker-restart

build-project:
	make build-backend
	make build-frontend

docker-up:
	make build-project
	make docker-backend-up
	make docker-frontend-up

docker-down:
	make docker-backend-down
	make docker-frontend-down

docker-restart:
	make docker-backend-restart
	make docker-frontend-restart

docker-erase:
	make docker-backend-erase
	make docker-frontend-erase

test-backend:
	cd backend && \
		make test

test-frontend:
	cd frontend && \
		make jest

test-project:
	make test-frontend
	make test-backend