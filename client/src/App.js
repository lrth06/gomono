import React from "react";
import { Layout } from "./components/layout";
import { Home } from "./components/home";
import './scss/App.scss'
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
export default function App() {
  return (
    <div className="App">
     <div className="App-header">
      <Router>
         <Layout />
        <Switch>
          <div>
        <Route path="/" exact component={Home} />
          </div>
        </Switch>
      </Router>
      </div>
    </div>
  );
}
