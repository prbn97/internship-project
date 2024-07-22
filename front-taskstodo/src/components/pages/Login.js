import { useState } from 'react';
import { useNavigate, useOutletContext } from 'react-router-dom';

function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();
    const { setUser, setJwtToken } = useOutletContext();

    const handleLogin = () => {
        // Simulate a login and set the JWT token and user info
        const fakeJwtToken = "fake-jwt-token";
        const fakeUser = { name: "Paulo", email: "paulo@example.com" };
        setJwtToken(fakeJwtToken);
        setUser(fakeUser);
        navigate("/");
    };

    return (
        <div className="container">
            <h2>Login</h2>
            <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Username"
            />
            <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Password"
            />
            <button onClick={handleLogin}>Login</button>
        </div>
    );
}

export default Login;
