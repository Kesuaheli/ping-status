.PHONY: mcversion

mcversion:
	@go run . https://piston-meta.mojang.com/mc/game/version_manifest_v2.json
