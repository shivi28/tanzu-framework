// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package packageclient

import (
	"github.com/pkg/errors"

	kappipkg "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/packaging/v1alpha1"
	kapppkg "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apiserver/apis/datapackaging/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/packageclients/pkg/packagedatamodel"
)

func (p *pkgClient) ListPackageInstalls(o *packagedatamodel.PackageOptions) (*kappipkg.PackageInstallList, error) {
	packageInstallList, err := p.kappClient.ListPackageInstalls(o.Namespace)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list installed packages")
	}
	return packageInstallList, nil
}

func (p *pkgClient) ListPackageMetadata(o *packagedatamodel.PackageAvailableOptions) (*kapppkg.PackageMetadataList, error) {
	packageList, err := p.kappClient.ListPackageMetadata(o.Namespace)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list packages")
	}
	return packageList, nil
}

func (p *pkgClient) ListPackages(o *packagedatamodel.PackageAvailableOptions) (*kapppkg.PackageList, error) {
	packageVersionList, err := p.kappClient.ListPackages(o.PackageName, o.Namespace)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list package versions")
	}
	return packageVersionList, nil
}
