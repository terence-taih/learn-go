Udemy course: 
    1. Environment: 
          - Install VS Code and Extension for Golang (View -> Extension, Search Golang)
          - Install Go Analysis tool (right bottom, change Select language mode to Go, then will get the message that we're missing the go tool)
    2. Run: 
          - Use terminal 
          - go run main.go
          - go build: build runnable file (but don't run, use ./main to run)
          - go fmt: format code 
          - go install: compile and install the dependency packages 
          - go get: get source from dependency package 
          - 

    3. Packages: http://golang.org/pkg/ 
        - Executable package: main 
        - Reuseable package: any name 
        - http://golang.org/pkg/ 

    4. Variables: 
      - Static type (java, C++, Go)
      - Dynamic type (Javascript, Python, Ruby)
      - Type: string, int, float, array (slice), map  

      Go is not an OOP language. It doesn't have class and so on. 
      - Custom type: type deck []string ==> same as extend on OOP 
      
    5. Testing 
      - Not mocha, Selenium, ... 
      - File must have _test.go, eg deck_test.go 
      - Run: go test 
      - Vs code supports to run test in package or in file (the hightlight in very top of test file)
      -   

    6. Pointer 
      - Go is language use "Pass by Value"
      - When we call a function of an struct, actualy, Go will create a copy of our object, then we will process on that. 
      - https://www.draw.io/#G15xXsN-49NNOX-hJKE4hLwtgz7kUz0QMh (VERY IMPORTANT)

      - & operator (read: Ampersand) Give me the memory address of the value this variable is pointing at 
            Example: jimPointer := &jim 
                                    RAM: 
                              Address     Value 
                              0001 
        jimPointer ==>        0002        person{firstName: "Jim", lastName: "Anderson"}        <== Jim 

      - * operator (star)  Give me the value this memory address is pointint at
            Example: func (pointerToPerson *person) updateName() {
                  *pointerToPerson.firstName = "Terence"
            }
            + star in before of a type : type description, it means we're working with a pointer to a person 
            + star in before of an actually pointer: this is an operator - it means we want to manupulate the value the pointer is referencing  
      - special case: when the receiver of a function is a pointer, we can call this function with either pointer or value 
            var thai person 
            (&thai).updateName and thai.updateName are both OK      

      - Value Type vs Reference Type (Slice)





   Question: 
      1. How to debug test in VS code 
