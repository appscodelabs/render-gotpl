/*
Copyright AppsCode Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"github.com/hashicorp/go-getter"
	flag "github.com/spf13/pflag"
	"sigs.k8s.io/yaml"
)

var (
	tplFile  = flag.String("template", "", "Path to Go template file (local file or url is accepted)")
	dataFile = flag.String("data", "", "Path to data file in JSON or YAML format (local file or url is accepted)")
)

func main() {
	flag.Parse()

	localTplFile := "/tmp/template.txt"
	err := getter.GetFile(localTplFile, *tplFile)
	if err != nil {
		panic(err)
	}

	localDataFile := "/tmp/template-data.txt"
	err = getter.GetFile(localDataFile, *dataFile)
	if err != nil {
		panic(err)
	}

	d, err := ioutil.ReadFile(localDataFile)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(d, &data)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("tpl").ParseFiles(localTplFile))
	err = tpl.Execute(os.Stdout, &data)
	if err != nil {
		panic(err)
	}
}
