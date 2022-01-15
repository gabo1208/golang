import React from 'react';
import usernameIcon from '../../static/img/username-icon.png'
import logoutIcon from '../../static/img/logout-icon.png'
import './Menu.css'

export function Menu(props) {
  return (
    <div className="menu-column">
      <div className="option-card selected">
        <div className="option-img">
          <img src={usernameIcon} className="option-img" alt="option-img.png" />
        </div>
        <div className="option-info">
          <div className="option">
            <h3 className="option-name"><b>{props.username}</b></h3>
          </div>
        </div>
      </div>
      <div className="option-card">
        <div className="option-img">
          <img src={logoutIcon} className="option-img" alt="option-img.png" />
        </div>
        <div className="option-info">
          <div className="option">
            <h3 className="option-name"><b>Logout</b></h3>
          </div>
        </div>
      </div>
    </div>
  );
}