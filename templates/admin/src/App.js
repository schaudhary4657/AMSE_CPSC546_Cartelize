import React from 'react';
import Header from './components/common/Navbar';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import Home from './views/Home';
import Login from './views/Login';

function App() {
  return (
    <div className="App">
      <Header></Header>
      <Router>
        <Route exact path="/" component={Home}></Route>
        <Route path="/login" component={Login}></Route>
      </Router>
    </div>
  );
}

export default App;
