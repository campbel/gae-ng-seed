package customers

import (
  "appengine"
  "appengine/datastore"
  "github.com/campbel/gore"
  "net/http"
  "strconv"
  "log"
)

func init() {
  gore.Get("/api/customers/:customerId", readById)
  //gore.Put("/api/customers/:customerId", update)
  //gore.Delete("/api/customers/:customerId", delete)
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
  log.Println(id)

  key := datastore.NewKey(c, "customers", "", id, nil)
  var customer Customer
  if err := datastore.Get(c, key, &customer); err != nil {
      res.Error(http.StatusNotFound, err)
      return
  }
  res.Send(customer)
}

/*
func update (req *gore.Request, res *gore.Response) {

}

func delete (req *gore.Request, res *gore.Response) {

}
*/
