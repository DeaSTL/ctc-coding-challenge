
type Props = {}

export default function Login({}: Props) {
  return (
    <div className="slide-in">
      <div className="mb-3">
        <label htmlFor="emailInput">Email</label>
        <input
          type="email" 
          id="emailInput" 
          className="form-control" 
          placeholder="name@example.com">
        </input>
      </div>
      <div className="mb-3">
        <label htmlFor="passwordInput">Password</label>
        <input
          type="password" 
          id="passwordInput" 
          className="form-control">
        </input>
      </div>
    </div>
  )
}
