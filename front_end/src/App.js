import React from 'react';
import { BrowserRouter, Switch, Route, useParams } from 'react-router-dom';
import './App.css';
import Login from "./component/auth/login.js";

class App extends React.Component{
  constructor(props){
    super(props)
     this.state = {
       id : null
    }
  }

  render(){
    console.log(this.props)
    return(
      <div>
        <BrowserRouter>
        <Switch>
           <Route path="/">
             <Login />
           </Route>
         </Switch>
         </BrowserRouter>
      </div>
    )
  }
}
export default App;
