package customers

import (
  "appengine"
  "appengine/datastore"
  "github.com/campbel/gore"
  "net/http"
  "strconv"
  "errors"
)

func init() {
  gore.Get("/api/customers/", read)
  gore.Post("/api/customers/", create)
  gore.Delete("/api/customers/", delete)

  gore.Get("/api/customers/:customerId/", readById)
  gore.Delete("/api/customers/:customerId/", deleteById)
}

const entity_kind string = "customers"

type Customer struct {
  Id int64
  Name string
}

func read (req *gore.Request, res *gore.Response) {
  c := appengine.NewContext(req.Raw)

  q := datastore.NewQuery(entity_kind)

  var customers []Customer
  _, err := q.GetAll(c, &customers)
  if err != nil {
    res.Error(http.StatusInternalServerError, err)
    return
  }

  res.Send(customers)
}

func readById (req *gore.Request, res *gore.Response) {
  c := appengine.NewContext(req.Raw)

  idStr, exists := req.Params["customerId"]
  if !exists {
    res.Error(http.StatusBadRequest, nil)
    return
  }

  id, err := strconv.ParseInt(idStr, 10, 0)
  if err != nil {
    res.Error(http.StatusInternalServerError, nil)
    return
  }

  key := datastore.NewKey(c, entity_kind, "", id, nil)
  var customer Customer
  if err := datastore.Get(c, key, &customer); err != nil {
    res.Error(http.StatusNotFound, err)
    return
  }
  res.Send(customer)
}

func create (req *gore.Request, res *gore.Response) {
  c := appengine.NewContext(req.Raw)

  var customer Customer
  req.Body(&customer)

  // get an Id
  id, _, err := datastore.AllocateIDs(c, entity_kind, nil, 1)
  if err != nil {
    res.Error(http.StatusInternalServerError, err)
    return
  }
  customer.Id = id

  key := datastore.NewKey(c, entity_kind, "", customer.Id, nil)
  if _, err := datastore.Put(c, key, &customer); err != nil {
    res.Error(http.StatusInternalServerError, err)
    return
  }

  var newCustomer Customer
  if err := datastore.Get(c, key, &newCustomer); err != nil {
      res.Error(http.StatusNotFound, err)
      return
  }
  res.Send(newCustomer)
}

func delete (req *gore.Request, res *gore.Response) {
  c := appengine.NewContext(req.Raw)

  var customers []Customer
  req.Body(&customers)

  keys := make([]*datastore.Key, len(customers))
  for i, customer := range customers {
    keys[i] = datastore.NewKey(c, entity_kind, "", customer.Id, nil)
  }

  if err := datastore.DeleteMulti(c, keys); err != nil {
    res.Error(http.StatusInternalServerError, err)
    return
  }

  res.Send(customers)
}

func deleteById (req *gore.Request, res *gore.Response) {
  c := appengine.NewContext(req.Raw)

  idStr, exists := req.Params["customerId"]
  if !exists {
    res.Error(http.StatusBadRequest, nil)
    return
  }

  id, err := strconv.ParseInt(idStr, 10, 0)
  if err != nil {
    res.Error(http.StatusInternalServerError, nil)
    return
  }

  key := datastore.NewKey(c, entity_kind, "", id, nil)
  if err := datastore.Delete(c, key); err != nil {
      res.Error(http.StatusNotFound, err)
      return
  }

  var customer Customer
  if err := datastore.Get(c, key, &customer); err != nil {
      res.Send("Entity deleted")
      return
  }
  res.Error(http.StatusInternalServerError, errors.New("Entity not deleted"))
}
