package pkg

import (
	"fmt"
	"strings"
)

func makeImage(countIndent int, repository string, tag string,pullPolicy string, containerPort int ) string {
	var template string
	template = template + "image:\n"
	template = template + fmt.Sprintf("%srepository:%s\n", strings.Repeat("\t", countIndent), repository)
	template = template + fmt.Sprintf("%stag:%s\n", strings.Repeat("\t", countIndent), tag)
	if pullPolicy != "" {
		template = template + fmt.Sprintf("%spullPolicy:%s\n", strings.Repeat("\t", countIndent), pullPolicy)
	}
	if containerPort != 0 {
		template = template + fmt.Sprintf("%sport:%d\n", strings.Repeat("\t", countIndent), containerPort)
	}
	return template
}

func makeService(countIndent int, serviceType string, port int, targetPort int) string {
	var template string
	template = template + "service:\n"
	template = template + fmt.Sprintf("%stype:%s\n", strings.Repeat("\t", countIndent), serviceType)
	template = template + fmt.Sprintf("%sport:%d\n", strings.Repeat("\t", countIndent), port)
	template = template + fmt.Sprintf("%stargetPort:%d\n", strings.Repeat("\t", countIndent), targetPort)
	return template
}


func MakeValues(replicas int,repository string, tag string, pullPolicy string,containerPort int, serviceType string, port int, targetPort int) string {
	var template string
	countIndent := 0

	template = template + fmt.Sprintf("replicas: %d\n", replicas)
	template = template + makeImage(countIndent+1, repository, tag, pullPolicy, containerPort)
	template = template + makeService(countIndent+1, serviceType, port, targetPort)

	return template
}