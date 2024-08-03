import { Link } from "react-router-dom";

const LoginReminder = () => {
    return (
        <div className="alert alert-warning" role="alert">
            Please <Link to="/login">log in</Link> to see your tasks.

        </div>
    )
};

export default LoginReminder;