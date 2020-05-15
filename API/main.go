package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
)

//Structs

type apiCocktail struct {
	IDDrink                     string    `json:"idDrink"`
	StrDrink                    string    `json:"strDrink"`
	StrDrinkAlternate           string    `json:"strDrinkAlternate,omitempty"`
	StrDrinkES                  string    `json:"strDrinkES,omitempty"`
	StrDrinkDE                  string    `json:"strDrinkDE,omitempty"`
	StrDrinkFR                  string    `json:"strDrinkFR,omitempty"`
	StrDrinkZHHANS              string    `json:"strDrinkZH-HANS,omitempty"`
	StrDrinkZHHANT              string    `json:"strDrinkZH-HANT,omitempty"`
	StrTags                     string    `json:"strTags,omitempty"`
	StrVideo                    string    `json:"strVideo,omitempty"`
	StrCategory                 string    `json:"strCategory,omitempty"`
	StrIBA                      string    `json:"strIBA,omitempty"`
	StrAlcoholic                string    `json:"strAlcoholic,omitempty"`
	StrGlass                    string    `json:"StrGlass,omitempty"`
	StrInstructions             string    `json:"StrInstructions"`
	StrInstructionsES           string    `json:"StrInstructionsES"`
	StrInstructionsDE           string    `json:"StrInstructionsDE"`
	StrInstructionsFR           string    `json:"StrInstructionsFR"`
	StrInstructionsZHHANS       string    `json:"strInstructionsZH-HANS"`
	StrInstructionsZHHANT       string    `json:"strInstructionsZH-HANT"`
	StrDrinkThumb               string    `json:"strDrinkThumb"`
	StrIngredient1              string    `json:"strIngredient1"`
	StrIngredient2              string    `json:"strIngredient2"`
	StrIngredient3              string    `json:"strIngredient3"`
	StrIngredient4              string    `json:"strIngredient4"`
	StrIngredient5              string    `json:"strIngredient5"`
	StrIngredient6              string    `json:"strIngredient6"`
	StrIngredient7              string    `json:"strIngredient7"`
	StrIngredient8              string    `json:"strIngredient8"`
	StrIngredient9              string    `json:"strIngredient9"`
	StrIngredient10             string    `json:"strIngredient10"`
	StrIngredient11             string    `json:"strIngredient11"`
	StrIngredient12             string    `json:"strIngredient12"`
	StrIngredient13             string    `json:"strIngredient13"`
	StrIngredient14             string    `json:"strIngredient14"`
	StrIngredient15             string    `json:"strIngredient15"`
	StrMeasure1                 string    `json:"strMeasure1"`
	StrMeasure2                 string    `json:"strMeasure2"`
	StrMeasure3                 string    `json:"strMeasure3"`
	StrMeasure4                 string    `json:"strMeasure4"`
	StrMeasure5                 string    `json:"strMeasure5"`
	StrMeasure6                 string    `json:"strMeasure6"`
	StrMeasure7                 string    `json:"strMeasure7"`
	StrMeasure8                 string    `json:"strMeasure8"`
	StrMeasure9                 string    `json:"strMeasure9"`
	StrMeasure10                string    `json:"strMeasure10"`
	StrMeasure11                string    `json:"strMeasure11"`
	StrMeasure12                string    `json:"strMeasure12"`
	StrMeasure13                string    `json:"strMeasure13"`
	StrMeasure14                string    `json:"strMeasure14"`
	StrMeasure15                string    `json:"strMeasure15"`
	StrCreativeCommonsConfirmed bool      `json:"strCreativeCommonsConfirmed"`
	DateModified                time.Time `json:"dateModified"`
}

type prefs struct {
	APIKey      string `json:"key"`
	WorkerCount int    `json:"worker_count"`
}

type cocktail struct {
	IDDrink      string       `json:"id"`
	Name         string       `json:"name"`
	Alcoholic    string       `json:"alcoholic"`
	Glass        string       `json:"glass"`
	Instructions string       `json:"instructions"`
	Thumbnail    string       `json:"thumbnail"`
	Ingredients  []ingredient `json:"ingredients"`
}

type ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"measure"`
}

type basicCocktail struct {
	StrDrink      string `json:"strDrink"`
	StrDrinkThumb string `json:"strDrinkThumb"`
	IDDrink       string `json:"idDrink"`
}

type drinkRes struct {
	Drinks []apiCocktail `json:"drinks"`
}

type drinkListRes struct {
	Drinks []basicCocktail `json:"drinks"`
}

type postData struct {
	Ingredients []string `json:"ingredients"`
}

type ingData struct {
	Ingredients []ingString `json:"drinks"`
}

type ingString struct {
	Ingredient string `json:"strIngredient1"`
}

//Program
var p prefs

func main() {
	pf, _ := ioutil.ReadFile("prefs.json")
	_ = json.Unmarshal(pf, &p)

	//getIngredientMap()

	m := martini.Classic()

	// CORS for * origins, allowing:
	// - POST methods
	// - Origin header
	// - Credentials share
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	m.Post("/", func(res http.ResponseWriter, req *http.Request) []byte {
		data, _ := ioutil.ReadAll(req.Body)
		var p postData
		json.Unmarshal(data, &p)
		if len(p.Ingredients) > 0 {
			results := make(chan []cocktail)
			go getCocktails(p.Ingredients, results)
			ds := <-results
			ret, _ := json.Marshal(ds)
			return ret
		} else {
			return nil
		}
	})

	m.Run()
}

func getIngredientMap() map[string]bool {
	ingredients := make(map[string]bool)

	req, err := http.Get("https://www.thecocktaildb.com/api/json/v2/" + p.APIKey + "/list.php?i=list")
	if err != nil {
		fmt.Println(err)
		return ingredients
	} else {
		data, _ := ioutil.ReadAll(req.Body)
		var d ingData
		json.Unmarshal(data, &d)
		for _, i := range d.Ingredients {
			i := i.Ingredient
			ingredients[strings.ToUpper(i)] = false
		}
	}

	return ingredients
}

func getCocktails(ing []string, r chan<- []cocktail) {
	counter := 0
	var cocktails []cocktail

	drinks := make(map[string]basicCocktail)
	ingredients := getIngredientMap()

	for _, i := range ing {
		ingredients[strings.ToUpper(i)] = true
		i = strings.Replace(i, " ", "%20", -1)
		resp, err := http.Get("https://www.thecocktaildb.com/api/json/v2/" + p.APIKey + "/filter.php?i=" + i)
		if err != nil {
			fmt.Println(err)
		} else {
			var dlr drinkListRes
			data, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(data, &dlr)
			for _, d := range dlr.Drinks {
				drinks[d.IDDrink] = d
			}
		}
	}

	size := len(drinks)
	jobs := make(chan string, size)
	results := make(chan cocktail, size)

	jobEnd := size - 1

	workers := (len(drinks) / p.WorkerCount)

	if workers > 20 {
		workers = 20
	}

	for i := 0; i < workers; i++ {
		go worker(jobs, results, jobEnd, &counter, ingredients)
	}

	for id := range drinks {
		jobs <- id
	}
	close(jobs)

	for c := range results {
		counter++
		if c.IDDrink != "" {
			cocktails = append(cocktails, c)
		}
	}

	r <- cocktails
}

func worker(j <-chan string, r chan<- cocktail, end int, curr *int, ingredients map[string]bool) {
	for c := range j {
		drink := getCocktail(c, ingredients)
		r <- drink
	}
	if *curr >= end {
		close(r)
	}
}

//DEPRICATED
//This was used to look through an array but due to hashmap's constant lookup times they serve as a better solution
func contains(a []string, x string) bool {
	for _, n := range a {
		if strings.ToUpper(x) == strings.ToUpper(n) {
			return true
		}
	}
	return false
}

func newIngredient(ing string, msr string) ingredient {
	var ingr ingredient
	if msr != "" {
		ingr = ingredient{ing, msr}
	} else {
		ingr = ingredient{ing, "To taste"}
	}
	return ingr
}

func getCocktail(id string, ingredients map[string]bool) cocktail {

	resp, err := http.Get("https://www.thecocktaildb.com/api/json/v2/" + p.APIKey + "/lookup.php?i=" + id)

	var c cocktail

	if err != nil {
		fmt.Printf("error %s\n", err)
		return c
	}

	data, _ := ioutil.ReadAll(resp.Body)
	var d drinkRes
	json.Unmarshal([]byte(data), &d)

	var df apiCocktail = d.Drinks[0]
	var i []ingredient

	if df.StrIngredient1 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient1)] {
			i = append(i, newIngredient(df.StrIngredient1, df.StrMeasure1))
		} else {
			return c
		}
	}
	if df.StrIngredient2 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient2)] {
			i = append(i, newIngredient(df.StrIngredient2, df.StrMeasure2))
		} else {
			return c
		}
	}
	if df.StrIngredient3 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient3)] {
			i = append(i, newIngredient(df.StrIngredient3, df.StrMeasure3))
		} else {
			return c
		}
	}
	if df.StrIngredient4 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient4)] {
			i = append(i, newIngredient(df.StrIngredient4, df.StrMeasure4))
		} else {
			return c
		}
	}
	if df.StrIngredient5 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient5)] {
			i = append(i, newIngredient(df.StrIngredient5, df.StrMeasure5))
		} else {
			return c
		}
	}
	if df.StrIngredient6 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient6)] {
			i = append(i, newIngredient(df.StrIngredient6, df.StrMeasure6))
		} else {
			return c
		}
	}
	if df.StrIngredient7 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient7)] {
			i = append(i, newIngredient(df.StrIngredient7, df.StrMeasure7))
		} else {
			return c
		}
	}
	if df.StrIngredient8 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient8)] {
			i = append(i, newIngredient(df.StrIngredient8, df.StrMeasure8))
		} else {
			return c
		}
	}
	if df.StrIngredient9 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient9)] {
			i = append(i, newIngredient(df.StrIngredient9, df.StrMeasure9))
		} else {
			return c
		}
	}
	if df.StrIngredient10 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient10)] {
			i = append(i, newIngredient(df.StrIngredient10, df.StrMeasure10))
		} else {
			return c
		}
	}
	if df.StrIngredient11 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient11)] {
			i = append(i, newIngredient(df.StrIngredient11, df.StrMeasure11))
		} else {
			return c
		}
	}
	if df.StrIngredient12 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient12)] {
			i = append(i, newIngredient(df.StrIngredient12, df.StrMeasure12))
		} else {
			return c
		}
	}
	if df.StrIngredient13 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient13)] {
			i = append(i, newIngredient(df.StrIngredient13, df.StrMeasure13))
		} else {
			return c
		}
	}
	if df.StrIngredient14 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient14)] {
			i = append(i, newIngredient(df.StrIngredient14, df.StrMeasure14))
		} else {
			return c
		}
	}
	if df.StrIngredient15 != "" {
		if ingredients[strings.ToUpper(df.StrIngredient15)] {
			i = append(i, newIngredient(df.StrIngredient15, df.StrMeasure15))
		} else {
			return c
		}
	}
	fmt.Println("added:", df.StrDrink, "👍")
	return cocktail{df.IDDrink, df.StrDrink, df.StrAlcoholic, df.StrGlass, df.StrInstructions, df.StrDrinkThumb, i}
}
