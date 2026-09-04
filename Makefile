run-all:
	$(MAKE) -j2 rf rs

rf:
	cd first && make run

rs:
	cd second && make run