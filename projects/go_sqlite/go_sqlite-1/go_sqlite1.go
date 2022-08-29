package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dixonwille/wmenu/v5"
	_ "github.com/mattn/go-sqlite3"
)
func main()  {

	db, err := sql.Open("sqlite3", "./workshop.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println()
		menu := wmenu.NewMenu("What would you like to do?")
		menu.Action(func(o []wmenu.Opt) error {handleFunc(db, o); return nil})
		
		menu.Option("Add a new recipe", 0, true, nil)
		menu.Option("Get a recipe", 1, false, nil)
		menu.Option("Edit a recipe", 2, false, nil)
		menu.Option("Delete a recipe", 3, false, nil)
		menu.Option("Exit", 4, false, nil)
		err = menu.Run()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func handleFunc(db *sql.DB,o []wmenu.Opt) {
	switch o[0].Value {
	case 0:
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter a name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSuffix(name,"\n")

		fmt.Println("Enter ingredients (delimeter by ','): ")
		ingredient, _ := reader.ReadString('\n')
		ingredient = strings.TrimSuffix(ingredient, "\n")
		ingredients := strings.Split(ingredient, ",")
		isHalal, _ := yesNo("isHalal ? ", reader)
		isVegetarian, _ := yesNo("isVegetarian ? ", reader)

		fmt.Println("Enter a description: ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSuffix(description,"\n")

		fmt.Println("Enter a rating: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input,"\n")
		input = strings.TrimSpace(input)
		rating, _ := strconv.ParseFloat(input, 64)

		recipe := recipe{
			id: 0,
			name: name,
			ingredients: strings.Join(ingredients, ","),
			isHalal: isHalal,
			isVegetarian: isVegetarian,
			description: description,
			rating: rating,
		}

		add(db, recipe)
	case 1:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a name to search for : ")
		searchString, _ := reader.ReadString('\n')
    	searchString = strings.TrimSuffix(searchString, "\n")
		recipes := search(db, searchString)


		for _, recipe := range recipes {
			fmt.Printf("\n----\nID: %d\nName: %s\ningredients: %s\nisHalal: %t\nisVegetarian: %t\ndescription: %s\nrating: %v\n",recipe.id, recipe.name, recipe.ingredients, recipe.isHalal, recipe.isVegetarian, recipe.description, recipe.rating)
		}

		fmt.Println()
		fmt.Println("---------------------------------")
		fmt.Printf("Found %v results\n", len(recipes))
	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an id to edit: ")
		id, _ := reader.ReadString('\n')
		currentRecipe := getById(db, id)

		fmt.Printf("name (currently `%s`): ", currentRecipe.name)
		name, _ := reader.ReadString('\n')
		if name != "\n" {
			currentRecipe.name = strings.TrimSuffix(name,"\n")
		}
		
		fmt.Printf("ingredients (currently `%s` \n): ", currentRecipe.ingredients)
		fmt.Println("Enter ingredients (delimeter by ','): ")
		ingredient, _ := reader.ReadString('\n')
		ingredient = strings.TrimSuffix(ingredient, "\n")
		ingredientList := strings.Split(ingredient, ",")
		ingredients := strings.Join(ingredientList, ",")
		if ingredient != "\n" {
			currentRecipe.ingredients = ingredients
		}

		fmt.Printf("isHalal (currently `%v`): ", currentRecipe.isHalal)
		isHalal, _ := yesNo("isHalal ? ", reader)
		currentRecipe.isHalal = isHalal

		fmt.Printf("isVegetarian (currently `%v`): ", currentRecipe.isVegetarian)
		isVegetarian, _ := yesNo("isHalal ? ", reader)
		currentRecipe.isVegetarian = isVegetarian

		fmt.Printf("description (currently `%s`): ", currentRecipe.description)
		description, _ := reader.ReadString('\n')
		if description != "\n" {
			currentRecipe.description = strings.TrimSuffix(description,"\n")
		}
		fmt.Printf("rating (currently `%v`): ", currentRecipe.rating)
		input, _ := reader.ReadString('\n')
		if input != "\n" {
			input = strings.TrimSuffix(input,"\n")
			input = strings.TrimSpace(input)
			rating, _ := strconv.ParseFloat(input, 64)
			currentRecipe.rating = rating
		}

		affected := update(db, currentRecipe)

		if affected == 1 {
			fmt.Println("one row affected (updated)")
		}
	case 3:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an id to delete: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSuffix(id, "\n")

		affected := deleteById(db, id)
		if affected == 1 {
			fmt.Println("one row affected (deleted)")
		}
	case 4:
		fmt.Println("Good Bye!")
		os.Exit(1)
	}
}

func yesNo(s string, r *bufio.Reader) (bool, error) {
	fmt.Printf("%s [y/n] ", s)
	input,err := r.ReadString('\n')
	input = strings.TrimSuffix(input,"\n")
	input = strings.TrimSpace(input)
	if err != nil {
		log.Fatal(err)
	}
	b, err := parse(input)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return b, nil
}

func parse(s string) (bool, error){
	switch strings.ToLower(s) {
	case "y":
		return true, nil
	case "n":
		return false, nil
	}
	return false, errors.New("invalid keypress, it does not match a boolean")
}

func deleteById(db *sql.DB, id string) int64{
	stmt, _ := db.Prepare("DELETE FROM recipes WHERE id=?;")
	res, err := stmt.Exec(id) 
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return affected
}

func update(db *sql.DB, recipe recipe) int {
	stmt, _ := db.Prepare("UPDATE recipes SET name=?, ingredients=json_array(?), isHalal=?, isVegetarian=?, description=?, rating=? WHERE id = ?;")
	res, err := stmt.Exec(recipe.name, recipe.ingredients, recipe.isHalal, recipe.isVegetarian, recipe.description, recipe.rating, recipe.id) 
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return int(affected)
}

func getById(db *sql.DB, id string) recipe {
	rows, err := db.Query("SELECT id, name, ingredients, isHalal, isVegetarian, description, rating FROM recipes WHERE id ='" + id + "'")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	recipe := recipe{}
	for rows.Next() {
		rows.Scan(&recipe.id, &recipe.name, &recipe.ingredients, &recipe.isHalal, &recipe.isVegetarian, &recipe.description, &recipe.rating)
	}
	return recipe
}

func search(db *sql.DB, searchString string) []recipe {
	
	rows, err := db.Query("SELECT id, name, ingredients, isHalal, isVegetarian, description, rating FROM recipes WHERE name like '%" + searchString + "%' OR description like '%" + searchString + "%';")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	recipes := make([]recipe, 0)

	for rows.Next() {
		recipe := recipe{}
		err = rows.Scan(&recipe.id, &recipe.name, &recipe.ingredients, &recipe.isHalal, &recipe.isVegetarian, &recipe.description, &recipe.rating)
		if err != nil {
			log.Fatal(err)
		}

		recipes = append(recipes, recipe)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return recipes
}

func add(db *sql.DB, recipe recipe) {
	stmt, _ := db.Prepare("INSERT INTO recipes (id, name, ingredients, isHalal, isVegetarian, description, rating) VALUES (?, ?, json_array(?), ?, ?, ?, ?)")
	ingredients := strings.Split(recipe.ingredients, ",")
	ingredient := strings.Join(ingredients, ",")
	stmt.Exec(nil, recipe.name, ingredient, recipe.isHalal, recipe.isVegetarian, recipe.description, recipe.rating) 
	defer stmt.Close()

	fmt.Println(SQLQueryDebugString("INSERT INTO recipes (id, name, ingredients, isHalal, isVegetarian, description, rating) VALUES (?, ?, ?, ?, ?, ?, ?)", nil, recipe.name, ingredient, recipe.isHalal, recipe.isVegetarian, recipe.description, recipe.rating))
	fmt.Printf("Added %v  \n", recipe.name)
}
type recipe struct {
	id int
	name, ingredients, description string
	isHalal, isVegetarian bool
	rating float64
}

// SQLQueryDebugString formats an sql query inlining its arguments
// The purpose is debug only - do not send this to the database!
// Sending this to the DB is unsafe and un-performant.
func SQLQueryDebugString(query string, args ...interface{}) string {
	var buffer bytes.Buffer
	nArgs := len(args)
	// Break the string by question marks, iterate over its parts and for each 
	// question mark - append an argument and format the argument according to
	// it's type, taking into consideration NULL values and quoting strings.
	for i, part := range strings.Split(query, "?") {
		buffer.WriteString(part)
		if i < nArgs {
			switch a := args[i].(type) {
			case int64:
				buffer.WriteString(fmt.Sprintf("%d", a))
			case float64:
				buffer.WriteString(fmt.Sprintf("%v", a))
			case bool:
				buffer.WriteString(fmt.Sprintf("%t", a))
			case sql.NullBool:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%t", a.Bool))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullInt64:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%d", a.Int64))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullString:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%q", a.String))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullFloat64:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%v", a.Float64))
				} else {
					buffer.WriteString("NULL")
				}
			default:
				buffer.WriteString(fmt.Sprintf("%q", a))
			}
		}
	}
	return buffer.String()
}
