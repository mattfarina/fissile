package model

import ()

type ReleaseLicense struct {
	Fingerprint string
	Sha1        string
	Contents    []string
	Release     *Release
}

type JobProperty struct {
	Name        string
	Description string
	Default     interface{}
	Job         *Job
}

type JobTemplate struct {
	SourcePath      string
	DestinationPath string
	Job             *Job
}