package handlers

import (
   "depositregws/api"
   "depositregws/dao"
   "depositregws/logger"
   "fmt"
   "net/http"
)

//
// OptionsGet -- get options request handler
//
func OptionsGet(w http.ResponseWriter, r *http.Request) {

   optionsSet, err := dao.DB.GetOptionsSet( )
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
      status := http.StatusInternalServerError
      encodeOptionsResponse(w, status,
         fmt.Sprintf("%s (%s)", http.StatusText(status), err),
         nil)
      return
   }

   options := createOptions( optionsSet )

   status := http.StatusOK
   encodeOptionsResponse(w, status, http.StatusText(status), options)
}

func createOptions( pairs [] dao.StringPair ) []api.Options {

   results := make([]api.Options, 0)
   for _, v := range pairs {
      ix := indexOf( results, v.A )
      if ix >= 0 {
         results[ ix ].Degrees = append( results[ ix ].Degrees, v.B )
      } else {
         results = append( results, api.Options{ Department: v.A, Degrees: []string{ v.B } })
      }
   }
   return( results )
}

func indexOf( options []api.Options, option string ) int {
   for ix, v := range options {

      if v.Department == option {
         return ix
      }
   }
   // not found
   return -1
}

//
// end of file
//
