import React, { useState } from 'react';
import './App.css';
import { Dashboard } from './components/dashboard/Dashboard'
import userImg from './static/img/user.png'

function App() {
  const [userContent, setuserContent] = useState({
    menuBarClass: '',
    userName: 'test'
  })

  const handleMenuClick = (e) => {
    var menuClass = ""
    if (userContent.menuBarClass === "") {
      menuClass = "change"
    }

    setuserContent({ menuBarClass: menuClass })
  }

  return (
    <div className="App">
      <header>
        <div className={`app-container ${userContent.menuBarClass}`} onClick={handleMenuClick}>
          <div className="container">
            <div className="bar1"></div>
            <div className="bar2"></div>
            <div className="bar3"></div>
          </div>
          <img className="header-img" alt="user-img.png" src={userImg} />
          <div className="app-title">What'sUp</div>
        </div>
      </header>
      <Dashboard menuBarStatus={userContent.menuBarClass !== ""} />
    </div>
  );
}

export default App;
