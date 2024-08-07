package lists

import "fmt"


type Product struct {

    Title string
    Id string
    price float64
}

func main() {

    hobbies := [3]string{"Football", "F1", "Cricket"}
    fmt.Println(hobbies)

    fmt.Printf("The first element is : %v \n", hobbies[0])
    fmt.Printf("The second and third element %v \n", hobbies[1:])

    mainHobbies := hobbies[0:2]
    fmt.Printf("The third task %v \n", mainHobbies)

    mainHobbies = mainHobbies[1:3]
    fmt.Print(mainHobbies)

    courseGoals := []string{"Be like prime", "Be awesome"}
    fmt.Println(courseGoals)

    courseGoals[1] = "Be super awesome"
    courseGoals = append(courseGoals, "Be a master")

    product := []Product{

        {
            "First Product",
            "1",
            12.89,
    },
        {
            "Second Product",
            "2",
            22.90,
        },
    }

    fmt.Println(product)

    newProduct := Product{
        "A third product",
        "3",
        49.99,
    }

    product = append(product, newProduct)

    discountedArray := []float64{10.9, 12.88}
    newArray := []float64{1.22, 5.99}
    newArray = append(newArray, discountedArray...)
    fmt.Println(newArray)

}
