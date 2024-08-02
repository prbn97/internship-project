import { useEffect, useState } from "react";
import { useParams, useOutletContext, useNavigate } from "react-router-dom";


// in this components we display the task by id
const Task = () => {
    const [task, setTask] = useState({}); // set state with empty object to store the task
    // get the id from the url to do the request
    // using the hook useParams to make this happen
    let { id } = useParams(); // this variable needs to match with the :id in index.js

    const { jwtToken, setAlertClassName, setAlertMessage } = useOutletContext()
    const navigate = useNavigate();
    // useEffect to make the request
    useEffect(() => {
        if (!jwtToken) {
            navigate("/login");
            return;
        } else {
            const headers = new Headers();
            headers.append("Content-Type", "application/json");

            const requestOptions = {
                method: "GET",
                headers: headers,
            }; // get task by id
            fetch(`${process.env.REACT_APP_BACKEND_ADDRESS}/tasks/${id}`, requestOptions)
                .then((response) => response.json())
                .then((data) => {
                    setTask(data);
                    setAlertClassName("d-none")
                    setAlertMessage("")
                })
                .catch(err => {
                    console.log(err);
                    setAlertClassName("alert-danger")
                    setAlertMessage("internal server error")
                });

        }
    }, [id, jwtToken, navigate, setAlertClassName, setAlertMessage]); // pass the id

    const getStatusBadgeClass = (status) => {
        switch (status) {
            case "ToDo":
                return "btn-danger";
            case "Doing":
                return "btn-warning";
            case "Done":
                return "btn-success";
            default:
                return "btn-secondary";
        }
    };

    const handleStatusClick = () => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "PUT",
            headers: headers,
        };

        fetch(`${process.env.REACT_APP_BACKEND_ADDRESS}/tasks/${id}/update`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setTask(data);
            })
            .catch(err => {
                console.log(err);
            });
    };

    return (
        <div className="card mt-3" style={{ boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)' }}
        >
            <div className="card-body">
                <h5 className="card-title">{task.title}</h5>
                <p className="card-text">{task.description}</p>
                <div className="mb-3">
                    <button
                        className={`btn ${getStatusBadgeClass(task.status)} cursor-pointer`}
                        onClick={handleStatusClick}> Status: {task.status}
                    </button>
                </div>
            </div>
        </div>
    )
}

export default Task;