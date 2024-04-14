export type APIMessage = {
  status: string;
  message: string;
  data: Map<string,any>;
}

export type APIUserRequest = {
  status: string;
  message: string;
  data: {
    exp: number;
    email: string;
  };
}

export type SessionUser = {
  email: string;
  loggedIn:boolean;
}


export function NewUser(email:string | "",loggedIn:boolean):SessionUser {
  return {email:email,loggedIn:loggedIn}
}

export type SendNotificationFunction = (message: string, type:string,timeout?:number) => void;

// class Input{
//   constructor(public value: string,public valid: boolean) {
//     this.value = value
//     this.valid = valid
//   }
// }
