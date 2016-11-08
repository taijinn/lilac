package main

import "net/http"

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "Supply",
    "GET",
    "/supply",
    Supply,
  },
  Route{
    "Demand",
    "GET",
    "/demand",
    Demand,
  },
  Route{
    "Ask",
    "POST",
    "/ask",
    Ask,
  },
  Route{
    "Offer",
    "POST",
    "/offer",
    Offer,
  },
  Route{
    "Login",
    "GET",
    "/login",
    Login,
  },
  Route{
    "Friends",
    "GET",
    "/friends",
    Friends,
  },
  Route{
    "Follow",
    "GET",
    "/follow",
    Follow,
  }
}
