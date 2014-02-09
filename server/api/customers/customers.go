package customers

import (
  "appengine"
  "appengine/datastore"
  "github.com/campbel/gore"
  "net/http"
)

func init() {
  gore.Get("/api/customers/", read)
  gore.Post("/api/customers/", create)
}

type Customer struct {
  Id int64
  Name string
}

func read (req *gore.Request, res *gore.Response) {
  c := appengine.NewContext(req.Raw)

  q := datastore.NewQuery("customers")

  var customers []Customer
  _, err := q.GetAll(c, &customers)
  if err != nil {
    res.Error(http.StatusInternalServerError, err)
    return
  }

  res.Send(customers)
}

func create (req *gore.Request, res *gore.Response) {
  c := appengine.NewContext(req.Raw)

  var customer Customer
  req.Body(&customer)

  // get an Id
  id, _, err := datastore.AllocateIDs(c, "customers", nil, 1)
  if err != nil {
    res.Error(http.StatusInternalServerError, err)
    return
  }
  customer.Id = id

  key := datastore.NewKey(c, "customers", "", customer.Id, nil)
  if _, err := datastore.Put(c, key, &customer); err != nil {
    res.Error(http.StatusInternalServerError, err)
    return
  }

  var newCustomer Customer
  if err := datastore.Get(c, key, &newCustomer); err != nil {
      res.Error(http.StatusInternalServerError, err)
      return
  }

  res.Send(newCustomer)
}
