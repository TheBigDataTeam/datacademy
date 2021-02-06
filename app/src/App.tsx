import React from 'react';
import { BrowserRouter as Router, Route} from 'react-router-dom';
import { Landing, SignUp } from './components/pages';
import { Courses } from './components/pages';

function App() {
  return (
    <Router>
      <Route exact path="/" component={Landing} />
      <Route path="/signup" component={SignUp} />
      <Route path="/courses" component={Courses} />
    </Router>
  );
}

export default App;
