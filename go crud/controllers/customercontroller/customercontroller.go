package customercontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/jeypc/go-crud/entities"
	"github.com/jeypc/go-crud/libraries"
	"github.com/jeypc/go-crud/models"
)

var validation = libraries.NewValidation()
var customermodel = models.NewCustomerModel()

func Index(response http.ResponseWriter, request *http.Request) {

	customer, _ := customermodel.FindAll()

	data := map[string]interface{}{
		"customer": customer,
	}

	temp, err := template.ParseFiles("views/customer/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/customer/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var customer entities.Customer
		customer.NamaLengkap = request.Form.Get("nama_lengkap")
		customer.Merk = request.Form.Get("merk")
		customer.Alamat = request.Form.Get("alamat")
		customer.NomorHp = request.Form.Get("nomor_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(customer)

		if vErrors != nil {
			data["customer"] = customer
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Customer berhasil disimpan"
			customermodel.Create(customer)
		}

		temp, _ := template.ParseFiles("views/customer/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var customer entities.Customer
		customermodel.Find(id, &customer)

		data := map[string]interface{}{
			"customer": customer,
		}

		temp, err := template.ParseFiles("views/customer/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var customer entities.Customer
		customer.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		customer.NamaLengkap = request.Form.Get("nama_lengkap")
		customer.Merk = request.Form.Get("merk")
		customer.Alamat = request.Form.Get("alamat")
		customer.NomorHp = request.Form.Get("nomor_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(customer)

		if vErrors != nil {
			data["customer"] = customer
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Customer berhasil diperbarui"
			customermodel.Update(customer)
		}

		temp, _ := template.ParseFiles("views/customer/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	customermodel.Delete(id)

	http.Redirect(response, request, "/customer", http.StatusSeeOther)

}
