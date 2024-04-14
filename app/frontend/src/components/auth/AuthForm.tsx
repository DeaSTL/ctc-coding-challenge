import { useEffect, useState } from 'react'
import Login from './Login'
import Register from './Register'

type Props = {
  successfulLogin():void
}

export default function AuthForm({successfulLogin}: Props) {
  const [selectedTab, setSelectedTab] = useState("login")


  return (
    <div className="mt-4">
      <div className="card p-4">
        <ul className="nav nav-underline"> 
          <li className="nav-item"> 
            <a 
            className={"nav-link " + (selectedTab == "login" ? "active" : "")}
            onClick={()=>{setSelectedTab("login")}}
            aria-current="page"
            href="#"
            >Login</a>
          </li>
          <li className="nav-item"> 
            <a 
            className={"nav-link " + (selectedTab == "register" ? "active" : "")} 
            onClick={()=>{setSelectedTab("register")}}
            aria-current="page"
            href="#"
            >Register</a>
          </li>
        </ul>
        <Login successfulLogin={successfulLogin} visible={selectedTab == "login"}/>
        <Register visible={selectedTab == "register"}/>
      </div>
    </div>
  )
}



