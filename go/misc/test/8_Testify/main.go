package main

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	var uniques []string

	check := make(map[string]struct{})
	for _, developer := range developers {
		check[developer.Name] = struct{}{}
	}

	for name := range check {
		uniques = append(uniques, name)
	}

	return uniques
}
