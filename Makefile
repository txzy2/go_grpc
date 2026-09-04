run-all:
	$(MAKE) -j2 rf rs

down-docker:
	$(MAKE) -j2 ddf dds

rf:
	cd first && docker compose -f compose.dev.yml up -d && make run

rs:
	cd second && docker compose -f compose.dev.yml up -d && make run

ddf:
	cd first && docker compose -f compose.dev.yml down

dds:
	cd second && docker compose -f compose.dev.yml down