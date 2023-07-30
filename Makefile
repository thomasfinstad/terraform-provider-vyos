
SHELL:=/usr/bin/bash
GO_IMPORT_ROOT:=github.com/thomasfinstad/terraform-provider-vyos

###
# Default helper target
help:
	@echo "Most targets are not ment for manual usage, these are the private targets starting with dot"
	@echo "Valid targets listed below"
	@echo "<target>: <dependency> <dependency ...>" | egrep --color '^[^ ]*:'
	@make -rpn | sed -n -e '/^$$/ { n ; /^[^ .#][^ ]*:/p ; }' | egrep --color '^[^ ]*:'

###
# VyOS ISO Upload timestamp

# Fetch and format newest ISO modification timestamp
data/vyos/rolling-iso-time.txt:
	curl -s --location --head \
		"https://s3-us.vyos.io/rolling/current/vyos-rolling-latest.iso" \
		> ".build/vyos/iso-headers.txt"

	mkdir -p data/vyos
	date \
		-d "$(shell sed -n '/last-modified:/s/last-modified: //p' ".build/vyos/iso-headers.txt")" \
		-Iseconds \
		> "data/vyos/rolling-iso-time.txt"

###
# VyOS src repo at correct commit
data/vyos/vyos-1x/submodule.log: data/vyos/rolling-iso-time.txt
	git submodule update --init --single-branch -- data/vyos/vyos-1x

	cd data/vyos/vyos-1x && \
	commit="$$(git rev-list --date=iso-strict -n 1 --before="$(shell cat "data/vyos/rolling-iso-time.txt")" "current")" && \
	git checkout "$$commit" && \
	echo "$$commit" > submodule.log

###
# Autogenerate Schemas

# Convert from relaxng to XSD
.build/vyos/schema/interface-definition.xsd: data/vyos/vyos-1x/submodule.log
	mkdir -p .build/vyos/schema/
	java -jar tools/trang-20091111/trang.jar -I rnc -O xsd data/vyos/vyos-1x/schema/interface_definition.rnc .build/vyos/schema/interface-definition.xsd

# Generate go structs from XSD
internal/vyos/schema/interfacedefinition/autogen-structs.go: .build/vyos/schema/interface-definition.xsd internal/vyos/schema/interfacedefinition/interface-definition.go

	@rm -v internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Generate structs from schema
	go run github.com/xuri/xgen/cmd/xgen -p interfacedefinition -i .build/vyos/schema/interface-definition.xsd -o internal/vyos/schema/interfacedefinition/autogen-structs.go -l Go

	# Ensure the nodes name atter will be properly unmarshaled from xml
	sed -i 's|\*NodeNameAttr.*|string `xml:"name,attr,omitempty"`|' internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Convert any undefined value as string type to stop unmarshaling from breaking
	sed -i 's|interface{}|string|' internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Add option to set if this is used as a root node
	sed -i 's|\(type [A-Za-z]*Node struct {\)|\1\nIsBaseNode bool|' internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Add a parent value to node structs
	sed -i 's|\(type [A-Za-z]*Node struct {\)|\1\nParent NodeParent|' internal/vyos/schema/interfacedefinition/autogen-structs.go

	# Format output
	gofumpt -w ./internal/vyos/schema/interfacedefinition/


###
# Terraform Resource Schemas

# Compile interface devfinitions
.build/vyos/interface-definitions: data/vyos/vyos-1x/submodule.log

	@rm -rf ".build/vyos/interface-definitions"
	python -m venv .build/vyos/vyos-1x/venv

	bash -c " \
		source .build/vyos/vyos-1x/venv/bin/activate && \
		cd data/vyos/vyos-1x && \
		pip install -r test-requirements.txt && \
		make interface_definitions \
	"

	mv data/vyos/vyos-1x/build/interface-definitions .build/vyos/interface-definitions

internal/vyos/vyosinterface/auto-package.go: .build/vyos/interface-definitions tools/build-vyos-infterface-definition-structs/main.go
	mkdir -p "internal/vyos/vyosinterface"

	@rm -fv internal/vyos/vyosinterface/auto-package.go
	@rm -fv internal/vyos/vyosinterface/autogen-*.go

	# Generate interfaces, skip xml component version metadata file
	for xmlFile in $(shell ls ".build/vyos/interface-definitions/" | grep -v "xml-component-version.xml"); do \
		echo -en "Input xml: '$${xmlFile}'\t"; \
		go run tools/build-vyos-infterface-definition-structs/*.go ".build/vyos/interface-definitions/$${xmlFile}" "internal/vyos/vyosinterface" "vyosinterface" || exit 1; \
	done

	echo -e "// Package vyosinterface generated by Makefile on $(shell date -u -Iseconds). DO NOT EDIT." > "internal/vyos/vyosinterface/auto-package.go"
	echo -e "package vyosinterface\n\n" >> "internal/vyos/vyosinterface/auto-package.go"

	echo -e 'import "$(GO_IMPORT_ROOT)/internal/vyos/schema/interfacedefinition"\n\n' >> "internal/vyos/vyosinterface/auto-package.go"

	echo -e "// GetInterfaces returns all autogenerated interface definitions" >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "func GetInterfaces() []interfacedefinition.InterfaceDefinition {" >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "return []interfacedefinition.InterfaceDefinition{" >> "internal/vyos/vyosinterface/auto-package.go"
	grep -r -o -h "func [a-z]*()" internal/vyos/vyosinterface/| sort | sed 's/func \(.*\)/\1,/g' >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "}" >> "internal/vyos/vyosinterface/auto-package.go"
	echo -e "}" >> "internal/vyos/vyosinterface/auto-package.go"

	gofumpt -w ./internal/vyos/vyosinterface/*.go

tf-resources:

	@rm -rfv internal/terraform/resource/autogen/named/
	mkdir -p "internal/terraform/resource/autogen/named"

	go run tools/build-terraform-resource-full/main.go \
		internal/terraform/resource/autogen/named \
		named \
		$(GO_IMPORT_ROOT) \
		"system/conntrack-timeout-custom-rule"

	gofumpt -w internal/terraform/resource/autogen/named/

	echo -e "// Package named generated by Makefile on $(shell date -u -Iseconds). DO NOT EDIT." > "internal/terraform/resource/autogen/named/autogen-package.go"
	echo -e "package named\n\n" >> "internal/terraform/resource/autogen/named/autogen-package.go"

	echo -e 'import (' >> "internal/terraform/resource/autogen/named/autogen-package.go"
	echo -e '"github.com/hashicorp/terraform-plugin-framework/resource"' >> "internal/terraform/resource/autogen/named/autogen-package.go"
	for f in $$(find internal/terraform/resource/autogen/named/ -name "resource.go"); do \
		grep -H "^package " "$$f" \
			| sed 's|^\(.*\)/[^/]\+:package \(.*\)$$|\2 "$(GO_IMPORT_ROOT)/\1"|' \
				>> "internal/terraform/resource/autogen/named/autogen-package.go"; \
	done
	echo -e ')\n\n' >> "internal/terraform/resource/autogen/named/autogen-package.go"

	echo -e "// GetResources returns all autogenerated resources" >> "internal/terraform/resource/autogen/named/autogen-package.go"
	echo -e "func GetResources() []func() resource.Resource {" >> "internal/terraform/resource/autogen/named/autogen-package.go"
	echo -e "return []func() resource.Resource{" >> "internal/terraform/resource/autogen/named/autogen-package.go"

	for f in $$(find internal/terraform/resource/autogen/named/ -name "resource.go"); do \
		sed -n \
			-e 's|^package \(.*\)$$|\1|p' \
			-e 's|func \(New.*\)().*|\1|p' \
			"$$f" \
				| tr '\n' '.' \
				| sed 's|.$$|,\n|'  \
					>> "internal/terraform/resource/autogen/named/autogen-package.go"; \
	done

	echo -e "}" >> "internal/terraform/resource/autogen/named/autogen-package.go"
	echo -e "}" >> "internal/terraform/resource/autogen/named/autogen-package.go"

	gofumpt -w "./internal/terraform/resource/autogen/named/autogen-package.go"

	goimports -w "./internal/terraform/resource/autogen/named"

	go generate main.go

install:
	go install .

.PHONY: clean
clean:
	rm -rfv .build
	git submodule deinit --all
