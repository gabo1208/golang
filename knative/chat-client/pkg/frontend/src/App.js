import './App.css';
import { Dashboard } from './components/dashboard/Dashboard'
import userImg from './user.png'

function App() {
  return (
    <div className="App">
      <header>
        <div class="container">
          <div class="bar1"></div>
          <div class="bar2"></div>
          <div class="bar3"></div>
        </div>
        <img className="user-img" alt="user-img.png" src={userImg} />
        What'sUp
      </header>
      <Dashboard />
    </div>
  );
}

export default App;
