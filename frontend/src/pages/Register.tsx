import { useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { useAuth } from '../auth/authProvider'
import axios from 'axios'

const Register = ()  => {
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")
    const [usernameError, setUsernameError] = useState("")
    const [passwordError, setPasswordError] = useState("")
    const [confirmPasswordError, setConfirmPasswordError] = useState("")
    
    const navigate = useNavigate();
    const auth = useAuth();
    const {state} = useLocation();
    
    const validateLogin = () => {
        setUsernameError("")
        setPasswordError("")
        setConfirmPasswordError("")

        if ("" === username) {
            setUsernameError("Please enter your username")
            return
        }

        if ("" === password) {
            setPasswordError("Please enter a password")
            return
        }

        if ("" === confirmPassword) {
            setConfirmPasswordError("Please confirm password")
            return
        }

        if (password.length < 7) {
            setPasswordError("The password must be 8 characters or longer")
            return
        }

        if (password !== confirmPassword) {
            setConfirmPasswordError("The passwords must match")
            return
        }

        register()
    }

    const register = async () => {
        let res: any = null
        try {
            res = await axios.post("/api/register", JSON.stringify({username, password}))
        } catch(e) {
            window.alert("Could not register")
            return
        }

        const resData = res.data
        if (res.status === 200) {
            auth?.setToken(resData)
            navigate(state?.path || "/")
        } else {
            window.alert("There was a problem with registering")
        }
    }

    return <div className="relative mainContainer bg-sky-200/90 h-full p-10 rounded-lg">
        <span onClick={() => {navigate('/')}} className="hover:bg-black/0 hover:translate-x-[-0.25em] transition duration-300 cursor-pointer absolute top-5 left-5 text-lg text-black bg-black/20 rounded-lg p-2">
            &#x25c0; Back
        </span>
        <div className="titleContainer text-4xl font-bold text-sky-500">
            <div>Sign in</div>
        </div>
        <br />
        <div className="inputContainer">
            <input
                value={username}
                placeholder="Username"
                onChange={ev => setUsername(ev.target.value)}
                className="inputBox focus:border-sky-600 text-sky-500 bg-sky-200/80 border-2 border-sky-500" />
            <label className="errorLabel">{usernameError}</label>
        </div>
        <br />
        <div className={"inputContainer"}>
            <input
                value={password}
                placeholder="Password"
                onChange={ev => setPassword(ev.target.value)}
                className="inputBox focus:border-sky-600 text-sky-500 border-2 bg-sky-200/80 border-sky-500" />
            <label className="errorLabel">{passwordError}</label>
        </div>
        <br />
        <div className={"inputContainer"}>
            <input
                value={confirmPassword}
                placeholder="Confirm password"
                onChange={ev => setConfirmPassword(ev.target.value)}
                className="inputBox focus:border-sky-600 text-sky-500 border-2 bg-sky-200/80 border-sky-500" />
            <label className="errorLabel">{confirmPasswordError}</label>
        </div>
        <br />
        <div className="">
            <input
                className="hover:bg-sky-700 bg-sky-500 px-20 py-4 text-xl rounded-lg cursor-pointer"
                type="button"
                onClick={validateLogin}
                value={"Sign in"} />
        </div>
        <div className="mt-2">
            <p className="text-body-2 text-black">
                Already have an account? 
                <a className="text-sky-500 cursor-pointer hover:text-sky-800" onClick={() => navigate('/login')}> Log in</a>
            </p>
        </div>
    </div>
}

export default Register
