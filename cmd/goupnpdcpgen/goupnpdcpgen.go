// Command to generate DCP package source from the XML specification.
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

func main() {
	var (
		dcpName  = flag.String("dcp_name", "", "Name of the DCP to generate.")
		specsDir = flag.String("specs_dir", ".", "Path to the specification storage directory. "+
			"This is used to find (and download if not present) the specification ZIP files.")
		useGofmt = flag.Bool("gofmt", true, "Pass the generated code through gofmt. "+
			"Disable this if debugging code generation and needing to see the generated code "+
			"prior to being passed through gofmt.")
	)
	flag.Parse()

	if err := run(*dcpName, *specsDir, *useGofmt); err != nil {
		log.Fatal(err)
	}
}

func run(dcpName, specsDir string, useGofmt bool) error {
	if err := os.MkdirAll(specsDir, os.ModePerm); err != nil {
		return errors.Wrapf(err, "could not create specs-dir %q: %v\n", specsDir)
	}

	for _, d := range dcpMetadata {
		if d.Name != dcpName {
			continue
		}
		specFilename := filepath.Join(specsDir, d.Name+".zip")
		err := acquireFile(specFilename, d.XMLSpecURL)
		if err != nil {
			return errors.Wrapf(err, "could not acquire spec for %s", d.Name)
		}
		dcp := newDCP(d)
		if err := dcp.processZipFile(specFilename); err != nil {
			return errors.Wrapf(err, "error processing spec for %s in file %q", d.Name, specFilename)
		}
		for i, hack := range d.Hacks {
			if err := hack(dcp); err != nil {
				return errors.Wrapf(err, "error with Hack[%d] for %s", i, d.Name)
			}
		}
		if err := dcp.writeCode(d.Name+".go", useGofmt); err != nil {
			return errors.Wrapf(err, "error writing package %q", dcp.Metadata.Name)
		}

		return nil
	}

	return errors.Errorf("could not find DCP with name %q", dcpName)
}

var (
	deviceURNPrefix  = regexp.MustCompile(`^urn\:schemas\-[a-z]+(\-[a-z]+)*\:device\:`)
	serviceURNPrefix = regexp.MustCompile(`^urn\:schemas\-[a-z]+(\-[a-z]+)*\:service\:`)
)
