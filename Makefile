.PHONY: submodule-init submodule-update submodule-status

# Path to the Bybit docs submodule
SUBMODULE_PATH := libs/bybit-docs

# Initialize the submodule (first-time clone or after fresh checkout)
submodule-init:
	git submodule update --init --recursive $(SUBMODULE_PATH)

# Update the submodule to the latest upstream commit on its tracked branch
# After running this, commit the updated pointer in the superproject:
#   git add $(SUBMODULE_PATH) && git commit -m "chore: update submodule"
submodule-update:
	git submodule sync -- $(SUBMODULE_PATH)
	git submodule update --init --remote $(SUBMODULE_PATH)

# Show current submodule revision(s)
submodule-status:
	git submodule status --recursive
