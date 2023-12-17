#!/bin/bash

# change this fields according to yours
packagename=github.com/sheikh-arman/test-crd-controller
groupname=arman.com
versionname=v1alpha1


depelopmentDir=$(pwd)
k8spath=$HOME/go/src/k8s.io

# getting the code-generator
#mkdir -p $k8spath
#cd $k8spath && git clone https://github.com/kubernetes/code-generator.git

 cd $depelopmentDir && go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5
 cd $depelopmentDir && go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5


# Generate clientset, informers & listers
execDir=$k8spath/code-generator
cd $depelopmentDir && "${execDir}"/generate-groups.sh all $packagename/pkg/client $packagename/pkg/apis $groupname:$versionname --go-header-file "${execDir}"/examples/hack/boilerplate.go.txt


# Get the dependancies
cd $depelopmentDir && go mod tidy


# Generate crd manifest file
cd $depelopmentDir && controller-gen crd:trivialVersions=true crd:crdVersions=v1 paths=./... output:crd:dir=$depelopmentDir/manifests output:stdout
#controller-gen paths=./... crd:trivialVersions=true crd:crdVersions=v1 rbac:roleName=controller-perms output:crd:artifacts:config=config/crd/bases


# controller-gen rbac:roleName=controller-perms crd paths=./... output:crd:dir=$depelopmentDir/manifests output:stdout
