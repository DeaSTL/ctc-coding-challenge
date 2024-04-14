import { createContext, useEffect, useState } from 'react';
import axios from 'axios'
import Cookies from 'js-cookie'
import './App.css'
import {AuthForm} from './components/auth'
import { APIMessage, APIUserRequest, NewUser as NewSessionUser, SessionUser } from './types';
import { useNotification } from './components/Notifications';





function App() {
  const sendNotification = useNotification();
  const [user, setUser] = useState<SessionUser>(NewSessionUser("",false))
  function getUser(){
    var bearerToken = Cookies.get('auth_token')

    if(!bearerToken){
      return 
    }
    
    axios.get("/api/user",{
        headers: {
          'Authorization' : 'Bearer ' + bearerToken
        }
      }).then(res=>{
      if(res.status == 200){
        var data = res.data as APIUserRequest
        
        if(data.data){
          if(data.data.email)
          setUser(NewSessionUser(data.data.email,true))
          sendNotification("Logging in...","success",3000)
        }else {
          // TODO delete the invalid token
        }
      }
    })

  }
  useEffect(() => {
    getUser()
  }, [])

  function successfulLogin(){
    getUser()
  }

  function logout() {
    sendNotification("Logging out...","success",3000)
    Cookies.remove('auth_token') 
    setUser(NewSessionUser("",false))
  }
  return (
    <div className="container">
      <div className="row">
        <div className="mt-5">
          {user.loggedIn ? <span className="d-flex slide-in">
          <span className="mr-2 my-auto">Logged in as: {user.email}</span>
          <button className="btn btn-outline-danger btn-sm ms-5" onClick={logout}>Logout</button>
          </span> : ``}
        </div>
      </div>
      <div className="row">
        <div className="col-md">
          <AuthForm successfulLogin={successfulLogin}/>
        </div>
      </div>
    </div>
  )
}

export default App
