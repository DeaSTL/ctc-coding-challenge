import { useState } from "react"
import axios from "axios"
import { APIMessage } from "../../types";
import { useNotification } from "../Notifications";

type Props = {
  successfulLogin():any;
  visible:boolean;
}

export default function Login({visible,successfulLogin}: Props) {

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const sendNotification = useNotification();
  function submitForm(){
    axios.post('/api/user/login',{email,password}).then(res=>{
      var data:APIMessage = res.data
      switch(data.status as string){
        case "success":
          sendNotification(data.message,"success")
          successfulLogin()
          break;
        case "failure":
          sendNotification(data.message,"error")
          break;
        default:
          sendNotification(data.message,"error")
          break;
      }
    })
    console.log(email,password);
     
  }
  
  return (
    <div className={`slide-in ${!visible ? "invisible" : ""}`} id="loginForm">
      <div className="mb-3">
        <label htmlFor="emailInput">Email</label>
        <input
          type="email" 
          id="emailInput" 
          className="form-control" 
          placeholder="name@example.com"
          onChange={(e)=>{setEmail(e.target.value)}}
          >
        </input>
      </div>
      <div className="mb-3">
        <label htmlFor="passwordInput">Password</label>
        <input
          type="password" 
          id="passwordInput" 
          className="form-control"
          onChange={(e)=>{setPassword(e.target.value)}}
          >
        </input>
      </div>
      <button className="btn btn-primary" onClick={submitForm}>Login</button>
    </div>
  )
}
