package main

import "fmt" 
import "flag"
import "github.com/Masterminds/semver"
import "github.com/mitchellh/go-homedir"
import "github.com/kylelemons/go-gypsy/yaml"
import "github.com/openshift/origin"
import "github.com/ant0ine/go-json-rest/rest"


func main() {
    fmt.Println("Intro to Go!")
    api := rest.NewApi()
    homedir.Dir() 
    //fmt.Println(v.Encode()) // will output: "q=foo&all=true&page=2"
    file := flag.String("file", "config.yaml", "(Simple) YAML file to read")
    config, err := yaml.ReadFile(*file)
    v, err := semver.NewVersion("1.2.3-beta.1+build345")
    
}
