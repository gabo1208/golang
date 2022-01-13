import React from 'react';
import './Dashboard.css'
import { Contacts } from '../contacts/Contacts'
import { Chat } from '../chat/Chat'

export function Dashboard() {
  return (
    <div className="dashboard">
      <div className="contacts">
        <Contacts />
      </div>
      <div className="chat">
        <Chat />
      </div>
    </div>
  );
}