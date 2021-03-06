package model

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pivotal-golang/archiver/extractor"
	"gopkg.in/yaml.v2"
)

// Job represents a BOSH job
type Job struct {
	Name        string
	Description string
	Templates   []*JobTemplate
	Packages    Packages
	Path        string
	Fingerprint string
	SHA1        string
	Properties  []*JobProperty
	Version     string
	Release     *Release

	jobReleaseInfo map[interface{}]interface{}
	jobSpec        map[interface{}]interface{}
}

// Jobs is an array of Job*
type Jobs []*Job

func newJob(release *Release, jobReleaseInfo map[interface{}]interface{}) (*Job, error) {
	job := &Job{
		Release: release,

		jobReleaseInfo: jobReleaseInfo,
	}

	if err := job.loadJobInfo(); err != nil {
		return nil, err
	}

	if err := job.loadJobSpec(); err != nil {
		return nil, err
	}

	return job, nil
}

func (j *Job) getProperty(name string) (*JobProperty, error) {
	for _, property := range j.Properties {
		if property.Name == name {
			return property, nil
		}
	}

	return nil, fmt.Errorf("Property %s not found in job %s", name, j.Name)
}

// ValidateSHA1 validates that the SHA1 of the actual job archive is the same
// as the one from the release manifest
func (j *Job) ValidateSHA1() error {
	file, err := os.Open(j.Path)
	if err != nil {
		return fmt.Errorf("Error opening the job archive %s for sha1 calculation", j.Path)
	}

	defer file.Close()

	h := sha1.New()

	_, err = io.Copy(h, file)
	if err != nil {
		return fmt.Errorf("Error copying job archive %s for sha1 calculation", j.Path)
	}

	computedSha1 := fmt.Sprintf("%x", h.Sum(nil))

	if computedSha1 != j.SHA1 {
		return fmt.Errorf("Computed sha1 (%s) is different than manifest sha1 (%s) for job archive %s", computedSha1, j.SHA1, j.Path)
	}

	return nil
}

// Extract will extract the contents of the job archive to destination
// It creates a directory with the name of the job
// Returns the full path of the extracted archive
func (j *Job) Extract(destination string) (string, error) {
	targetDir := filepath.Join(destination, j.Name)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", err
	}

	if err := extractor.NewTgz().Extract(j.Path, targetDir); err != nil {
		return "", err
	}

	return targetDir, nil
}

func (j *Job) loadJobInfo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error trying to load job information: %s", r)
		}
	}()

	j.Name = j.jobReleaseInfo["name"].(string)
	j.Version = j.jobReleaseInfo["version"].(string)
	j.Fingerprint = j.jobReleaseInfo["fingerprint"].(string)
	j.SHA1 = j.jobReleaseInfo["sha1"].(string)
	j.Path = j.jobArchivePath()

	return nil
}

func (j *Job) loadJobSpec() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error trying to load job spec: %s", r)
		}
	}()

	tempJobDir, err := ioutil.TempDir("", "fissile-job-dir")
	defer func() {
		if cleanupErr := os.RemoveAll(tempJobDir); cleanupErr != nil && err != nil {
			err = fmt.Errorf("Error loading job spec: %v,  cleanup error: %v", err, cleanupErr)
		} else if cleanupErr != nil {
			err = fmt.Errorf("Error cleaning up after load job spec: %v", cleanupErr)
		}
	}()
	if err != nil {
		return err
	}

	jobDir, err := j.Extract(tempJobDir)
	if err != nil {
		return fmt.Errorf("Error extracting archive (%s) for job %s: %s", j.Path, j.Name, err.Error())
	}

	specFile := filepath.Join(jobDir, "job.MF")

	specContents, err := ioutil.ReadFile(specFile)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal([]byte(specContents), &j.jobSpec); err != nil {
		return err
	}

	if j.jobSpec["description"] != nil {
		j.Description = j.jobSpec["description"].(string)
	}

	if j.jobSpec["packages"] != nil {
		for _, pkg := range j.jobSpec["packages"].([]interface{}) {
			pkgName := pkg.(string)
			dependency, err := j.Release.LookupPackage(pkgName)
			if err != nil {
				return fmt.Errorf("Cannot find dependency for job %s: %v", j.Name, err.Error())
			}

			j.Packages = append(j.Packages, dependency)
		}
	}

	if j.jobSpec["templates"] != nil {
		for source, destination := range j.jobSpec["templates"].(map[interface{}]interface{}) {
			templateFile := filepath.Join(jobDir, "templates", source.(string))

			templateContent, err := ioutil.ReadFile(templateFile)
			if err != nil {
				return err
			}

			template := &JobTemplate{
				SourcePath:      source.(string),
				DestinationPath: destination.(string),
				Job:             j,
				Content:         string(templateContent),
			}

			j.Templates = append(j.Templates, template)
		}
	}

	if j.jobSpec["properties"] != nil {
		for propertyName, propertyDefinition := range j.jobSpec["properties"].(map[interface{}]interface{}) {

			property := &JobProperty{
				Name: propertyName.(string),
				Job:  j,
			}

			if propertyDefinition != nil {
				propertyDefinitionMap := propertyDefinition.(map[interface{}]interface{})

				if propertyDefinitionMap["description"] != nil {
					property.Description = propertyDefinitionMap["description"].(string)
				}
				if propertyDefinitionMap["default"] != nil {
					property.Default = propertyDefinitionMap["default"]
				}
			}

			j.Properties = append(j.Properties, property)
		}
	}

	return nil
}

// Len implements the Len function to satisfy sort.Interface
func (slice Jobs) Len() int {
	return len(slice)
}

// Less implements the Less function to satisfy sort.Interface
func (slice Jobs) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

// Swap implements the Swap function to satisfy sort.Interface
func (slice Jobs) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (j *Job) jobArchivePath() string {
	return filepath.Join(j.Release.DevBOSHCacheDir, j.SHA1)
}
