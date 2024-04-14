
import { createContext, useContext, useCallback, ReactNode, useState, useEffect } from 'react';
import { SendNotificationFunction } from '../types';

type Props = {
  children:ReactNode
}

const NotificationContext = createContext<SendNotificationFunction>(
(_:string,__:string,___:number=8000)=>{
  console.log("sendNotification not defined")
});

export function NotificationProvider({ children }:Props) {

  const [msg, setMsg] = useState("")
  const [type, setType] = useState("")
  const [delay, setDelay] = useState(8000)

  const sendNotification = useCallback<SendNotificationFunction>((message:string,type:string,delay:number=8000) => {
    console.log("Notification:", message,"Type: ",type);
    setType(type)
    setMsg(message)
    setDelay(delay)
  }, []);

  useEffect(() => {
    if(msg != ""){
      const timeout = setTimeout(()=>{
        setMsg("")
      },delay)

      return ()=>{
        clearTimeout(timeout)
      }
    }
  }, [msg])

  return (
    <NotificationContext.Provider value={sendNotification}>
      <div className={`notification-container slide-in ${msg == "" ? "invisible" : ""}`} id="notifications">
        <div className={`notification ${type}`}>
          {msg}
        </div>
      </div>

      {children}
    </NotificationContext.Provider>
  );
}

export function useNotification() {
  return useContext(NotificationContext);
}
