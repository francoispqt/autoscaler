// Copyright 2018 Drone.IO Inc
// Use of this software is governed by the Business Source License
// that can be found in the LICENSE file.

package amazon

import (
	"io/ioutil"

	"github.com/drone/autoscaler/drivers/internal/userdata"
)

// Option configures a Digital Ocean provider option.
type Option func(*provider)

// WithImage returns an option to set the image.
func WithImage(image string) Option {
	return func(p *provider) {
		p.image = image
	}
}

// WithRegion returns an option to set the target region.
func WithRegion(region string) Option {
	return func(p *provider) {
		p.region = region
	}
}

// WithSecurityGroup returns an option to set the instance size.
func WithSecurityGroup(group ...string) Option {
	return func(p *provider) {
		p.groups = group
	}
}

// WithSize returns an option to set the instance size.
func WithSize(size string) Option {
	return func(p *provider) {
		p.size = size
	}
}

// WithSSHKey returns an option to set the ssh key.
func WithSSHKey(key string) Option {
	return func(p *provider) {
		p.key = key
	}
}

// WithSubnet returns an option to set the subnet id.
func WithSubnet(id string) Option {
	return func(p *provider) {
		p.subnet = id
	}
}

// WithTags returns an option to set the image.
func WithTags(tags map[string]string) Option {
	return func(p *provider) {
		p.tags = tags
	}
}

// WithUserData returns an option to set the cloud-init
// template from text.
func WithUserData(text string) Option {
	return func(p *provider) {
		if text != "" {
			p.userdata = userdata.Parse(text)
		}
	}
}

// WithUserDataFile returns an option to set the cloud-init
// template from file.
func WithUserDataFile(filepath string) Option {
	return func(p *provider) {
		if filepath != "" {
			b, err := ioutil.ReadFile(filepath)
			if err != nil {
				panic(err)
			}
			p.userdata = userdata.Parse(string(b))
		}
	}
}
