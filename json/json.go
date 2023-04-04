package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json : "name" ` // déférencier le type de la structure avec l'attribut du fichier json
	Age   int    `json : "age"`   // on peut dire qu'on fait une assignation de champ selon mon raisonnement
	Email string `json : "email"`
	Phone string `json : "phone"`
}

func main() {
	// exemple de récupération e fichier json
	jsonfromapi := ` 
[
{
	"name" : "yao",
	"age" : 23,
	"email" : "yao@gmail.com",
	"phone" : "065463738"
},
{
	"name" : "mian",
	"age" : 28,
	"email" : "mian@gmail.com",
	"phone" : "095463738"
}
]
	`
	// passer du format json à un slice en définissant un slice de User
	var users []User
	//récupréré les informations via le fichier json
	err := json.Unmarshal([]byte(jsonfromapi), &users)

	if err != nil {
		fmt.Println("erreur Unmarshal json :", err)
	}
	// lecture dans notre terminal
	fmt.Printf("json : %v\n", users)

	//---------------------- // encoder

	var Userstruct []User // assignatiion d'un slice de type *User* à ma variable struct

	var user_one User // user_one est de type user

	user_one.Name = "kouakou"

	user_one.Age = 28

	user_one.Email = "kouakou@gmail.com"

	user_one.Phone = "0835626282"

	var user_two User // user_one est de type user

	user_two.Name = "anou"

	user_two.Age = 48

	user_two.Email = "anou@gmail.com"

	user_two.Phone = "567399393"

	Userstruct = append(Userstruct, user_one, user_two) // ajouter les informations des users a notre slice

	// jsonfromslice, err := json.Marshal(Userstruct) // encoder notre slice en format json sans indentation

	jsonfromslice, err := json.MarshalIndent(Userstruct, "", "  ") // encoder avec MarshatIndent permet de mieux formater notre fichier json

	if err != nil {
		fmt.Println("erreur Unmarshal json :", err)
	}
	fmt.Println(string(jsonfromslice)) // afichage des données en json avec clé et valeur
}
