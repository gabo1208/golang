import React, { useState, useCallback } from 'react';
import './App.css';
import { Dashboard } from './components/dashboard/Dashboard'
import userImg from './static/img/user.png'

function App() {
  const [chatState, setChatState] = useState({
    menubarClass: "change",
    username: "",
    selectedChat: { messages: [] },
    connectedUsers: []
  })

  const usernameCallback = useCallback(username => {
    if (username !== "") {
      setChatState({ ...chatState, username: username })
    }
  }, [setChatState, chatState])


  const handleMenuClick = () => {
    setChatState({ ...chatState, menuBarClass: getNextChatClass() })
  }

  const getNextChatClass = () => {
    var menuClass = ""
    if (chatState.menuBarClass === "" || !chatState.username) {
      menuClass = "change"
    }

    return menuClass
  }

  return (
    <div className="App">
      <header>
        <div className={`app-container ${chatState.menuBarClass}`} onClick={handleMenuClick}>
          <div className="container cursor-pointer">
            <div className="bar1"></div>
            <div className="bar2"></div>
            <div className="bar3"></div>
          </div>
          <img className="header-img" alt="user-img.png" src={userImg} />
          <div className="app-title">What'sUp</div>
        </div>
      </header>
      <Dashboard
        menuBarStatus={chatState.menuBarClass !== ""}
        chatState={chatState}
        usernameCallback={usernameCallback}
      />
    </div>
  );
}

export default App;
