import axios from "axios"

import {  useState } from "react"
import FormInput from "../FormInput"

type Props = {}


export default function Register({}: Props) {
  const [password, setPassword] = useState("")
  const [repeatPassword, setRepeatPassword] = useState("")

  console.log(password);
  console.log(repeatPassword);
  
  

  async function checkEmail(value:string) : Promise<string> {
    return new Promise((resolve, _)=>{
      axios.get(`/api/user/check-email?value=${value}`).then(res=>{

        var message:APIMessage = res.data


        switch (message.status as string) {
          case "success":
            resolve("")
            break;
          case "missing-values":
            resolve("")
            break;
          case "failure":
            resolve(message.message)
            break;
          default:
            resolve(message.message)
            break;
        }
      })
    })
  }
  


  return (
    <div className="slide-in needs-validation">
      <div className="mb-3">
        <FormInput id={"emailInput"} type={"email"} label={"Email"} validate={checkEmail}/>
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
      <div className="mb-3">
        <label htmlFor="passwordRepeatInput">Repeat Password</label>
        <input
          type="password" 
          id="passwordRepeatInput" 
          className="form-control"
          onChange={(e)=>{setRepeatPassword(e.target.value)}}
          >
        </input>
      </div>
    </div>
  )
}
