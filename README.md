# MVC iris golang starter

#### Depedency
Install dep

		$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
		$ dep ensure -update
		$ rm -irf vendor/ # There's still bug on go dep for macOS, you should omit / remove completely the vendor directory


#### Setting Env

Buat file .env berdasarkan file ENVIRONMENT_EXAMPLE

#### Membuat Model
Untuk membuat model , masuk ke folder model dian buat file namamodel.go, Lihat contoh 

	package model

	import (
		"time"
	)

	type User struct {
		ID int `json:"id" gorm:"primary_key"`

		FirstName string     `json:"firstname, omitempty" gorm:"not null; type:varchar(100)"`
		LastName  string     `json:"lastname, omitempty" gorm:"not null; type:varchar(100)"`
		Email     string     `json:"email, omitempty" gorm:"not null; type:varchar(100)"`
		CreatedAt *time.Time `json:"createdAt, omitempty"`
		UpdatedAt *time.Time `json:"updatedAt, omitempty"`
		DeletedAt *time.Time `json:"deletedAt, omitempty" sql:"index"`
	}

	func (User) TableName() string {
		return "users" // table name when succesfully migrate
	} 
	
untuk lebih jelasnya silahkan baca documentasi [Gorm Model](http://doc.gorm.io/models.html) 

#### Setting Migrate

Untuk setting automigrate bisa di lihat pada main.go pada fungsi AutoMigrate

	func DBMigrate() {
	fmt.Println("[::] Migration Databases .....")
	db := config.GetDatabaseConnection() // check connection to Databases
	db.AutoMigrate(&model.User{})        // Migrate Model
	//db.AutoMigrate(&model.Profile{})        // Migrate Model
	....
	fmt.Println("[::] Migration Databases Done")
	}
	
####  Setting Route
untuk membuat router silahkan ke folder router ke function Routers

	func Routers() {
		db := config.GetDatabaseConnection()
		inDB := &controller.InDB{DB: db }
		app := iris.Default()
		// for / endpoint
		app.Get("/", controller.WelcomeController)

		// example group: v1
		v1:= app.Party("/v1")
		{
			v1.Post("/user", inDB.CreteUser)
			v1.Get("/user", inDB.GetAll)
			v1.Get("/user/{id : int}", inDB.GetById)
			v1.Put("/user/{id : int}", inDB.UpdateUser)
			v1.Delete("/user/{id : int}", inDB.DeleteUser)
		}

	app.Run(iris.Addr( ":"+os.Getenv("API_PORT"))) // starter handler untuk route
	}
	
Untuk lebih jelasnya silahkan baca documentasi [iris router](https://docs.iris-go.com/routing.html) 

#### Cara running program

	go run main.go // untuk running
	go build main.go // untuk build
	

#### Silahkan coba endpoint berikut ini untuk testing

1. Create user  [POST] localhost:3000/v1/user
	body post 
	
		{
			"firstname" : "Rahmat Wahyu",
			"lastname" : "Hadi",
			"email"		: "hrahmatwahyu@gmial.com"
		}
		
2. Ger all user [GET] localhost:3000/v1/user
3. Get by Id [GET] localhost:3000/v1/user/1{iduser}
4. Update User [PUT] localhost:3000/v1/user/1{iduser}

	body post 
	
		{
			"firstname" : "Rahmat W",
			"lastname" : "Hadi",
			"email"		: "rhmt@gmial.com"
		}
		
5 Delete user [DELETE] localhost:3000/v1/user/1{iduser}



#### Contact
email : hrahmatwahyu@gmail.com, LINE : rahmatwahyuhadi, WA : 6285205039835

