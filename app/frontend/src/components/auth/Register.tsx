import axios from "axios"

import {  useState } from "react"
import FormInput from "../FormInput"
import { APIMessage } from "../../types"
import { useNotification } from "../Notifications"

type Props = {
  visible:boolean;
}


export default function Register({visible}: Props) {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [__, setRepeatPassword] = useState("")

  const [emailValid, setEmailValid] = useState(false)
  const [passwordValid, setPasswordValid] = useState(false)
  const [passwordRepeatValid, setPasswordRepeatValid] = useState(false)

  const sendNotification = useNotification()


  function submitForm(){
    axios.post(`/api/user/register`,{email,password}).then(res=>{
      if(res.status == 200){
        var data = res.data as APIMessage
        if(data.status == "success") {
          sendNotification(data.message,"success")    
        }else{
          sendNotification(data.message,"error")    
        }
      }
    })
  }

  async function checkEmail(value:string) : Promise<string> {
    return new Promise((resolve, _)=>{
      axios.get(`/api/user/check-email?value=${value}`).then(res=>{

        var message:APIMessage = res.data


        switch (message.status as string) {
          case "success":
            resolve("")
            setEmailValid(true)
            break;
          case "missing-values":
            resolve("")
            setEmailValid(false)
            break;
          case "failure":
            resolve(message.message)
            setEmailValid(false)
            break;
          default:
            resolve(message.message)
            setEmailValid(false)
            break;
        }
      })
    })
  }
  async function checkPassword(value:string) : Promise<string> {
    return new Promise((resolve, _) => { 
      if(value.length == 0 || value.length > 11){
        resolve("")
        setPasswordValid(true)
      }else {
        setPasswordValid(false)
        resolve("Your password must be at least 12 characters long")
      }
    })
  }

  async function checkRepeatPassword(value:string) : Promise<string> {
    return new Promise((resolve, _) => {
      console.log(value,password);
      
      if(value != password){
        resolve("Your password do not match")
        setPasswordRepeatValid(false)
      }else {
        resolve("")
        setPasswordRepeatValid(true)
      }
    })
  }
  


  return (
    <div className={`slide-in needs-validation ${!visible ? "invisible" : ""}`} id="registrationForm">
      <div className="mb-3">
        <FormInput id={"emailInput"} type={"email"} label={"Email"} validate={checkEmail} set={setEmail}/>
      </div>
      <div className="mb-3">
        <FormInput id={"passwordInput"} type={"password"} label={"Password"} validate={checkPassword} set={setPassword}/>
      </div>
      <div className="mb-3">
        <FormInput id={"repeatPasswordInput"} type={"password"} label={"Repeat Password"} validate={checkRepeatPassword} set={setRepeatPassword}/>
      </div>

      <button className="btn btn-primary" disabled={!(passwordValid && passwordRepeatValid && emailValid)} onClick={submitForm}>Register</button>
    </div>
  )
}
