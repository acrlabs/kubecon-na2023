KIND_CLUSTER_NAME=production

.PHONY: default all kind prometheus clean kube-prometheus/manifests

default: all

all: kind prometheus

kind:
	-kind create cluster --name $(KIND_CLUSTER_NAME) --config=kind-config.yml
	CLUSTER_NAME=${KIND_CLUSTER_NAME} ./registry-certs.sh
	kubectl apply -f local-registry-hosting.yml
	kubectl patch -n kube-system ds kindnet --patch-file kindnet-patch.yml
	kubectl apply -f 0001-conjob-crd.yml
	kubectl apply -f 0002-sk-tracer.yml
	kubectx production=kind-production

kube-prometheus/manifests: kube-prometheus/kind.jsonnet
	rm -rf kube-prometheus/manifests
	cd kube-prometheus && ./build.sh kind.jsonnet

prometheus:
	kubectl apply --server-side -f kube-prometheus/manifests/setup
	kubectl wait --for condition=Established --all CustomResourceDefinition --namespace=monitoring
	kubectl apply -f kube-prometheus/manifests
	kubectl patch -n monitoring ds node-exporter --patch-file kube-prometheus/patches/node-exporter.yml

clean:
	kubectx -d production
	kind delete cluster --name $(KIND_CLUSTER_NAME)
