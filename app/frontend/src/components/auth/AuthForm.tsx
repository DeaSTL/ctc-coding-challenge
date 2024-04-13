import { useEffect, useState } from 'react'
import Login from './Login'
import Register from './Register'

type Props = {}

export default function AuthForm({}: Props) {
  const [selectedTab, setSelectedTab] = useState("login")

  useEffect(()=>{
    console.log("set selected changed",selectedTab);
    
  },[selectedTab])
  // big no no, I know ....
  function TabPageControl(){
    switch (selectedTab) {
      case "login":
        return <Login/>
      case "register":
        return <Register/>
      default:
        return ``
    }
  }


  return (
    <div>
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
        <TabPageControl/>
      </div>
    </div>
  )
}



