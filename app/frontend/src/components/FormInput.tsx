import {  useEffect, useState } from "react"

type Props = {
  id: string;
  type: string;
  label: string;
  set(value:string):any;
  validate(value:string) : Promise<string>
  placeholder?: string | ""
}

export default function FormInput({type,validate,placeholder,id,label,set}: Props) {
  const [error, setError] = useState("")
  const [input, setInput] = useState("")
  
  useEffect(() => {
    set(input) 
  }, [input])

  function RenderError({errorMessage}:any){
    if(errorMessage != ""){
      return <div className="invalid-feedback">
        {errorMessage} 
      </div>
    }
  }
  useEffect(() => {
    // some debounce
    const timeout = setTimeout(()=>{
      validate(input).then((message)=>{
        console.log("validation message: ",message);
        
        if(message != ""){
          setError(message)
        }else{
          setError("")
        }
      })
    },500)

    return ()=>{
      clearTimeout(timeout)
    }
  }, [input])

  return (
    <div className="mb-3">
      <label htmlFor={id}>{label}</label>
      <input
        type={type}
        id={id}
        className={`form-control ${error != "" ? "is-invalid" : ""}`}
        placeholder={placeholder}
        onChange={(e)=>{setInput(e.target.value)}}
        >
      </input>
      <RenderError errorMessage={error}/> 
    </div>
  )
}
