package main

import "fmt"

type student struct {
	id      int
	name    string
	surname string
	age     int
	phone   string
	fix     string
}

func createStudent(id int, name string, surname string, age int, phone string, fix string) student {

	newStudent := student{

		id:      id,
		name:    name,
		surname: surname,
		age:     age,
		phone:   phone,
		fix:     fix,
	}

	return newStudent

}

func main() {

	studentRegistration := createStudent(1, "siaka", "siahoué", 25, "0759211358", "2224255252")

	fmt.Printf("le nom de l'étudiant est %v\nle prénom de l'étudiant est %v\nl'age de l'étudiant est %v\nle numéro de téléphone est %v\nle numéro de téléphone est %v", studentRegistration.id, studentRegistration.name, studentRegistration.surname, studentRegistration.phone, studentRegistration.fix)

}
