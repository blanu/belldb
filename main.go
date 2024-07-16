package main

import (
	"context"
	"fmt"
	"github.com/edgedb/edgedb-go"
	"log"
	"os"
	"strings"
)

type Module struct {
	Id           edgedb.UUID  `edgedb:"id"`
	Name         string       `edgedb:"name"`
	Dependencies []Module     `edgedb:"dependencies"`
	Definitions  []Definition `edgedb:"definitions"`
}

type Definition struct {
	Id         edgedb.UUID        `edgedb:"id"`
	Name       string             `edgedb:"name"`
	Label      string             `edgedb:"label"`
	Expression edgedb.OptionalStr `edgedb:"expression"`
}

func main() {
	ctx := context.Background()
	db, clientError := edgedb.CreateClient(ctx, edgedb.Options{})
	if clientError != nil {
		log.Fatal(clientError)
	}
	defer func() {
		_ = db.Close()
	}()

	if len(os.Args) < 2 {
		fmt.Println("Usage: belldb <subcommand>")
		os.Exit(1)
	}

	if os.Args[1] == "module" && len(os.Args) >= 3 {
		if os.Args[2] == "list" {
			if len(os.Args) == 3 {
				results := make([]Module, 0)
				queryError := db.Query(ctx, "SELECT `Module` {name}", &results)
				if queryError != nil {
					log.Fatal(queryError)
				}

				if len(results) == 0 {
					fmt.Println("No modules.")
				} else {
					fmt.Println("Modules:")
					for _, result := range results {
						fmt.Printf("- %v\n", result.Name)
					}
				}

				os.Exit(0)
			} else {
				fmt.Println("Unsupported")
			}
		} else if len(os.Args) == 3 {
			name := os.Args[2]

			fmt.Printf("Adding module %v.\n", name)
			insertError := db.Execute(ctx, "INSERT `Module` { name := <str>$0 };", name)
			if insertError != nil {
				log.Fatal(insertError)
			}

			os.Exit(0)

		} else if os.Args[3] == "needs" {
			name := os.Args[2]
			dependenciesText := os.Args[4]
			dependencies := strings.Split(dependenciesText, ",")

			for _, dependency := range dependencies {
				// FIXME - check for any size of circular dependency loop
				if dependency == name {
					fmt.Println("Circular dependencies are not allowed.")
					os.Exit(3)
				}

				fmt.Printf("Adding dependency: %v to module %v\n", dependency, name)
				insertError := db.Execute(ctx, "with dependency := ( select `Module` filter .name = <str>$0 ) UPDATE `Module` filter .name = <str>$1 set { dependencies += dependency }", dependency, name)
				if insertError != nil {
					log.Fatal(insertError)
				}
			}

			os.Exit(0)
		} else if os.Args[3] == "builtin" {
			name := os.Args[2]
			label := os.Args[4]

			fmt.Printf("Adding builtin %v to module %v\n", label, name)

			insertError := db.Execute(ctx, "Insert Definition { `module` := (select `Module` filter .name = <str>$0), label := <str>$1 }", name, label)
			if insertError != nil {
				log.Fatal(insertError)
			}

			os.Exit(0)
		} else if os.Args[3][len(os.Args[3])-1:] == ":" {
			name := os.Args[2]
			label := os.Args[3][0 : len(os.Args[3])-1]
			expression := os.Args[4]

			fmt.Printf("Setting %v.%v to \"%v\"\n", name, label, expression)

			insertError := db.Execute(ctx, "Insert Definition { `module` := (select `Module` filter .name = <str>$0), label := <str>$1, expression := <str>$2 }", name, label, expression)
			if insertError != nil {
				log.Fatal(insertError)
			}

			os.Exit(0)
		} else if os.Args[3] == "list" && len(os.Args) >= 5 {
			switch os.Args[4] {
			case "dependencies":
				name := os.Args[2]

				var result Module
				queryError := db.QuerySingle(ctx, "SELECT `Module` {**} filter .name = <str>$0", &result, name)
				if queryError != nil {
					fmt.Println("No such module.")
					os.Exit(2)
				}

				if len(result.Dependencies) == 0 {
					fmt.Println("No dependencies.")
					os.Exit(0)
				}

				fmt.Printf("module %v needs ", name)

				for index, dependency := range result.Dependencies {
					fmt.Printf("%v", dependency.Name)

					if index < len(result.Dependencies)-1 {
						fmt.Printf(",")
					} else {
						fmt.Printf("\n")
					}
				}

				os.Exit(0)
			case "definitions":
				name := os.Args[2]

				results := make([]Definition, 0)
				queryError := db.Query(ctx, "SELECT Definition {*} filter .`module` = ( select `Module` filter .name = <str>$0 )", &results, name)
				if queryError != nil {
					log.Fatal(queryError)
				}

				if len(results) == 0 {
					fmt.Println("No definitions.")
					os.Exit(0)
				}

				foundOne := false
				for _, definition := range results {
					expression, hasExpression := definition.Expression.Get()
					if hasExpression {
						fmt.Printf("%v: %v\n", definition.Label, expression)
						foundOne = true
					}
				}

				if !foundOne {
					fmt.Println("No definitions.")
					os.Exit(0)
				}

				os.Exit(0)
			case "builtins":
				name := os.Args[2]

				results := make([]Definition, 0)
				queryError := db.Query(ctx, "SELECT Definition {*} filter .`module` = ( select `Module` filter .name = <str>$0 )", &results, name)
				if queryError != nil {
					log.Fatal(queryError)
				}

				if len(results) == 0 {
					fmt.Println("No builtins.")
					os.Exit(0)
				}

				foundOne := false
				for _, definition := range results {
					_, hasExpression := definition.Expression.Get()
					if !hasExpression {
						fmt.Printf("builtin %v\n", definition.Label)
						foundOne = true
					}
				}

				if !foundOne {
					fmt.Println("No builtins.")
					os.Exit(0)
				}

				os.Exit(0)
			}
		}
	}

	fmt.Println("Usage: ")
	fmt.Println("belldb module list\t\t\t\t\tList existing modules")
	fmt.Println("belldb module <modulename>\t\t\t\t\tCreate a new module")
	fmt.Println("belldb module <modulename> needs <modulename>,<modulename>,...\tAdd dependencies to a module")
	fmt.Println("belldb module <modulename> builtin <label>\t\t\tAdd a builtin to a module")
	fmt.Println("belldb module <modulename> <label>: <expression>\t\tAdd a definition to a module")
	fmt.Println("belldb module <modulename> list dependencies\t\t\t\t\tList module depedencies")
	fmt.Println("belldb module <modulename> list definitions\t\t\t\t\tList module definitions")
	fmt.Println("belldb module <modulename> list builtins\t\t\t\t\tList module builtins")
	os.Exit(1)
}
