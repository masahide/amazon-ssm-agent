// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the AWS Customer Agreement (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/agreement/

// Package platform provides instance information
package platform

import (
	"github.com/aws/amazon-ssm-agent/agent/managedInstances/registration"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

// dependency for managed instance registration
var managedInstance instanceRegistration = instanceInfo{}

type instanceRegistration interface {
	InstanceID() string
	Region() string
}

type instanceInfo struct{}

// ServerID returns the managed instance ID
func (instanceInfo) InstanceID() string { return registration.InstanceID() }

// Region returns the managed instance region
func (instanceInfo) Region() string { return registration.Region() }

// dependency for metadata
var metadata metadataClient = instanceMetadata{
	Client: ec2metadata.New(session.New(aws.NewConfig().WithMaxRetries(5))),
}

type metadataClient interface {
	GetMetadata(p string) (string, error)
	Region() (string, error)
}

type instanceMetadata struct {
	Client *ec2metadata.EC2Metadata
}

// GetMetadata uses the path provided to request
func (c instanceMetadata) GetMetadata(p string) (string, error) {
	return c.Client.GetMetadata(p)
}

// Region returns the region the instance is running in.
func (c instanceMetadata) Region() (string, error) { return c.Client.Region() }
