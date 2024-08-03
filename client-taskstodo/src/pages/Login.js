import { useState } from "react";
import Input from "../components/form/Input"
import { useNavigate, useOutletContext } from "react-router-dom";

const Login = () => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    //with this hook we call the info set in App component
    const { setJwtToken } = useOutletContext();
    const { setAlertClassName } = useOutletContext();
    const { setAlertMessage } = useOutletContext();
    // navigate hook
    const navigate = useNavigate();




    const handleSubmit = (event) => {
        event.preventDefault();
        console.log("email/pass", email, password)

        if (email === "admin@example.com") {
            setJwtToken("abc")
            setAlertClassName("d-none")
            setAlertMessage("")
            navigate("/")
        } else {
            setAlertClassName("alert-danger")
            setAlertMessage("invalid credentials")
        }
    }


    return (
        <div className="col-md-6 offset-md-3">
            <h2>Login</h2>

            <form onSubmit={handleSubmit}>
                <Input
                    title="Email Address"
                    type="email"
                    className="form-control"
                    name="email"
                    autoComplete="email-new"
                    onChange={(event) => setEmail(event.target.value)}
                />

                <Input
                    title="Password"
                    type="password"
                    className="form-control"
                    name="password"
                    autoComplete="password-new"
                    onChange={(event) => setPassword(event.target.value)}
                />

                <input
                    type="submit"
                    className="btn btn-primary"
                    value="Login"
                />





            </form>
        </div>
    )
}

export default Login;