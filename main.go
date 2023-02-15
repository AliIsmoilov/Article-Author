package main

import (
	"fmt"
	"log"
)

type Author struct{
	id int
	name string
	age int
}

type Article struct{

	Id int
	Title string
	Description string
	Author_id int
}

type InMemoryAuthors struct{
	Authors []Author
}



func (INM *InMemoryAuthors) CreateAuthor(input_id int, input_name string, input_age int){

	err := false

	for _, v := range INM.Authors{
		
		if input_id == v.id{
			err = true

			fmt.Println("Such id is already exists")
			return
		}
	}

	if !err{

		newAuthor := Author{

			id: input_id,
			name: input_name,
			age: input_age,
		}

		INM.Authors = append(INM.Authors, newAuthor)
		fmt.Println("New author has been created")
	}
}

func (INM *InMemoryAuthors) UpdateAuthor(newauthor Author){

	err := true

	for i, v := range INM.Authors{

		if v.id == newauthor.id{

			INM.Authors[i] = newauthor
			err = false
		}
	}

	if err{
		fmt.Println("Not found such user!!!")
	}
}

func (INM *InMemoryAuthors) DeleteAuthor(id int){

	err := true

	for _, v := range INM.Authors{

		if v.id == id{
			err = false
		}
	}

	if !err{
		INM.Authors = append(INM.Authors[:id-1], INM.Authors[id:]...)
		fmt.Println("Author has been deleted")
	} else {
		fmt.Println("Not found such user for deleting!!!")
	}
}

type InMemoryArticles struct{
	Articles []Article
}

func (INM *InMemoryArticles) CreateArticle(newArticle Article){


	INM.Articles = append(INM.Articles, newArticle)
}

func (INM *InMemoryArticles) GetAll() ([]Article, bool){

	err := false

	if len(INM.Articles) == 0{
		err = true
		return nil, err
	}

	return INM.Articles, err

}

func (INM *InMemoryArticles) UpdateArticle(id int, Newarticle Article) (bool){
	
	exist := false

	for ind, article := range INM.Articles{

		if article.Id == id{
			
			exist = true

			INM.Articles[ind] = Newarticle

			return exist
		}
	}

	
	return exist
} 

func (INM *InMemoryArticles) Delete(id int) bool{

	exists := false

	for _, article := range INM.Articles{
		
		if article.Id == id{

			exists = true
			INM.Articles = append(INM.Articles[:id], INM.Articles[id+1:]...)
			return exists
		}
	}

	return exists
}

func (INM *InMemoryArticles) GetById(id int) (bool, Article){

	exists := false
	var found_article Article

	for _, article := range INM.Articles{
		if article.Id == id{
			
			exists = true
			found_article = article
			
			return exists, found_article

		}
	}
	return exists, found_article
}

func main(){

	InmAuthors := InMemoryAuthors{}
	InmArticles := InMemoryArticles{}

	InmAuthors.CreateAuthor(1, "Ali", 21)
	fmt.Println(InmAuthors)

	// newAuthor := Author{
	// 	id: 2,
	// 	name: "Aziz",
	// 	age: 22,
	// }

	// InmAuthors.UpdateAuthor(newAuthor)
	// fmt.Println(InmAuthors)

	// Authors.DeleteAuthor(1)
	// fmt.Println(Authors)

	newArticle := Article{
		Id: 1,
		Title: "New York Times",
		Description: "Russia attaked to Ukraine",
		Author_id: InmAuthors.Authors[0].id,
	}

	Articles, err := InmArticles.GetAll()

	if !err {
		log.Println(err)
	}

	fmt.Println(Articles)

	InmArticles.CreateArticle(newArticle)

	Articles, err = InmArticles.GetAll()

	if !err {
		log.Println(err)
	}

	fmt.Println(Articles)

	err, article := InmArticles.GetById(2)

	fmt.Println(article)
}