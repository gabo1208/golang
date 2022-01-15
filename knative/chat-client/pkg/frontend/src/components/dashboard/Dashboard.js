import React from 'react';
import './Dashboard.css'
import { Contacts } from '../contacts/Contacts'
import { Chat } from '../chat/Chat'
import { Menu } from '../menu/Menu'

export function Dashboard(props) {
  return (
    <div className="dashboard">
      <div className="app-container bg">
        <div className="side-bar">
          {props.menuBarStatus ? <Menu /> : <Contacts />}
        </div>
        <div className="chat">
          <Chat />
        </div>
      </div>
    </div>
  );
}