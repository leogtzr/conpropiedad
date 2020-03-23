import React, { Component } from 'react';
import './App.css';
import Home from './Home';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Insultos from './Insultos';

class App extends Component {

  componentDidMount(){
    document.title = "Menú"
  }

  render() {
    return (
      <Router>
        <Switch>
          <Route path='/' exact component={Home}/>
          <Route path='/insultos' component={Insultos} exact />
        </Switch>
      </Router>
    )
  }
}

export default App;